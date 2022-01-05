package main

import (
	"fmt"
	"mall/common/cryptx"
)

func main()  {
	dk:= cryptx.PasswordEncrypt("HWVOFkGgPTryzICwd7qnJaZR9KQ2i8xe", "123456")
	 //fmt.Sprintf("%x", string(dk))
	 fmt.Printf(string(dk))
}




