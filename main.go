package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PhiBrandon/packagemanagement/nannyhelper"
	"github.com/PhiBrandon/packagemanagement/utilities"
)

func main() {
	nannyhelper.GetHelper()
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	utilities.CheckError(err)

	resp, err := client.Do(req)
	utilities.CheckError(err)

	if resp.StatusCode == 200 {
		fmt.Println("Success.")
	} else {
		fmt.Println("Failure.")
	}
}
