package main

import (
	"aws_sdk_example/pkg"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	// Warning: Do not embed credentials inside an application. Use this method only for testing purposes. please ~ hearts man.
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIA4LSRNN5GQYZAUHES", "j/WukQxgTRr9S6sM6EyBh3+ltGCoCAksWvKyfIfW", ""),
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Aws connected:", *sess.Config.Region)

	//pkg.AwsSendEmail(sess, "unicatpm@gmail.com")
	//pkg.AwsSendNotification(sess)
	//pkg.DeleteRegisteredEndpoint(sess, "arn:aws:sns:ap-southeast-1:849500532557:endpoint/GCM/vinbrace/f257e549-7012-31d4-8be5-c6716b4bbaa0")
	pkg.AwsSendSms(sess, "+886979421551")
}
