package main

import (
    "fmt"
    
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)


func main() {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    svc := sqs.New(sess) // `sess`はセッション
    
	//作ったキューのURLをここに代入
    qURL := "https://sqs.ap-northeast-1.amazonaws.com/047827674873/Go_Queue_Test2"
    
    // message送信
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		// メッセージ属性
	        DelaySeconds: aws.Int64(10),
	        MessageAttributes: map[string]*sqs.MessageAttributeValue{
	            "Title": &sqs.MessageAttributeValue{
	                DataType:    aws.String("String"),
	                StringValue: aws.String("The Whistler"),
	            },
	            "Author": &sqs.MessageAttributeValue{
	                DataType:    aws.String("String"),
	                StringValue: aws.String("John Grisham"),
	            },
	            "WeeksOn": &sqs.MessageAttributeValue{
	                DataType:    aws.String("Number"),
	                StringValue: aws.String("6"),
	            },
	        },
	        // メッセージ本文
	        MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 12/11/2016."),
	        QueueUrl:    &qURL,
	    })
	
	    if err != nil {
	        fmt.Println("Error", err)
	        return
	    }
		// 成功したらサクセスメッセージ
	    fmt.Println("Success", *result.MessageId)
}