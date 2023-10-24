package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/article"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    articleReq "github.com/flipped-aurora/gin-vue-admin/server/model/article/request"
)

type ArticleService struct {
}

// CreateArticle 创建文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (aService *ArticleService) CreateArticle(a *article.Article) (err error) {
	err = global.GVA_DB.Create(a).Error
	return err
}

// DeleteArticle 删除文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (aService *ArticleService)DeleteArticle(a article.Article) (err error) {
	err = global.GVA_DB.Delete(&a).Error
	return err
}

// DeleteArticleByIds 批量删除文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (aService *ArticleService)DeleteArticleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]article.Article{},"id in ?",ids.Ids).Error
	return err
}

// UpdateArticle 更新文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (aService *ArticleService)UpdateArticle(a article.Article) (err error) {
	err = global.GVA_DB.Save(&a).Error
	return err
}

// GetArticle 根据id获取文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (aService *ArticleService)GetArticle(id uint) (a article.Article, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&a).Error
	return
}

// GetArticleInfoList 分页获取文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (aService *ArticleService)GetArticleInfoList(info articleReq.ArticleSearch) (list []article.Article, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&article.Article{})
    var as []article.Article
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Title != "" {
        db = db.Where("title LIKE ?","%"+ info.Title+"%")
    }
    if info.Author != "" {
        db = db.Where("author LIKE ?","%"+ info.Author+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&as).Error
	return  as, total, err
}
