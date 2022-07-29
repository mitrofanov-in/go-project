package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type JiraUpdater interface {
	checkStatus()
	updateStatus()
}

type Transit struct {
	Id Id `json:"transition"`
}

type Id struct {
	Id int `json:"id"`
}

var idProjGitlab string = ""
var tasks []string
var username string = ""
var password string = ""

//var urlJira string = "https://jira.interfax.ru/rest/api/2/issue/" + tasks[0] + "/transitions"
//var urlGitlab string = "https://gitlab.interfax.ru/api/v4/projects/" + idProjGitlab + "/repository/tags"

func checkStatus(url string) string {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(bodyText))
	return string(bodyText)
}

func updateStatus(url string, jstr []byte) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jstr))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bodyText))
}

func getGithubRelease(url string, id string) string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("PRIVATE-TOKEN", "")
	//req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(bodyText))
	return string(bodyText)
}

func ConvertToMap(m []map[string]interface{}, body string) string {
	json.Unmarshal([]byte(body), &m)
	f := m[0]["release"]
	moveToString := fmt.Sprintf("%v", f)
	return moveToString
}

func main() {
	jsonData := Transit{
		Id: Id{
			Id: 31,
		},
	}
	project := os.Args
	idProjGitlab := project[1]
	if len(project) < 1 {
		fmt.Println("Введено аргументов больше, чем ожидается")
		os.Exit(1)
	}
	var m []map[string]interface{}

	jsonDataSort, _ := json.Marshal(jsonData)
	jStr := []byte(jsonDataSort)

	fmt.Println(string(jStr))

	var urlGitlab string = "https://gitlab.interfax.ru/api/v4/projects/" + idProjGitlab + "/repository/tags"
	var body string = getGithubRelease(urlGitlab, idProjGitlab)

	parse := ConvertToMap(m, body)

	strParse := strings.SplitAfter(parse, "|") //обрезаем разделитель |
	lnghtStr := len(strParse)                  //определяем длину

	delSymbN := strings.ReplaceAll(parse, "\n", "")
	regX := regexp.MustCompile("XCO-[0-9]+")
	tasksGitlab := regX.FindAllString(delSymbN, lnghtStr)
	//fmt.Println(tasks)

	for _, uri := range tasksGitlab {
		var urlJira string = "https://jira.interfax.ru/rest/api/2/issue/" + uri + "/?fields=status"
		val := checkStatus(urlJira)

		var j map[string]interface{}
		json.Unmarshal([]byte(val), &j)
		firstLev := j["fields"]
		two := firstLev.(map[string]interface{})
		twoLev := two["status"]
		three := twoLev.(map[string]interface{})
		threeLev := three["name"]
		//fmt.Println(firstLev, "вложенность внутри", twoLev, threeLev)

		moveToString := fmt.Sprintf("%v", threeLev)
		if moveToString != "В работе" {
			tasks = append(tasks, uri)
		}

	}
	fmt.Println(tasks)

	for _, uriJira := range tasks {
		var urlJira string = "https://jira.interfax.ru/rest/api/2/issue/" + uriJira + "/transitions"
		updateStatus(urlJira, jStr)
	}

}
