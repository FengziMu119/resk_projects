package main

import (
	"encoding/json"
	"fmt"
	"resk_projects/services"
)

func main() {

	d, e := json.Marshal(&services.AccountTransferDTO{})
	fmt.Println(e)
	fmt.Println(string(d))
}
