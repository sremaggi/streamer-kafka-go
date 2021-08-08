package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"syscall"
)


var (
	BROKER = "127.0.0.1:9092"
	INPUT_TOPIC = []string{"productos-topic"}
	OUTPUT_TOPIC = "portabilidad-topic"
	CONSUMER_GROUP = "consumer_group"

)




func main() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": BROKER,
		// Avoid connecting to IPv6 brokers:
		// This is needed for the ErrAllBrokersDown show-case below
		// when using localhost brokers on OSX, since the OSX resolver
		// will return the IPv6 addresses first.
		// You typically don't need to specify this configuration property.
		"broker.address.family": "v4",
		"group.id":              CONSUMER_GROUP,
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "latest"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	err = c.SubscribeTopics(INPUT_TOPIC, nil)
it
	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",e.TopicPartition, string(e.Value))
				//PRODUCER
				p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": BROKER})
				if err != nil {
					panic(err)
				}
				fmt.Printf("Created Producer %v\n", p)
				defer p.Close()
				//b, err := json.Marshal(e.Value)
				//if err != nil {
				//	fmt.Printf("Error: %s", err)

				//}
				deliveryChan := make(chan kafka.Event)

				err = p.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &OUTPUT_TOPIC, Partition: kafka.PartitionAny},
					Value:          []byte(string(e.Value)),
					Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
				}, deliveryChan)

				ch := <-deliveryChan
				m := ch.(*kafka.Message)

				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}

				close(deliveryChan)


				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()
}