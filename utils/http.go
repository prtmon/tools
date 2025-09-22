package utils

import "strings"

func IsActive(menuPath, currentPath string) bool {
	// 精确匹配
	if menuPath == currentPath {
		return true
	}
	// 通配符匹配（如 /files/*filepath）
	if strings.HasPrefix(menuPath, "/*") && strings.HasPrefix(currentPath, menuPath[:len(menuPath)-1]) {
		return true
	}
	return false
}
