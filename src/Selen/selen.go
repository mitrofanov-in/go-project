package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {

	caps := selenium.Capabilities{"browserName": "chrome", "browserVersion": "103.0"}

	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--headless", // <<<
			"--no-sandbox",
			//"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
			"--proxy-server=https://193.31.101.137:9749",
		},
	}
	caps.AddChrome(chromeCaps)

	driver, err := selenium.NewRemote(caps, "http://51.250.102.170:4444/wd/hub")
	if err != nil {
		fmt.Println(err)
	}
	driver.Refresh()
	fmt.Println(driver.SessionId())
	driver.Refresh()
	time.Sleep(5 * time.Second)
	x := driver.Get("https://www.boostra.ru")
	fmt.Println(x)
	driver.Refresh()
	form, err := driver.FindElement(selenium.ByXPATH, "//button[. = 'Найти']")
	if err := form.Click(); err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	defer driver.Quit()

}
