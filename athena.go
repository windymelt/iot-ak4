package iotak4

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
)

func InsertAthena(operation string) {
	now := time.Now()
	iso8601 := now.UTC().Format("YYYY-MM-DDTHH:MM:SSZ")
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
	fmt.Println(query_wait_result.QueryExecution.Status.StateChangeReason)
	if err != nil {
		fmt.Println(err)
		return
	}
}