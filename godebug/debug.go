package godebug

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// GetFuncName 获取调用方法的名称;f为目标方法;onlyFun为true时仅返回方法,不包括包名.
func GetFuncName(f interface{}, onlyFun ...bool) string {
	var funcObj *runtime.Func
	if f == nil {
		// Skip this function, and fetch the PC and file for its parent
		pc, _, _, _ := runtime.Caller(1)
		// Retrieve a Function object this functions parent
		funcObj = runtime.FuncForPC(pc)
	} else {
		funcObj = runtime.FuncForPC(reflect.ValueOf(f).Pointer())
	}

	name := funcObj.Name()
	if len(onlyFun) > 0 && onlyFun[0] == true {
		// extract just the function name (and not the module path)
		return strings.TrimPrefix(filepath.Ext(name), ".")
	}

	return name
}

// GetFuncLine 获取调用方法的行号.
func GetFuncLine() int {
	// Skip this function, and fetch the PC and file for its parent
	_, _, line, _ := runtime.Caller(1)
	return line
}

// GetFuncFile 获取调用方法的文件路径.
func GetFuncFile() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// DumpStacks 打印堆栈信息.
func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}

// HasMethod 检查对象是否具有某方法.
func HasMethod(t interface{}, method string) bool {
	_, ok := reflect.TypeOf(t).MethodByName(method)
	if ok {
		return true
	}
	return false
}
