package utils

import "path/filepath"

// DropPathLevels 路径降级，如c:/aaa/bbb,降一级为c:/aaa
func DropPathLevels(path string, levels int) string {
	for i := 0; i < levels; i++ {
		path = filepath.Dir(path)
	}
	return path
}
