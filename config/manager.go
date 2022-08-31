package config

import (
	"encoding/json"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/LittleGriseo/GriseoProxy/common/set"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	Config     configMain
	Lists      map[string]*set.StringSet
	reloadLock sync.Mutex
)

func LoadConfig() {
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("配置文件不存在。 生成一个新的配置文件中...")
			generateDefaultConfig()
			goto success
		} else {
			log.Panic(color.HiRedString("加载配置时出现意外错误: %s", err.Error()))
		}
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		log.Panic(color.HiRedString("配置格式错误: %s", err.Error()))
	}

success:
	LoadLists(false)
	log.Println(color.HiYellowString("已成功载入配置文件!"))
}

func generateDefaultConfig() {
	file, err := os.Create("config.json")
	if err != nil {
		log.Panic("创建配置文件失败：", err.Error())
	}
	Config = configMain{
		Services: []*ConfigProxyService{
			{
				Name:          "HypixelDefault",
				TargetAddress: "mc.hypixel.net",
				TargetPort:    25565,
				Listen:        25565,
				Flow:          "auto",
				Minecraft: minecraft{
					EnableHostnameRewrite: true,
					IgnoreFMLSuffix:       true,
					OnlineCount: onlineCount{
						Max:            114514,
						Online:         -1,
						EnableMaxLimit: false,
					},
					MotdFavicon:     "{DEFAULT_MOTD}",
					MotdDescription: "§d{NAME}§e service is working on §a§o{INFO}§r\n§c§lProxy for §6§n{HOST}:{PORT}§r",
				},
			},
		},
		Lists: map[string][]string{
			//"test": {"foo", "bar"},
		},
	}
	newConfig, _ :=
		json.MarshalIndent(Config, "", "    ")
	_, err = file.WriteString(strings.ReplaceAll(string(newConfig), "\n", "\r\n"))
	file.Close()
	if err != nil {
		log.Panic("保存配置文件失败:", err.Error())
	}
}

func LoadLists(isReload bool) bool {
	reloadLock.Lock()
	if isReload {
		configFile, err := os.ReadFile("config.json")
		if err != nil {
			if os.IsNotExist(err) {
				log.Println(color.HiRedString("重新加载失败：配置文件不存在。"))
			} else {
				log.Println(color.HiRedString("重新加载配置时出现意外错误: %s", err.Error()))
			}
			reloadLock.Unlock()
			return false
		}

		err = json.Unmarshal(configFile, &Config)
		if err != nil {
			log.Println(color.HiRedString("无法重新加载：配置格式错误: %s", err.Error()))
			reloadLock.Unlock()
			return false
		}
	}
	//log.Println("Lists:", Config.Lists)
	if l := len(Config.Lists); l == 0 { // if nothing in Lists
		Lists = map[string]*set.StringSet{} // empty map
	} else {
		Lists = make(map[string]*set.StringSet, l) // map size init
		for k, v := range Config.Lists {
			//log.Println("List: Loading", k, "value:", v)
			set := set.NewStringSetFromSlice(v)
			Lists[k] = &set
		}
	}
	Config.Lists = nil // free memory
	reloadLock.Unlock()
	runtime.GC()
	return true
}

func MonitorConfig(watcher *fsnotify.Watcher) error {
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					continue
				}
				if event.Op&fsnotify.Write == fsnotify.Write { // config reload
					log.Println(color.HiMagentaString("配置重新加载：检测到文件更改。 正在重新加载..."))
					if LoadLists(true) { // reload success
						log.Println(color.HiMagentaString("配置重新加载：成功重新加载列表。"))
					} else {
						log.Println(color.HiMagentaString("配置重新加载：无法重新加载列表。"))
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					continue
				}
				log.Println(color.HiRedString("配置重新加载错误 : ", err))
			}
		}
	}()

	return watcher.Add("config.json")
}
