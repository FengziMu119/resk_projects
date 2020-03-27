package main

import (
	"encoding/json"
	"fmt"
	"resk_projects/services"
)

func main() {
	data, _ := json.Marshal(&services.RedEnvelopeSendingDTO{})
	fmt.Println(string(data))
}
