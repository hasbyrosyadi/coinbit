package main

import (
	"coinbit-test/config"
	pb "coinbit-test/gen/proto"
	"coinbit-test/server/usecase"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Kafka struct {
		Brokers   []string `yaml:"brokers"`
		Group     string   `yaml:"group"`
		Stream    string   `yaml:"stream"`
		Redis     string   `yaml:"redis"`
		Namespace string   `yaml:"namespace"`
	} `yaml:"kafka"`
}

var (
	filename = flag.String("config", "config.yaml", "path to config file")
)

func main() {
	// flag.Parse()

	// conf, err := readConfig(*filename)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("data", conf.Kafka)

	// // consuming
	// go func() {
	// 	err := kafka.Consume(new(nopPublisher), conf.Kafka.Brokers, conf.Kafka.Group,
	// 		conf.Kafka.Stream, conf.Kafka.Redis, conf.Kafka.Namespace)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// // producing
	// producer, err := kafka.NewProducer(conf.Kafka.Brokers, conf.Kafka.Stream)
	// if err != nil {
	// 	log.Fatal("disni", err)
	// }
	// for {
	// 	// for testing
	// 	event := &config.Event{
	// 		WalletID: 1,
	// 		Amount:   10000,
	// 	}
	// 	fmt.Printf("emit -> key:`%v` ->event: `%v`\n", event.WalletID, event)
	// 	strWalletID := strconv.Itoa(event.WalletID)
	// 	err = producer.Emit(strWalletID, event)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	time.Sleep(5 * time.Second)
	// }

	mux := runtime.NewServeMux()

	go func() {
		var server = usecase.DepositServiceServer{}
		pb.RegisterDepositServiceHandlerServer(context.Background(), mux, &server)

	}()

	log.Fatal(http.ListenAndServe("localhost:8081", mux))
}

func readConfig(filename string) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	conf := new(Config)
	err = yaml.Unmarshal(b, conf)
	return conf, err
}

type nopPublisher struct{}

func (p *nopPublisher) Publish(ctx context.Context, key string, event *config.Event) error {
	fmt.Printf("published ->key: `%v` ->event: `%v`\n", key, event)
	return nil
}

func (p *nopPublisher) Close() error {
	return nil
}
