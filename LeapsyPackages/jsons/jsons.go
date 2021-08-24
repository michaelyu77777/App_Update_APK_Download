package jsons

import (
	"encoding/json"

	"leapsy.com/packages/logings"

	"github.com/sirupsen/logrus"
)

// JSONBytes - 取得JSON位元組陣列
/**
 * @param interface{} inputObject 輸入物件
 * @return []byte returnJSONBytes 取得JSON位元組陣列
 */
func JSONBytes(inputObject interface{}) (returnJSONBytes []byte) {

	returnJSONBytes, jsonMarshalError := json.Marshal(inputObject) // 轉成JSON

	logings.SendLog(
		[]string{`%+v 轉成JSON位元組陣列 %s `},
		[]interface{}{inputObject, string(returnJSONBytes)},
		jsonMarshalError,
		logrus.InfoLevel,
	)

	if nil != jsonMarshalError { // 若轉成JSON錯誤
		return // 回傳
	}

	return // 回傳
}

// JSONString - 取得JSON字串
/**
 * @param interface{} inputObject 輸入物件
 * @return string returnJSONString 取得JSON字串
 */
func JSONString(inputObject interface{}) (returnJSONString string) {
	returnJSONString = string(JSONBytes(inputObject)) // 回傳JSON字串
	return                                            // 回傳
}
