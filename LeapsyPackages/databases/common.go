package databases

import (
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"leapsy.com/packages/configurations"
)

const (
	equalToConstString            = `$eq`  // =
	greaterThanConstString        = `$gt`  // >
	greaterThanEqualToConstString = `$gte` // >=
	lessThanConstString           = `$lt`  // <
	lessThanEqualToConstString    = `$lte` // <=
	modConstString                = `$mod`
)

const (
	sizeConstString = `$size`
)

const (
	lookupConstString         = `$lookup`
	fromConstString           = `from`
	localFieldConstString     = `localField`
	foreignFieldConstString   = `foreignField`
	pipelineConstString       = `pipeline`
	matchConstString          = `$match`
	sortConstString           = `$sort`
	groupConstString          = `$group`
	setConstString            = `$set`
	unsetConstString          = `$unset`
	asConstString             = `as`
	unwindConstString         = `$unwind`
	firstConstString          = `$first`
	lastConstString           = `$last`
	outConstString            = `$out`
	subtractConstString       = `$subtract`
	projectConstString        = `$project`
	mergeConstString          = `$merge`
	intoConstString           = `into`
	onConstString             = `on`
	whenMatchedConstString    = `whenMatched`
	whenNotMatchedConstString = `whenNotMatched`
	concatConstString         = `$concat`
	exprConstString           = `$expr`
	andConstString            = `$and`
	orConstString             = `$or`
	limitConstString          = `$limit`
)

var (
	eventSortPrimitiveM = bson.M{
		`time`: 1,
	}
	codeSortPrimitiveM = bson.M{
		`employeeid`: 1,
	}
)

var (
	appRecordSortPrimitiveD = bson.D{
		{
			Key:   `time`,
			Value: 1,
		},
	}
)

var ( // 記錄器
	batchSize = configurations.GetConfigPositiveIntValueOrPanic(`local`, `batch-size`) // 取得預設批次大小
	accountRWMutex,
	deviceRWMutex,
	eventRWMutex,
	appRecordRWMutex,
	appRecordTypeRWMutex,
	settingRWMutex,
	employeeRWMutex,
	codeRWMutex,
	applicationRWMutex sync.RWMutex // 讀寫鎖
)
