package model

import "gorm.io/gorm"

// 模型定义
type ResourcePerms struct {
	gorm.Model
	Name     string `json:"name" description:"资源中文名"`
	Module   string `json:"module" description:"模块"`
	Resource string `json:"resource" description:"资源路径"`
	Action   string `json:"action" description:"资源权限动作"`
	Label    string `json:"label" description:"说明"`
}
