package main

import (
	"fmt"
	"io/ioutil"
	"os"

	httpDeliver "github.com/shadialtarsha/config-service/http"
	"github.com/shadialtarsha/config-service/service"
)

func main() {
	config := service.Config{}

	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	data = []byte(os.ExpandEnv(string(data)))
	err = config.SetFromBytes(data)
	if err != nil {
		panic(err)
	}

	server := httpDeliver.NewServer(&config)

	fmt.Println("Starting http server")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err, "HTTP connection terminated")
	}

}
