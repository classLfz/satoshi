package siri

import (
	"log"
	"os"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/classlfz/satoshi/cmd/config"
)

// Start 开启模拟 Siri 可控设备
func Start(configPath string) {
	const CurrentAPI = "Siri Start"
	cfg, _, loadErr := config.Load(configPath)
	if loadErr != nil {
		log.Fatal(CurrentAPI, " config.load \"", configPath, "\" err:\n", loadErr)
		return
	}

	acArr := make([]*accessory.Accessory, 0)

	// bridge
	info := accessory.Info{
		ID:   1,
		Name: "SatoshiBridge",
	}
	bridge := accessory.NewBridge(info)

	for _, v := range cfg.Satoshi.Siri.Devices {
		switch v.Type {
		case "switch":
			ac := InitSwitch(v, InitSwitchOption{
				LircdPath: cfg.Lircd.LircdPath,
			})
			if ac != nil {
				acArr = append(acArr, ac.Accessory)
			}
			break
		case "lightbulb":
			ac := InitLightbulb(v, InitLightbulbOption{
				LircdPath: cfg.Lircd.LircdPath,
			})
			if ac != nil {
				acArr = append(acArr, ac.Accessory)
			}
			break
		}
	}

	// configure the ip transport
	config := hc.Config{
		Pin: cfg.Satoshi.Siri.PinCode,
	}
	t, err := hc.NewIPTransport(config, bridge.Accessory, acArr...)
	if err != nil {
		log.Panic(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
		os.Exit(1)
	})

	t.Start()
}
