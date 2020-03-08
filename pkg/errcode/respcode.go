/*
Package errcode 是外部(返回给端的)错误码

  规范如下:

  1.外部错误码常量名为R+大驼峰

  2.外部错误码数值共6位,格式为1-ZZ-Y-XX
       其中ZZ为大类序号,范围[01,99],
       Y为小类序号,范围[1,9],
       XX为具体错误码,范围[01,99]


  3.具体分类如下(!常量名统一采用 以R开头的 大驼峰格式!)
   3.1 成功(0|RSuccess)

   3.2 通用错误码(101yxx|RCommon*)
       3.2.1 参数(1011xx): 参数空、非法、无效等
       3.2.2 服务端(1012xx): 服务器出错、第三方服务异常、数据(库)错误、缓存错误等
       3.2.3 网络/连接(1013xx): 网络不稳定、连接超时等
       3.2.4 签名错误(101401-101420): 验证错误等

   3.3 订单相关(102yxx|ROrder*)
       3.3.1 状态(1021xx): 订单超时、订单不存在、进行中、已支付、未支付、已关单、价格已更新、已过期等
       3.3.2 取消/应答(1022xx): (被)取消成功/失败、被应答等

   3.4 场景相关(103yxx|RScene*)
       3.4.1 航班信息(1031xx): 航班、起飞、到达、取消、延误等
       3.4.2 接送机(1032xx): 接机口、机场、预约时间等
       3.4.3 拼车(1033xx): 站点、拼友、线路、跨城等

   3.5 支付相关(104yxx|RPay*)
       3.5.1 支付方式(1041xx): 线上、线下、微信、支付宝、余额、企业支付等
       3.5.2 支付结果(1042xx): 成功、失败、重试、异常、繁忙等

   3.6 提现相关(105yxx|RWithdraw*)
       3.6.1 提现规则(1051xx): 提现日、余额、次数、申请等
       3.6.2 提现结果(1052xx): 成功、失败、超时、冻结、关闭等

   3.7 语音消息相关(107yxx|RVoice*)
       3.7.1 录制(1071xx): 录制过短、过长、重新录制等
       3.7.2 传送(1072xx): 保存成功/失败、上传成功/失败等

   3.8 券相关(108yxx|RCoupon*)
       3.8.1 领取(1081xx): 重复领取、已领取、领取失败/成功、券发送成功/失败等
       3.8.2 使用(1082xx): 不符合使用规则、获取券信息失败、查询券异常等

   3.9 管理平台相关(109yxx|RMis*)
       3.9.1 司机信息(1091xx): 司机注册、查询、修改完善信息等
       3.9.2 管理员操作(1092xx): 编辑/管理员 添加配置、查询处理case等
       3.9.3 司机考试(1093xx): 司机申请/预约/登录/查询考试等

   3.10 运营活动(110yxx|RActi*)
       3.10.1 活动时效范围(1101xx): 过期、即将展开、未开通、不存在、该城市不参与等
       3.10.2 活动规则(1102xx): 具体规则、限制、建议等

   3.11 状态和服务(111yxx|RStat*)
       3.11.1 登录状态(1111xx): 成功、过期、失败、(失效/验证失败)重新登录、繁忙等
       3.11.2 定位状态(1112xx): 定位异常、起始/终点定位错误等
       3.11.3 开城状态(1113xx): xx功能未开通、起点/终点/该城市未开通xx等
       3.11.4 司乘状态(1114xx): 司机已到达/迟到、乘客迟到、获取不到周围司机、没有司机出车等
       3.11.5 行程服务(1115xx): 行程前/中/后服务、开具发票等

   3.12 验证判责(112yxx|RJudge*)
       3.12.1 判责(1121xx): 判责相关
       3.12.2 反作弊(1122xx): 反作弊相关
       3.12.3 soter(1123xx): soter相关

   3.13 开放api(113yxx|ROpen*)


   3.14 其他信息(114yxx|RWarn*):匹配不到上面的分类,且不符合独自建立一个分类的


  4.添加错误码时,优先匹配上面的分类(3.1~3.13) 若都不符合,且又不属于单独分类,则请放入3.14.
*/
package errcode


//RespCode : passenger module response code
type RespCode int


/*--------------------------按新规范添加的错误码(如下)--------------------------------*/
const (
    /**********************3.1 成功(0|RSuccess)****************************/

    //RSuccess :成功
    RSuccess RespCode = 0

    /*****************3.2 通用错误码(101yxx|RCommon*)***********************/
    //3.2.1 参数(1011xx): 参数空、非法、无效等

    //RCommonParamEmpty :参数空
    RCommonParamEmpty RespCode = 101101
    //RCommonParamIllegal :参数非法、类型不对
    RCommonParamIllegal RespCode = 101102
    //RCommonParamInvalid :参数无效、不在范围内
    RCommonParamInvalid RespCode = 101103
    //RCommonTokenInvalid :token失效
    RCommonTokenInvalid RespCode = 101104

    //3.2.2 服务端(1012xx): 服务器出错、第三方服务异常、数据(库)错误、缓存错误等

    //RCommonServerError :服务器出错
    RCommonServerError RespCode = 101201
    //RCommonMysqlError :mysql出错
    RCommonMysqlError  RespCode = 101202
    //RCommonCurlError :curl出错
    RCommonCurlError   RespCode = 101203
    //RCommonThriftError :thrift出错
    RCommonThriftError RespCode = 101204
    //RCommonCacheError :cache出错
    RCommonCacheError  RespCode = 101205

    //3.2.3 网络/连接(1013xx): 网络不稳定、连接超时等

    //RCommonNetworkError :网络异常、不稳定
    RCommonNetworkError   RespCode = 101301
    //RCommonConnectTimeout :连接超时
    RCommonConnectTimeout RespCode = 101302

    //3.2.4 签名错误(101401-101420): 验证错误等

    //RCommonDsigInvalid :开启验签阻拦后,签名验证错误
    RCommonDsigInvalid RespCode = 101401

    /**********************3.3 订单相关(102yxx|ROrder*)*********************/
    //3.3.1 状态(1021xx): 订单超时、订单不存在、进行中、已支付、未支付、已关单、价格已更新、已过期等
    //3.3.2 取消/应答(1022xx): (被)取消成功/失败、被应答等

    /**********************3.4 场景相关(103yxx|RScene*)*********************/
    //3.4.1 航班信息(1031xx): 航班、起飞、到达、取消、延误等
    //3.4.2 接送机(1032xx): 接机口、机场、预约时间等
    //3.4.3 拼车(1033xx): 站点、拼友、线路、跨城等

    /**********************3.5 支付相关(104yxx|RPay*)***********************/
    //3.5.1 支付方式(1041xx): 线上、线下、微信、支付宝、余额、企业支付等
    //3.5.2 支付结果(1042xx): 成功、失败、重试、异常、繁忙等

    /**********************3.6 提现相关(105yxx|RWithdraw*)******************/
    //3.6.1 提现规则(1051xx): 提现日、余额、次数、申请等
    //3.6.2 提现结果(1052xx): 成功、失败、超时、冻结、关闭等

    /**********************3.7 语音消息相关(107yxx|RVoice*)******************/
    //3.7.1 录制(1071xx): 录制过短、过长、重新录制等
    //3.7.2 传送(1072xx): 保存成功/失败、上传成功/失败等

    /**********************3.8 券相关(108yxx|RCoupon*)**********************/
    //3.8.1 领取(1081xx): 重复领取、已领取、领取失败/成功、券发送成功/失败等
    //3.8.2 使用(1082xx): 不符合使用规则、获取券信息失败、查询券异常等

    /**********************3.9 管理平台相关(109yxx|RMis*)********************/
    //3.9.1 司机信息(1091xx): 司机注册、查询、修改完善信息等
    //3.9.2 管理员操作(1092xx): 编辑/管理员 添加配置、查询处理case等
    //3.9.3 司机考试(1093xx): 司机申请/预约/登录/查询考试等

    /**********************3.10 运营活动(110yxx|RActi*)**********************/
    //3.10.1 活动时效范围(1101xx): 过期、即将展开、未开通、不存在、该城市不参与等
    //3.10.2 活动规则(1102xx): 具体规则、限制、建议等

    /**********************3.11 状态和服务(111yxx|RStat*)*********************/
    //3.11.1 登录状态(1111xx): 成功、过期、失败、(失效/验证失败)重新登录、繁忙等
    //3.11.2 定位状态(1112xx): 定位异常、起始/终点定位错误等
    //3.11.3 开城状态(1113xx): xx功能未开通、起点/终点/该城市未开通xx等
    //3.11.4 司乘状态(1114xx): 司机已到达/迟到、乘客迟到、获取不到周围司机、没有司机出车等
    //3.11.5 行程服务(1115xx): 行程前/中/后服务、开具发票等

    /**********************3.12 验证判责(112yxx|RJudge*)*********************/
    //3.12.1 判责(1121xx): 判责相关
    //3.12.2 反作弊(1122xx): 反作弊相关
    //3.12.3 soter(1123xx): soter相关

    /**********************3.13 开放api(113yxx|ROpen*)**********************/

    /**3.14 其他信息(114yxx|RWarn*):匹配不到上面的分类,且不符合独自建立一个分类的**/
)



/*----------------------------现有的旧错误码(如下)-------------------------------------*/
const (
	//Success : 服务器端错误 Server
	Success RespCode = 0

	//SuccessInternalError : internal error
	SuccessInternalError RespCode = -1
	//SuccessJSONDecodeError :json decode error
	SuccessJSONDecodeError RespCode = -2

	// 乘客端 Passenger 错误码

	//PassengerSuccess :成功
	PassengerSuccess RespCode = 0

	//PassengerTokenError :token错误
	PassengerTokenError RespCode = 101
	//PassengerVoiceUploadError :上传语音错误
	PassengerVoiceUploadError RespCode = 1001
	//PassengerVoiceTooShort :语音太短
	PassengerVoiceTooShort RespCode = 1002
	//PassengerVoiceTooLong :语音太长
	PassengerVoiceTooLong RespCode = 1003
	//PassengerInputError :输入错误
	PassengerInputError RespCode = 1004
	//PassengerAddressEmpty :地址为空
	PassengerAddressEmpty RespCode = 1005
	//PassengerOrderMoreFail :其他错误
	PassengerOrderMoreFail RespCode = 1006
	//PassengerGPSError :gps错误
	PassengerGPSError RespCode = 1007
	//PassengerParamsError :参数错误
	PassengerParamsError RespCode = 1010
	//PassengerParamsError505000 :参数错误505000
	PassengerParamsError505000 RespCode = 505000
	//PassengerPhoneError :手机号码错误
	PassengerPhoneError RespCode = 1012
	//PassengerSetCacheFail :设置cache错误
	PassengerSetCacheFail RespCode = 1013
	//PassengerSaveVoiceFail :保存语音错误
	PassengerSaveVoiceFail RespCode = 1014
	//PassengerOrderCreate :创建订单错误
	PassengerOrderCreate RespCode = 1015
	//PassengerServiceNotOpen :服务为开放
	PassengerServiceNotOpen RespCode = 1016
	//PassengerRecallTimeError :passenger recall time error
	PassengerRecallTimeError RespCode = 1017
	//PassengerRecallCallError :passenger recall call error
	PassengerRecallCallError RespCode = 1018
	//PassengerPayCardNotBind :未绑定支付卡
	PassengerPayCardNotBind RespCode = 1019
	//PassengerPayOverdraft :passenger pay overdraft
	PassengerPayOverdraft RespCode = 1020
	//PassengerPayNoOpenid :没有支付openid
	PassengerPayNoOpenid RespCode = 1021
	//PassengerPayHavePaid :已支付
	PassengerPayHavePaid RespCode = 1022
	//PassengerBalanceGetFail :获取余额错误
	PassengerBalanceGetFail RespCode = 1023
	//PassengerBalanceInfoError :余额信息错误
	PassengerBalanceInfoError RespCode = 1024
	//PassengerDepartDestSame :起始地相同
	PassengerDepartDestSame RespCode = 1025
	//PassengerSystemError :系统错误
	PassengerSystemError RespCode = 1026
	//PassengerCancelTripFail :取消行程错误
	PassengerCancelTripFail RespCode = 1027
	//PassengerSpamDeteched :passenger spam deteched
	PassengerSpamDeteched RespCode = 1028
	//PassengerInvoiceInfoError :语音信息错误
	PassengerInvoiceInfoError RespCode = 1029
	//PassengerInvoiceCacheError :语音缓存错误
	PassengerInvoiceCacheError RespCode = 1229
	//PassengerRecancelOrder :重新取消订单
	PassengerRecancelOrder RespCode = 1030
	//PassengerTencentMapFail :腾讯地图错误
	PassengerTencentMapFail RespCode = 1031
	//PassengerLockFail :passenger lock fail
	PassengerLockFail RespCode = 1032
	//PassengerChargeOrderError :passenger charge order error
	PassengerChargeOrderError RespCode = 1033
	//PassengerCurlPassportFail :请求passport错误
	PassengerCurlPassportFail RespCode = 1034
	//PassengerTripHadClosed :行程已关闭
	PassengerTripHadClosed RespCode = 1035
	//PassengerChannelNotB2B :channel不是B2B
	PassengerChannelNotB2B RespCode = 1036
	//PassengerChannelNotSame :channel不一致
	PassengerChannelNotSame RespCode = 1037
	//PassengerAPIWebTimeError :时间错误
	PassengerAPIWebTimeError RespCode = 1038
	//PassengerFastcarNotOpen :未开通开车
	PassengerFastcarNotOpen RespCode = 1039
	//PassengerUnpaidOrderExist :存在未支付订单
	PassengerUnpaidOrderExist RespCode = 1040
	//PassengerProductIDError :验证支付能力时，传递错误的产品号
	PassengerProductIDError RespCode = 1041
	//PassengerPayCheatDetected :验证支付能力时,调用反作弊接口获取支付真身
	PassengerPayCheatDetected RespCode = 1042
	//PassengerCarpoolNotOpen :验证是否开通拼车
	PassengerCarpoolNotOpen RespCode = 1043
	//PassengerDriverReloadOrder :取消行程时司机正在改派中的提示
	PassengerDriverReloadOrder RespCode = 1044
	//PassengerPayGetStatusError :sdk获取支付状态错误
	PassengerPayGetStatusError RespCode = 1045
	//PassengerSDKDevSignError :sdk开放平台的签名校验错误
	PassengerSDKDevSignError RespCode = 1046
	//PassengerCarpoolFreqExceed :拼车次数达到限制
	PassengerCarpoolFreqExceed RespCode = 1047
	//PassengerCarPhoneCallError :代叫车乘车人手机号码错误
	PassengerCarPhoneCallError RespCode = 1056
	//PassengerCarNameCallError :代叫车乘车人姓名错误
	PassengerCarNameCallError RespCode = 1057
	//PassengerAwardCallFail :调用积分系统出错
	PassengerAwardCallFail RespCode = 1101
	//PassengerDynamicPriceError :动态调价客户端和服务端不一致
	PassengerDynamicPriceError RespCode = 1102
	//PassengerCarlevelCouponGetFail :乘客获取车型优惠券错误
	PassengerCarlevelCouponGetFail RespCode = 1110
	//PassengerHistoryOrderGetFail :乘客获取历史订单错误
	PassengerHistoryOrderGetFail RespCode = 1120
	//PassengerNearDriverGetFail :乘客获取附近司机错误
	PassengerNearDriverGetFail RespCode = 1121
	//PassengerCarpoolRefuseBooking :预约单不支持拼车
	PassengerCarpoolRefuseBooking RespCode = 1122
	//PassengerCarpoolPriceError :拼车单一口价小于等于0
	PassengerCarpoolPriceError RespCode = 1123
	//PassengerBookingDepartTimeError :预约单出发时间无效
	PassengerBookingDepartTimeError RespCode = 1124
	//PassengerOpenAPICallFail :openApi调用失败
	PassengerOpenAPICallFail RespCode = 1160
	//PassengerGetComboFail :获取套餐数据错误
	PassengerGetComboFail RespCode = 10161
	//PassengerFreezeAmoutFail :发单冻结企业用户余额失败
	PassengerFreezeAmoutFail RespCode = 10625
	//PassengerChargeIllegal :发单冻结企业用户余额失败
	PassengerChargeIllegal RespCode = 10631

	//PassengerTagCallFail :高价值、活跃用户请求TAG
	PassengerTagCallFail RespCode = 12210
	//PassengerTagResultError :返回信息错误
	PassengerTagResultError RespCode = 12211

	//获取用户基础信息请求公共服务

	//PassengerCommonInfoGetFail :请求失败
	PassengerCommonInfoGetFail RespCode = 12310
	//PassengerCommonInfoRetError :返回信息错误
	PassengerCommonInfoRetError RespCode = 12311

	//导流系统

	//PassengerGuideSystemGetFail :请求导流失败
	PassengerGuideSystemGetFail RespCode = 13210
	//PassengerGuideSystemParamError :请求导流系统参数格式错误
	PassengerGuideSystemParamError RespCode = 13211
	//PassengerGuideSystemProcFail :导流系统内部错误
	PassengerGuideSystemProcFail RespCode = 13212
	//PassengerGuideSystemCacheFail :缓存返回信息错误
	PassengerGuideSystemCacheFail RespCode = 13213
	//PassengerRecommendCallFail :请求导流穿透服务失败
	PassengerRecommendCallFail RespCode = 13220
	//PassengerActivityMisEstimateReq :预估请求乘客运营错误,返回数据格式错误
	PassengerActivityMisEstimateReq RespCode = 1070

	// 订单错误码

	//OrderCancelSucc :订单取消成功
	OrderCancelSucc RespCode = 3001
	//OrderCancelFail :订单取消失败
	OrderCancelFail RespCode = 3002
	//OrderCanceled :订单已取消
	OrderCanceled RespCode = 3003
	//OrderDriverGrabbed :订单被司机抢单
	OrderDriverGrabbed RespCode = 3004
	//OrderExpired :订单过期
	OrderExpired RespCode = 3005
	//OrderNoPlusTip :订单未加小费
	OrderNoPlusTip RespCode = 3006
	//OrderTransferError :订单转换错误
	OrderTransferError RespCode = 3007
	//OrderBilllIDError :账单id错误
	OrderBilllIDError RespCode = 3008
	//OrderStatusError :订单状态错误
	OrderStatusError RespCode = 3009
	//OrderTimeout :订单超时
	OrderTimeout RespCode = 3010
	//OrderAntispamFail :反作弊失败
	OrderAntispamFail RespCode = 3011
	//OrderWithdawError :提现错误
	OrderWithdawError RespCode = 3012
	//OrderDriverGrabbedWEBAPP :司机通过webapp抢单
	OrderDriverGrabbedWEBAPP RespCode = 3013
	//OrderHideAddress :隐藏地址
	OrderHideAddress RespCode = 3014
	//OrderStrategyTokenError :策略id错误
	OrderStrategyTokenError RespCode = 3015
	//OrderComboTypeError :套餐类型错误
	OrderComboTypeError RespCode = 3016
	//OrderAirportError :机场信息匹配错误
	OrderAirportError RespCode = 3017
	//OrderAirportInfoError :机场信息拉取失败
	OrderAirportInfoError RespCode = 3018
	//OrderComboInfoError :套餐信息错误
	OrderComboInfoError RespCode = 3019
	//OrderLBSAddError :添加到LBS错误
	OrderLBSAddError RespCode = 3020
	//OrderLBSDelError :删除LBS错误
	OrderLBSDelError RespCode = 3021
	//OrderTrafDepTimeError :航班出发日期错误
	OrderTrafDepTimeError RespCode = 3022
	//OrderPOIServiceError :检查坐标是否经过机场服务异常
	OrderPOIServiceError RespCode = 3030
	//OrderNoPaidRecode :openapi关闭订单，没找到orderispay记录
	OrderNoPaidRecode RespCode = 3031
	//OrderFinishError :结束计费出错
	OrderFinishError RespCode = 3032
	//OrderCompensatedriverFail :乘客取消平台补偿司机转账失败
	OrderCompensatedriverFail RespCode = 3033
	//OrderDynamicPriceEmpty :动态调价信息为空
	OrderDynamicPriceEmpty RespCode = 3034
	//OrderOrderSystemNetError :请求订单系统网络异常
	OrderOrderSystemNetError RespCode = 3035
	//OrderOrderInfoGetFail :从订单系统获取订单信息失败
	OrderOrderInfoGetFail RespCode = 3036
	//OrderOrderExtraInfoGetFail :从订单系统获取订单信息失败
	OrderOrderExtraInfoGetFail RespCode = 3037
	//OrderLineupQueryError :查询排队订单
	OrderLineupQueryError RespCode = 505121
	//OrderStatusCheckFail :检查订单状态失败
	OrderStatusCheckFail RespCode = 505122
	//OrderLineupUpdateFail :更新排队信息失败
	OrderLineupUpdateFail RespCode = 505123

	// API公用

	//APISigEmpty :api签名为空
	APISigEmpty RespCode = 8001
	//APISigError :api签名有误
	APISigError RespCode = 8002
	//APIReqError :api请求有误
	APIReqError RespCode = 8003
	//APIPhoneTimeError :api请求有误
	APIPhoneTimeError RespCode = 8004
	//APISigTimeout :api签名超时
	APISigTimeout RespCode = 8005

	//其它错误

	//CreditReqFail :查询信用失败
	CreditReqFail RespCode = 9714
	//EmptyFastCoupon :获取智能派券失败
	EmptyFastCoupon RespCode = 21530
)
