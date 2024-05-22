package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/1Panel-dev/1Panel/backend/app/dto/request"
	"github.com/1Panel-dev/1Panel/backend/app/dto/response"
	"github.com/1Panel-dev/1Panel/backend/app/model"
	"github.com/1Panel-dev/1Panel/backend/global"
	"github.com/1Panel-dev/1Panel/backend/utils/copier"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type MpHostService struct {
}

type IMpHostService interface {
	CreateMpHost(mpHost *model.MpHost) (err error)
	DeleteMpHostByIds(IDs []uint) (err error)
	UpdateMpHost(mpHost model.MpHost) (err error)
	GetMpHost(ID string) (mpHost model.MpHost, err error)
	GetMpHostInfoList(info request.MpHostSearch) (list []response.MpHost, total int64, err error)
	SetDefaultHost(ID string) (err error)
	LoginHost(ID string) (mpHost *model.MpHost, err error)
	LoginDefaultHost() (*model.MpHost, error)
}

func NewIMpHostService() IMpHostService {
	return &MpHostService{}
}

// CreateMpHost 创建mpHost表记录
func (mpHostService *MpHostService) CreateMpHost(mpHost *model.MpHost) (err error) {
	err = global.DB.Create(mpHost).Error
	return err
}

// DeleteMpHostByIds 批量删除mpHost表记录
func (mpHostService *MpHostService) DeleteMpHostByIds(IDs []uint) (err error) {
	err = global.DB.Delete(&[]model.MpHost{}, "id in ?", IDs).Error
	return err
}

// UpdateMpHost 更新mpHost表记录
func (mpHostService *MpHostService) UpdateMpHost(mpHost model.MpHost) (err error) {
	err = global.DB.Model(&model.MpHost{}).Where("id = ?", mpHost.ID).Updates(&mpHost).Error
	return err
}

// GetMpHost 根据ID获取mpHost表记录
func (mpHostService *MpHostService) GetMpHost(ID string) (mpHost model.MpHost, err error) {
	err = global.DB.Where("id = ?", ID).First(&mpHost).Error
	return
}

// GetMpHostInfoList 分页获取mpHost表记录
func (mpHostService *MpHostService) GetMpHostInfoList(info request.MpHostSearch) (list []response.MpHost, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.MpHost{})
	var mpHosts []model.MpHost
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.Name) > 0 {
		db = db.Where("name like ?", info.Name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&mpHosts).Error

	var resMpHosts []response.MpHost
	_ = copier.Copy(&resMpHosts, &mpHosts)

	setting, err := settingRepo.Get(settingRepo.WithByKey("defaultHostId"))
	if err != nil {
		return resMpHosts, total, nil
	}
	if len(setting.Value) > 0 {
		for i, resMpHost := range resMpHosts {
			fmt.Printf(strconv.Itoa(int(resMpHost.ID)))
			fmt.Printf(setting.Value)
			if strings.Compare(strconv.Itoa(int(resMpHost.ID)), setting.Value) == 0 {
				resMpHosts[i].IsDefault = "1"
			}
		}
	}

	return resMpHosts, total, err
}

func (mpHostService *MpHostService) SetDefaultHost(ID string) (err error) {
	return settingRepo.Update("defaultHostId", ID)
}

func (mpHostService *MpHostService) LoginDefaultHost() (*model.MpHost, error) {
	setting, err := settingRepo.Get(settingRepo.WithByKey("defaultHostId"))
	if err != nil {
		return nil, err
	}
	if len(setting.Value) > 0 {
		return mpHostService.LoginHost(setting.Value)
	}
	return nil, nil
}

func (mpHostService *MpHostService) LoginHost(ID string) (mpHost *model.MpHost, err error) {
	err = global.DB.Where("id = ?", ID).First(&mpHost).Error
	url := mpHost.Host + "/api/v1/auth/login"
	jsonData, _ := json.Marshal(
		map[string]interface{}{
			"name":          mpHost.User,
			"password":      mpHost.Pwd,
			"ignoreCaptcha": true,
			"captcha":       "",
			"captchaID":     "",
			"authMethod":    "jwt",
			"language":      "zh",
		},
	)

	client := http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", "127.0.0.1")
	req.Host = "127.0.0.1"

	// resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		token := result["data"].(map[string]interface{})["token"]

		err = global.CACHE.Set("mp-token:"+strconv.Itoa(int(mpHost.ID)), token)
		err = global.CACHE.Set("mp-host:"+strconv.Itoa(int(mpHost.ID)), mpHost.Host)
		if err != nil {
			return
		}
		return mpHost, err
	} else {
		err = errors.New(resp.Status)
		return
	}
}
