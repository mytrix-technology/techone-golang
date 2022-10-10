package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	contextParent := context.Background()
	ctx1 := context.WithValue(contextParent, "key1", "hello world")
	ctx2 := context.WithValue(ctx1, "key2", "hello girls")
	ctx3 := context.WithValue(ctx2, "key3", "hello Boys")
	ctx4 := context.WithValue(contextParent, "key4", "hello children")
	ctx5 := context.WithValue(ctx1, "key5", "hello Later")
	fmt.Println(ctx5.Value("key5"))
	fmt.Println(ctx5.Value("key4"))
	fmt.Println(ctx5.Value("key3"))
	fmt.Println(ctx5.Value("key2"))
	fmt.Println(ctx5.Value("key1"))
	fmt.Println("========")
	fmt.Println(ctx4.Value("key5"))
	fmt.Println(ctx4.Value("key4"))
	fmt.Println(ctx4.Value("key3"))
	fmt.Println(ctx4.Value("key2"))
	fmt.Println(ctx4.Value("key1"))
	fmt.Println("========")
	fmt.Println(ctx3.Value("key5"))
	fmt.Println(ctx3.Value("key4"))
	fmt.Println(ctx3.Value("key3"))
	fmt.Println(ctx3.Value("key2"))
	fmt.Println(ctx3.Value("key1"))
	fmt.Println("========")
	fmt.Println(ctx2.Value("key5"))
	fmt.Println(ctx2.Value("key4"))
	fmt.Println(ctx2.Value("key3"))
	fmt.Println(ctx2.Value("key2"))
	fmt.Println(ctx2.Value("key1"))
	fmt.Println("========")
	fmt.Println(ctx1.Value("key5"))
	fmt.Println(ctx1.Value("key4"))
	fmt.Println(ctx1.Value("key3"))
	fmt.Println(ctx1.Value("key2"))
	fmt.Println(ctx1.Value("key1"))

	c := context.Background()
	ctx, cancel := context.WithTimeout(c, time.Second*2) //change this second value, and see the different. Try with 7, 15 and see the different
	defer cancel()

	fbreceiver := make(chan string)
	dbreceiver := make(chan string)
	go GetDataFromFacebook(ctx, fbreceiver)
	go GetDataFromDatabase(ctx, dbreceiver)
	defer close(fbreceiver)
	defer close(dbreceiver)
	for i := 0; i < 2; i++ {
		select {
		case fb := <-fbreceiver:
			fmt.Println(">>>>>> Data Received From: ", fb)
		case db := <-dbreceiver:
			fmt.Println(">>>>>> Data Received From: ", db)
		case <-ctx.Done():
			fmt.Println(">>>>> Timeout to process")
		}
	}

	time.Sleep(time.Second * 50) // Just To simulate the context
}

func GetDataFromFacebook(ctx context.Context, datareceiver chan string) {
	startTime := time.Now()
	ticker := time.NewTicker(time.Second * 1)

	for _ = range ticker.C {
		fmt.Println("Still Processing From FB")
		if time.Since(startTime).Seconds() >= 5 { // Ex:Fetch Data needs 5 seconds
			ticker.Stop()
			break
		}
	}

	if ctx.Err() == nil { // ctx.Err() will filled when the ctx is timeout or the ctx canceled
		datareceiver <- "facebook"
	}

	return
}

func GetDataFromDatabase(ctx context.Context, datareceiver chan string) {
	startTime := time.Now()
	ticker := time.NewTicker(time.Second * 1)

	for _ = range ticker.C {
		fmt.Println("Still Processing From DB")
		if time.Since(startTime).Seconds() >= 10 { // Example: Fetch Data needs 10 detik
			ticker.Stop()
			break
		}
	}

	if ctx.Err() == nil { // ctx.Err() will filled when the ctx is timeout or the ctx canceled
		datareceiver <- "database"
	}

	return
}
