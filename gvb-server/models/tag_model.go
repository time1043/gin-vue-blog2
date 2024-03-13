package models

type TagModel struct {
	MODEL
	Title    string         `json:"title" gorm:"size:16"`                  // 标签的名称
	Articles []ArticleModel `json:"-" gorm:"many2many:article_tag_models"` // 关联该标签的文章列表
}
