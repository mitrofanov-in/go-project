package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	for i := 0; i < 1000; i++ {
		url := "http://octo.prod.wd.xco.devel.ifx/api/v2/companies/search"
		method := "POST"

		payload := strings.NewReader(`{"query":"ООО \"ДЛЛ ЛИЗИНГ\"","filter":{"min_score":97.0,"source_types":["*"]}}`)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body), res.StatusCode)
	}
}
