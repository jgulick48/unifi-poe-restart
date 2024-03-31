package main

import (
	"encoding/json"
	"flag"
	"fmt"
	controller2 "github.com/jgulick48/unifi-poe-restart/internal/controller"
	"github.com/jgulick48/unifi-poe-restart/internal/models"
	"github.com/jgulick48/unifi-poe-restart/internal/unificlient"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var configLocation = flag.String("configFile", "./config.json", "Location for the configuration file.")

func main() {
	config := LoadClientConfig(*configLocation)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	uniClient := unificlient.NewClient(config.Controller)
	controller := controller2.NewClient(config.Devices, uniClient)

	log.Println("Running application")
	controller.Start()
	select {
	case <-done:
		controller.Stop()
	}
	log.Println("exiting")
}

func LoadClientConfig(filename string) models.UnifiPoeRestartConfiguration {
	if filename == "" {
		filename = "./config.json"
	}
	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("No config file found. Making new IDs")
		panic(err)
	}
	var config models.UnifiPoeRestartConfiguration
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Printf("Invliad config file provided")
		panic(err)
	}
	return config
}
