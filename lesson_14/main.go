package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main(){
	//ctx := context.Background()
	//context.TODO()
	//ctx, cancel := context.WithCancel(ctx)
	//cancel()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	go func(){
	err := cancelRequest(ctx)
	if err != nil{
		cancel()
		}
	}()

	doREquest(ctx, "https://ya.ru")
}

func cancelRequest(ctx context.Context) error{
	time.Sleep(5000 * time.Millisecond)
	return fmt.Errorf("fal request")
}

func doREquest(ctx context.Context, requestStr string){
	req, _ := http.NewRequest(http.MethodGet, requestStr, nil)
	req = req.WithContext(ctx)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("Responce completed, status code = %d", res.StatusCode)
	case <-ctx.Done():
		fmt.Println("request too long")
	}
}



