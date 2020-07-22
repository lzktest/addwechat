package main

import (
	"fmt"
	"io/ioutil"

	"github.com/go-vgo/robotgo"
)

func main() {
	var files_path = "C:\\Users\\yw\\Desktop\\表情图片"
	fmt.Print("test")
	var send_file_mleft = robotgo.AddEvent("mleft")
	if send_file_mleft {
		position_send_file_x, position_send_file_y := robotgo.GetMousePos()
		fmt.Printf("position_send_file: %v %v\n", position_send_file_x, position_send_file_y)
	}
	fmt.Print("click1 close")
	new_photo_mleft := robotgo.AddEvent("mleft")
	if new_photo_mleft {
		position_new_photo_x, position_new_photo_y := robotgo.GetMousePos()
		fmt.Printf("position_send_file: %v %v\n", position_new_photo_x, position_new_photo_y)
	}
	fmt.Print("click2 close")
	add_emoji_mleft := robotgo.AddEvent("mleft")
	if add_emoji_mleft {
		position_add_emoji_x, position_add_emoji_y := robotgo.GetMousePos()
		fmt.Printf("position_send_file: %v %v\n", position_add_emoji_x, position_add_emoji_y)
	}
	fmt.Print("click3 close")
	file_name_list, err := ioutil.ReadDir(files_path)
	if err != nil {
		fmt.Print("error: %v", err)
	}

	fmt.Printf("test %v", file_name_list)
}
