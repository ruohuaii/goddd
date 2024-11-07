package application

import (
	"context"
	"time"

	"github.com/ruohuaii/goddd/domain/merchant"
	"github.com/ruohuaii/goddd/infrastructure/faults"
	"github.com/ruohuaii/goddd/utils"
)

//应用层：用例处理和服务管理
//应用层包含应用服务，处理具体用例逻辑。应用服务利用领域层接口，协调多个实体的操作和领域逻辑。

type MerchantService struct {
	jwtSvc *JwtService
	repo   merchant.Repository
}

func NewMerchantService(repo merchant.Repository, jwtSvc *JwtService) *MerchantService {
	return &MerchantService{repo: repo, jwtSvc: jwtSvc}
}

func (s *MerchantService) Signup(ctx context.Context, data *merchant.Merchant) (*merchant.Merchant, error) {
	//生成merchant_id
	merchantID, err := utils.RandomNumber()
	if err != nil {
		return nil, err
	}
	data.MerchantID = merchantID
	//密码加密
	salt, err := utils.RandomString(32)
	if err != nil {
		return nil, err
	}
	password := data.Password
	cipherPassword, err := utils.HmacSha256(password, salt)
	if err != nil {
		return nil, err
	}
	data.Salt = salt
	data.Password = cipherPassword
	data.CreatedAt = time.Now()
	return s.repo.Save(ctx, data)
}

func (s *MerchantService) Signin(ctx context.Context, name string, password string) (*JwtPayload, error) {
	data, err := s.repo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	cmp1, err := utils.HmacSha256(password, data.Salt)
	if err != nil {
		return nil, err
	}
	cmp2 := data.Password
	if cmp1 != cmp2 {
		return nil, faults.MerchantPasswordError
	}
	//颁发令牌
	accessToken, err := s.jwtSvc.Generate(data.MerchantID)
	if err != nil {
		return nil, err
	}
	return &JwtPayload{TokenType: "Bearer", AccessToken: accessToken, ExpiresIn: 24 * 60 * 60}, nil
}
