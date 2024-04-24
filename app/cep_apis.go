package app

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func CallAPIs(cep string) {
	chViaCEP := make(chan string)
	chBrasilAPI := make(chan string)

	go func() {
		url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
		res := reqToAPIs(url)
		chViaCEP <- res
	}()

	go func() {
		url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
		res := reqToAPIs(url)
		chBrasilAPI <- res
	}()

	select {
	case res := <-chViaCEP:
		msg := fmt.Sprintf("API %s \nResponse: %s", "ViaCEP", res)
		fmt.Println(msg)
	case res := <-chBrasilAPI:
		msg := fmt.Sprintf("API %s \nResponse: %s", "BrasilAPI", res)
		fmt.Println(msg)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}

func reqToAPIs(url string) string {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
