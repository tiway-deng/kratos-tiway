package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	v1 "kratos-tiway/api/user/v1"
)

var userClient v1.UserClient
var conn *grpc.ClientConn

func main() {
	Init()

	//TestCreateUser() // 创建用户
	TestUpdateUser()
	//TestUserList()
	//TestGetUserByMobile()
	//TestCheckPassword()
	//TestGetById()
	//TestGetByMobile()
	//TestGetByNickname()

	conn.Close()
}

// Init 初始化 grpc 链接 注意这里链接的 端口
func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	userClient = v1.NewUserClient(conn)
}

func TestCreateUser() {

	rsp, err := userClient.CreateUser(context.Background(), &v1.CreateUserInfo{
		Mobile:   "13566666666",
		Password: "123456",
		Nickname: "tiway",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rsp.Id)
}

func TestUpdateUser() {
	rsp, err := userClient.UpdateUser(context.Background(), &v1.UpdateUserInfo{
		Id:       1,
		Mobile:   "13536999999",
		Email:    "5011125@qq.com",
		Name:     "天天",
		Nickname: "tiwayD",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rsp.Message)
}

func TestUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &v1.SearchUser{
		//Id:2,
		//Mobile:"13536999999",
		//Email:"5011125@qq.com",
		Name:     "天",
		Nickname: "tian",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.Data)
}

func TestGetUserByMobile() {
	rsp, err := userClient.GetUserByMobile(context.Background(), &v1.MobileRequest{
		Mobile: "13536999999",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rsp.Password)
}

func TestCheckPassword() {
	rsp, err := userClient.CheckPassword(context.Background(), &v1.PasswordCheckInfo{
		Mobile:   "13566666666",
		Password: "123456",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rsp.Success)
	//options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
	//check := password.Verify("123456", "qDuy3h5nDjhFPJPg", "655fdf283ab63b5bb25074cb79d3d19bb1838d6f8a183bee096e7778872eb557", options)
	//
	//fmt.Println(check)
}

func TestGetById() {
	rsp, err := userClient.GetUserById(context.Background(), &v1.IdRequest{
		Id:2,
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rsp.NickName)
}

func TestGetByMobile() {
	rsp, err := userClient.GetUserByMobile(context.Background(), &v1.MobileRequest{
		Mobile:"13566666666",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rsp.NickName)
}

func TestGetByNickname() {
	rsp, err := userClient.GetUserByNickname(context.Background(), &v1.NicknameRequest{
		Nickname:"encrypty",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rsp.NickName)
}