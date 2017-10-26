package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create a SQS service client.
    // `sess`はセッション
    svc := sqs.New(sess)

    result, err := svc.ListQueues(nil)
    if err != nil {
        fmt.Println("Error", err)
        return
    }
    fmt.Println("Success")
    // As these are pointers, printing them out directly would not be useful.
    
    for i, urls := range result.QueueUrls {
        // Avoid dereferencing a nil pointer.
        if urls == nil {
            continue
        }
        fmt.Printf("%d: %s\n", i, *urls)
    }
}
