package models

type CommentModel struct {
	/*
		评论表与用户表关联
		评论表与文章表关联
		评论表自关联：子评论
	*/
	MODEL
	SubComments        []*CommentModel `json:"sub_comments" gorm:"foreignKey:ParentCommentID"`  // 子评论列表
	ParentCommentModel *CommentModel   `json:"comment_model" gorm:"foreignKey:ParentCommentID"` // 父级评论
	ParentCommentID    *uint           `json:"parent_comment_id"`                               // 父评论id
	Content            string          `json:"content" gorm:"size:256"`                         // 评论内容
	DiggCount          int             `json:"digg_count" gorm:"size:8;default:0"`              // 点赞数
	CommentCount       int             `json:"comment_count" gorm:"size:8;default:0"`           // 子评论数
	Article            ArticleModel    `json:"-" gorm:"foreignKey:ArticleID"`                   // 关联文章
	ArticleID          uint            `json:"article_id"`                                      // 文章id
	User               UserModel       `json:"user"`                                            // 关联用户
	UserID             uint            `json:"user_id"`                                         // 评论用户
}
