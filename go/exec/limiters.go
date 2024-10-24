package main

import (
	"context"
	"exec/go/exec/limit"
	mhttp "exec/go/exec/limit/middleware/stdlib"
	sredis "exec/go/exec/limit/redis"
	"fmt"
	"log"
	"net/http"
	"time"

	libredis "github.com/redis/go-redis/v9"
)

type Tele interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *libredis.StatusCmd
}

func NewTele(tele Tele) {
	fmt.Println(tele.Set(context.Background(), "huyufan", "1212", 0))
}

func main() {
	// rate := limit.Rate{
	// 	Period: 1 * time.Second,
	// 	Limit:  10,
	// }

	rate, err := limit.NewRateFromFormatted("2-M")

	if err != nil {
		log.Fatal(err)
		return
	}
	option, err := libredis.ParseURL("redis://localhost:6379/0")

	if err != nil {
		log.Fatal(err)
		return
	}

	client := libredis.NewClient(option)
	// f := client.Set(context.Background(), "test", 1, 0)
	// fmt.Println(f)

	NewTele(client)

	store, err := sredis.NewStoreWithOptions(client, limit.StoreOptions{
		Prefix:   "limit-http",
		MaxRetry: 3,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	//store := memory.NewStore()
	middleware := mhttp.NewMiddleware(limit.New(store, rate, limit.WithTrustForwardHeader(true)))
	http.Handle("/", middleware.Handler(http.HandlerFunc(index)))
	fmt.Println("Server is running on port 7777...")
	http.ListenAndServe(":7777", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;chartset=utf-8")
	_, err := w.Write([]byte(`{"meaasge":"ok}`))
	if err != nil {
		log.Fatalln(err)
	}
}
