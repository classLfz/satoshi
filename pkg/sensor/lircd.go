package sensor

import (
	"log"
	"time"

	"github.com/chbmuc/lirc"
)

type CallLircdOptions struct {
	LircdSocketPath string
	SendStr         string
	SendLongStr     string
	SendLongDelay   time.Duration
}

// CallLircd lircd控制函数
func CallLircd(options CallLircdOptions) (err error) {
	const CurrentAPI = "Lircd Call"

	LircdSocketPath := options.LircdSocketPath

	// initialize
	ir, err := lirc.Init(LircdSocketPath)
	if err != nil {
		log.Panic(CurrentAPI, " lirc.Init err:\n", err)
		return err
	}

	if options.SendStr != "" {
		err = ir.Send(options.SendStr)
		if err != nil {
			log.Panic(CurrentAPI, " ir.Send err:\n", err)
			return err
		}
		// 防止两次发送过于接近
		if options.SendLongStr != "" {
			time.Sleep(time.Second)
		}
	}

	if options.SendLongStr != "" {
		err = ir.SendLong(options.SendLongStr, options.SendLongDelay)
		if err != nil {
			log.Panic(CurrentAPI, " ir.SendLong err:\n", err)
			return err
		}
	}

	return nil
}
