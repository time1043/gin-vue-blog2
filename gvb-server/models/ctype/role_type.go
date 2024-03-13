package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       = 1 // 管理员
	PermissionUser        = 2 // 普通登录者
	PermissionVisitor     = 3 // 游客
	PermissionDisableUser = 4 // 被封禁的用户
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "普通登录者"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisableUser:
		str = "被封禁的用户"
	default:
		str = "其他"
	}
	return str
}
