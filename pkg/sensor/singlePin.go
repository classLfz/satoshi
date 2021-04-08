package sensor

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// CallSinglePinOptions 开关控制函数参数结构
type CallSinglePinOptions struct {
	PinNum           uint8
	Read             bool
	Flashing         bool
	FlashingInterval time.Duration
	FlashingCount    int
	Toggle           bool
	State            bool
}

// CallSinglePin 单个GPIO口控制函数
func CallSinglePin(options CallSinglePinOptions) (state rpio.State, err error) {
	const CurrentAPI = "Sensors Switch Call"
	openErr := rpio.Open()

	if openErr != nil {
		log.Panicf("%s rpio.Open error.\n", CurrentAPI)
		return rpio.Low, openErr
	}

	PinNum := options.PinNum
	Read := options.Read
	Toggle := options.Toggle
	State := options.State
	Flashing := options.Flashing
	FlashingInterval := options.FlashingInterval
	FlashingCount := options.FlashingCount

	// 仅读取状态
	if Read {
		return readPinState(PinNum, false)
	}

	// 设置闪烁
	if Flashing && FlashingInterval > 0 && FlashingCount > 0 {
		setFlashing(PinNum, FlashingInterval, FlashingCount, false)
		return readPinState(PinNum, false)
	}

	pin := rpio.Pin(PinNum)
	pin.Output()

	// 切换状态
	if Toggle {
		pin.Toggle()
		return readPinState(PinNum, false)
	}

	// 设置开关
	var newState rpio.State
	if State {
		newState = rpio.High
	} else {
		newState = rpio.Low
	}
	pin.Write(newState)

	return readPinState(PinNum, false)
}

// setFlashing 设置闪烁
func setFlashing(pinNum uint8, interval time.Duration, count int, open bool) {
	const CurrentAPI = "Sensors Laser Call setFlashing"
	if open {
		openErr := rpio.Open()

		if openErr != nil {
			log.Panicf("%s rpio.Open error.\n", CurrentAPI)
			return
		}
	}

	go func() {
		pin := rpio.Pin(pinNum)
		timer := time.NewTimer(interval)
		defer timer.Stop()
		i := 0
		for {
			<-timer.C
			pin.Toggle()
			if i < count {
				i++
				timer.Reset(interval)
			} else {
				pin.Low()
			}
		}
	}()
}
