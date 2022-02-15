package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type IndexELASTIC struct {
	Source Source `json:"source"`
	Dest   Dest   `json:"dest"`
}
type Source struct {
	Index string `json:"index"`
}
type Dest struct {
	Index string `json:"index"`
}

func basicAuth(nameIndex string, DataIndex string) string {

	ElasticReidex := IndexELASTIC{
		Source: Source{
			Index: nameIndex + "-" + DataIndex + "*",
		},
		Dest: Dest{
			Index: nameIndex,
		},
	}

	jsonDataSort, _ := json.Marshal(ElasticReidex)

	jStr := []byte(jsonDataSort)
	var username string = os.Getenv("USER_ES")
	var passwd string = os.Getenv("PASS_ES")
	var url string = os.Getenv("URL_ES")
	client := &http.Client{}
	req, err := http.NewRequest("POST", url+"/_reindex", bytes.NewBuffer(jStr))
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

func deleteIndex(nameIndex string, DataIndex string) string {

	/*
	   ElasticReidex := IndexELASTIC{
	           Source: Source{
	                   Index: nameIndex + "-" + DataIndex + "*",
	           },
	           Dest: Dest{
	                   Index: nameIndex,
	           },
	   }

	   jsonDataSort, _ := json.Marshal(ElasticReidex)

	   jStr := []byte(jsonDataSort)
	*/
	var username string = os.Getenv("USER_ES")
	var passwd string = os.Getenv("PASS_ES")
	var url string = os.Getenv("URL_ES")
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url+"/"+nameIndex+"-"+DataIndex+"*", nil)
	//req.Header.Add("Content-Type", "application/json")
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
	DataIndex := arguments[2]

	S := basicAuth(nameIndex, DataIndex)
	var x string = S
	Y := deleteIndex(nameIndex, DataIndex)
	var y string = Y

	fmt.Println(x)
	fmt.Println(y)
}
