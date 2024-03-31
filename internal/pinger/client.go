package pinger

import (
	"github.com/prometheus-community/pro-bing"
	"log"
	"time"
)

func PingHost(address string) (*probing.Statistics, error) {
	pinger, err := probing.NewPinger(address)
	if err != nil {
		return nil, err
	}
	pinger.Count = 10
	pinger.Timeout = time.Second * 20
	pinger.OnFinish = func(stats *probing.Statistics) {
		log.Printf("--- %s ping statistics ---\n", stats.Addr)
		log.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		log.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}
	err = pinger.Run()
	if err != nil {
		return nil, err
	}
	stats := pinger.Statistics()
	return stats, nil
}
