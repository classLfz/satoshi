package siri

import (
	"log"
	"time"

	"github.com/brutella/hc/accessory"
	"github.com/classlfz/satoshi/cmd/config"
	"github.com/classlfz/satoshi/pkg/sensor"
)

type InitSwitchOption struct {
	LircdPath string
}

// InitSwitch 初始化开关
func InitSwitch(config config.SiriDevice, option InitSwitchOption) *accessory.Switch {
	// create an accessory
	info := accessory.Info{
		ID:   config.ID,
		Name: config.Name,
	}
	ac := accessory.NewSwitch(info)

	ac.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			// by pin
			if config.SwitchConfig.OnPin != 0 {
				newState, err := sensor.CallSinglePin(sensor.CallSinglePinOptions{
					PinNum: config.SwitchConfig.OnPin,
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
			if config.SwitchConfig.OnLircdCmd != "" {
				err := sensor.CallLircd(sensor.CallLircdOptions{
					LircdSocketPath: option.LircdPath,
					SendStr:         config.SwitchConfig.OnLircdCmd,
				})
				if err != nil {
					log.Panic("Call sensor err: \n", err)
					return
				}
			}
		} else {
			// by pin
			if config.SwitchConfig.OnPin != 0 {
				newState, err := sensor.CallSinglePin(sensor.CallSinglePinOptions{
					PinNum: config.SwitchConfig.OnPin,
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
			if config.SwitchConfig.OffLircdCmd != "" {
				err := sensor.CallLircd(sensor.CallLircdOptions{
					LircdSocketPath: option.LircdPath,
					SendStr:         config.SwitchConfig.OffLircdCmd,
				})
				if err != nil {
					log.Panic("Call sensor err: \n", err)
					return
				}
			}
		}
	})

	// polling update state
	if config.SwitchConfig.UpdateInterval != 0 {
		go func() {
			d := time.Duration(time.Second * time.Duration(config.SwitchConfig.UpdateInterval))

			t := time.NewTicker(d)
			defer t.Stop()

			for {
				<-t.C
				// by pin
				if config.SwitchConfig.OnPin != 0 {
					state, err := sensor.CallSinglePin(sensor.CallSinglePinOptions{
						PinNum: config.SwitchConfig.OnPin,
						Read:   true,
					})
					if err != nil {
						log.Panic("Switch polling update state err: \n", err)
					}
					stateVal := false
					if state == 1 {
						stateVal = true
					}
					ac.Switch.On.SetValue(stateVal)
				}
			}
		}()
	}

	return ac
}
