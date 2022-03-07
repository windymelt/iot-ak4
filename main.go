package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

const TYPE_ON = 11
const TYPE_OFF = 12

func HandleRequest(ctx context.Context, event interface{}) (map[string]interface{}, error) {
	click_type := event.(map[string]interface{})["deviceEvent"].(map[string]interface{})["buttonClicked"].(map[string]interface{})["clickType"].(string)
	fmt.Println(click_type)

	var error bool = false
	switch click_type {
	case "SINGLE":
		InsertAthena("in")
		error = doPunch(TYPE_ON)
	case "DOUBLE":
		InsertAthena("out")
		error = doPunch(TYPE_OFF)
	case "LONG":
	default:
	}

	sc := 200
	if error {
		sc = 500
	}

	j, err := json.Marshal(click_type)
	if err != nil {
		panic(err)
	}

	return map[string]interface{}{
		"statusCode": sc,
		"body":       j,
	}, nil
}

func prepareUrl(endpoint string, query map[string]string) *url.URL {
	ak_coop_id := os.Getenv("AK_COOP_ID")
	u, err := url.Parse("https://atnd.ak4.jp/api/cooperation/" + ak_coop_id + endpoint)
	if err != nil {
		panic(err)
	}

	base_query := map[string]string{
		"token": os.Getenv("AK_TOKEN"),
	}

	final_query := map[string]string{}

	// merge base_query and query
	for k, v := range base_query {
		final_query[k] = v
	}
	for k, v := range query {
		final_query[k] = v
	}

	q := u.Query()
	for k, v := range final_query {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	return u
}

func preparePunchBody(punch_type int, token string) []byte {
	body, err := json.Marshal(map[string]interface{}{
		"type":  punch_type,
		"token": token,
	})
	if err != nil {
		panic(err)
	}
	return body
}

func doPunch(punch_type int) bool {
	u := prepareUrl("/stamps", map[string]string{})
	body := preparePunchBody(punch_type, os.Getenv("AK_TOKEN"))

	res, err := http.Post(u.String(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	return res.StatusCode != http.StatusOK
}

func main() {
	lambda.Start(HandleRequest)
}
