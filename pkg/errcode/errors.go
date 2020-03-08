package errcode

//BizErr :Biz module err type
type BizErr int

/**
* 通用错误码 520000 ~ 529999
 */
// [520000 ~ 520200) 000  Param
const (
	//ErrCommonParamInvalidType:520000
	ErrCommonParamInvalidType BizErr = 520000 + iota
	//ErrCommonParamInvalidValue:520001
	ErrCommonParamInvalidValue
	//ErrCommonParamJSONEncodeGetNull:520002
	ErrCommonParamJSONEncodeGetNull
	//ErrCommonParamJSONEncodeFail:520003
	ErrCommonParamJSONEncodeFail
	//ErrCommonParamJSONDecodeFail:520004
	ErrCommonParamJSONDecodeFail
	//ErrCommonParamCheckSignFail:520005
	ErrCommonParamCheckSignFail
)

// [520200 ~ 520600) 200  Redis
const (
	//ErrCommonRedisLoadError:520200
	ErrCommonRedisLoadError BizErr = 520200 + iota
	//ErrCommonRedisSetFail:520201
	ErrCommonRedisSetFail
	//ErrCommonRedisGetFail:520202
	ErrCommonRedisGetFail
	//ErrCommonRedisDeleteError:520203
	ErrCommonRedisDeleteError
	// ErrCommonRedisQueueOutLimit:520204
	ErrCommonRedisQueueOutLimit
	// ErrCommonRedisCkvError:520205
	ErrCommonRedisCkvError
	// ErrCommonRedisCkvLockProcess:520206
	ErrCommonRedisCkvLockProcess
	// ErrCommonRedisKeySetexFail:520207
	ErrCommonRedisKeySetexFail
)

// [520600 ~ 521000) 400  DB(MySQL...) 数据库
const (
	// ErrCommonDBLoadError:520600
	ErrCommonDBLoadError BizErr = 520600 + iota
	// ErrCommonDBTableFieldsError:520601
	ErrCommonDBTableFieldsError
	// ErrCommonDBInsertError:520602
	ErrCommonDBInsertError
	// ErrCommonDBDeleteError:520603
	ErrCommonDBDeleteError
	// ErrCommonDBSelectError:520604
	ErrCommonDBSelectError
	// ErrCommonDBUpdateError:520605
	ErrCommonDBUpdateError
	// ErrCommonDBGetStrategyConfError:520606
	ErrCommonDBGetStrategyConfError
	// ErrCommonDBProcessStrategyError:520607
	ErrCommonDBProcessStrategyError
	// ErrCommonDBSyncStrategyConfError:520608
	ErrCommonDBSyncStrategyConfError
	// ErrCommonDBUpdateNewGUIDError:520609
	ErrCommonDBUpdateNewGUIDError
)

// [521000 ~ 521600) 600  HTTP RPC
const (
	//ErrCommonHTTPReadFail:521000
	ErrCommonHTTPReadFail BizErr = 521000 + iota
	//ErrCommonHTTPConnectTimeOut:521001
	ErrCommonHTTPConnectTimeOut
	//ErrCommonHTTPReadTimeOut:521002
	ErrCommonHTTPReadTimeOut
	//ErrCommonHTTPBillFinishError:521003
	ErrCommonHTTPBillFinishError
	//ErrCommonHTTPPassportPassengerError:521004
	ErrCommonHTTPPassportPassengerError
)

// [521600 ~ 522600) 1000  通用业务错误码
const (
	//ErrCommonAreaNotOpenService:521600
	ErrCommonAreaNotOpenService BizErr = 521600 + iota
	//ErrCommonConfigNotFount:521601
	ErrCommonConfigNotFount
	//ErrCommonCoordinateInvalid:521602
	ErrCommonCoordinateInvalid
	//ErrCommonCurrencyInvalid:521603
	ErrCommonCurrencyInvalid
	//ErrCommonPointsReqError:521604
	ErrCommonPointsReqError
	//ErrCommonPointsReqStatusError:521605
	ErrCommonPointsReqStatusError
	//ErrCommonTokenInvalid:521606
	ErrCommonTokenInvalid
	//ErrCommonOrderIDInvalid:521607
	ErrCommonOrderIDInvalid
	//ErrCommonPassengerIDInvalid:521608
	ErrCommonPassengerIDInvalid
	//ErrCommonPassengerIDUnmatch:521609
	ErrCommonPassengerIDUnmatch
	//ErrCommonDriverIDInvalid:521610
	ErrCommonDriverIDInvalid
	//ErrCommonNetworkException:521611
	ErrCommonNetworkException
)

// [523000 ~ 523050) PHP Error Handler 错误码
const (
	//ErrCommonErrorHandlerGetFatal：523000
	ErrCommonErrorHandlerGetFatal BizErr = 523000 + iota
	//ErrCommonErrorHandlerGetWarn :523001
	ErrCommonErrorHandlerGetWarn
	//ErrCommonErrorHandlerGetException :523002
	ErrCommonErrorHandlerGetException
	//ErrCommonErrorHandlerGetUncaughtException :523003
	ErrCommonErrorHandlerGetUncaughtException
)

// [530000 ~ 530102) passenger
const (
	//ErrPassengerActivityMisEstimateReq :530000
	ErrPassengerActivityMisEstimateReq BizErr = 530000 + iota
	//ErrPassengerActivityMisEstimateReqStatusError :530001
	ErrPassengerActivityMisEstimateReqStatusError
	//ErrPassengerFenceServiceRetDataError :530006
	ErrPassengerFenceServiceRetDataError BizErr = 530006
	//ErrPassengerCancelTripExpire :530007
	ErrPassengerCancelTripExpire BizErr = 530007
)

// [530102 ~ 535000) passenger
const (
	//ErrPassengerStationServiceReq :530102
	ErrPassengerStationServiceReq BizErr = 530102 + iota
	//ErrPassengerStationForecastReq :530103
	ErrPassengerStationForecastReq
	//ErrPassengerStationFenceFailed :530104
	ErrPassengerStationFenceFailed
	//ErrPassengerStationFenceNotHit :530105
	ErrPassengerStationFenceNotHit
	//ErrPassengerStationMatchReq :530106
	ErrPassengerStationMatchReq
)

// [535000 ~ 540000) driver
// [540000 ~ 545000) order
// [545000 ~ 550000) themis
// [550000 ~ 555000) carpool
// [555000 ~ 560000) mis

// [560000 ~ 561000) Athena
const (
	//ErrAthenaRequestFail :560001
	ErrAthenaRequestFail BizErr = 560001 + iota
)

// [561000 ~ 562000) minos
// [562000 ~ 563000) minos_async
// [563000 ~ 565000) develop
// [565000 ~ 570000) hermes
// [570000 ~ 571000) soter
// [571000 ~ 572000) horae
// [572000 ~ 574000) internal-api
// [574000 ~ 575000) dmc
// [575000 ~ 576000) price-api
// [576000 ~ 580000) openapi\openapi-inner
// [580000 ~ 582000) csi


//Code :get error code
func (pErr BizErr) Code() int {
	return (int)(pErr)
}

//Error :get error msg
func (pErr BizErr) Error() string {
	switch pErr {
	case ErrCommonParamInvalidType:
		return "param type is error"
	case ErrCommonParamInvalidValue:
		return "param value is error"
	case ErrCommonParamJSONEncodeGetNull:
		return "json encode param get a null"
	case ErrCommonParamJSONEncodeFail:
		return "json encode data fail"
	case ErrCommonParamJSONDecodeFail:
		return "json decode data fail"
	case ErrCommonParamCheckSignFail:
		return "check sign data fail"

	// [520200 ~ 520400) 200  Redis 缓存
	case ErrCommonRedisLoadError:
		return "redis init error"
	case ErrCommonRedisSetFail:
		return "redis set error"
	case ErrCommonRedisGetFail:
		return "redis read error"
	case ErrCommonRedisDeleteError:
		return "redis delete error"
	case ErrCommonRedisQueueOutLimit:
		return "redis queue out of limit error"
	case ErrCommonRedisCkvError:
		return "redis ckv error"
	case ErrCommonRedisCkvLockProcess:
		return "redis ckv lock process error"
	case ErrCommonRedisKeySetexFail:
		return "redis setex fail"

	// [520400 ~ 520600) 200  Memcache

	// [520600 ~ 521000) 400  DB(MySQL...) 数据库
	case ErrCommonDBLoadError:
		return "db init error"
	case ErrCommonDBTableFieldsError:
		return "db table fields error"
	case ErrCommonDBInsertError:
		return "db insert error"
	case ErrCommonDBDeleteError:
		return "db delete error"
	case ErrCommonDBSelectError:
		return "db select error"
	case ErrCommonDBUpdateError:
		return "db update error"
	case ErrCommonDBGetStrategyConfError:
		return "db get strategy conf error"
	case ErrCommonDBProcessStrategyError:
		return "db process strategy error"
	case ErrCommonDBSyncStrategyConfError:
		return "db sync strategy conf error"
	case ErrCommonDBUpdateNewGUIDError:
		return "db update new guid error"

	// [521000 ~ 521600) 600  HTTP RPC
	case ErrCommonHTTPReadFail:
		return "HTTP response data read fail"
	case ErrCommonHTTPConnectTimeOut:
		return "HTTP connect timeout"
	case ErrCommonHTTPReadTimeOut:
		return "HTTP read response data timeout"
	case ErrCommonHTTPBillFinishError:
		return "HTTP call finish bill fail"
	case ErrCommonHTTPPassportPassengerError:
		return "HTTP call passport fail"

	// [521600 ~ 522600) 1000  通用业务错误码
	case ErrCommonAreaNotOpenService:
		return "current area not open service"
	case ErrCommonConfigNotFount:
		return "config not found"
	case ErrCommonCoordinateInvalid:
		return "invalid coordinate data"
	case ErrCommonCurrencyInvalid:
		return "currency error"
	case ErrCommonPointsReqError:
		return "points request error"
	case ErrCommonPointsReqStatusError:
		return "points request status error"
	case ErrCommonTokenInvalid:
		return "invalid token"
	case ErrCommonOrderIDInvalid:
		return "invalid order id"
	case ErrCommonPassengerIDInvalid:
		return "invalid passenger id"
	case ErrCommonPassengerIDUnmatch:
		return "passenger id not match"
	case ErrCommonDriverIDInvalid:
		return "invalid driver id"
	case ErrCommonNetworkException:
		return "network exception"

	// [522600 ~ 523000) 400  Kafka

	// [523000 ~ 523050) PHP Error Handler 错误码
	case ErrCommonErrorHandlerGetFatal:
		return "error handler got Fatal error"
	case ErrCommonErrorHandlerGetWarn:
		return "error handler got Warn error"
	case ErrCommonErrorHandlerGetException:
		return "exception handler got Exception"
	case ErrCommonErrorHandlerGetUncaughtException:
		return "exception handler got uncaught Exception"

	/**
	 * Passenger 模块
	 *
	 * [530000 ~ 535000) 5000
	 */
	case ErrPassengerActivityMisEstimateReq:
		return "estimate request passenger activity error"
	case ErrPassengerActivityMisEstimateReqStatusError:
		return "estimate request passenger activity result error"
	case ErrPassengerStationServiceReq:
		return "estimate request station service error"
	case ErrPassengerFenceServiceRetDataError:
		return "estimate request fence service ret data error"
	case ErrPassengerCancelTripExpire:
		return "cancel trip expire"
	case ErrPassengerStationForecastReq:
		return "passenger get station status error"
	case ErrPassengerStationFenceFailed:
		return "passenger get station fence failed"
	case ErrPassengerStationFenceNotHit:
		return "passenger get station fence not hit"
	case ErrPassengerStationMatchReq:
		return "passenger get station match degree failed"

	/**
	 * driver 模块
	 *
	 * [535000 ~ 540000)
	 */
	// const E_DRIVER_

	/**
	 * order 模块
	 *
	 * [540000 ~ 545000)
	 */
	// const E_ORDER_

	/**
	 * themis 模块
	 *
	 * [545000 ~ 550000)
	 */
	// const E_THEMIS_

	/**
	 * carpool 模块
	 *
	 * [550000 ~ 555000)
	 */
	// const E_CARPOOL_

	/**
	 * mis 模块
	 *
	 * [555000 ~ 560000) 5000
	 */

	/**
	 * Athena 模块
	 *
	 * [560001 ~ 560010) 10
	 */
	case ErrAthenaRequestFail:
		return "estimate request athena result fail"
	default:
		return "unknow err"
	}
}
