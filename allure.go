package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var project string
var url string = os.Getenv("URL_ALLURE")
var login string = os.Getenv("USER_ALLURE")
var pass string = os.Getenv("PASS_ALLURE")
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

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetAuthParmas(payload_lgn []byte) []*http.Cookie {
	url_lgn := url + "/login"
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

	url_send := url + "/send-results?project_id=" + project
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

func GenerateReport(project string) {

	execution_name := "execution-from-my-script"
	execution_from := "http://google.com"
	execution_type := "teamcity"

	url_gen := url + "/generate-report?project_id=" + project + "&execution_name=" + execution_name + "&execution_from=" + execution_from + "&execution_type=" + execution_type
	method_gen := "GET"

	fmt.Println(url_gen)
	client := &http.Client{}
	req_gen, err := http.NewRequest(method_gen, url_gen, nil)
	if err != nil {
		panic(err)
	}

	for i := range cookie {
		req_gen.AddCookie(cookie[i])

	}
	req_gen.Header.Add("X-CSRF-TOKEN", tmpCSRF.Value)

	resp_gen, err := client.Do(req_gen)
	if err != nil {
		panic(err)
	}
	//defer req_gen.Body.Close()

	/*
		body_gen, err := ioutil.ReadAll(resp_gen.Body)
		if err != nil {
			panic(err)
		}
	*/
	fmt.Println("INFO ABOUT SEND FILES", resp_gen.StatusCode)
}

func main() {

	arguments := os.Args
	project = arguments[1]

	lgn := Auth{
		Username: login,
		Password: pass,
	}
	payloadDataSort, _ := json.Marshal(lgn)
	payload_lgn := []byte(payloadDataSort)
	fmt.Println(string(payload_lgn))

	//Авторизация

	GetAuthParmas(payload_lgn)

	// Сохраняем куки
	tmpCSRF = cookie[1]

	var filesToSend []string
	var base64ToSend []string
	var path string
	path = "./allure_results/"

	// Формируем входной массив json
	var bs64 string
	tmp, _ := ioutil.ReadDir(path)
	for _, t := range tmp {
		if !t.IsDir() {
			filesToSend = append(filesToSend, t.Name())
			content, _ := ioutil.ReadFile(path + t.Name())
			bs64 = base64.StdEncoding.EncodeToString([]byte(content))
			base64ToSend = append(base64ToSend, string(bs64))
		}
	}

	xLen := len(filesToSend)

	for x := 0; x < xLen; x++ {
		content = append(content, Results{FileName: filesToSend[x], ContentBase64: base64ToSend[x]})
	}

	Allure := Allure{
		Results: content,
	}

	// Передаем полученный json
	jsonDataSort, _ := json.Marshal(Allure)

	jStr := []byte(jsonDataSort)

	fmt.Println(string(jStr))

	// Вызываем функцию передачи
	SendResult(jStr)

	// Генерируем результат
	GenerateReport(project)

}
