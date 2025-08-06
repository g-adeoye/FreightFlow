package main

import (
	"context"
	"encoding/json"
	"kafka_producer/producer"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	kafka_broker := "host.docker.internal:9092"
	log.Printf("Dialing: %s", kafka_broker)
	writer := kafka.Writer{
		Addr: kafka.TCP(kafka_broker),
		Topic: "truck-data",
		Balancer: &kafka.CRC32Balancer{},
		Logger: kafka.LoggerFunc(producer.Logf),
		ErrorLogger: kafka.LoggerFunc(producer.Logf),
		AllowAutoTopicCreation: true,

	}

	trucks := producer.MakeTrucks(5)

	for {
		for _, truck := range trucks {
			truck.Update()
			msg := truck.ToMessage()
			jsonBytes, _ := json.Marshal(msg)
			log.Printf("%s",jsonBytes)
			err := writer.WriteMessages(
				context.Background(),
				kafka.Message{
					Key: []byte(truck.TruckID),
					Value: jsonBytes,
				},
			)

			if err != nil {
				log.Fatalf("Error writing message to kafka broker, %s", err)
			}
		}

		time.Sleep(time.Second * time.Duration(300))
	}

}