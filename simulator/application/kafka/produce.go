package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	route2 "github.com/ThalesGabriel/code-delivery/simulator/application/route"
	"github.com/ThalesGabriel/code-delivery/simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// Produce is responsible to publish the positions of each request
// Example of a json request:
//{"clientId":"1","routeId":"1"}
//{"clientId":"2","routeId":"2"}
//{"clientId":"3","routeId":"3"}
func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	fmt.Println(msg.Value)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		fmt.Println(p)
		time.Sleep(time.Millisecond * 500)
	}
}
