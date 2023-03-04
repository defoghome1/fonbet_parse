package main

import (
	"awesomeProject1/scrap"
	send_message "awesomeProject1/send-message"
	"time"
)

func main() {
	N := 1.0
	N = send_message.TakeK(N)
	for{
		scrap.FonbetParse(N)
		time.Sleep(time.Second*10)
	}
}