package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/JWT"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) error {
	//判断用户是否存在
	err := mysql.QueryUserByUserName(p.UserName)
	if err != nil {
		//查询出错
		return err
	}
	//生成UID
	Uid := snowflake.GetID()
	//保存进数据库
	u := models.User{
		UserID:   Uid,
		Password: p.Password,
		Username: p.UserName,
	}
	mysql.InsertUser(&u)
	return nil
}

func LogIn(p *models.ParamLogin) (user *models.User, rtoken string, err error) {
	var u = models.User{
		Username: p.UserName,
		Password: p.Password,
	}

	err = mysql.Login(&u)
	if err != nil {
		return nil, "", err
	}
	//生成JWT
	var token string
	token, JWT.FreshToken, err = JWT.GenToken(u.UserID, u.Username)
	if err != nil {
		return nil, "", err
	}
	u.Token = token
	return &u, JWT.FreshToken, nil
}

//
