	package main

	import (
		"context"
		"fmt"
		"os"
		"github.com/aws/aws-lambda-go/lambda"
	)


	func Handler(ctx context.Context) (string, error) {
		return "Hello from Lambda!", nil
	}

	func test() {

		if os.Getenv("AWS_LAMBDA_RUNTIME_API") == "" {
	

		result, err := Handler(context.Background())

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(result)
	} else {
		lambda.Start(Handler)
	}
	}