package sensor

import (
	"fmt"
	"log"

	"github.com/d2r2/go-dht"
	"github.com/stianeikeland/go-rpio"
)

// Dht11 温湿度传感器控制函数
func Dht11() {
	const CurrentAPI = "Sensors Dht11 Call"
	openErr := rpio.Open()

	if openErr != nil {
		log.Panicf("%s rpio.Open error.\n", CurrentAPI)
		return
	}
	log.Printf("%s rpio is ready.\n", CurrentAPI)

	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 4, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
}
