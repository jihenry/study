package ent

type Account struct {
	ID         int64  //用户ID
	Nick       string //昵称
	User       string //用户名
	Pwd        string //密码
	Level      int32  //超v等级
	Permission int64  //权限
}
