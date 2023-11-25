package notice

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/notice"
	noticeReq "github.com/flipped-aurora/gin-vue-admin/server/model/notice/request"
)

type NoticeService struct {
}

// CreateNotice 创建公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (noService *NoticeService) CreateNotice(no *notice.Notice) (err error) {
	err = global.GVA_DB.Create(no).Error
	return err
}

// DeleteNotice 删除公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (noService *NoticeService) DeleteNotice(no notice.Notice) (err error) {
	err = global.GVA_DB.Delete(&no).Error
	return err
}

// DeleteNoticeByIds 批量删除公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (noService *NoticeService) DeleteNoticeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]notice.Notice{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateNotice 更新公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (noService *NoticeService) UpdateNotice(no notice.Notice) (err error) {
	err = global.GVA_DB.Save(&no).Error
	return err
}

// GetNotice 根据id获取公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (noService *NoticeService) GetNotice(id uint) (no notice.Notice, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&no).Error
	return
}

// GetNoticeInfoList 分页获取公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (noService *NoticeService) GetNoticeInfoList(info noticeReq.NoticeSearch) (list []notice.Notice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&notice.Notice{})
	var nos []notice.Notice
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id DESC").Find(&nos).Error
	return nos, total, err
}
