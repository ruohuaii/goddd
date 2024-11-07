package merchant

type (
	SignupReq struct {
		Name     string `json:"name" validate:"required,min=3,max=32"`
		Email    string `json:"email" validate:"required,email"`
		Mobile   string `json:"mobile" validate:"required,min=7,max=11"`
		Password string `json:"password" validate:"required,min=8,max=32"`
	}
	SignupResp struct {
		Name       string `json:"name"`
		MerchantID string `json:"merchant_id"`
	}
)

type (
	SigninReq struct {
		Name     string `json:"name" validate:"required,min=3,max=32"`
		Password string `json:"password" validate:"required,min=8,max=32"`
	}
	SigninResp struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
)
