package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/models"
)

const secret = "suxianjin"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUSerNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

func QueryUserByUserName(username string) error {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	err := db.Get(&count, sqlStr, username)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

//密码加密,将一条用户信息插入数据库
func InsertUser(u *models.User) {
	password := encryptPassword(u.Password)
	sqlStr := `insert into user(user_id,username,password) values (?,?,?)`
	db.Exec(sqlStr, u.UserID, u.Username, password)
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func Login(u *models.User) (err error) {
	oPassword := u.Password

	sqlStr := `select user_id ,username ,password from user where username = ?`
	err = db.Get(u, sqlStr, u.Username)
	if err == sql.ErrNoRows {
		return ErrorUSerNotExist
	}
	if err != nil {
		return err
	}
	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != u.Password {
		return ErrorInvalidPassword
	}
	return nil
}

func GetUserByID(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id ,username from user where user_id = ?`
	db.Get(user, sqlStr, uid)
	return user, nil
}
