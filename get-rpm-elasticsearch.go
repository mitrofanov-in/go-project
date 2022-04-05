package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var jStr = []byte(`{"track_total_hits": true, "size": 0, "aggs": {"2": {"date_histogram": {"field": "@timestamp","fixed_interval": "1s","time_zone": "Europe/Samara","min_doc_count": 1}}},"fields": [{"field": "@timestamp","format": "date_time"}],"script_fields": {},"stored_fields": ["*"],"runtime_mappings": {},"_source": {"excludes": []},"query": {"bool": {"must": [],"filter": [{"range": {"@timestamp": {"format": "strict_date_optional_time","gte": "now-1m","lte": "now"}}}],"should": [],"must_not": []}}}`)

var bodyStat int = 0
var url string

func basicAuth(nameIndex string, url string) float64 {

	//var username string = os.Getenv("USER_ES")
	//var passwd string = os.Getenv("PASS_ES")
	//var url string = os.Getenv("URL_ES")
	var username string = "xco"
	var passwd string
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://"+url+"/"+nameIndex+"*"+"/_search", bytes.NewBuffer(jStr))
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	//s := string(bodyText)

	var m map[string]interface{}
	json.Unmarshal([]byte(bodyText), &m)

	f := m["hits"]
	j := f.(map[string]interface{})
	x := j["total"]
	val := x.(map[string]interface{})
	var str1 string = ""
	str1 = fmt.Sprintf("%v", val["value"])
	l_int, _ := strconv.ParseFloat(str1, 64)
	bodyStat = resp.StatusCode

	return l_int

}

func main() {

	//e_string := ""
	arguments := os.Args
	nameIndex := arguments[1]
	url := arguments[2]

	S := basicAuth(nameIndex, url)

	fmt.Println(S)
}
