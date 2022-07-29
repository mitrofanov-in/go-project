package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	chromeCaps := chrome.Capabilities{
		Args: []string{
			// "--headless",
			"--no-sandbox",
			"--start-maximized",
			"--window-size=1920,1080",
			"--disable-crash-reporter",
			"--hide-scrollbars",
			"--disable-gpu",
			"--test-type=browser",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
			"--lang=ru",
		},
	}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, "http://51.250.102.170:4444/wd/hub")
	if err != nil {
		panic(err)
	}

	Url := "http://www.formy-i-blanki.ru/dogovor-arendy-kvartiry-skachat-obrazec"
	Usel := "http://51.250.102.170:4444/download"

	getValSes := wd.SessionId()

	wd.Refresh()

	fmt.Printf(wd.SessionId())
	wd.SwitchFrame("news_globalize_translations_attributes_ru_content_ifr")
	wd.Get(Url)

	form, err := wd.FindElement(selenium.ByXPATH, "//a [text()=\"Договор аренды квартиры юридическим лицом у физ. лица\"]")
	if err := form.Click(); err != nil {
		panic(err)
	}

	wd.SetImplicitWaitTimeout(4 * time.Second)
	wd.Refresh()

	form2, err := wd.FindElement(selenium.ByXPATH, "//a [text()=\"Договор аренды квартиры\"]")
	if err := form2.Click(); err != nil {
		panic(err)
	}

	wd.SetImplicitWaitTimeout(4 * time.Second)
	wd.Refresh()

	form3, err := wd.FindElement(selenium.ByXPATH, "//a [text()=\"Договор аренды квартиры с правом выкупа\"]")
	if err := form3.Click(); err != nil {
		panic(err)
	}
	wd.SetImplicitWaitTimeout(4 * time.Second)
	wd.Refresh()

	dwn_url := Usel + "/" + getValSes

	client := &http.Client{}

	req, err := http.NewRequest("GET", dwn_url, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// записываем вывод

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	//strng := rData.String()
	strng := string(bodyText)
	fmt.Println(strng)

	// режем тэги

	stripped := bluemonday.StripTagsPolicy()
	html := stripped.Sanitize(strng)
	fmt.Println(html)

	// формируем список

	file, err := os.Create("list.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(html)

	f_str := []string{}

	// заполняем массив

	rfile, err := os.Open("list.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
	}
	defer rfile.Close()

	scanner := bufio.NewScanner(rfile)
	for scanner.Scan() {
		f_str = append(f_str, scanner.Text())
	}
	fmt.Println(f_str)
	fmt.Println(len(f_str))

	// Убираем все пустые строки

	f_str_cl := []string{}

	for i := 0; i < len(f_str); i++ {
		if f_str[i] != "" {
			f_str_cl = append(f_str_cl, f_str[i])
		}
	}
	fmt.Println(f_str_cl)
	fmt.Println(len(f_str_cl))

	for i := 0; i < len(f_str_cl); i++ {

		dwn_url_f := Usel + "/" + getValSes + "/" + f_str_cl[i]

		errs := downloadFile(f_str_cl[i], dwn_url_f)
		if errs != nil {
			panic(errs)
		}

	}

}
