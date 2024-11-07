package faults

const Unknown = "UNKNOWN"

var (
	MerchantExists         = &Faults{Code: "MERCHANT_EXISTS", Message: "商户信息已存在"}
	MerchantNotFound       = &Faults{Code: "MERCHANT_NOT_FOUND", Message: "商户信息不存在"}
	MerchantCreationFailed = &Faults{Code: "MERCHANT_CREATION_FAILED", Message: "创建商户失败"}
	MerchantInfoError      = &Faults{Code: "MERCHANT_INFO_ERROR", Message: "商户信息错误"}
	MerchantPasswordError  = &Faults{Code: "MERCHANT_PASSWORD_ERROR", Message: "商户密码错误"}
)
