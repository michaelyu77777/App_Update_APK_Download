package logings

import (
	"strings"
	"time"

	filename "github.com/keepeye/logrus-filename"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// LogData - 紀錄資料
type LogData struct {
	formatSlices []string
	args         []interface{}
	error        error
	level        logrus.Level
}

const (
	fileExtensionConstString     = `.%Y%m%d%H`        // log檔副檔名
	rotateTimeConstDuration      = time.Hour          // 切割期間
	errorsMaxAgeConstDuration    = time.Hour * 24 * 7 // 錯誤保留最大期間
	nonErrorsMaxAgeConstDuration = time.Hour * 24 * 7 // 非錯誤保留最大期間
	channelSize                  = 32768
)

var (
	logger         = logrus.New() // 記錄器
	logDataChannel = make(chan LogData, channelSize)
)

// init - 初始函式
func init() {
	setLogFiles() // 設定記錄檔
}

// setLogFiles - 設定記錄檔
func setLogFiles() {
	logger.SetLevel(logrus.TraceLevel) // 設定最高層級
	logger.AddHook(filename.NewHook()) // 加上行數鉤
	logger.AddHook(getLogFilesHook())  // 加上記錄檔鉤
}

// getLogFilesHook - 回傳記錄檔鉤
/**
 * @return  logrus.Hook  記錄檔鉤
 */
func getLogFilesHook() logrus.Hook {

	formatStringItemSlices := []string{`新增輪流`, `記錄檔`} // 記錄器格式片段
	defaultArgs := []interface{}{}                    // 記錄器預設參數

	errorRotateLogsName := `logs/errors/error` // 錯誤記錄檔名

	// 新增輪流記錄檔
	errorRotateLogs, rotatelogsNewError := rotatelogs.New(
		errorRotateLogsName+fileExtensionConstString,
		rotatelogs.WithLinkName(errorRotateLogsName),
		rotatelogs.WithRotationTime(rotateTimeConstDuration),
		rotatelogs.WithMaxAge(errorsMaxAgeConstDuration),
	)

	// 取得記錄器格式字串與參數
	formatString, args := GetLogFuncFormatAndArguments(
		[]string{strings.Join(formatStringItemSlices, `錯誤`)},
		defaultArgs,
		rotatelogsNewError,
	)

	if nil != rotatelogsNewError { // 若新增輪流記錄檔錯誤，則記錄錯誤並逐層結束程式
		logger.Panicf(formatString, args...) // 記錄錯誤並逐層結束程式
	} else { // 若新增輪流錯誤記錄檔成功，則記錄資訊
		go logger.Infof(formatString, args...) // 記錄資訊
	}

	warnRotateLogsName := `logs/warns/warn` // 警告記錄檔名

	// 新增輪流記錄檔
	warnRotateLogs, rotatelogsNewError := rotatelogs.New(
		warnRotateLogsName+fileExtensionConstString,
		rotatelogs.WithLinkName(warnRotateLogsName),
		rotatelogs.WithRotationTime(rotateTimeConstDuration),
		rotatelogs.WithMaxAge(nonErrorsMaxAgeConstDuration),
	)

	// 取得記錄器格式字串與參數
	formatString, args = GetLogFuncFormatAndArguments(
		[]string{strings.Join(formatStringItemSlices, `警告`)},
		defaultArgs,
		rotatelogsNewError,
	)

	if nil != rotatelogsNewError { // 若新增輪流記錄檔錯誤，則記錄錯誤並逐層結束程式
		logger.Panicf(formatString, args...) // 記錄錯誤並逐層結束程式
	} else { // 若新增輪流警告記錄檔成功，則記錄資訊
		go logger.Infof(formatString, args...) // 記錄資訊
	}

	infoRotateLogsName := `logs/infos/info` // 資訊記錄檔名

	// 新增輪流記錄檔
	infoRotateLogs, rotatelogsNewError := rotatelogs.New(
		infoRotateLogsName+fileExtensionConstString,
		rotatelogs.WithLinkName(infoRotateLogsName),
		rotatelogs.WithRotationTime(rotateTimeConstDuration),
		rotatelogs.WithMaxAge(nonErrorsMaxAgeConstDuration),
	)

	// 取得記錄器格式字串與參數
	formatString, args = GetLogFuncFormatAndArguments(
		[]string{strings.Join(formatStringItemSlices, `資訊`)},
		defaultArgs,
		rotatelogsNewError,
	)

	if nil != rotatelogsNewError { // 若新增輪流記錄檔錯誤，則記錄錯誤並逐層結束程式
		logger.Panicf(formatString, args...) // 記錄錯誤並逐層結束程式
	} else { // 若新增輪流記錄檔成功，則記錄資訊
		go logger.Infof(formatString, args...) // 記錄資訊
	}

	// 建立記錄檔鉤
	logFilesHook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.PanicLevel: errorRotateLogs,
			logrus.FatalLevel: errorRotateLogs,
			logrus.ErrorLevel: errorRotateLogs,
			logrus.WarnLevel:  warnRotateLogs,
			logrus.InfoLevel:  infoRotateLogs,
		},
		&logrus.TextFormatter{},
	)

	return logFilesHook // 回傳記錄檔鉤
}

// GetLogger - 取得記錄器
/**
 * @return  *logrus.Logger  記錄器
 */
func GetLogger() *logrus.Logger {
	return logger // 回傳記錄器
}

// GetLogFuncFormatAndArguments - 取得記錄器格式與參數
/**
 * @param  []string formatStringSlices 格式字串片段
 * @param  []interface{} args 參數
 * @param  error err 錯誤
 * @return string 格式字串
 * @return []interface{} 參數
 */
func GetLogFuncFormatAndArguments(formatStringSlices []string, args []interface{}, err error) (string, []interface{}) {

	if nil != err { // 若有錯誤
		return strings.Join(append(formatStringSlices, `: %+v`), `失敗`), append(args, err) // 回傳錯誤訊息格式字串與參數
	}

	return strings.Join(append(formatStringSlices, ``), `成功`), args // 回傳成功訊息格式字串與參數

}

// SendLog - 傳送紀錄
/**
 * @param  string inputFormatSlices 輸入格式片段
 * @param []interface{} inputArgs 輸入參數
 * @param error inputError 輸入錯誤
 * @param logrus.Level inputLevel 輸入錯誤層級
 */
func SendLog(inputFormatSlices []string, inputArgs []interface{}, inputError error, inputLevel logrus.Level) {
	logDataChannel <- LogData{
		formatSlices: inputFormatSlices,
		args:         inputArgs,
		error:        inputError,
		level:        inputLevel,
	}
}

// StartLogging - 開始記錄
func StartLogging() {

	// 取得記錄器格式和參數
	formatString, args := GetLogFuncFormatAndArguments(
		[]string{`啟動 %s 更新 `},
		[]interface{}{`logings`},
		nil,
	)

	go logger.Infof(formatString, args...) // 記錄資訊

	for {

		logData := <-logDataChannel

		go func() {

			logDataError := logData.error

			// 取得記錄器格式和參數
			formatString, args := GetLogFuncFormatAndArguments(
				logData.formatSlices,
				logData.args,
				logDataError,
			)

			if nil != logDataError {

				logDataLevel := logData.level

				switch logDataLevel {

				case logrus.PanicLevel:
					logger.Panicf(formatString, args...) // 記錄資訊

				case logrus.ErrorLevel:
					logger.Errorf(formatString, args...) // 記錄資訊

				case logrus.WarnLevel:
					logger.Warnf(formatString, args...) // 記錄資訊

				case logrus.InfoLevel:
					logger.Infof(formatString, args...) // 記錄資訊

				default:
					logger.Errorf(formatString, args...) // 記錄資訊

				}

			} else {
				logger.Infof(formatString, args...) // 記錄資訊
			}

		}()

	}

}
