package data

import (
	"context"
	"crypto/sha512"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-tiway/app/user/internal/biz"
	"time"
)

type User struct {
	ID        int64     `gorm:"primarykey"`
	Mobile    string    `gorm:"index:idx_mobile;unique;type:varchar(11) comment '手机号码，用户唯一标识';not null"`
	Name      string    `gorm:"type:varchar(32) comment '名字'"`
	Nickname  string    `gorm:"type:varchar(64) comment '昵称'"`
	Password  string    `gorm:"type:varchar(100);not null "` // 用户密码的保存需要注意是否加密
	Salt      string    `gorm:"type:varchar(100);not null "`
	Email     string    `gorm:"type:varchar(255);not null "`
	Avatar    string    `gorm:"column:avatar;type:varchar(255) comment '头像'"`
	Status    int32     `gorm:"column:status;default:1;type:int comment '1:普通用户，2:失效'"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

type PasswordEn struct {
	Salt         string
	EncryptedStr string
}

func (User) TableName() string {
	return "users"
}

func getEncryptOption() *password.Options {
	return &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUser .
func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	// 验证是否已经创建
	var user User
	result := r.data.db.Where(&User{Mobile: u.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, errors.New(500, "USER_EXIST", "用户已存在"+u.Mobile)
	}

	pswInfo := encrypt(u.Password)
	user.Mobile = u.Mobile
	user.Nickname = u.Nickname
	user.Password = pswInfo.EncryptedStr
	user.Salt = pswInfo.Salt
	res := r.data.db.Create(&user)
	if res.Error != nil {
		return nil, errors.New(500, "CREAT_USER_ERROR", "用户创建失败")
	}
	userInfoRes := modelToResponse(user)
	return &userInfoRes, nil
}

// Password encryption
func encrypt(psd string) PasswordEn {
	options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(psd, options)

	return PasswordEn{
		Salt:         salt,
		EncryptedStr: encodedPwd,
	}
}

// ModelToResponse 转换 user 表中所有字段的值
func modelToResponse(user User) biz.User {
	userInfoRsp := biz.User{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Name:      user.Name,
		Mobile:    user.Mobile,
		Password:  user.Password,
		Salt:      user.Salt,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
	}
	return userInfoRsp
}

// ListUser .
func (r *userRepo) ListUser(ctx context.Context, ureq *biz.UserListReq) ([]*biz.User, int, error) {
	var users []User

	//创建查询
	query := r.data.db
	if ureq.Id != 0 {
		query = query.Where("id = ?", ureq.Id)
	}
	if ureq.Name != "" {
		query = query.Where("name like ?", ureq.Name+"%")
	}
	if ureq.Nickname != "" {
		query = query.Where("nickname like ?", ureq.Nickname+"%")
	}
	if ureq.Email != "" {
		query = query.Where("email = ?", ureq.Email)
	}
	if ureq.Mobile != "" {
		query = query.Where("mobile = ?", ureq.Mobile)
	}
	total := int(query.Find(&users).RowsAffected)
	result := query.Limit(ureq.Limit).Offset((ureq.Page - 1) * ureq.Limit).Find(&users)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, 0, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	if result.Error != nil {
		return nil, 0, errors.New(500, "FIND_USER_ERROR", "find user error")
	}

	rv := make([]*biz.User, 0)
	for _, user := range users {
		rv = append(rv, &biz.User{
			ID:        user.ID,
			Mobile:    user.Mobile,
			Password:  user.Password,
			Email:     user.Email,
			Avatar:    user.Avatar,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
		})
	}
	return rv, total, nil
}

// UserByMobile .
func (r *userRepo) UserByMobile(ctx context.Context, mobile string) (*biz.User, error) {
	var user User
	result := r.data.db.Where(&User{Mobile: mobile}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	if result.Error != nil {
		return nil, errors.New(500, "FIND_USER_ERROR", "find user error")
	}

	if result.RowsAffected == 0 {
		return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	re := modelToResponse(user)
	return &re, nil
}

func (r *userRepo) UserByNickname(ctx context.Context, nickname string) (*biz.User, error) {
	var user User
	result := r.data.db.Where(&User{Nickname: nickname}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	if result.Error != nil {
		return nil, errors.New(500, "FIND_USER_ERROR", "find user error")
	}

	if result.RowsAffected == 0 {
		return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	re := modelToResponse(user)
	return &re, nil
}

// UpdateUser .
func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (bool, error) {
	var userInfo User
	result := r.data.db.Where(&User{ID: user.ID}).First(&userInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, errors.NotFound("USER_NOT_FOUND", "user not found")
	}

	if result.RowsAffected == 0 {
		return false, errors.NotFound("USER_NOT_FOUND", "用户不存在")
	}

	userInfo.Nickname = user.Nickname
	userInfo.Name = user.Name
	userInfo.Mobile = user.Mobile
	userInfo.Email = user.Email

	if err := r.data.db.Save(&userInfo).Error; err != nil {
		return false, errors.New(500, "USER_NOT_FOUND", err.Error())
	}

	return true, nil
}

// GetUserById .
func (r *userRepo) GetUserById(ctx context.Context, Id int64) (*biz.User, error) {
	var user User
	if err := r.data.db.Where(&User{ID: Id}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER_NOT_FOUND", err.Error())
	}
	re := modelToResponse(user)
	return &re, nil
}

// CheckPassword .
func (r *userRepo) CheckPassword(ctx context.Context, mobile string, psw string) (bool, error) {
	user, _ := r.UserByMobile(ctx, mobile)
	options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
	check := password.Verify(psw, user.Salt, user.Password, options)
	return check, nil
}
