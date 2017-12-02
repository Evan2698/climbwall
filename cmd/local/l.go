package main

import (
	"syscall"
	"os/signal"
	"path/filepath"
	"os"
    "climbwall/utils"
    "climbwall/core"
)


func run_local (config * core.AppConfig) {
	core.Run_Local_routine(config)
}


func wait_s(){
	var system_signal = make(chan os.Signal, 2)
	signal.Notify(system_signal, syscall.SIGINT, syscall.SIGHUP)
	for sig := range system_signal {
		if sig == syscall.SIGHUP || sig == syscall.SIGINT{
			utils.Logger.Printf("caught signal %v, exit", sig)
			os.Exit(0)
			
		} else {

			utils.Logger.Printf("XXX caught signal %v, exit", sig)
			os.Exit(0)
		}
	}
}



func main(){

	utils.Logger.Print(".................")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		utils.Logger.Fatal("can not combin config file path!")
		os.Exit(1)
	}

	config, err := core.Parse(dir + "/config.json")
	if err != nil {
		utils.Logger.Fatal("load config file failed!")
		os.Exit(1)
	}

	go run_local(config)

	wait_s()
}



