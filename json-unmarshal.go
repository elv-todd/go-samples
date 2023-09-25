package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type EmailResponse struct {
	EmailVerified bool `json:"email_verified"`
	Data          struct {
		Id        string    `json:"id"`
		Status    string    `json:"status"`
		Type      string    `json:"type"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"data"`
}

func main() {
	var resp EmailResponse
	js := `{"email_verified":false,"data":{"status":"pending","type":"verification_email","created_at":"2023-05-09T21:21:22.528Z","id":"job_TY9ULWy4zP8HCjHl"}}`
	json.Unmarshal([]byte(js), &resp)
	fmt.Printf("%+v", resp)
}
