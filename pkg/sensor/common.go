package sensor

import (
	"log"

	"github.com/stianeikeland/go-rpio"
)

// readPinState 读取端口状态
func readPinState(pinNum uint8, open bool) (state rpio.State, err error) {
	const CurrentAPI = "Sensors Laser Call readPinState"
	if open {
		openErr := rpio.Open()

		if openErr != nil {
			log.Panicf("%s rpio.Open error.\n", CurrentAPI)
			return rpio.Low, openErr
		}
	}

	pin := rpio.Pin(pinNum)
	readRes := pin.Read()

	return readRes, nil
}
