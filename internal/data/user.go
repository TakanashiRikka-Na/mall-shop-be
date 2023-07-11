package data

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"mall-shop-be/internal/biz"
	errs "mall-shop-be/internal/errors"
	"mall-shop-be/internal/utils"
	"regexp"
	"strings"
	"time"
)

type User struct {
	ID           int32
	UserName     string
	HashedPasswd string
}
type userRepo struct {
	data *Data
	log  *log.Helper
	erc  errs.ErrorController
}

func unknownErr(err error, msg string, repo *userRepo) *errors.Error {
	return repo.erc.UnKnownError().WithMetadata(map[string]string{msg: err.Error()})
}

func NewUserRepo(data *Data, logger log.Logger, controller errs.ErrorController) biz.UserRepo {
	return &userRepo{data: data, log: log.NewHelper(logger), erc: controller}
}

func (u *userRepo) Register(user biz.User) error {
	//只需要到数据库中查找 因为 只读缓存在数据库中是最新的数据
	existed, err := CheckUserExist(u, user.UserName)
	if err != nil {
		return err
	}
	if existed {
		return u.erc.UserExist()
	}
	//检查密码是否符合规范
	if !isValidPassword(user.Password) {
		return u.erc.PasswdValid()
	}
	//进入注册逻辑
	hashedPassword, err := generateHashedPassword(user.Password)
	if err != nil {
		u.log.Error(err)
		return u.erc.UnKnownError()
	}
	var userInfo = User{
		UserName:     user.UserName,
		HashedPasswd: hashedPassword,
	}
	tx := u.data.DB.Begin()
	if err = tx.Model(&User{}).Create(&userInfo).Error; err != nil {
		u.log.Error(err)
		tx.Rollback()
		return unknownErr(err, UserError, u)
	}
	tx.Commit()
	return nil
}

func (u *userRepo) Login(user biz.User) (string, error) {
	// 先到redis拿一下
	token, err := u.data.Rdb.Get(user.UserName).Result()
	if err == nil && len(token) != 0 {
		return token, nil
	}
	// 如果错误原因只是空值就不打印
	if err.Error() != "redis: nil" {
		u.log.Error("read redis error reason:", err)
	}

	//到数据库查 hashPassword
	var userInfo User
	if err = u.data.DB.Model(&User{}).Where("user_name=?", user.UserName).Find(&userInfo).Error; err != nil {
		u.log.Error(err)
		return "", unknownErr(err, UserError, u)
	}
	if !compareHashPasswd(user.Password, userInfo.HashedPasswd) {
		return "", u.erc.PasswdError()
	}
	token, err = utils.GenerateToken(user.UserName)
	if err != nil {
		return "", unknownErr(err, "generate token err", u)
	}
	if err = u.data.Rdb.Set(user.UserName, token, time.Hour).Err(); err != nil {
		u.log.Error(err)
	}
	return token, nil

}

func CheckUserExist(repo *userRepo, username string) (isExisted bool, err error) {
	err = repo.data.DB.Table("users").Select("1").Where("user_name=?", username).Find(&isExisted).Error
	return
}

func isValidPassword(password string) bool {
	// 密码长度至少为 6 位
	if len(password) < 6 {
		return false
	}

	// 密码包含数字
	hasDigit, _ := regexp.MatchString(`\d`, password)
	if !hasDigit {
		return false
	}

	// 密码包含字母
	hasLetter, _ := regexp.MatchString(`[a-zA-Z]`, password)
	if !hasLetter {
		return false
	}

	// 密码包含标点符号
	hasPunctuation := false
	for _, char := range password {
		if strings.ContainsAny(string(char), "!@#$%^&*()_+-=.") {
			hasPunctuation = true
			break
		}
	}

	return hasPunctuation
}

func generateHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func compareHashPasswd(passwd, hashPasswd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPasswd), []byte(passwd))
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}
