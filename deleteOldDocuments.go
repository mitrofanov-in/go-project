package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var jStr = []byte(`{
	"query": {
	  "bool": {
		"filter": [
		  {
			"range": {
			  "@timestamp": {
				"lt": "now-30d"
			  }
			}
		  }
		]
	  }
	}
  }`)

var bodyStat int = 0
var url string

func basicAuth(nameIndex string) string {

	var username string = os.Getenv("USER_ES")
	var passwd string = os.Getenv("PASS_ES")
	var url string = os.Getenv("URL_ES")
	//var username string = "xco"
	//var passwd string
	client := &http.Client{}
	req, err := http.NewRequest("POST", url+"/"+nameIndex+"/_delete_by_query?pretty", bytes.NewBuffer(jStr))
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s

}

func main() {

	//e_string := ""
	arguments := os.Args
	nameIndex := arguments[1]

	S := basicAuth(nameIndex)

	fmt.Println(string(S))
}
