package models

type FadeBackModel struct {
	MODEL
	Email        string `json:"email" gorm:"size:64"`
	Content      string `json:"content"gorm:"size:128"`
	ApplyContent string `json:"apply_content" gorm:"size:128"` // 回复的内容
	IsApply      bool   `json:"is_apply"`                      // 是否回复
}
