package times

import (
	"sync"

	"leapsy.com/databases"
	"leapsy.com/packages/logings"
	"leapsy.com/records"

	"github.com/robfig/cron"
)

var (
	mongoDB databases.MongoDB
	startOfficeHourOfDayRWMutex,
	startOfficeMinuteOfDayRWMutex,
	checkInBufferStartHourRWMutex,
	checkInBufferStartMinuteRWMutex,
	checkInBufferEndHourRWMutex,
	checkInBufferEndMinuteRWMutex,
	checkOutStartHourRWMutex,
	checkOutStartMinuteRWMutex sync.RWMutex
	startOfficeHourOfDay     = int(mongoDB.FindSettingsByID(`startOfficeHourOfDay`)[0].Value.(int32))
	startOfficeMinuteOfDay   = int(mongoDB.FindSettingsByID(`startOfficeMinuteOfDay`)[0].Value.(int32))
	checkInBufferStartHour   = int(mongoDB.FindSettingsByID(`checkInBufferStartHour`)[0].Value.(int32))
	checkInBufferStartMinute = int(mongoDB.FindSettingsByID(`checkInBufferStartMinute`)[0].Value.(int32))
	checkInBufferEndHour     = int(mongoDB.FindSettingsByID(`checkInBufferEndHour`)[0].Value.(int32))
	checkInBufferEndMinute   = int(mongoDB.FindSettingsByID(`checkInBufferEndMinute`)[0].Value.(int32))
	checkOutStartHour        = int(mongoDB.FindSettingsByID(`checkOutStartHour`)[0].Value.(int32))
	checkOutStartMinute      = int(mongoDB.FindSettingsByID(`checkOutStartMinute`)[0].Value.(int32))
)

// StartUpdating - 開始更新
func StartUpdating() {

	logings.SendLog(
		[]string{`啟動 %s 更新 `},
		[]interface{}{`times`},
		nil,
		0,
	)

	startOfficeHourOfDayChannel := make(chan int, 1)
	startOfficeMinuteOfDayChannel := make(chan int, 1)
	checkInBufferStartHourChannel := make(chan int, 1)
	checkInBufferStartMinuteChannel := make(chan int, 1)
	checkInBufferEndHourChannel := make(chan int, 1)
	checkInBufferEndMinuteChannel := make(chan int, 1)
	checkOutStartHourChannel := make(chan int, 1)
	checkOutStartMinuteChannel := make(chan int, 1)

	newCron := cron.New() // 新建一個定時任務物件

	newCron.AddFunc(
		`@every 30s`,
		func() {

			settings := records.SettingArray(mongoDB.FindSettings())

			go func() {
				startOfficeHourOfDayChannel <- int(settings.GetValueOfID(`startOfficeHourOfDay`).(int32))
			}()

			go func() {
				startOfficeMinuteOfDayChannel <- int(settings.GetValueOfID(`startOfficeMinuteOfDay`).(int32))
			}()

			go func() {
				checkInBufferStartHourChannel <- int(settings.GetValueOfID(`checkInBufferStartHour`).(int32))
			}()

			go func() {
				checkInBufferStartMinuteChannel <- int(settings.GetValueOfID(`checkInBufferStartMinute`).(int32))
			}()

			go func() {
				checkInBufferEndHourChannel <- int(settings.GetValueOfID(`checkInBufferEndHour`).(int32))
			}()

			go func() {
				checkInBufferEndMinuteChannel <- int(settings.GetValueOfID(`checkInBufferEndMinute`).(int32))
			}()

			go func() {
				checkOutStartHourChannel <- int(settings.GetValueOfID(`checkOutStartHour`).(int32))
			}()

			go func() {
				checkOutStartMinuteChannel <- int(settings.GetValueOfID(`checkOutStartMinute`).(int32))
			}()

			newStartOfficeHourOfDay := <-startOfficeHourOfDayChannel

			if GetStartOfficeHourOfDay() != newStartOfficeHourOfDay {
				setStartOfficeHourOfDay(newStartOfficeHourOfDay)
			}

			newStartOfficeMinuteOfDay := <-startOfficeMinuteOfDayChannel

			if GetStartOfficeMinuteOfDay() != newStartOfficeMinuteOfDay {
				setStartOfficeMinuteOfDay(newStartOfficeMinuteOfDay)
			}

			newCheckInBufferStartHour := <-checkInBufferStartHourChannel

			if GetCheckInBufferStartHour() != newCheckInBufferStartHour {
				setCheckInBufferStartHour(newCheckInBufferStartHour)
			}

			newCheckInBufferStartMinute := <-checkInBufferStartMinuteChannel

			if GetCheckInBufferStartMinute() != newCheckInBufferStartMinute {
				setCheckInBufferStartMinute(newCheckInBufferStartMinute)
			}

			newCheckInBufferEndHour := <-checkInBufferEndHourChannel

			if GetCheckInBufferEndHour() != newCheckInBufferEndHour {
				setCheckInBufferEndHour(newCheckInBufferEndHour)
			}

			newCheckInBufferEndMinute := <-checkInBufferEndMinuteChannel

			if GetCheckInBufferEndMinute() != newCheckInBufferEndMinute {
				setCheckInBufferEndMinute(newCheckInBufferEndMinute)
			}

			newCheckOutStartHour := <-checkOutStartHourChannel

			if GetCheckOutStartHour() != newCheckOutStartHour {
				setCheckOutStartHour(newCheckOutStartHour)
			}

			newCheckOutStartMinute := <-checkOutStartMinuteChannel

			if GetCheckOutStartMinute() != newCheckOutStartMinute {
				setCheckOutStartMinute(newCheckOutStartMinute)
			}

		},
	) // 給物件增加定時任務

	newCron.Start()

	select {}

}

// setStartOfficeHourOfDay - 設定辦公起始小時
func setStartOfficeHourOfDay(inputStartOfficeHourOfDay int) {
	startOfficeHourOfDayRWMutex.Lock()
	startOfficeHourOfDay = inputStartOfficeHourOfDay
	startOfficeHourOfDayRWMutex.Unlock()
}

// GetStartOfficeHourOfDay - 取得辦公起始小時
/*
 * @return int 結果
 */
func GetStartOfficeHourOfDay() int {
	startOfficeHourOfDayRWMutex.RLock()
	gotStartOfficeHourOfDay := startOfficeHourOfDay
	startOfficeHourOfDayRWMutex.RUnlock()
	return gotStartOfficeHourOfDay
}

// setStartOfficeMinuteOfDay - 設定辦公起始分鐘
func setStartOfficeMinuteOfDay(inputStartOfficeMinuteOfDay int) {
	startOfficeMinuteOfDayRWMutex.Lock()
	startOfficeMinuteOfDay = inputStartOfficeMinuteOfDay
	startOfficeMinuteOfDayRWMutex.Unlock()
}

// GetStartOfficeMinuteOfDay - 取得辦公起始分鐘
/*
 * @return int 結果
 */
func GetStartOfficeMinuteOfDay() int {
	startOfficeMinuteOfDayRWMutex.RLock()
	gotStartOfficeMinuteOfDay := startOfficeMinuteOfDay
	startOfficeMinuteOfDayRWMutex.RUnlock()
	return gotStartOfficeMinuteOfDay
}

// setCheckInBufferStartHour - 設定打卡緩衝起始小時
/*
 * @param int inputCheckInBufferStartHour 輸入打卡緩衝起始小時
 */
func setCheckInBufferStartHour(inputCheckInBufferStartHour int) {
	checkInBufferStartHourRWMutex.Lock()
	checkInBufferStartHour = inputCheckInBufferStartHour
	checkInBufferStartHourRWMutex.Unlock()
}

// GetCheckInBufferStartHour - 取得打卡緩衝起始小時
/*
 * @return int 結果
 */
func GetCheckInBufferStartHour() int {
	checkInBufferStartHourRWMutex.RLock()
	gotCheckInBufferStartHour := checkInBufferStartHour
	checkInBufferStartHourRWMutex.RUnlock()
	return gotCheckInBufferStartHour
}

// setCheckInBufferStartMinute - 設定打卡緩衝起始分鐘
/*
 * @param int inputCheckInBufferStartMinute 輸入打卡緩衝起始分鐘
 */
func setCheckInBufferStartMinute(inputCheckInBufferStartMinute int) {
	checkInBufferStartMinuteRWMutex.Lock()
	checkInBufferStartMinute = inputCheckInBufferStartMinute
	checkInBufferStartMinuteRWMutex.Unlock()
}

// GetCheckInBufferStartMinute - 取得打卡緩衝起始分鐘
/*
 * @return int 結果
 */
func GetCheckInBufferStartMinute() int {
	checkInBufferStartMinuteRWMutex.RLock()
	gotCheckInBufferStartMinute := checkInBufferStartMinute
	checkInBufferStartMinuteRWMutex.RUnlock()
	return gotCheckInBufferStartMinute
}

// setCheckInBufferEndHour - 設定打卡緩衝結束小時
/*
 * @param int inputCheckInBufferEndHour 輸入打卡緩衝結束小時
 */
func setCheckInBufferEndHour(inputCheckInBufferEndHour int) {
	checkInBufferEndHourRWMutex.Lock()
	checkInBufferEndHour = inputCheckInBufferEndHour
	checkInBufferEndHourRWMutex.Unlock()
}

// GetCheckInBufferEndHour - 取得打卡緩衝結束小時
/*
 * @return int 結果
 */
func GetCheckInBufferEndHour() int {
	checkInBufferEndHourRWMutex.RLock()
	gotCheckInBufferEndHour := checkInBufferEndHour
	checkInBufferEndHourRWMutex.RUnlock()
	return gotCheckInBufferEndHour
}

// setCheckInBufferEndMinute - 設定打卡緩衝結束分鐘
/*
 * @param int inputCheckInBufferEndMinute 輸入打卡緩衝結束分鐘
 */
func setCheckInBufferEndMinute(inputCheckInBufferEndMinute int) {
	checkInBufferEndMinuteRWMutex.Lock()
	checkInBufferEndMinute = inputCheckInBufferEndMinute
	checkInBufferEndMinuteRWMutex.Unlock()
}

// GetCheckInBufferEndMinute - 取得打卡緩衝結束分鐘
/*
 * @return int 結果
 */
func GetCheckInBufferEndMinute() int {
	checkInBufferEndMinuteRWMutex.RLock()
	gotCheckInBufferEndMinute := checkInBufferEndMinute
	checkInBufferEndMinuteRWMutex.RUnlock()
	return gotCheckInBufferEndMinute
}

// setCheckOutStartHour - 設定打卡起始小時
/*
 * @param int inputCheckOutStartHour 輸入打卡起始小時
 */
func setCheckOutStartHour(inputCheckOutStartHour int) {
	checkOutStartHourRWMutex.Lock()
	checkOutStartHour = inputCheckOutStartHour
	checkOutStartHourRWMutex.Unlock()
}

// GetCheckOutStartHour - 取得打卡起始小時
/*
 * @return int 結果
 */
func GetCheckOutStartHour() int {
	checkOutStartHourRWMutex.RLock()
	gotCheckOutStartHour := checkOutStartHour
	checkOutStartHourRWMutex.RUnlock()
	return gotCheckOutStartHour
}

// setCheckOutStartMinute - 設定打卡起始分鐘
/*
 * @param int inputCheckOutStartMinute 輸入打卡起始分鐘
 */
func setCheckOutStartMinute(inputCheckOutStartMinute int) {
	checkOutStartMinuteRWMutex.Lock()
	checkOutStartMinute = inputCheckOutStartMinute
	checkOutStartMinuteRWMutex.Unlock()
}

// GetCheckOutStartMinute - 取得打卡起始分鐘
/*
 * @return int 結果
 */
func GetCheckOutStartMinute() int {
	checkOutStartMinuteRWMutex.RLock()
	gotCheckOutStartMinute := checkOutStartMinute
	checkOutStartMinuteRWMutex.RUnlock()
	return gotCheckOutStartMinute
}
