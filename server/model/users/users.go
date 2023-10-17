// 自动生成模板Users
package users

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 用户 结构体  Users
type Users struct {
      global.GVA_MODEL
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:129;"`  //手机号 
      Password  string `json:"password" form:"password" gorm:"column:password;comment:密码;"`  //密码 
      Nickname  string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:昵称;"`  //昵称 
      Avatar  string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;"`  //头像 
}


// TableName 用户 Users自定义表名 users
func (Users) TableName() string {
  return "users"
}

