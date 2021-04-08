package devicecallerrouter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/bitly/go-simplejson"
	"github.com/classlfz/satoshi/cmd/config"
	"github.com/classlfz/satoshi/pkg/sensor"
)

type caller struct {
	ID     string `json:"id"`
	Read   bool   `json:"read"`
	Toggle bool   `json:"toggle"`
	State  bool   `json:"state"`
}

// CallDeviceByID 调用设备
func CreateDeviceCaller(writer http.ResponseWriter, req *http.Request) {
	var c caller
	var d config.HttpDevice
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&c); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("c: %v", c)

	cfg, _, configErr := config.Load("~/.satoshi/config.yaml")
	if configErr != nil {
		log.Panicf("config.Load err\n%v", configErr)
		http.Error(writer, configErr.Error(), http.StatusInternalServerError)
		return
	}

	devices := cfg.Satoshi.Http.Devices

	for i := 0; i < len(devices); i++ {
		if devices[i].ID == c.ID {
			d = devices[i]
		}
	}

	if d.ID == "" {
		http.Error(writer, "device not found", http.StatusNotFound)
		return
	}

	switch d.Type {
	case "switch":
		// by pin
		if d.SwitchConfig.OnPin != 0 {
			newState, callErr := sensor.CallSinglePin(sensor.CallSinglePinOptions{
				PinNum: d.SwitchConfig.OnPin,
				Read:   c.Read,
				State:  c.State,
				Toggle: c.Toggle,
			})
			if callErr != nil {
				log.Panicf("sensor.CallSwitch err %v", callErr)
				http.Error(writer, configErr.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(writer, "Call success with new state %v", newState)
			return
		}
		// by lircd
		if d.SwitchConfig.OnLircdCmd != "" && d.SwitchConfig.OffLircdCmd != "" {
			var SendStr string
			if c.State {
				SendStr = d.SwitchConfig.OnLircdCmd
			} else {
				SendStr = d.SwitchConfig.OffLircdCmd
			}
			err := sensor.CallLircd(sensor.CallLircdOptions{
				LircdSocketPath: cfg.Lircd.LircdPath,
				SendStr:         SendStr,
			})
			if err != nil {
				log.Panic("Call sensor err: \n", err)
				return
			}
		}
	}
}
