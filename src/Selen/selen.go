package main

import (
	"fmt"

	"github.com/tebeka/selenium"
)

func main() {

	// "github.com/tebeka/selenium"
	caps := selenium.Capabilities{"browserName": "chrome", "browserVersion": "100.0"}
	driver, err := selenium.NewRemote(caps, "http://selenoid.edge.wd.xco.devel.ifx/wd/hub")
	if err != nil {
		fmt.Println(err)
	}
	driver.Get("https://yandex.ru")
	defer driver.Quit()

}
