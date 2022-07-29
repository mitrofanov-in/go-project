package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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

	caps := selenium.Capabilities{"browserName": "chrome", "browserVersion": "103.0"}

	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--headless", // <<<
			"--no-sandbox",
			//"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
			//"--proxy-server=https://193.31.101.137:9749",
		},
	}
	caps.AddChrome(chromeCaps)

	url := "http://51.250.102.170:4444/"
	urlApi := "http://51.250.102.170:4444/wd/hub"
	driver, err := selenium.NewRemote(caps, urlApi)
	if err != nil {
		fmt.Println(err)
	}
	driver.Refresh()

	sesId := driver.SessionId()

	dwn_url := url + "download/" + sesId

	fmt.Println(sesId)
	time.Sleep(5 * time.Second)

	x := driver.Get("https://www.boostra.ru/info#info")

	fmt.Println(x)

	form, err := driver.FindElement(selenium.ByXPATH, "//body/div[1]/section[1]/div[1]/div[1]/div[3]/ul/li[1]/a[text()='Устав']")
	if err = form.Click(); err != nil {
		panic(err)
	}
	driver.SetImplicitWaitTimeout(4 * time.Second)
	time.Sleep(5 * time.Second)
	driver.Refresh()

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
	fmt.Println(string(bodyText))

	defer driver.Quit()

}
