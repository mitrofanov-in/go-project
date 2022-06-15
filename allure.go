package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var cookie []*http.Cookie

var tmpCSRF *http.Cookie

type Allure struct {
	Results []Results `json:"results"`
}

type Results struct {
	FileName      string `json:"file_name"`
	ContentBase64 string `json:"content_base64"`
}

var content []Results

func GetAuthParmas(payload_lgn []byte) []*http.Cookie {
	url_lgn := "http://allure.back.wd.xco.devel.ifx/login"
	method := "POST"

	client := &http.Client{}
	req_lgn, err := http.NewRequest(method, url_lgn, bytes.NewBuffer(payload_lgn))

	if err != nil {
		fmt.Println(err)
	}
	req_lgn.Header.Add("Content-Type", "application/json")

	res_lgn, err := client.Do(req_lgn)
	if err != nil {
		fmt.Println(err)
	}
	cookie = res_lgn.Cookies() //save cookies
	defer res_lgn.Body.Close()

	fmt.Println("PRINT ALL COOCKIE")
	for _, c := range cookie {
		fmt.Println(c.Name, c.Value)

	}

	//SAVE COOCKIE CSRF

	return cookie
}

func SendResult(jStr []byte) {

	url_send := "http://allure.back.wd.xco.devel.ifx/send-results?project_id=default"
	method_send := "POST"

	client := &http.Client{}
	req_send, err := http.NewRequest(method_send, url_send, bytes.NewBuffer(jStr))
	if err != nil {
		panic(err)
	}
	for i := range cookie {
		req_send.AddCookie(cookie[i])

	}
	req_send.Header.Add("X-CSRF-TOKEN", tmpCSRF.Value)
	req_send.Header.Add("Content-Type", "application/json")
	resp_send, err := client.Do(req_send)
	if err != nil {
		panic(err)
	}
	defer req_send.Body.Close()

	body_send, err := ioutil.ReadAll(resp_send.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("INFO ABOUT SEND FILES", string(body_send), resp_send.StatusCode)

}

func main() {

	payload_lgn := []byte(`{ "username": "xco","password": "xco_interf@x" }`)
	fmt.Println(string(payload_lgn))
	GetAuthParmas(payload_lgn)

	tmpCSRF = cookie[1]
	fmt.Println("COOCKA", tmpCSRF.Value)

	var filesToSend []string
	var base64ToSend []string

	var bs64 string
	tmp, _ := ioutil.ReadDir("./allure-result/")
	for _, t := range tmp {
		if !t.IsDir() {
			filesToSend = append(filesToSend, t.Name())
			bs64 = base64.StdEncoding.EncodeToString([]byte(t.Name()))
			base64ToSend = append(base64ToSend, bs64)
		}
	}
	fmt.Println("1строка", filesToSend)
	fmt.Println("2строка", base64ToSend)

	xLen := len(filesToSend)

	for x := 0; x < xLen; x++ {
		content = append(content, Results{FileName: filesToSend[x], ContentBase64: base64ToSend[x]})
	}

	Allure := Allure{
		Results: content,
	}

	jsonDataSort, _ := json.Marshal(Allure)

	jStr := []byte(jsonDataSort)

	fmt.Println(string(jStr))

	SendResult(jStr)

}
