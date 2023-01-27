package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"learngoframe/internal/model/entity"
)

type UserProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"UserService" summary:"获取当前用户的个人资料"`
}

type UserProfileRes struct {
	*entity.User
}

type UserSignUpReq struct {
	g.Meta    `path:"/user/sign-up" method:"post" tags:"UserService" summary:"注册一个新的用户帐户"`
	Passport  string `v:"required|length:6,16"`
	Password  string `v:"required|length:6,16"`
	Password2 string `v:"required|length:6,16|same:Password"`
	Nickname  string
}

type UserSignUpRes struct {
}

type UserSignInReq struct {
	g.Meta   `path:"/user/sign-in" method:"post" tags:"UserService" summary:"使用现有帐户登录"`
	Passport string `v:"required"`
	Password string `v:"required"`
}

type UserSignInRes struct {
}

type UserCheckPassportReq struct {
	g.Meta   `path:"/user/check-passport" method:"post" tags:"UserService" summary:"检查护照可用"`
	Passport string `v:"required"`
}

type UserCheckPassportRes struct {
}

type UserCheckNickNameReq struct {
	g.Meta   `path:"/user/check-nickname" method:"post" tags:"UserService" summary:"检查可用的昵称"`
	Nickname string `v:"required"`
}

type UserCheckNickNameRes struct {
}

type UserIsSignedInReq struct {
	g.Meta `path:"/user/is-signed-in" method:"post" tags:"UserService" summary:"检查当前用户是否已经登录"`
}

type UserIsSignedInRes struct {
	OK bool `dc:"True if current user is signed in; or else false"`
}

type UserSignOutReq struct {
	g.Meta `path:"/user/sign-out" method:"post" tags:"UserService" summary:"注销当前用户"`
}

type UserSignOutRes struct {
}
