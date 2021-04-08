package siri

import (
	"log"
	"time"

	"github.com/brutella/hc/accessory"
	"github.com/classlfz/satoshi/cmd/config"
	"github.com/classlfz/satoshi/pkg/sensor"
)

type InitLightbulbOption struct {
	LircdPath string
}

// InitLightbulb 初始化灯泡
func InitLightbulb(config config.SiriDevice, option InitLightbulbOption) *accessory.Lightbulb {
	// create an accessory
	info := accessory.Info{
		ID:   config.ID,
		Name: config.Name,
	}
	ac := accessory.NewLightbulb(info)

	ac.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			// by pin
			if config.LightbulbConfig.OnPin != 0 {
				newState, err := sensor.CallSinglePin(sensor.CallSinglePinOptions{
					PinNum: config.LightbulbConfig.OnPin,
					State:  true,
				})
				if err != nil {
					log.Panic("Call sensor err: \n", err)
					return
				}
				log.Printf("Switch state: %v", newState)
				return
			}
			// by lircd
			if config.LightbulbConfig.OnLircdCmd != "" {
				err := sensor.CallLircd(sensor.CallLircdOptions{
					LircdSocketPath: option.LircdPath,
					SendStr:         config.LightbulbConfig.OnLircdCmd,
				})
				if err != nil {
					log.Panic("Call sensor err: \n", err)
					return
				}
			}
		} else {
			// by pin
			if config.LightbulbConfig.OnPin != 0 {
				newState, err := sensor.CallSinglePin(sensor.CallSinglePinOptions{
					PinNum: config.LightbulbConfig.OnPin,
					State:  false,
				})
				if err != nil {
					log.Panic("Call sensor err: \n", err)
					return
				}
				log.Printf("Switch state: %v", newState)
				return
			}
			// by lircd
			if config.LightbulbConfig.OffLircdCmd != "" {
				err := sensor.CallLircd(sensor.CallLircdOptions{
					LircdSocketPath: option.LircdPath,
					SendStr:         config.LightbulbConfig.OffLircdCmd,
				})
				if err != nil {
					log.Panic("Call sensor err: \n", err)
					return
				}
			}
		}
	})

	// polling update state
	if config.LightbulbConfig.UpdateInterval != 0 {
		go func() {
			d := time.Duration(time.Second * time.Duration(config.LightbulbConfig.UpdateInterval))

			t := time.NewTicker(d)
			defer t.Stop()

			for {
				<-t.C
				// by pin
				if config.LightbulbConfig.OnPin != 0 {
					state, err := sensor.CallSinglePin(sensor.CallSinglePinOptions{
						PinNum: config.LightbulbConfig.OnPin,
						Read:   true,
					})
					if err != nil {
						log.Panic("Lightbulb polling update state err: \n", err)
					}
					stateVal := false
					if state == 1 {
						stateVal = true
					}
					ac.Lightbulb.On.SetValue(stateVal)
				}
			}
		}()
	}

	return ac
}
