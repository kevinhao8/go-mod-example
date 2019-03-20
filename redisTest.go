package main

import (
	"net/http"
	"time"

	"dbredis"

	"github.com/gin-gonic/gin"
)

func func1(c *gin.Context) {
	// 回复一个200OK,在client的http-get的resp的body中获取数据
	c.String(http.StatusOK, "test1 OK")
}

func main() {
	dbcon := dbredis.DBConnetion{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DBIndex:  0,  // use default DB
		PoolSize: 5,
	}

	dbcon.InitConnetion()

	router := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	router.GET("/test1", func1)
	s.ListenAndServe()

	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "127.0.0.1:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// 	PoolSize: 5,
	// })

	// pong, err := client.Ping().Result()
	// fmt.Println(pong, err)

	// err = client.Set("key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := client.Get("key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// val2, err := client.Get("key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exists")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
}
