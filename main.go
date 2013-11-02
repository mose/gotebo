package main

import (
	// "bufio"
	"fmt"
	"github.com/jessevdk/go-flags"
	// "github.com/nsf/termbox-go"
	"log"
	// "math/rand"
	"os"
	// "time"
)

type Opts struct {
	Configfile string `short:"c" long:"config" description:"Config file." default: "~/.gotebo"`
}

var opts Opts

// func draw() {
// 	w, h := termbox.Size()
// 	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
// 	for y := 0; y < h; y++ {
// 		for x := 0; x < w; x++ {
// 			termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.Attribute(rand.Int()%8)+1)
// 		}
// 	}
// 	termbox.Flush()
// }

func main() {
	logfile, _ := os.OpenFile("gotebo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(logfile)

	args, _ := flags.Parse(&opts)
	if args != nil {
		fmt.Printf("args: %s\n", args)
	}

	// fileinfo, err := os.Stat("config.yml")
	// if err != nil {
	// 	panic(err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("args: %s\n", fileinfo)
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
}
