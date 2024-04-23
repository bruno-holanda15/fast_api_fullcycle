package main

import (
	// "encoding/json"
	// "errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	chViaCEP := make(chan string)
	chBrasilAPI := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		res := reqToAPIs("https://viacep.com.br/ws/01310100/json/")
		chViaCEP <- res
	}()

	go func() {
		time.Sleep(time.Second * 2)
		res := reqToAPIs("https://brasilapi.com.br/api/cep/v1/01310100")
		chBrasilAPI <- res
	}()

	select {
	case res := <-chViaCEP:
		msg := fmt.Sprintf("API %s \nResponse: %s", "ViaCEP", res)
		fmt.Println(msg)
	case res := <-chBrasilAPI:
		msg := fmt.Sprintf("API %s \nResponse: %s", "BrasilAPI", res)
		fmt.Println(msg)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
	}
}

func reqToAPIs(url string) string {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// if res.StatusCode != 200 {
	// 	statusCode := res.StatusCode
	// 	panic(errors.New(fmt.Sprintf("Status code is %d", statusCode)))
	// }

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
