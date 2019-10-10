package main

import (
	"fmt"
	"github.com/robfig/cron"
	"os"
	"os/signal"
	"time"
)

func main() {
	crontab := cron.New()
	crontab.AddFunc("1 */2 21-23 * * *", task)
	fmt.Println("run...")
	defer crontab.Stop()
	crontab.Start()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	go listen()
	select {
	case <-ch:
		fmt.Println("interrupt receive...")
		crontab.Stop()
	}
}

func listen() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	for {
		select {
		case <-ch:
			fmt.Println("anothre listen...")
		}
	}
}

func task() {
	fmt.Println("task...")
	fmt.Println(time.Now().Unix())
}
