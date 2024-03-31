package models

import (
	"encoding/json"
	"errors"
	"time"
)

type UnifiPoeRestartConfiguration struct {
	Devices    []DeviceMapper `json:"devices"`
	Controller Controller     `json:"controller"`
}

type DeviceMapper struct {
	Name          string   `json:"name"`
	SwitchMac     string   `json:"switchMac"`
	PortIDX       int      `json:"portIDX"`
	IPAddress     string   `json:"ipAddress"`
	PingInterval  Duration `json:"pingInterval"`
	RetryCount    int      `json:"retryCount"`
	LossThreshold float64  `json:"lossThreshold"`
	RestartDelay  Duration `json:"restartDelay"`
}

type Controller struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}
