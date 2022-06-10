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

	for i := 0; i < 130; i++ {
		url := "https://www.boostra.ru"
		method := "GET"

		//payload := strings.NewReader(`{"query":"ООО \"ДЛЛ ЛИЗИНГ\"","filter":{"min_score":97.0,"source_types":["*"]}}`)
		//////payload := strings.NewReader(`{ "query": "Забродин Геннадий Анатольевич", "birthdates": ["", "1970.07.18"] }`)
		//payload := strings.NewReader(`{ "query": "Магомедов Чеэрмах Абдулаевич", "roles": ["sanction","ex_sanction","watch_list","ex_watch_list"], "birthdates": ["", "1986.10.03"] }`)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.AddCookie(ck)
		req.Header.Add("Content-Type", "application/json")
		//req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")
		//req.Header.Add("referer", "https://m.avito.ru/")

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

		//body, err := ioutil.ReadAll(res.Body)

		if res.StatusCode == 200 {

			fmt.Println(res.StatusCode) //, string(body), res.Request.UserAgent(), "steps", i)
		}
		for _, cookie := range res.Cookies() {
			fmt.Println("Found a cookie named:", cookie.Name)
		}

	}
}
