package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/buger/jsonparser"
	"github.com/e-commerce-microservices/tiki-crawler-service/httpclient"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	hc := httpclient.NewClient(httpclient.WithMaxConns(1000))
	ctx := context.Background()
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	allProducts := make(map[int]Product)

	for i := 1; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resp := hc.Do(ctx, httpclient.Args{
				URL:    fmt.Sprintf("https://tiki.vn/api/personalish/v1/blocks/listings?limit=50&category=1789&page=%d", i),
				Method: http.MethodGet,
			})
			if resp.Err != nil {
				log.Println(resp.Err)
			}
			offset, err := jsonparser.ArrayEach(resp.Body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				product := Product{}
				err = jsoniter.Unmarshal(value, &product)
				if err == nil {
					mu.Lock()
					allProducts[product.ID] = product
					mu.Unlock()
				}

			}, "data")
			if err != nil {
				log.Printf("%v occur at offset %d", err, offset)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println(allProducts)
}
