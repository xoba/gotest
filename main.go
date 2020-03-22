package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	s, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	check(err)
	svc := s3.New(s)
	o, err := svc.ListBuckets(&s3.ListBucketsInput{})
	check(err)
	fmt.Println(o)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
