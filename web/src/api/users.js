import service from '@/utils/request'

// @Tags Users
// @Summary 创建用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "创建用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /u/createUsers [post]
export const createUsers = (data) => {
  return service({
    url: '/u/createUsers',
    method: 'post',
    data
  })
}

// @Tags Users
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /u/deleteUsers [delete]
export const deleteUsers = (data) => {
  return service({
    url: '/u/deleteUsers',
    method: 'delete',
    data
  })
}

// @Tags Users
// @Summary 批量删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /u/deleteUsers [delete]
export const deleteUsersByIds = (data) => {
  return service({
    url: '/u/deleteUsersByIds',
    method: 'delete',
    data
  })
}

// @Tags Users
// @Summary 更新用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "更新用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /u/updateUsers [put]
export const updateUsers = (data) => {
  return service({
    url: '/u/updateUsers',
    method: 'put',
    data
  })
}

// @Tags Users
// @Summary 用id查询用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Users true "用id查询用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /u/findUsers [get]
export const findUsers = (params) => {
  return service({
    url: '/u/findUsers',
    method: 'get',
    params
  })
}

// @Tags Users
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /u/getUsersList [get]
export const getUsersList = (params) => {
  return service({
    url: '/u/getUsersList',
    method: 'get',
    params
  })
}
