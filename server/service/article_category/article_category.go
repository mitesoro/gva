package article_category

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/article_category"
	article_categoryReq "github.com/flipped-aurora/gin-vue-admin/server/model/article_category/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ArticleCategoryService struct {
}

// CreateArticleCategory 创建文章分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (acService *ArticleCategoryService) CreateArticleCategory(ac *article_category.ArticleCategory) (err error) {
	err = global.GVA_DB.Create(ac).Error
	return err
}

// DeleteArticleCategory 删除文章分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (acService *ArticleCategoryService) DeleteArticleCategory(ac article_category.ArticleCategory) (err error) {
	err = global.GVA_DB.Delete(&ac).Error
	return err
}

// DeleteArticleCategoryByIds 批量删除文章分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (acService *ArticleCategoryService) DeleteArticleCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]article_category.ArticleCategory{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateArticleCategory 更新文章分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (acService *ArticleCategoryService) UpdateArticleCategory(ac article_category.ArticleCategory) (err error) {
	err = global.GVA_DB.Save(&ac).Error
	return err
}

// GetArticleCategory 根据id获取文章分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (acService *ArticleCategoryService) GetArticleCategory(id uint) (ac article_category.ArticleCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ac).Error
	return
}

// GetArticleCategoryInfoList 分页获取文章分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (acService *ArticleCategoryService) GetArticleCategoryInfoList(info article_categoryReq.ArticleCategorySearch) (list []article_category.ArticleCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&article_category.ArticleCategory{})
	var acs []article_category.ArticleCategory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id DESC").Find(&acs).Error
	return acs, total, err
}
