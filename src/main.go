package main

import (
	"fmt"
	"log"
	"prefix-matcher/src/config"
	"prefix-matcher/src/repository"
	"prefix-matcher/src/server"
	"runtime/debug"
	"sync"
)

func main() {

	defer func() {
		r := recover()
		if r != nil {
			log.Println(r, debug.Stack())
		}
	}()
	fileName, err := config.LoadConfig()
	if err != nil {
		log.Println(err, "error loading config")
		return
	}

	go config.SetupConfig("src/config/" + fileName)
	wg := &sync.WaitGroup{}
	r := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go repository.Data.Insert(config.PrefixChan, wg, r)
	}
	fmt.Println("Wait")
	wg.Wait()
	log.Println("INFO : Stating Server")
	server.LoadRoute()

}
