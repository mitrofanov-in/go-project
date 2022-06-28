package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	login := ""
	pass := ""

	/*
		data, err := os.Open("1.pdf")
		if err != nil {
			log.Fatal(err)
		}
		//data.Close()
	*/

	server := "http://nextcloud.storage.wd.xco.devel.ifx"
	path := "/API/v1/CSV/"
	url := server + "/remote.php/dav/files/" + login + path
	fmt.Println(url)
	method := "PROPFIND"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(login, pass)
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	req.Header.Set("OCS-APIRequest", "true")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//defer req_gen.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("INFO ABOUT SEND FILES", string(body), resp.StatusCode)
}
