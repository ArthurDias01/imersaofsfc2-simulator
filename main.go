package main

import (
	"fmt"
	"github.com/ArthurDias01/imersaofsfc2-sumulator/application/route"
	"github.com/ArthurDias01/imersaofsfc2-simulator/infra/kafka"

	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal
	}
}


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
