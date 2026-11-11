package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func request(ctx context.Context, uri string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil) // muitas funções no go tem a variante "WithContext"
	// na vdd muitas delas são abtrações que chamam sua própria variante withContext por baixo dos panos
	if err != nil {
		return "", err
	}

	time.Sleep(time.Second * 6)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	uri := "https://jsonplaceholder.typicode.com/posts/1"
	response, err := request(ctx, uri)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("context cancelado por deadline")
		} else if ctx.Err() == context.Canceled {
			fmt.Println("context cancelado manualmente")
		} else {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(response)
}
