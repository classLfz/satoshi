package config

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Satoshi *SatoshiConfig `yaml:"satoshi"`
	Lircd   *LircdConfig   `yaml:"lircd"`
}

// -- lircd配置 --
type LircdConfig struct {
	LircdPath string `yaml:"lircd_path"`
}

// -- Satoshi配置 --
type SatoshiConfig struct {
	Http *HttpConfig `yaml:"http"`
	Siri *SiriConfig `yaml:"siri"`
}

type SiriConfig struct {
	PinCode string       `yaml:"pin_code"`
	Devices []SiriDevice `yaml:"devices"`
}

// -- Siri设备配置 --
type SiriSwitchConfig struct {
	OnPin          uint8  `yaml:"on_pin"`
	UpdateInterval uint8  `yaml:"update_interval"`
	OnLircdCmd     string `yaml:"on_lircd_cmd"`
	OffLircdCmd    string `yaml:"off_lircd_cmd"`
}

type SiriLightbulbConfig struct {
	OnPin          uint8  `yaml:"on_pin"`
	UpdateInterval uint8  `yaml:"update_interval"`
	OnLircdCmd     string `yaml:"on_lircd_cmd"`
	OffLircdCmd    string `yaml:"off_lircd_cmd"`
}

type SiriDevice struct {
	ID              uint64               `yaml:"id"`
	Type            string               `yaml:"type"`
	Name            string               `yaml:"name"`
	SwitchConfig    *SiriSwitchConfig    `yaml:"switch_config"`
	LightbulbConfig *SiriLightbulbConfig `yaml:"lightbulb_config"`
}

// -- http接口配置 --
type HttpConfig struct {
	Port    string       `yaml:"port"`
	Devices []HttpDevice `yaml:"devices"`
}

type HttpDevice struct {
	ID           string        `yaml:"id"`
	Type         string        `yaml:"type"`
	Name         string        `yalm:"name"`
	SwitchConfig *SwitchConfig `yaml:"switch_config"`
}

type SwitchConfig struct {
	OnPin       uint8  `yaml:"on_pin"`
	OnLircdCmd  string `yaml:"on_lircd_cmd"`
	OffLircdCmd string `yaml:"off_lircd_cmd"`
}

func (c *Config) Dump(p string) error {
	var err error

	if p, err = homedir.Expand(p); err != nil {
		return err
	}

	if err = os.MkdirAll(path.Dir(p), 0755); err != nil {
		return err
	}

	buf, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(p, buf, 0644); err != nil {
		return err
	}

	return nil
}

func Load(p string) (config *Config, loaded bool, err error) {
	if p, err = homedir.Expand(p); err != nil {
		return nil, false, err
	}

	config = NewDefaultConfig()

	buf, err := ioutil.ReadFile(p)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, false, err
		}
		return config, false, nil
	}

	if err = yaml.Unmarshal(buf, config); err != nil {
		return nil, false, err
	}

	return config, true, nil
}

func NewDefaultConfig() *Config {
	return &Config{
		Lircd: &LircdConfig{
			LircdPath: "/var/run/lirc/lircd",
		},
		Satoshi: &SatoshiConfig{
			Http: &HttpConfig{
				Port: "8234",
				Devices: []HttpDevice{
					{
						ID:   "",
						Type: "switch",
						Name: "device_name",
						SwitchConfig: &SwitchConfig{
							OnPin: 0,
						},
					},
				},
			},
			Siri: &SiriConfig{
				PinCode: "00102003",
				Devices: []SiriDevice{
					{
						ID:   2,
						Type: "switch",
						Name: "Switch001",
						SwitchConfig: &SiriSwitchConfig{
							OnPin:          0,
							UpdateInterval: 0,
						},
						LightbulbConfig: &SiriLightbulbConfig{
							OnPin:          0,
							UpdateInterval: 0,
						},
					},
				},
			},
		},
	}
}
