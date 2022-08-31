package version

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func printErr(err error) {
	log.Printf("检查更新时发生错误, caution: %v.", err.Error())
	log.Println(`你可以尝试在这里查询最新版本 https://github.com/Eden-CN/ZBProxy-Griseo/releases`)
}

func CheckUpdate() {
	resp, err := http.Get(`https://cdn.jsdelivr.net/gh/Eden-CN/ZBProxy-Griseo@master/version/version.go`)
	if err != nil {
		printErr(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		printErr(err)
		return
	}
	if strings.Contains(string(body), Version) {
		fmt.Println("Your ZBProxy is up-to-date. Have fun!")
	} else {
		fmt.Println("Your ZBProxy is out of date! Check for the latest version at https://github.com/layou233/ZBProxy/releases")
	}
}
