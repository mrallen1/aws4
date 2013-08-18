package aws4

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func ExampleSign() {
	data := strings.NewReader("{}")
	r, _ := http.NewRequest("POST", "https://dynamodb.us-east-1.amazonaws.com/", data)
	r.Header.Set("Host", "dynamodb.us-east-1.amazonaws.com")
	r.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))
	r.Header.Set("Content-Type", "application/x-amz-json-1.0")
	r.Header.Set("X-Amz-Target", "DynamoDB_20111205.ListTables")

	tk := &Keys{
		AccessKey: os.Getenv("AWS_ACCESS_KEY"),
		SecretKey: os.Getenv("AWS_SECRET_KEY"),
	}

	sv := &Service{
		Name:   "dynamodb",
		Region: "us-east-1",
	}

	if err := sv.Sign(tk, r); err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)
	// Output:
	// 200
}

func ExampleSignGlacier() {
	r, _ := http.NewRequest("GET", "https://glacier.us-east-1.amazonaws.com/-/vaults", nil)
	r.Header.Set("Host", "glacier.us-east-1.amazonaws.com")
	r.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))
	r.Header.Set("X-Amz-Glacier-Version", "2012-06-01")

	tk := &Keys{
		AccessKey: os.Getenv("AWS_ACCESS_KEY"),
		SecretKey: os.Getenv("AWS_SECRET_KEY"),
	}

	sv := &Service{
		Name:   "glacier",
		Region: "us-east-1",
	}

	if err := sv.Sign(tk, r); err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)

	// Output:
	// 200
}
