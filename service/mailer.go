package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendMail(email, address, productName string) {
	url := "http://localhost:8080/mailer"
	mailBody := map[string]string{
		"buyer_email":   email,
		"buyer_address": address,
		"product_name":  productName,
	}

	mailBodyMarshal, _ := json.Marshal(mailBody)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(mailBodyMarshal))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	fmt.Println("response status", response.Status)
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body), "response body")

	defer response.Body.Close()
}
