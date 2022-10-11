package auth

import (
	"subs/core/response"
	"subs/service"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// @Summary 小程序登录
// @Id 006
// @Tags 小程序管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param signin_info body WxSigninRequest true "登录类型"
// @Success 200 object response.SuccessRes{data=WxSigninResponse} 登录成功
// @Failure 400 object response.ErrorRes 内部错误
// @Failure 401 object response.ErrorRes 登录失败
// @Router /wx/signin [POST]
func WxSignin(c *gin.Context) {
	var signinInfo WxSigninRequest
	err := c.ShouldBindJSON(&signinInfo)
	if err != nil {
		response.ResponseError(c, "数据格式错误", err)
		return
	}
	authService := NewAuthService()
	userInfo, err := authService.VerifyWechatSignin(signinInfo.Code)
	if err != nil {
		if err.Error() == "新用户" {
			err = authService.CreateWxUser(userInfo.OpenID)
			if err != nil {
				response.ResponseUnauthorized(c, "创建小程序用户失败", err)
				return
			}
			userInfo, err = authService.GetWxUserByOpenID(userInfo.OpenID)
			if err != nil {
				response.ResponseUnauthorized(c, "用户不存在", err)
				return
			}
		} else {
			response.ResponseUnauthorized(c, "用户不存在", err)
			return
		}
	}

	claims := service.CustomClaims{
		OpenID:   userInfo.OpenID,
		UserName: userInfo.Name,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + 72000,
			Issuer:    "subs",
		},
	}
	jwtServices := service.JWTAuthService()
	generatedToken := jwtServices.GenerateToken(claims)
	var res WxSigninResponse
	res.Token = generatedToken
	res.User = *userInfo
	response.Response(c, res)
}
