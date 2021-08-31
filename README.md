# Helper
##### 1.一些常用的助手函数工具包

##### 2.获取 
go get -u github.com/applytogetintoyourlife/Helper
##### 3. 示例 for example
```Golang
package helper

import (
	"fmt"
	"github.com/applytogetintoyourlife/Helper"
)

func example() {
	// string
	s1 := `123456`
	md5 := helper.TStr.MD5(s1)
	fmt.Println(md5) // E10ADC3949BA59ABBE56E057F20F883E

	s2 := `hello world`
	ucFirst := helper.TStr.UcFirst(s2)
	fmt.Println(ucFirst) // Hello world

	sfx := helper.TStr.Shuffle(s1)
	fmt.Println(sfx) // 125436

	// array
	arr := [5]int{1, 2, 3, 4, 5}
	i := 2
	if helper.TArr.InArray(i, arr) {
		fmt.Printf(" %v in %v \n", i, arr)
	}

	// debug
	funcName := helper.TDebug.GetFuncName(helper.TArr.ArrayDiff, true) // ...ArrayDiff-fm
	fmt.Println(funcName) // ArrayDiff-fm

	// int
	round := helper.TInt.Round(4.56)
	fmt.Println(round) // 4

	// json
	jsonArr := `[{"email_address":"test1@email.com"},{"email_address":"test2@email.com"}]`
	m := helper.TJson.JsonToMapArr(jsonArr)
	fmt.Println(m) // [map[email_address:test1@email.com] map[email_address:test2@email.com]]
	// ...
}
```

##### 4. 使用过程如有疑问欢迎issue