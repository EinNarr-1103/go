package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Print("Mouse hack processing... \n")
	// ここの数を変更
	mouseMoveByMinute(5)
}

func mouseMoveByMinute(minute int) {
	milliSecond := float64(minute) * 60.0 * 1000.0 * 0.99999
	for {
		now := time.Now()
		time.Sleep(time.Millisecond * time.Duration(milliSecond - 5000))
		fmt.Print("経過時間(秒)" + strconv.FormatInt(int64(time.Since(now).Seconds()), 10) + "\n")
		// マウスの現在位置を取得
		x,y := robotgo.GetMousePos();
		robotgo.Move(x + 1, y + 1)
	}
}