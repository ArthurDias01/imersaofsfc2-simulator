package main

import (
	"fmt"
	"github.com/ArthurDias01/imersaofsfc2-simulator/application/route"
	"github.com/ArthurDias01/imersaofsfc2-simulator/tree/main/infra/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal
	}
}Art13393679thur!


func main() {
 producer := kafka.NewKafkaProducer()
 kafka.Publish("Ola", "readtest", producer)
}

// route := route.Route{
// 	ID:        "1",
// 	ClientID:  "1",
// }
// route.LoadPositions()

// stringjson,_ := route.ExportJsonPositions()
// fmt.Println(stringjson[1])
