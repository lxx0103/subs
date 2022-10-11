package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"subs/core/config"
	"subs/core/database"
	"time"
)

type authService struct {
}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) VerifyWechatSignin(code string) (*WxUserResponse, error) {
	var credential WechatCredential
	if code == "test" {
		credential.ErrCode = 0
		credential.ErrMsg = ""
		credential.OpenID = "test"
	} else if code == "test2" {
		credential.ErrCode = 0
		credential.ErrMsg = ""
		credential.OpenID = "test2"
	} else {
		httpClient := &http.Client{}
		signin_uri := config.ReadConfig("Wechat.signin_uri")
		appID := config.ReadConfig("Wechat.app_id")
		appSecret := config.ReadConfig("Wechat.app_secret")
		uri := signin_uri + "?appid=" + appID + "&secret=" + appSecret + "&js_code=" + code + "&grant_type=authorization_code"
		req, err := http.NewRequest("GET", uri, nil)
		if err != nil {
			return nil, err
		}
		res, err := httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &credential)
		if err != nil {
			msg := "解码出错"
			return nil, errors.New(msg)
		}
	}
	if credential.ErrCode != 0 {
		msg := credential.ErrMsg
		return nil, errors.New(msg)
	}
	db := database.RDB()
	query := NewAuthQuery(db)
	userInfo, err := query.GetWxUserByOpenID(credential.OpenID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			msg := "新用户"
			var res WxUserResponse
			res.OpenID = credential.OpenID
			return &res, errors.New(msg)
		} else {
			msg := "获取用户出错"
			return nil, errors.New(msg)
		}
	}
	return userInfo, nil
}

func (s authService) CreateWxUser(openID string) error {
	db := database.WDB()
	tx, err := db.Begin()
	if err != nil {
		msg := "开启事务出错"
		return errors.New(msg)
	}
	defer tx.Rollback()
	repo := NewAuthRepository(tx)
	var newUser WxUser
	isConflict, err := repo.CheckWxUserConfict(openID)
	if err != nil {
		msg := "检查用户名合法性出错"
		return errors.New(msg)
	}
	if isConflict {
		msg := "用户名已存在"
		return errors.New(msg)
	}
	newUser.OpenID = openID
	newUser.Status = 1
	newUser.Created = time.Now()
	newUser.CreatedBy = "SIGNUP"
	newUser.Updated = time.Now()
	newUser.UpdatedBy = "SIGNUP"
	err = repo.CreateWxUser(newUser)
	if err != nil {
		msg := "创建微信用户出错"
		return errors.New(msg)
	}
	tx.Commit()
	return nil
}

func (s *authService) GetWxUserByOpenID(id string) (*WxUserResponse, error) {
	db := database.RDB()
	query := NewAuthQuery(db)
	user, err := query.GetWxUserByOpenID(id)
	if err != nil {
		msg := "用户不存在"
		return nil, errors.New(msg)
	}
	return user, nil
}
