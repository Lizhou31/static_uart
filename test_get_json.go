package main
/*
#cgo LDFLAGS: -L. -lgetcontrollerinfo
#include "control_get_info.h"
*/
import "C"

import (
	"encoding/json"
	"fmt"
)
type controller_info struct {
    Timestamp int
    Velocity int
	Alarm int
	Brake int
	Meas int
	Rssi int
	Mode int
	Ob_dis int
}

func main(){
	var info controller_info
	error := C.init_serial()
	if error != 0{
		fmt.Println("error")
		return
	}
	for true{
		error := C.read_data(100)
		if error != 0{
			fmt.Println("error")
			return
		}
		str := C.GoString(&C.read_buf[0])
		fmt.Println(str)
		json.Unmarshal([]byte(str), &info)
    	fmt.Println(info)
		fmt.Println("")
	}
}
