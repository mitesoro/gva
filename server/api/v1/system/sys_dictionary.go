package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/article_category"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type DictionaryApi struct{}

// CreateSysDictionary
// @Tags      SysDictionary
// @Summary   创建SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionary           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{msg=string}  "创建SysDictionary"
// @Router    /sysDictionary/createSysDictionary [post]
func (s *DictionaryApi) CreateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.CreateSysDictionary(dictionary)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysDictionary
// @Tags      SysDictionary
// @Summary   删除SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionary           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{msg=string}  "删除SysDictionary"
// @Router    /sysDictionary/deleteSysDictionary [delete]
func (s *DictionaryApi) DeleteSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.DeleteSysDictionary(dictionary)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSysDictionary
// @Tags      SysDictionary
// @Summary   更新SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionary           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{msg=string}  "更新SysDictionary"
// @Router    /sysDictionary/updateSysDictionary [put]
func (s *DictionaryApi) UpdateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.UpdateSysDictionary(&dictionary)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysDictionary
// @Tags      SysDictionary
// @Summary   用id查询SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysDictionary                                       true  "ID或字典英名"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "用id查询SysDictionary"
// @Router    /sysDictionary/findSysDictionary [get]
func (s *DictionaryApi) FindSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindQuery(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if dictionary.Type == "user" { // 查询用户
		t := true
		sysDictionary := system.SysDictionary{
			Name:   dictionary.Type,
			Status: &t,
			Type:   dictionary.Type,
		}
		var us []users.Users
		if err = global.GVA_DB.Find(&us).Error; err == nil {
			var details []system.SysDictionaryDetail
			for _, user := range us {
				details = append(details, system.SysDictionaryDetail{
					Value:           int(user.ID),
					Label:           fmt.Sprintf("%s(%s)", user.Nickname, user.Phone),
					SysDictionaryID: 999,
				})
			}
			sysDictionary.SysDictionaryDetails = details
		}
		response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "查询成功", c)
		return
	}
	if strings.Contains(dictionary.Type, "#user_") { // 查询用户
		t := true
		sysDictionary := system.SysDictionary{
			Name:   dictionary.Type,
			Status: &t,
			Type:   dictionary.Type,
		}
		arr := strings.Split(dictionary.Type, "_")
		db := global.GVA_DB
		if len(arr) == 2 {
			db = db.Where("admin_id", arr[1])
		}
		var us []users.Users
		if err = db.Find(&us).Error; err == nil {
			var details []system.SysDictionaryDetail
			for _, user := range us {
				details = append(details, system.SysDictionaryDetail{
					Value:           int(user.ID),
					Label:           fmt.Sprintf("%s(%s)", user.Nickname, user.Phone),
					SysDictionaryID: 999,
				})
			}
			sysDictionary.SysDictionaryDetails = details
		}
		response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "查询成功", c)
		return
	}
	if dictionary.Type == "article_category" { // 文章分类
		t := true
		sysDictionary := system.SysDictionary{
			Name:   dictionary.Type,
			Status: &t,
			Type:   dictionary.Type,
		}
		var aa []article_category.ArticleCategory
		if err = global.GVA_DB.Find(&aa).Error; err == nil {
			var details []system.SysDictionaryDetail
			for _, a := range aa {
				details = append(details, system.SysDictionaryDetail{
					Value:           int(a.ID),
					Label:           a.Name,
					SysDictionaryID: 999,
				})
			}
			sysDictionary.SysDictionaryDetails = details
		}
		response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "查询成功", c)
		return
	}
	if dictionary.Type == "admin" { // 查询后台用户
		t := true
		sysDictionary := system.SysDictionary{
			Name:   dictionary.Type,
			Status: &t,
			Type:   dictionary.Type,
		}
		var us []system.SysUser
		if err = global.GVA_DB.Find(&us).Error; err == nil {
			var details []system.SysDictionaryDetail
			for _, user := range us {
				details = append(details, system.SysDictionaryDetail{
					Value:           int(user.ID),
					Label:           fmt.Sprintf("%s", user.NickName),
					SysDictionaryID: 999,
				})
			}
			sysDictionary.SysDictionaryDetails = details
		}
		response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "查询成功", c)
		return
	}

	sysDictionary, err := dictionaryService.GetSysDictionary(dictionary.Type, dictionary.ID, dictionary.Status)
	if err != nil {
		global.GVA_LOG.Error("字典未创建或未开启!", zap.Error(err))
		response.FailWithMessage("字典未创建或未开启", c)
		return
	}
	response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "查询成功", c)
}

// GetSysDictionaryList
// @Tags      SysDictionary
// @Summary   分页获取SysDictionary列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.SysDictionarySearch                             true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取SysDictionary列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysDictionary/getSysDictionaryList [get]
func (s *DictionaryApi) GetSysDictionaryList(c *gin.Context) {
	var pageInfo request.SysDictionarySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dictionaryService.GetSysDictionaryInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
