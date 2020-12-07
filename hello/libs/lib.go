package lib

import (
	//"container/list"
	"context"
	"fmt"
	"sync"

	//"sync"
	"time"
)
func A(){
   //l:=list.New()
   //var once sync.Once
   //once.Do(B)
   //context.TODO()
   //withValue:=context.WithValue(context.Background(),1,1)
   ct,f1:=context.WithTimeout(context.Background(),time.Second*3)
   ct2,_:=context.WithTimeout(ct,time.Second*3)
   p:=sync.Pool{}
   p.New=func() interface{}{
   	  return 6
	}
	
   f1()
   //withValue:=context.WithValue(ct,1,1)
   go func(){
	   for{
		   select {
		   case <-ct.Done():
		   default:
			   fmt.Println("context not end")
			   //time.Sleep(time.Second)
		   }
		   select {
		   case <-ct2.Done():
		   default:
			   fmt.Println("withValue not end")
			   //time.Sleep(time.Second)
		   }
		   time.Sleep(time.Second)
	   }
   }()

   time.Sleep(time.Second*5)

}
func  B(){

}
