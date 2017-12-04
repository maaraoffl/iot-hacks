package main

import (
	"encoding/json"
	"fmt"
	"github.com/danward79/go.wemo"
	"log"
	"strconv"
	"time"
)

func hello() {

	listenerAddress := "192.168.1.3:6767"
	timeout := 300

	api, _ := wemo.NewByInterface("en0")

	devices, _ := api.DiscoverAll(3 * time.Second)

	subscriptions := make(map[string]*wemo.SubscriptionInfo)

	for _, device := range devices {
		_, err := device.ManageSubscription(listenerAddress, timeout, subscriptions)
		if err != 200 {
			log.Println("Initial Error Subscribing: ", err)
		}
	}

	cs := make(chan wemo.SubscriptionEvent)

	go wemo.Listener(listenerAddress, cs)

	for m := range cs {
		if _, ok := subscriptions[m.Sid]; ok {
			data, err := json.Marshal(m)
			if err != nil {
				fmt.Println(err)
			}

			log.Println("state change event info: ", string(data))

			binaryState, _ := strconv.Atoi(m.Deviceevent.BinaryState)
			InsertState("192.168.1.8", "Wemo Insight Proj1", binaryState)

			// fmt.Println(string(data))
			// log.Println(subscriptions)
			// subscriptions[m.Sid].State = m.State
			// log.Println("---Current state: ", m.Deviceevent.BinaryState)
			// log.Println("---Subscriber Event: ", subscriptions[m.Sid])
		} else {
			log.Println("Does'nt exist, ", m.Sid)
		}
	}

}

func main() {
	InsertState("192.168.1.7", "Living room light", 1)

}
