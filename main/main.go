package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"leapsy.com/packages/configurations"
	"leapsy.com/packages/logings"
)

const ()

// APK檔路徑基底
var apkBasicPath = ""

// APK檔名
var apkFileName = []string{}

// APK檔路徑
var apkFilePath = []string{}

// URL基底
var urlBasicPath = ""

// URL路徑
var downloadURL = []string{}

func main() {

	// 設定 APK路徑
	setAPKFilePath()

	// 打開APK下載Server
	openServer()

}

// 設定APK檔路徑
func setAPKFilePath() {

	// APK檔路徑基底
	apkBasicPath = configurations.GetConfigValueOrPanic("local", "apkBasicPath")

	// APK檔名
	apkFileName = []string{
		"",                       //[0] 保留
		"camera.apk",             //1
		"album.apk",              //2
		"webBrowser.apk",         //3
		"throne.apk",             //4
		"leapsyStore.apk",        //5
		"intelligenceSystem.apk", //6
		"environment.apk",        //7
		"palace.apk",             //8
		"FH1.apk",                //9
		"settings.apk",           //10
		"programs.apk",           //11
		"expert.apk",             //12
		"faceRecognize.apk",      //13
		"posture.apk",            //14
		"gesture.apk"}            //15

	// APK檔路徑
	apkFilePath = []string{
		"",                                     //[0] 保留
		apkBasicPath + "1/" + apkFileName[1],   //1 "apk/1/camera.apk"
		apkBasicPath + "2/" + apkFileName[2],   //2 "apk/2/album.apk"
		apkBasicPath + "3/" + apkFileName[3],   //3
		apkBasicPath + "4/" + apkFileName[4],   //4
		apkBasicPath + "5/" + apkFileName[5],   //5
		apkBasicPath + "6/" + apkFileName[6],   //6
		apkBasicPath + "7/" + apkFileName[7],   //7
		apkBasicPath + "8/" + apkFileName[8],   //8
		apkBasicPath + "9/" + apkFileName[9],   //9
		apkBasicPath + "10/" + apkFileName[10], //10
		apkBasicPath + "11/" + apkFileName[11], //11
		apkBasicPath + "12/" + apkFileName[12], //12
		apkBasicPath + "13/" + apkFileName[13], //13
		apkBasicPath + "14/" + apkFileName[14], //14
		apkBasicPath + "15/" + apkFileName[15]} //15

}

// 設定Server URL路徑
func setURLPath() {

	// URL基底
	urlBasicPath = configurations.GetConfigValueOrPanic("local", "urlBasicPath")

	// URL路徑
	downloadURL = []string{
		"",                 //[0] 保留不用
		urlBasicPath + "1", // "/appUpdate/download/1"
		urlBasicPath + "2",
		urlBasicPath + "3",
		urlBasicPath + "4",
		urlBasicPath + "5",
		urlBasicPath + "6",
		urlBasicPath + "7",
		urlBasicPath + "8",
		urlBasicPath + "9",
		urlBasicPath + "10",
		urlBasicPath + "11",
		urlBasicPath + "12",
		urlBasicPath + "13",
		urlBasicPath + "14",
		urlBasicPath + "15"}
}

// 打開APK下載Server
func openServer() {

	setURLPath()

	http.HandleFunc(downloadURL[1], downloadFile1) // /appUpdate/download/1
	http.HandleFunc(downloadURL[2], downloadFile2)
	http.HandleFunc(downloadURL[3], downloadFile3)
	http.HandleFunc(downloadURL[4], downloadFile4)
	http.HandleFunc(downloadURL[5], downloadFile5)
	http.HandleFunc(downloadURL[6], downloadFile6)
	http.HandleFunc(downloadURL[7], downloadFile7)
	http.HandleFunc(downloadURL[8], downloadFile8)
	http.HandleFunc(downloadURL[9], downloadFile9)
	http.HandleFunc(downloadURL[10], downloadFile10)
	http.HandleFunc(downloadURL[11], downloadFile11)
	http.HandleFunc(downloadURL[12], downloadFile12)
	http.HandleFunc(downloadURL[13], downloadFile13)
	http.HandleFunc(downloadURL[14], downloadFile14)
	http.HandleFunc(downloadURL[15], downloadFile15)

	// 啟動log紀錄
	go logings.StartLogging()

	// api port
	port := configurations.GetConfigValueOrPanic("local", "port")

	// log
	logings.SendLog(
		[]string{`啟動伺服器 Port:%s 提供APK下載服務`},
		[]interface{}{port},
		nil,
		logrus.InfoLevel,
	)

	// print
	fmt.Printf("啟動伺服器 Port:%s 提供APK下載服務\n", port)

	// 開啟Port
	http.ListenAndServe(":"+port, nil)

}

// 檔案一下載
func downloadFile1(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[1]
	apkPath := apkFilePath[1]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)

	// print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	// 服務設定
	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案二下載
func downloadFile2(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[2]
	apkPath := apkFilePath[2]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案三下載
func downloadFile3(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[3]
	apkPath := apkFilePath[3]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案四下載
func downloadFile4(w http.ResponseWriter, r *http.Request) {
	apkName := apkFileName[4]
	apkPath := apkFilePath[4]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案五下載
func downloadFile5(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[5]
	apkPath := apkFilePath[5]
	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案六下載
func downloadFile6(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[6]
	apkPath := apkFilePath[6]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案七下載
func downloadFile7(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[7]
	apkPath := apkFilePath[7]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案八下載
func downloadFile8(w http.ResponseWriter, r *http.Request) {
	apkName := apkFileName[8]
	apkPath := apkFilePath[8]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案9
func downloadFile9(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[9]
	apkPath := apkFilePath[9]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案10
func downloadFile10(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[10]
	apkPath := apkFilePath[10]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案11
func downloadFile11(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[11]
	apkPath := apkFilePath[11]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案12
func downloadFile12(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[12]
	apkPath := apkFilePath[12]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案13
func downloadFile13(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[13]
	apkPath := apkFilePath[13]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案14
func downloadFile14(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[14]
	apkPath := apkFilePath[14]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}

// 檔案15
func downloadFile15(w http.ResponseWriter, r *http.Request) {

	apkName := apkFileName[15]
	apkPath := apkFilePath[15]

	// log
	logings.SendLog(
		[]string{`收到HOST %s 要求路徑 %s 下載檔案 %s`},
		[]interface{}{r.Host, r.URL, apkName},
		nil,
		logrus.InfoLevel,
	)
	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", apkName)

	w.Header().Set("Content-Disposition", "attachment; filename="+apkName) // 下載檔名
	http.ServeFile(w, r, apkPath)                                          // 檔案路徑
}
