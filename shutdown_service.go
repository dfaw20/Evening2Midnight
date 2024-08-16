package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func isLidClosed() bool {
	// /proc/acpi/button/lid/LID0/state を読み込む
	data, err := os.ReadFile("/proc/acpi/button/lid/LID0/state")

	fmt.Println(strings.ToLower(string(data)))

	if err != nil {
		fmt.Println("Err reading lid state:", err)
		return false
	}

	// "closed" という単語が含まれているかどうかを確認
	return strings.Contains(strings.ToLower(string(data)), "closed")
}

func isShutdownPeriod() bool {
	// 日本の現在時刻を取得
	jst := time.FixedZone("JST", 9*60*60)
	now := time.Now().In(jst)

	fmt.Println(now)

	// 19:00~1:00の間
	evening2MidnightPeriod := 19 < now.Hour() || now.Hour() < 1
	fmt.Println("Evening 2 Midnight: ", evening2MidnightPeriod)
	return !evening2MidnightPeriod
}

func main() {
	isProd := true

	for {
		if isShutdownPeriod() && isLidClosed() {

			// 10分待機
			if isProd {
				time.Sleep(10 * time.Minute)
			} else {
				time.Sleep(1 * time.Second)
			}

			// 再度蓋が閉じたままか確認
			if isShutdownPeriod() && isLidClosed() {
				fmt.Println("Shutting down ... ")
				cmd := exec.Command("/sbin/shutdown", "-h", "now")
				if err := cmd.Run(); err != nil {
					fmt.Println("Error executing shutdown:", err)
				}
				return
			}
		}
		// 30分ごとにチェック
		if isProd {
			time.Sleep(30 * time.Minute)
		} else {
			time.Sleep(3 * time.Second)
		}
		fmt.Println()
	}
}
