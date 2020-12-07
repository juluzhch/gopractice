package tobuildin

import "fmt"

func Protect() {

	defer func() {
		fmt.Println("第一个defer")
	}()
	defer func() {
		fmt.Println("recover 之前，defer 里面")
		if x := recover(); x != nil {
			fmt.Printf("panic :%v \n", x)
		}
		fmt.Println("recover 后，defer 里面")
	}()
	defer func() {
		fmt.Println("第三个defer")
	}()
	throwFun()
}
func throwFun() {
	panic("主动抛出panic")
}
