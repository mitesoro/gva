package users

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	usersReq "github.com/flipped-aurora/gin-vue-admin/server/model/users/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UsersApi struct {
}

var uService = service.ServiceGroupApp.UsersServiceGroup.UsersService

// CreateUsers 创建用户
// @Tags Users
// @Summary 创建用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body users.Users true "创建用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /u/createUsers [post]
func (uApi *UsersApi) CreateUsers(c *gin.Context) {
	var u users.Users
	err := c.ShouldBindJSON(&u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Phone":    {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
		"Nickname": {utils.NotEmpty()},
		// "Avatar":   {utils.NotEmpty()},
	}
	if err := utils.Verify(u, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := uService.CreateUsers(&u); err != nil {
		global.GVA_LOG.Error("创建失败!"+err.Error(), zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUsers 删除用户
// @Tags Users
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body users.Users true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /u/deleteUsers [delete]
func (uApi *UsersApi) DeleteUsers(c *gin.Context) {
	var u users.Users
	err := c.ShouldBindJSON(&u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := uService.DeleteUsers(u); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUsersByIds 批量删除用户
// @Tags Users
// @Summary 批量删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /u/deleteUsersByIds [delete]
func (uApi *UsersApi) DeleteUsersByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := uService.DeleteUsersByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUsers 更新用户
// @Tags Users
// @Summary 更新用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body users.Users true "更新用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /u/updateUsers [put]
func (uApi *UsersApi) UpdateUsers(c *gin.Context) {
	var u users.Users
	err := c.ShouldBindJSON(&u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Phone":    {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
		"Nickname": {utils.NotEmpty()},
		// "Avatar":{utils.NotEmpty()},
	}
	if err := utils.Verify(u, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := uService.UpdateUsers(u); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUsers 用id查询用户
// @Tags Users
// @Summary 用id查询用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query users.Users true "用id查询用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /u/findUsers [get]
func (uApi *UsersApi) FindUsers(c *gin.Context) {
	var u users.Users
	err := c.ShouldBindQuery(&u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reu, err := uService.GetUsers(u.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reu": reu}, c)
	}
}

// GetUsersList 分页获取用户列表
// @Tags Users
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query usersReq.UsersSearch true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /u/getUsersList [get]
func (uApi *UsersApi) GetUsersList(c *gin.Context) {
	var pageInfo usersReq.UsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := uService.GetUsersInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
