package service

import (
	"context"
	"kratos-tiway/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "kratos-tiway/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func toUserInfoResponse(user *biz.User) pb.UserInfoResponse{
	return pb.UserInfoResponse{
		Id:user.ID,
		NickName:user.Nickname,
		Name:user.Name,
		Email:user.Email,
		Mobile:user.Mobile,
		Password:user.Password,
		Status:user.Status,
		CreatedAt:user.CreatedAt.Unix(),
	}
}

func (s *UserService) GetUserList(ctx context.Context, req *pb.SearchUser) (*pb.UserListResponse, error) {
	listUser, total, _ := s.uc.List(ctx, &biz.UserListReq{
		Page:       int(req.Page),
		Limit:      int(req.Limit),
		Id:         req.Id,
		Name:       req.Name,
		Nickname:   req.Nickname,
		Email:      req.Email,
		Mobile:     req.Mobile,
		Status:     req.Status,
		Created_at: req.CreatedAt,
	})

	list := make([]*pb.UserInfoResponse, 0)
	for _, item := range listUser {
		userInfoResponse := toUserInfoResponse(item)
		list = append(list, &userInfoResponse)
	}
	return &pb.UserListResponse{Total: int32(total), Data: list}, nil
}
func (s *UserService) GetUserByMobile(ctx context.Context, req *pb.MobileRequest) (*pb.UserInfoResponse, error) {
	user,err := s.uc.UserByMobile(ctx,req.Mobile)
	if user == nil {
		return nil,err
	}
	userInfoResponse := toUserInfoResponse(user)
	return &userInfoResponse, nil
}

//get by nickname
func (s *UserService) GetUserByNickname(ctx context.Context, req *pb.NicknameRequest) (*pb.UserInfoResponse, error) {
	user,err := s.uc.UserByNickname(ctx,req.Nickname)
	if user == nil {
		return nil,err
	}
	userInfoResponse := toUserInfoResponse(user)
	return &userInfoResponse, nil
}

//get by user id
func (s *UserService) GetUserById(ctx context.Context, req *pb.IdRequest) (*pb.UserInfoResponse, error) {
	user,err := s.uc.UserById(ctx,req.Id)
	if user == nil {
		return nil,err
	}
	userInfoResponse := toUserInfoResponse(user)
	return &userInfoResponse, nil
}

//创建用户
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserInfo) (*pb.UserInfoResponse, error) {
	user, err := s.uc.Create(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}

	userInfoResponse := toUserInfoResponse(user)
	return &userInfoResponse, nil
}

//更新用户信息
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserInfo) (*pb.UpdateUserResponse, error) {
	_, err := s.uc.UpdateUser(ctx, &biz.User{
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Email:    req.Email,
		Name:     req.Name,
	})
	success := true
	message := "更新成功"
	if err != nil {
		success = false
		message = "更新失败"
	}

	return &pb.UpdateUserResponse{Success: success, Message: message}, nil
}

func (s *UserService) CheckPassword(ctx context.Context, req *pb.PasswordCheckInfo) (*pb.CheckResponse, error) {
	isCheck,err := s.uc.CheckPassword(ctx,req.Mobile,req.Password)

	success := false
	if isCheck {
		success = true
	}
	return &pb.CheckResponse{
		Success:success,
	}, err
}
