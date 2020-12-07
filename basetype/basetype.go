package basetype

import (
	"fmt"
)

type S struct {
	a int
}

func Point() {
	t := S{a: 1}
	q := S{a: 2}
	t2 := S{a: 3}
	p := &t
	q = *p
	fmt.Printf("t:%v,q:%v,t2:%v \n", t, q, t2)
	*p = t2
	fmt.Printf("t:%v,q:%v,t2:%v \n", t, q, t2)
	t2.a = 4
	fmt.Printf("t:%v,q:%v,t2:%v \n", t, q, t2)
}

func Slince() {
	//切片赋值
	//切片参数
	s := make([]string, 3, 10)
	//s:=[]string{3:"a"}
	doSlince(&s)
	printSlince(s)
	//s2:=s
	//s2[0]="s2"
	//s2 = append(s2, "2")
	//s2[1]="3"
	//s[1]="33"//覆盖s2[1],虽然s2与s不是相同的切片，len不相同
	//printSlince(s)
	//printSlince(s2)
	//s3:=s[0:1]
	//s3[0]="00"
	//s3=append(s3, "s3")
	//printSlince(s)
	//printSlince(s3)
}
func doSlince(s *[]string) {
	(*s)[0] = "1"
	(*s)[2] = "2"
	s2 := append(*s, "a")
	*s = s2
}
func printSlince(s []string) {
	fmt.Printf("slince：size=%d,len=%d ,s=%v \n", cap(s), len(s), s)
}
