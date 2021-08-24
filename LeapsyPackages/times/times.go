package times

import (
	"time"

	"leapsy.com/packages/logings"
)

// IsMinute - 判斷時間是否為整分
/**
 * @param  time.Time dateTime  時間
 * @return bool 判斷是否為整秒
 */
func IsMinute(dateTime time.Time) bool {
	return dateTime.Local().Second() == 0 // 秒為零
}

// GetHourlyBounds - 取得時間上下限小時
/**
 * @param  time.Time dateTime  時間
 * @return time.Time lower 下限小時
 * @return time.Time upper 上限小時
 */
func GetHourlyBounds(dateTime time.Time) (low time.Time, upper time.Time) {

	localDateTime := dateTime.Local()

	low = ConvertToHourlyDateTime(localDateTime)
	upper = low.Add(time.Hour)

	logings.SendLog(
		[]string{`取得時間 %+v 下限小時 %+v 上限小時 %+v`},
		[]interface{}{localDateTime, low, upper},
		nil,
		0,
	)

	return
}

// IsHour - 判斷時間是否為整點
/**
 * @param  time.Time dateTime  時間
 * @return bool 判斷是否為整點
 */
func IsHour(dateTime time.Time) bool {
	localDateTime := dateTime.Local()
	return localDateTime.Minute() == 0 && IsMinute(localDateTime) // 分秒為零
}

// ConvertToHourlyDateTime - 轉成整點時間
/**
 * @param time.Time dateTime 時間
 * @return time.Time returnHourlyDateTime 回傳小時時間
 */
func ConvertToHourlyDateTime(dateTime time.Time) (returnHourlyDateTime time.Time) {

	localDateTime := dateTime.Local()

	returnHourlyDateTime = time.Date(localDateTime.Year(), localDateTime.Month(), localDateTime.Day(), localDateTime.Hour(), 0, 0, 0, time.Local)

	logings.SendLog(
		[]string{`將時間 %+v 轉成整點時間 %+v `},
		[]interface{}{localDateTime, returnHourlyDateTime},
		nil,
		0,
	)

	return
}

// GetDailyBounds - 取得時間上下限日
/**
 * @param  time.Time dateTime  時間
 * @return  time.Time lower  下限日
 * @return  time.Time upper  上限日
 */
func GetDailyBounds(dateTime time.Time) (low time.Time, upper time.Time) {

	localDateTime := dateTime.Local()

	low = ConvertToDailyDateTime(localDateTime) // 今日
	upper = low.AddDate(0, 0, 1)                // 明日

	logings.SendLog(
		[]string{`取得時間 %+v 下限日 %+v 上限日 %+v`},
		[]interface{}{localDateTime, low, upper},
		nil,
		0,
	)

	return // 回傳
}

// IsDay - 判斷是否為整日
/**
 * @param  time.Time dateTime  時間
 * @return  bool 判斷是否為整日
 */
func IsDay(dateTime time.Time) bool {

	localDateTime := dateTime.Local()

	return localDateTime.Hour() == GetStartOfficeHourOfDay() &&
		localDateTime.Minute() == GetStartOfficeMinuteOfDay() &&
		IsMinute(localDateTime) // 秒為零
}

// ConvertToDailyDateTime - 轉成整日時間
/**
 * @param time.Time dateTime 時間
 * @return time.Time returnDailyDateTime 回傳小時時間
 */
func ConvertToDailyDateTime(dateTime time.Time) (returnDailyDateTime time.Time) {

	localDateTime := dateTime.Local()

	// 修改時間
	returnDailyDateTime = time.Date(
		localDateTime.Year(),
		localDateTime.Month(),
		localDateTime.Day(),
		GetStartOfficeHourOfDay(),
		GetStartOfficeMinuteOfDay(),
		0,
		0,
		time.Local,
	)

	if !localDateTime.IsZero() && localDateTime.Before(returnDailyDateTime) {
		returnDailyDateTime = returnDailyDateTime.AddDate(0, 0, -1)
	}

	logings.SendLog(
		[]string{`將時間 %+v 轉成整日時間 %+v `},
		[]interface{}{localDateTime, returnDailyDateTime},
		nil,
		0,
	)

	return
}

// GetMonthlyBounds - 取得時間上下限月
/**
 * @param  time.Time dateTime  時間
 * @return  time.Time lower  下限月
 * @return  time.Time upper  上限月
 */
func GetMonthlyBounds(dateTime time.Time) (low time.Time, upper time.Time) {

	localDateTime := dateTime.Local()

	low = ConvertToMonthlyDateTime(localDateTime) // 這個月
	upper = low.AddDate(0, 1, 0)                  // 下個月

	logings.SendLog(
		[]string{`取得時間 %+v 下限月 %+v 上限月 %+v `},
		[]interface{}{localDateTime, low, upper},
		nil,
		0,
	)

	return // 回傳
}

// IsMonth - 判斷是否為整月
/**
 * @param  time.Time dateTime  時間
 * @return  bool 判斷是否為整月
 */
func IsMonth(dateTime time.Time) bool {
	localDateTime := dateTime.Local()
	// 日為一、時分為換日時分、秒為零
	return localDateTime.Day() == 1 && IsDay(localDateTime)
}

// ConvertToMonthlyDateTime - 轉成整月時間
/**
 * @param time.Time dateTime 時間
 * @return time.Time returnMonthlyDateTime 回傳小時時間
 */
func ConvertToMonthlyDateTime(dateTime time.Time) (returnMonthlyDateTime time.Time) {

	localDateTime := dateTime.Local()

	// 修改時間
	returnMonthlyDateTime = time.Date(
		localDateTime.Year(),
		localDateTime.Month(),
		1,
		GetStartOfficeHourOfDay(),
		GetStartOfficeMinuteOfDay(),
		0,
		0,
		time.Local,
	)

	if !localDateTime.IsZero() && localDateTime.Before(returnMonthlyDateTime) {
		returnMonthlyDateTime = returnMonthlyDateTime.AddDate(0, -1, 0)
	}

	logings.SendLog(
		[]string{`將時間 %+v 轉成整月時間 %+v `},
		[]interface{}{localDateTime, returnMonthlyDateTime},
		nil,
		0,
	)

	return
}

// GetWeeklyBounds - 取得時間上下限星期日
/**
 * @param  time.Time dateTime  時間
 * @return time.Time lower 下限星期日
 * @return time.Time upper 上限星期日
 */
func GetWeeklyBounds(dateTime time.Time) (low time.Time, upper time.Time) {

	localDateTime := dateTime.Local()

	weekDayInt := (int(localDateTime.Weekday()))

	low = dateTime.AddDate(0, 0, -((weekDayInt + 6) % 7))
	upper = dateTime.AddDate(0, 0, (7-weekDayInt)%7)

	logings.SendLog(
		[]string{`取得時間 %+v 下限星期日 %+v 上限星期日 %+v`},
		[]interface{}{localDateTime, low, upper},
		nil,
		0,
	)

	return
}
