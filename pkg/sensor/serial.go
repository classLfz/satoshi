package sensor

import (
	"log"

	"go.bug.st/serial"
)

type CallSerialOptions struct {
	// mode
	PortName string
	BaudRate int
	DataBits int
	StopBits serial.StopBits

	// write
	Write bool
	Data  []byte

	// read
	Read     bool
	BuffBits int
}

// CallSerial 调用串口通信函数
func CallSerial(options CallSerialOptions) (readChan chan []byte, err error) {
	const CurrentAPI = "Sensors Serial Call"

	PortName := options.PortName
	BaudRate := options.BaudRate
	DataBits := options.DataBits
	StopBits := options.StopBits
	Write := options.Write
	Data := options.Data
	Read := options.Read
	BuffBits := options.BuffBits

	if !Write || !Read {
		return nil, nil
	}

	mode := &serial.Mode{
		BaudRate: BaudRate,
		DataBits: DataBits,
		StopBits: StopBits,
	}

	port, openErr := serial.Open(PortName, mode)
	if openErr != nil {
		log.Fatal(CurrentAPI, " serial.Open error\n", openErr)
		return nil, openErr
	}

	if Write {
		n, writeErr := port.Write(Data)
		if writeErr != nil {
			log.Fatal(CurrentAPI, " port.Write error\n", writeErr)
			return nil, writeErr
		}

		log.Printf("%v Sent %v bytes", CurrentAPI, n)
	}

	if Read {
		for {
			buff := make([]byte, BuffBits)
			_, readErr := port.Read(buff)
			if readErr != nil {
				log.Fatal(CurrentAPI, " port.Read error\n", readErr)
			} else {
				readChan <- buff
			}
		}
	}

	return nil, nil
}
