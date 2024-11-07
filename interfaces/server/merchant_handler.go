package server

import (
	"net/http"
	"strconv"

	"github.com/ruohuaii/goddd/application"
	"github.com/ruohuaii/goddd/domain/merchant"
	"github.com/ruohuaii/goddd/infrastructure/faults"
	dto "github.com/ruohuaii/goddd/interfaces/dto/merchant"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MerchantHandler struct {
	svc *application.MerchantService
}

func NewMerchantHandler(svc *application.MerchantService) *MerchantHandler {
	return &MerchantHandler{svc: svc}
}

func (h *MerchantHandler) Signup(c *gin.Context) {
	var req dto.SignupReq
	var resp dto.SignupResp
	if err := c.ShouldBind(&req); err != nil {
		//请求参数错误
		c.JSON(http.StatusBadRequest, Response{Errors: []*faults.Faults{faults.As(err)}})
		return
	}
	if err := validator.New().Struct(&req); err != nil {
		//请求参数错误
		c.JSON(http.StatusBadRequest, Response{Errors: []*faults.Faults{faults.As(err)}})
		return
	}
	ctx := c.Request.Context()
	data, err := h.svc.Signup(ctx, &merchant.Merchant{
		Name:     req.Name,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		//数据持久化错误
		c.JSON(http.StatusInternalServerError, Response{Errors: []*faults.Faults{faults.As(err)}})
		return
	}
	resp = dto.SignupResp{
		Name:       data.Name,
		MerchantID: strconv.FormatUint(data.MerchantID, 10),
	}
	c.JSON(http.StatusOK, Response{Data: resp})
}

func (h *MerchantHandler) Signin(c *gin.Context) {
	var req dto.SigninReq
	var resp dto.SigninResp
	if err := c.ShouldBind(&req); err != nil {
		//请求参数错误
		c.JSON(http.StatusBadRequest, Response{Errors: []*faults.Faults{faults.As(err)}})
		return
	}
	if err := validator.New().Struct(&req); err != nil {
		//请求参数错误
		c.JSON(http.StatusBadRequest, Response{Errors: []*faults.Faults{faults.As(err)}})
		return
	}
	ctx := c.Request.Context()
	token, err := h.svc.Signin(ctx, req.Name, req.Password)
	if err != nil {
		//登录失败
		c.JSON(http.StatusUnauthorized, Response{Errors: []*faults.Faults{faults.As(err)}})
		return
	}
	resp = dto.SigninResp{
		TokenType:   token.TokenType,
		AccessToken: token.AccessToken,
		ExpiresIn:   token.ExpiresIn,
	}
	c.JSON(http.StatusOK, Response{Data: resp})
}
