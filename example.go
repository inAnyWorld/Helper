package helper

import "fmt"

func example() {
	sub := TStr.GetBetweenStr(`@abc456%`, "@", "%")
	fmt.Println(sub) // @abc456
}
