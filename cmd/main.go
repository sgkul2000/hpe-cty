package main

import (
	"fmt"

	ilorest "github.com/sgkul2000/hpe-cty"
)

func main() {
	client := ilorest.NewRFClient((ilorest.RFClient{
		BaseUrl:       "http://localhost:8000/redfish/v1",
		UserName:      "",
		Password:      "",
		BiosPassword:  "",
		SessionKey:    "",
		DefaultPrefix: "",
		Timeout:       0,
		Proxy:         "",
	}))
	fmt.Println(client)

	response, err := client.Get("/UpdateService", map[string]string{}, map[string]string{}, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
