package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
)

func InsertAthena(operation string) {
	now := time.Now()
	iso8601 := now.UTC().Format(time.RFC3339) // a profile of ISO8601
	year := now.Year()
	month := now.Month()

	mySession := session.Must(session.NewSession())
	client := athena.New(mySession)
	database := "kintairecord"
	output_location := "s3://athenaresult.3qe.us"
	catalog := "AwsDataCatalog"
	wg := "primary"

	query_string := fmt.Sprintf(`
	INSERT INTO kintaitable (created_at, operation, year, month)
	VALUES (
		CAST(from_iso8601_timestamp('%s') AS TIMESTAMP),
		'%s',
		CAST(%d AS SMALLINT),
		CAST(%d AS TINYINT)
	);
	`, iso8601, operation, year, month)

	query_result, err := client.StartQueryExecution(&athena.StartQueryExecutionInput{
		QueryString: &query_string,
		QueryExecutionContext: &athena.QueryExecutionContext{
			Database: &database,
			Catalog:  &catalog,
		},
		ResultConfiguration: &athena.ResultConfiguration{
			OutputLocation: &output_location,
		},
		WorkGroup: &wg,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(1 * time.Second)
	query_wait_result, err := client.GetQueryExecution(&athena.GetQueryExecutionInput{
		QueryExecutionId: query_result.QueryExecutionId,
	})
	if state := query_wait_result.QueryExecution.Status.State; state != nil {
		fmt.Printf("State: %s\n", *state)
	}
	if reason := query_wait_result.QueryExecution.Status.StateChangeReason; reason != nil {
		fmt.Printf("Reason: %s\n", *reason)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}
