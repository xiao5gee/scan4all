package main

import (
	"context"
	_ "embed"
	"github.com/hktalent/scan4all/lib"
)

// 测试正则表达式是否正确
func main() {
	// 中途控制关闭当前目标所有fuzz
	_, stop := context.WithCancel(lib.Ctx_global)
	stop()
	stop()
	stop()
	stop()

}