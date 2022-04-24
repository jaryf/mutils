package runtime

import "runtime"

// RunFuncInfo 获取正在运行的函数名
func RunFuncInfo() (fileName string, line int) {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	fileName, line = f.FileLine(pc[0])
	return
}
