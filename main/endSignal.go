package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	loopChan := make(chan bool, 1)
	defer close(loopChan)

	sigInt := make(chan bool, 1)
	defer close(sigInt)

	// interrupt 시그널 수신 대기
	sigs := make(chan os.Signal, 1)
	defer close(sigs)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// interrupt 시그널을 수신하면 프로그램을 종료한다.
	go func() {
		sig := <-sigs
		fmt.Println("Receive signal: ", sig)
		sigInt <- true
	}()

	// 8초에 한 번씩 Working 을 호출함
	go func() {
		for {
			time.Sleep(time.Second * time.Duration(8))
			loopChan <- true
		}
	}()

	workingProcess(loopChan, sigInt)
}

func workingProcess(loopChan chan bool, sigInt chan bool) {
	// loop 를 돌면서 channel 에 데이터를 Receive 때마다 해당 로직을 수행
	for { // loop
		select {
		case <-loopChan:
			log.Println("==========Working============")
		case <-sigInt: // 프로세스 종료 시그널
			log.Println("==========Cancel============")
			os.Exit(0)
			return
		default:
			time.Sleep(time.Second * 1)
		}
	}
}
