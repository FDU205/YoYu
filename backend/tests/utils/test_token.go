/*token的测试文件*/
package main

import (
	"YOYU/backend/utils"
	"fmt"
	"time"
)

func main() {

	token := utils.GenToken(999)

	fmt.Println("Token = ", token)
	fmt.Println(len(token))

	time.Sleep(time.Second * 2)

	my_claim, err := utils.ParseToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Println("my claim = ", my_claim)

}
