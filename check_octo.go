package main

import (
	"fmt"
	"net/http"
)

var cookie []*http.Cookie

func main() {

	ck := &http.Cookie{
		Name:   "gip",
		Value:  "123",
		MaxAge: 300,
	}

	for i := 0; i < 80; i++ {
		url := "https://xco.news"
		method := "GET"

		//payload := strings.NewReader(`{"query":"ООО \"ДЛЛ ЛИЗИНГ\"","filter":{"min_score":97.0,"source_types":["*"]}}`)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.AddCookie(ck)
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		cookie = res.Cookies() //save cookies
		defer res.Body.Close()
		/*
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
		*/
		fmt.Println(res.StatusCode)
		for _, cookie := range res.Cookies() {
			fmt.Println("Found a cookie named:", cookie.Name)
		}

	}
}
