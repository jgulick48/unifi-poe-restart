package unificlient

import (
	"context"
	"encoding/json"
	"github.com/jgulick48/unifi-poe-restart/internal/models"
	"github.com/unpoller/unifi"
	"log"
	"time"
)

type Client interface {
	RestartDevice(ctx context.Context, mac string, port int) error
}
type client struct {
	unifiClient *unifi.Unifi
	lastLogin   time.Time
}

func NewClient(config models.Controller) Client {
	c := unifi.Config{
		User: config.Username,
		Pass: config.Password,
		URL:  config.URL,
		// Log with log.Printf or make your own interface that accepts (msg, fmt)
		ErrorLog: log.Printf,
		DebugLog: log.Printf,
	}
	uni, err := unifi.NewUnifi(&c)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	return &client{
		uni,
		time.Now(),
	}
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *client) RestartDevice(ctx context.Context, mac string, port int) error {
	if time.Now().After(c.lastLogin.Add(2 * time.Hour)) {
		err := c.unifiClient.Login()
		if err != nil {
			log.Printf("error logging into unifi controller skipping restart: %s", err)
			return err
		}
	}
	command := SwitchCommand{
		Mac:     mac,
		PortIDX: port,
		Cmd:     "power-cycle",
	}
	body, err := json.Marshal(command)
	if err != nil {
		return err
	}
	response, err := c.unifiClient.PostJSON("/api/s/default/cmd/devmgr", string(body))
	log.Printf("Got repsone on port restart of %s\n", response)
	return err
}
