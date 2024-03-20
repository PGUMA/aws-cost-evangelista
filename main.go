package main

import (
	"fmt"

	"github.com/PGUMA/aws-cost-evangelista/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	fmt.Print("start lambda!")
	lambda.Start(handler.HandleRequest)
	fmt.Print("end lambda!")
}
