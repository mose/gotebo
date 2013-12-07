package main

import (
	// "bufio"
	"fmt"
	//"github.com/jessevdk/go-flags"
	// "github.com/nsf/termbox-go"
	"log"
	// "math/rand"
	"os"
	// "time"
)

func Run() error {
	logfile, _ := os.OpenFile("gotebo.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	log.SetOutput(logfile)

	if err := Flags(); err != nil {
		log.Fatal(err)
		return err
	}

	fileinfo, err := os.Stat(ConfigFile)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("args: %s\n", fileinfo)
	// if err != nil {
	// 	fmt.Printf("Config file not found: %s\n", opts.Configfile)
	// 	fmt.Printf("Do you want to create one ? [Y/n]")
	// 	bio := bufio.NewReader(os.Stdin)
	// 	line, _, _ := bio.ReadLine()
	// 	if string(line[:]) == "y" {
	// 		fmt.Printf("ok\n")
	// 	}
	// 	err := termbox.Init()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer termbox.Close()
	// 	event_queue := make(chan termbox.Event)
	// 	go func() {
	// 		for {
	// 			event_queue <- termbox.PollEvent()
	// 		}
	// 	}()
	// 	draw()
	// loop:
	// 	for {
	// 		select {
	// 		case ev := <-event_queue:
	// 			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
	// 				break loop
	// 			}
	// 		default:
	// 			draw()
	// 			time.Sleep(10 * time.Millisecond)
	// 		}
	// 	}
	// }
	return nil
}
