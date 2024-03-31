package controller

import (
	"context"
	"github.com/jgulick48/unifi-poe-restart/internal/models"
	"github.com/jgulick48/unifi-poe-restart/internal/pinger"
	"github.com/jgulick48/unifi-poe-restart/internal/unificlient"
	"log"
	"sync"
	"time"
)

type device struct {
	models.DeviceMapper
	LastRestarted time.Time `json:"-"`
	CurrentTries  int       `json:"-"`
	mu            sync.Mutex
	done          chan (bool)
}

func (d *device) pingDevice() bool {
	results, err := pinger.PingHost(d.IPAddress)
	if err != nil {
		log.Printf("error trying to ping host %s:%v %s\n", d.Name, d.IPAddress, err)
		d.mu.Lock()
		d.CurrentTries++
		d.mu.Unlock()
		return true
	}
	if results == nil {
		log.Printf("error trying to ping host %s: Results set was nil\n", d.IPAddress)
		return false
	}
	if results.PacketLoss > d.LossThreshold {
		log.Printf("device %s had packet loss higher than threshold %v of %v, adding failed retry\n", d.Name, d.LossThreshold, results.PacketLoss)
		d.mu.Lock()
		d.CurrentTries++
		d.mu.Unlock()
		log.Printf("failed retry count for device %s is currently at %v threshold for failed retries is %v\n", d.Name, d.CurrentTries, d.RetryCount)
		return true
	}
	if d.CurrentTries > 0 {
		log.Printf("got successful ping for device %s: after %v tries, resetting to 0", d.Name, d.CurrentTries)
		d.mu.Lock()
		d.CurrentTries = 0
		d.mu.Unlock()
	}
	return false
}

func (d *device) checkShouldRestart() bool {
	if d.CurrentTries <= d.RetryCount {
		return false
	}
	log.Printf("device %s is currently at %v tries which is above its retry count of %v, checking to see if it needs to be restarted\n", d.Name, d.CurrentTries, d.RetryCount)
	if time.Now().Add(1 * d.RestartDelay.Duration).Before(d.LastRestarted) {
		log.Printf("device %s was last restarted at %s, must wait till at least after %s to attempt another restart\n", d.Name, d.LastRestarted, d.LastRestarted.Add(d.RestartDelay.Duration))
		return false
	}
	return true
}

func (d *device) restart(uiClient unificlient.Client) {
	err := uiClient.RestartDevice(context.Background(), d.SwitchMac, d.PortIDX)
	if err == nil {
		d.mu.Lock()
		d.LastRestarted = time.Now()
		d.mu.Unlock()
		return
	}
	log.Printf("error trying to restart device %s:%s\n", d.Name, err)
}
