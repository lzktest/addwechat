package main

import (
	"fmt"
	"io/ioutil"

	"github.com/go-vgo/robotgo"
)

func main() {
	var files_path = "C:\\Users\\yw\\Desktop\\表情图片"
	var position_send_file_x, position_send_file_y, position_new_photo_x, position_new_photo_y, position_add_emoji_x, position_add_emoji_y int
	//ch := make(chan [2]int)
	//fmt.Print("test")
	//go get_position(ch)
	//fmt.Println(<-ch)
	fmt.Printf("test3")
	var send_file_mleft = robotgo.AddEvent("a")
	if send_file_mleft {
		position_send_file_x, position_send_file_y = robotgo.GetMousePos()
		fmt.Printf("position_send_file: %v %v\n", position_send_file_x, position_send_file_y)
	}
	robotgo.SetMouseDelay(30)
	fmt.Print("click1 close")
	new_photo_mleft := robotgo.AddEvent("b")
	if new_photo_mleft {
		position_new_photo_x, position_new_photo_y = robotgo.GetMousePos()
		fmt.Printf("position_send_file: %v %v\n", position_new_photo_x, position_new_photo_y)
	}

	fmt.Print("click2 close")
	robotgo.SetMouseDelay(30)
	add_emoji_mleft := robotgo.AddEvent("c")
	if add_emoji_mleft {
		position_add_emoji_x, position_add_emoji_y = robotgo.GetMousePos()
		fmt.Printf("position_send_file: %v %v\n", position_add_emoji_x, position_add_emoji_y)
	}

	file_name_list, err := ioutil.ReadDir(files_path)
	if err != nil {
		fmt.Print("error: %v", err)
	}
	for _, file_name := range file_name_list {
		fmt.Print(file_name.Name())
		robotgo.MoveClick(position_send_file_x, position_send_file_y, "left", false)
		robotgo.TypeStr(files_path + "\\" + file_name.Name())
		robotgo.KeyTap("enter")
		robotgo.KeyTap("enter")
		robotgo.MoveClick(position_new_photo_x, position_new_photo_y, "right", false)
		robotgo.MoveClick(position_add_emoji_x, position_add_emoji_y, "left", false)
	}

	//fmt.Printf("test %v", file_name_list)

	fmt.Print("end")

}

/*
func get_position(ch chan<- [2]int) {
	fmt.Printf("in")
	var send_file_mleft = robotgo.AddEvent("mleft")
	if send_file_mleft {
		//position_send_file_x, position_send_file_y := robotgo.GetMousePos()
		var array1 [2]int
		array1[0], array1[1] = robotgo.GetMousePos()
		//fmt.Printf("position_send_file: %v %v\n", position_send_file_x, position_send_file_y)
		fmt.Printf("position_send_file: %v %v\n", array1[0], array1[1])
		ch <- array1
	}
	robotgo.StopEvent()
	fmt.Printf("test1")
}
*/

/*

 */
