package main

import (
	"encoding/json"
	"fmt"

	ilorest "github.com/sgkul2000/hpe-cty"
)

func PrettyPrint(data interface{}) {
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal JSON!")
		return
	}
	fmt.Println("\n" + string(dataJSON) + "\n\n")
}

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
	fmt.Println("Initializing client...")
	PrettyPrint(client)

	fmt.Println("Getting root object...")
	response, err := client.Get("/", map[string]string{}, map[string]string{}, 10)
	if err != nil {
		fmt.Println(err)
	}
	PrettyPrint(response)

	fmt.Println("Getting operator role info...")
	operatorInfo, err := client.Get("/AccountService/Roles/Operator", map[string]string{}, map[string]string{}, 10)
	if err != nil {
		fmt.Println(err)
	}
	PrettyPrint(operatorInfo)

	fmt.Println("Patching operator role description...")
	patchResp, err := client.Patch("/AccountService/Roles/Operator", map[string]interface{}{"Description": "this is the new patched description"}, map[string]string{}, map[string]string{}, 10)
	if err != nil {
		fmt.Println(err)
	}

	PrettyPrint(patchResp)

	fmt.Println("Getting operator role info again...")
	newOperatorInfo, err := client.Get("/AccountService/Roles/Operator", map[string]string{}, map[string]string{}, 10)
	if err != nil {
		fmt.Println(err)
	}
	PrettyPrint(newOperatorInfo)
	fmt.Println("Operator role description: " + operatorInfo["Description"].(string) + " => " + newOperatorInfo["Description"].(string))
}
