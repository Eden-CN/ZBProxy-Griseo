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
		fmt.Println("你的是最新版本，好耶！")
	} else {
		fmt.Println("你当前运行的版本不是最新版本，请前往 https://github.com/Eden-CN/ZBProxy-Griseo/releases 获取最新版本！")
	}
}
