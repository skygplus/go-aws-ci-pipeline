package main

import (
	"fmt"

	"go-aws-ci-pipeline/internal/greeting"
)

func main() {
	fmt.Println(greeting.Message("AWS DevOps"))
}
