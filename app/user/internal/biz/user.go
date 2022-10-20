package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Nickname  string
	Avatar    string
	Password  string
	Salt      string
	Email     string
	Mobile    string
	Status    int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserListReq struct {
	Page       int
	Limit      int
	Id         int64
	Name       string
	Nickname   string
	Email      string
	Mobile     string
	Status     int32
	Created_at int64
}

type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	ListUser(ctx context.Context, req *UserListReq) ([]*User, int, error)
	UserByMobile(ctx context.Context, mobile string) (*User, error)
	UserByNickname(ctx context.Context, nickname string) (*User, error)
	GetUserById(ctx context.Context, id int64) (*User, error)
	UpdateUser(context.Context, *User) (bool, error)
	CheckPassword(ctx context.Context, mobile string, password string) (bool, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUsecase) List(ctx context.Context, ureq *UserListReq) ([]*User, int, error) {
	return uc.repo.ListUser(ctx, ureq)
}

func (uc *UserUsecase) UserByMobile(ctx context.Context, mobile string) (*User, error) {
	return uc.repo.UserByMobile(ctx, mobile)
}

func (uc *UserUsecase) UserByNickname(ctx context.Context, nickname string) (*User, error) {
	return uc.repo.UserByNickname(ctx, nickname)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *User) (bool, error) {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUsecase) CheckPassword(ctx context.Context, mobile string, password string) (bool, error) {
	return uc.repo.CheckPassword(ctx, mobile, password)
}

func (uc *UserUsecase) UserById(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUserById(ctx, id)
}
