package servers

import (
	"leapsy.com/packages/logings"
)

// Server - 伺服器
type Server struct {
}

// StartServer - 啟動伺服器
func StartServer() {

	go logings.StartLogging()
	go StartUpdatingEvents()

	var (
		apiServer APIServer // API伺服器
	)

	defer func() {
		apiServer.stop() // 結束API伺服器
		StopServer()     // 結束伺服器
	}()

	logings.SendLog(
		[]string{`啟動 伺服器 `},
		[]interface{}{},
		nil,
		0,
	)

	apiServer.start() // 啟動API伺服器

}

// StopServer - 結束伺服器
func StopServer() {

	logings.SendLog(
		[]string{`結束 伺服器 `},
		[]interface{}{},
		nil,
		0,
	)

}
