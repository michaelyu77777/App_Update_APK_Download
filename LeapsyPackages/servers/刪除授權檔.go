package servers

// import (
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// // deleteMacAddressCybLicenseBinAPIHandler - 刪除授權檔
// /**
//  * @param  *APIServer apiServer API伺服器指標
//  * @param  *gin.Context ginContextPointer  gin Context 指標
//  */
// func deleteMacAddressCybLicenseBinAPIHandler(apiServer *APIServer, ginContextPointer *gin.Context) {

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
// 				deleteCybLicenseBin(parametersMacAddress)
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
