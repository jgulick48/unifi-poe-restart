package controller

import (
	"github.com/jgulick48/unifi-poe-restart/internal/models"
	"github.com/jgulick48/unifi-poe-restart/internal/unificlient"
	"sync"
	"time"
)

type Client interface {
	Start()
	Stop()
}

type client struct {
	uiClient unificlient.Client
	devices  []device
	done     chan (bool)
}

func NewClient(config []models.DeviceMapper, uiClient unificlient.Client) Client {
	deviceList := make([]device, 0, len(config))
	for _, deviceConfig := range config {
		deviceList = append(deviceList, device{
			DeviceMapper: deviceConfig,
			mu:           sync.Mutex{},
			done:         make(chan bool),
		})
	}
	return &client{
		uiClient: uiClient,
		devices:  deviceList,
		done:     make(chan bool),
	}
}

func (c *client) Start() {
	for i := range c.devices {
		index := i
		go func() {
			ticker := time.NewTicker(c.devices[i].PingInterval.Duration)
			for {
				select {
				case <-ticker.C:
					if c.devices[index].pingDevice() {
						if c.devices[index].checkShouldRestart() {
							c.devices[index].restart(c.uiClient)
						}
					}
				case <-c.devices[index].done:
					return
				}
			}
		}()
	}
	for {
		select {
		case <-c.done:
			for i := range c.devices {
				c.devices[i].done <- true
			}
		}
	}
}

func (c *client) Stop() {
	c.done <- true
}
