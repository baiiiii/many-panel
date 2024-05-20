package v1

import (
	"github.com/1Panel-dev/1Panel/backend/app/api/v1/helper"
	"github.com/1Panel-dev/1Panel/backend/app/dto"
	"github.com/1Panel-dev/1Panel/backend/app/dto/request"
	"github.com/1Panel-dev/1Panel/backend/app/model"
	"github.com/1Panel-dev/1Panel/backend/constant"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CreateMpHost 创建mpHost表
// @Tags MpHost
// @Summary 创建mpHost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mp.MpHost true "创建mpHost表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mpHost/createMpHost [post]
func (b *BaseApi) CreateMpHost(c *gin.Context) {
	var req request.CreateMpHost
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	mpHost := &model.MpHost{Name: req.Name, Host: req.Host, User: req.User, Pwd: req.Pwd}
	if err := mpHostService.CreateMpHost(mpHost); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
	} else {
		helper.SuccessWithData(c, nil)
	}
}

// DeleteMpHostByIds 批量删除mpHost表
// @Tags MpHost
// @Summary 批量删除mpHost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /mpHost/deleteMpHostByIds [delete]
func (b *BaseApi) DeleteMpHostByIds(c *gin.Context) {
	var req request.DeleteMpHost
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}
	if err := mpHostService.DeleteMpHostByIds(req.Ids); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
	} else {
		helper.SuccessWithData(c, nil)
	}
}

// UpdateMpHost 更新mpHost表
// @Tags MpHost
// @Summary 更新mpHost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mp.MpHost true "更新mpHost表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mpHost/updateMpHost [put]
func (b *BaseApi) UpdateMpHost(c *gin.Context) {
	var req request.UpdateMpHost
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	mpHost := model.MpHost{BaseModel: model.BaseModel{ID: req.ID}, Name: req.Name, Host: req.Host, User: req.User, Pwd: req.Pwd}

	if err := mpHostService.UpdateMpHost(mpHost); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
	} else {
		helper.SuccessWithData(c, nil)
	}
}

// FindMpHost 用id查询mpHost表
// @Tags MpHost
// @Summary 用id查询mpHost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mp.MpHost true "用id查询mpHost表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mpHost/findMpHost [get]
func (b *BaseApi) FindMpHost(c *gin.Context) {
	ID := c.Query("ID")
	if rempHost, err := mpHostService.GetMpHost(ID); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
	} else {
		helper.SuccessWithData(c, rempHost)
	}
}

// GetMpHostList 分页获取mpHost表列表
// @Tags MpHost
// @Summary 分页获取mpHost表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mpReq.MpHostSearch true "分页获取mpHost表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mpHost/getMpHostList [get]
func (b *BaseApi) GetMpHostList(c *gin.Context) {
	var pageInfo request.MpHostSearch
	if err := helper.CheckBindAndValidate(&pageInfo, c); err != nil {
		return
	}
	if list, total, err := mpHostService.GetMpHostInfoList(pageInfo); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
	} else {
		helper.SuccessWithData(c, dto.PageResult{
			Items: list,
			Total: total,
		})
	}
}

func (b *BaseApi) SetDefaultHost(c *gin.Context) {
	id, err := helper.GetParamID(c)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInternalServer, nil)
		return
	}
	err = mpHostService.SetDefaultHost(strconv.Itoa(int(id)))
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInternalServer, nil)
		return
	}
	helper.SuccessWithData(c, nil)
}

func (b *BaseApi) LoginMpHost(c *gin.Context) {
	id, err := helper.GetParamID(c)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInternalServer, nil)
		return
	}
	mpHost, err := mpHostService.LoginHost(strconv.Itoa(int(id)))
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInternalServer, nil)
		return
	}
	helper.SuccessWithData(c, mpHost)
}
