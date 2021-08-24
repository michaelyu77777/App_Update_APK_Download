package servers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// getMacAddressCybLicenseBinAPIHandler - 取得授權檔
/**
 * @param  *APIServer apiServer API伺服器指標
 * @param  *gin.Context ginContextPointer  gin Context 指標
 */
func getAPPsAPIHandler(apiServer *APIServer, ginContextPointer *gin.Context) {

	eventTime := time.Now()

	isStatusBadRequestErrorChannel := make(chan bool, 1)

	isStatusForbiddenErrorChannel := make(chan bool, 1)

	isStatusNotFoundErrorChannel := make(chan bool, 1)

	httpStatusChannel := make(chan int, 1)

	var parameters Parameters

	bindError := ginContextPointer.ShouldBind(&parameters)

	bindURIError := ginContextPointer.ShouldBindUri(&parameters)

	isError := nil != bindError || nil != bindURIError
	isStatusBadRequestErrorChannel <- isError

	if !isError {

		isToWorkChannel := make(chan bool, 1)

		parametersDownladKeyword := parameters.DownloadKeyword

		// go func() {

		// 	isError = !isLowerCaseOrDigit(parametersDownladKeyword)

		// 	isToWorkChannel <- !isError

		// 	isStatusBadRequestErrorChannel <- isError

		// }()

		// go func() {

		// 	isError = !isAuthorized(ginContextPointer)

		// 	isToWorkChannel <- !isError
		// 	isStatusForbiddenErrorChannel <- isError

		// }()

		apkDirectoryName := parametersDownladKeyword // apk資料夾名稱
		apkFileName := ""                            //apk檔名

		// 看檔案是否存在
		go func() {

			// 檔案存在會取得APK檔名
			isError, apkFileName = isFileNotExistedAndGetApkFileName(apkDirectoryName)
			// isError = isFileNotExisted(parametersDownladKeyword)

			isToWorkChannel <- !isError

			isStatusNotFoundErrorChannel <- isError

		}()

		go func() {

			isToWork := true

			for counter := 1; counter <= 1; counter++ {
				isToWork = isToWork && <-isToWorkChannel
			}

			if isToWork {

				// 下載檔案
				attachApkFile(ginContextPointer, apkDirectoryName, apkFileName)
				// attachCybLicenseBin(ginContextPointer, parametersDownladKeyword)

				httpStatusChannel <- http.StatusOK
			}

		}()

	}

	go func() {

		for {

			if <-isStatusBadRequestErrorChannel {
				httpStatusChannel <- http.StatusBadRequest
			}

		}

	}()

	go func() {

		for {

			if <-isStatusForbiddenErrorChannel {
				httpStatusChannel <- http.StatusForbidden
			}

		}

	}()

	go func() {

		for {

			if <-isStatusNotFoundErrorChannel {
				httpStatusChannel <- http.StatusNotFound
			}

		}

	}()

	for {

		httpStatus := <-httpStatusChannel

		SendEvent(
			eventTime,
			ginContextPointer,
			parameters,
			nil,
			httpStatus,
			APIResponse{},
		)

		ginContextPointer.Status(httpStatus)

		return

	}

}

// getMacAddressCybLicenseBinAPIHandler - 取得授權檔
/**
 * @param  *APIServer apiServer API伺服器指標
 * @param  *gin.Context ginContextPointer  gin Context 指標
 */
// func getMacAddressCybLicenseBinAPIHandler(apiServer *APIServer, ginContextPointer *gin.Context) {

// 	eventTime := time.Now()

// 	isStatusBadRequestErrorChannel := make(chan bool, 1)

// 	isStatusForbiddenErrorChannel := make(chan bool, 1)

// 	isStatusNotFoundErrorChannel := make(chan bool, 1)

// 	httpStatusChannel := make(chan int, 1)

// 	var parameters Parameters

// 	bindError := ginContextPointer.ShouldBind(&parameters)

// 	bindURIError := ginContextPointer.ShouldBindUri(&parameters)

// 	isError := nil != bindError || nil != bindURIError
// 	isStatusBadRequestErrorChannel <- isError

// 	if !isError {

// 		isToWorkChannel := make(chan bool, 1)

// 		parametersMacAddress := parameters.MacAddress

// 		go func() {

// 			isError = !isLowerCaseOrDigit(parametersMacAddress)

// 			isToWorkChannel <- !isError

// 			isStatusBadRequestErrorChannel <- isError

// 		}()

// 		go func() {

// 			isError = !isAuthorized(ginContextPointer)

// 			isToWorkChannel <- !isError
// 			isStatusForbiddenErrorChannel <- isError

// 		}()

// 		go func() {

// 			isError = isFileNotExisted(parametersMacAddress)

// 			isToWorkChannel <- !isError

// 			isStatusNotFoundErrorChannel <- isError

// 		}()

// 		go func() {

// 			isToWork := true

// 			for counter := 1; counter <= 3; counter++ {
// 				isToWork = isToWork && <-isToWorkChannel
// 			}

// 			if isToWork {
// 				attachCybLicenseBin(ginContextPointer, parametersMacAddress)
// 				httpStatusChannel <- http.StatusOK
// 			}

// 		}()

// 	}

// 	go func() {

// 		for {

// 			if <-isStatusBadRequestErrorChannel {
// 				httpStatusChannel <- http.StatusBadRequest
// 			}

// 		}

// 	}()

// 	go func() {

// 		for {

// 			if <-isStatusForbiddenErrorChannel {
// 				httpStatusChannel <- http.StatusForbidden
// 			}

// 		}

// 	}()

// 	go func() {

// 		for {

// 			if <-isStatusNotFoundErrorChannel {
// 				httpStatusChannel <- http.StatusNotFound
// 			}

// 		}

// 	}()

// 	for {

// 		httpStatus := <-httpStatusChannel

// 		SendEvent(
// 			eventTime,
// 			ginContextPointer,
// 			parameters,
// 			nil,
// 			httpStatus,
// 			APIResponse{},
// 		)

// 		ginContextPointer.Status(httpStatus)

// 		return

// 	}

// }
