package main

import (
	"encoding/json"
	"fmt"
	"github.com/danward79/go.wemo"
	"log"
	"time"
)

func main() {

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
			fmt.Println(string(data))

			// subscriptions[m.Sid].State = m.State
			log.Println("---Current state: ", m.Deviceevent.BinaryState)
			log.Println("---Subscriber Event: ", subscriptions[m.Sid])
		} else {
			log.Println("Does'nt exist, ", m.Sid)
		}
	}

}
