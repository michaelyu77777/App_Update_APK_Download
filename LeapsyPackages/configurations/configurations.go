package configurations

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"

	"leapsy.com/packages/logings"
)

const (
	configFileNameConstString string = `config.ini` // 設定檔名
)

var configMap map[string]map[string]string // 設定資料

// init - 初始函式
func init() {
	initializeConfigMapOrPanic() // 初始化設定資料或逐層結束程式
}

// initializeConfigMapOrPanic - 初始化設定資料或逐層結束程式
func initializeConfigMapOrPanic() {

	if nil == configMap { // 若沒有設定資料
		loadConfigFileOrPanic()          // 載入設定檔或逐層結束程式
		loadConfigMapFromOSArgsOrPanic() //  從執行檔參數載入設定Map或逐層結束程式
	}

}

// loadConfigFileOrPanic - 載入設定檔或逐層結束程式
func loadConfigFileOrPanic() {

	configFile, iniLoadError := ini.Load(configFileNameConstString) // 載入設定檔

	logings.SendLog(
		[]string{`載入設定檔 %s `},
		[]interface{}{configFileNameConstString},
		iniLoadError,
		logrus.PanicLevel,
	)

	configMap = make(map[string]map[string]string) // 位設定資料建立空間

	for _, section := range configFile.Sections() { // 針對每一個設定檔區塊

		configMap[section.Name()] = make(map[string]string) // 為設定檔區塊建立空間

		for _, key := range section.KeyStrings() { // 針對設定檔區塊下每一個關鍵字
			configMap[section.Name()][key] = section.Key(key).String() // 設定設定檔區塊下關鍵字對應的值
		}

	}

}

// loadConfigMapFromOSArgsOrPanic - 從執行檔參數載入設定Map或逐層結束程式
func loadConfigMapFromOSArgsOrPanic() {

	for index, value := range os.Args {

		if index > 0 {

			if results := regexp.MustCompile(`([^\[\]=]+)\[([^\[\]=]+)\]=([^\[\]=]+)`).FindStringSubmatch(value); len(results) == 0 {
				logings.SendLog(
					[]string{`第 %d 個參數 %s 解析 `},
					[]interface{}{index, value},
					fmt.Errorf(`參數格式應為section[key]=value`),
					logrus.PanicLevel,
				)
			} else {
				configMap[results[1]][results[2]] = results[3]
			}

		}

	}

}

// GetConfigValueOrPanic - 取得設定值否則結束程式
/**
 * @param  string sectionName  區塊名
 * @param  string key  關鍵字
 * @return string returnConfigValue 回傳設定資料區塊下關鍵字對應的值
 */
func GetConfigValueOrPanic(sectionName, key string) (returnConfigValue string) {

	if configValue, ok := configMap[sectionName][key]; !ok { // 若取得設定檔區塊下關鍵字對應的值失敗

		logings.SendLog(
			[]string{`設定檔 %s 設定`},
			[]interface{}{configFileNameConstString},
			fmt.Errorf(fmt.Sprintf(`[ %s ] %s 值錯誤`, sectionName, key)),
			logrus.PanicLevel,
		)

	} else {
		returnConfigValue = configValue
	}

	return // 回傳取得的設定檔區塊下關鍵字對應的值
}

// GetConfigPositiveIntValueOrPanic - 取得正整數設定值否則結束程式
/**
 * @param  string sectionName  區塊名
 * @param  string key  關鍵字
 * @return string returnValue 回傳設定資料區塊下關鍵字對應的正整數值
 */
func GetConfigPositiveIntValueOrPanic(sectionName string, key string) (returnValue int) {

	configValue := GetConfigValueOrPanic(sectionName, key)

	if value, valueError := strconv.Atoi(configValue); valueError != nil || value <= 0 { // 若取得設定檔區塊下關鍵字對應的整數值非正整數

		logings.SendLog(
			[]string{`設定檔 %s 設定`},
			[]interface{}{configFileNameConstString},
			fmt.Errorf(fmt.Sprintf(`[ %s ] %s 值 %s 應為正整數`, sectionName, key, configValue)),
			logrus.PanicLevel,
		)

	} else {
		returnValue = value
	}

	return // 回傳取得的設定檔區塊下關鍵字對應的正整數值
}
