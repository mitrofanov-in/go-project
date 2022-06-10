package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var cookie []*http.Cookie

var tmpCSRF *http.Cookie

type Allure struct {
	Results []struct {
		FileName      string `json:"file_name"`
		ContentBase64 string `json:"content_base64"`
	} `json:"results"`
}

func main() {

	url_lgn := "http://allure.back.wd.xco.devel.ifx/login"
	method := "POST"

	payload_lgn := strings.NewReader(`{ "username": "xco","password": "xco_interf@x" }`)

	client := &http.Client{}
	req_lgn, err := http.NewRequest(method, url_lgn, payload_lgn)

	if err != nil {
		fmt.Println(err)
		return
	}
	req_lgn.Header.Add("Content-Type", "application/json")

	res_lgn, err := client.Do(req_lgn)
	if err != nil {
		fmt.Println(err)
		return
	}
	cookie = res_lgn.Cookies() //save cookies
	defer res_lgn.Body.Close()

	for _, c := range cookie {
		fmt.Println(c.Name, c.Value)

	}

	url_proj := "http://allure.back.wd.xco.devel.ifx/projects"
	method_proj := "POST"

	payload_proj := strings.NewReader(`{ "id": "default" }`)

	req_proj, err := http.NewRequest(method_proj, url_proj, payload_proj)
	if err != nil {
		panic(err)
	}
	for i := range cookie {
		req_proj.AddCookie(cookie[i])

	}
	tmpCSRF := cookie[1]
	req_proj.Header.Add("X-CSRF-TOKEN", tmpCSRF.Value)
	req_proj.Header.Add("Content-Type", "application/json")
	resp_proj, err := client.Do(req_proj)
	if err != nil {
		panic(err)
	}
	defer resp_proj.Body.Close()

	body, err := ioutil.ReadAll(resp_proj.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body), resp_proj.StatusCode)

	url_send := "http://allure.back.wd.xco.devel.ifx/send-results?project_id=default"
	method_send := "POST"

	var filesToSend []string
	//var base64ToSend []string

 var x string
 var yx string
	tmp, _ := ioutil.ReadDir("./allure-result/")
	for _, t := range tmp {
		if !t.IsDir() {
			filesToSend = append(filesToSend, t.Name())
		}
	}

	sendResult:=Allure{
		Results:[]Results{
        FileName: x,
		ContentBase64: yx
		}
		}
		

	req_send, err := http.NewRequest(method_send, url_send, payload_send)
	if err != nil {
		panic(err)
	}
	for i := range cookie {
		req_send.AddCookie(cookie[i])

	}
	req_send.Header.Add("X-CSRF-TOKEN", tmpCSRF.Value)
	req_send.Header.Add("Content-Type", "multipart/form-data")
	resp_send, err := client.Do(req_send)
	if err != nil {
		panic(err)
	}
	defer resp_proj.Body.Close()

	body_send, err := ioutil.ReadAll(resp_send.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body_send), resp_proj.StatusCode)

}
