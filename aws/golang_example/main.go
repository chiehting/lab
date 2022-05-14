package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	// Warning: Do not embed credentials inside an application. Use this method only for testing purposes. please ~ hearts man.
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Aws connected:", *sess.Config.Region)

	pkg.AwsSendEmail(sess, "test@gmail.com")
	// pkg.AwsSendNotification(sess)
	// pkg.AwsSendSms(sess, "+886111111111")
}
