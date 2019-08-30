package dao

import (
	"go-blog/dto/inputDto"
	"go-blog/entity"
)


func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []entity.Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&entity.Article{}).Where(maps).Count(&count)

	return
}

func GetArticle(id int) (article entity.Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&entity.Article{}).Where("id = ?", id).Update(data)

	return true
}

func AddArticle(articleCreateDto inputDto.ArticleInputDto) bool {
	db.Create(&entity.Article{
		TagID:   articleCreateDto.TagId,
		Title:   articleCreateDto.Title,
		Desc:    articleCreateDto.Desc,
		Content: articleCreateDto.Content,
		Creator: articleCreateDto.CreatedBy,
		State:   articleCreateDto.State,
	})

	return true
}
func ExistArticleByID(id int) bool {
	var article entity.Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(entity.Article{})

	return true
}