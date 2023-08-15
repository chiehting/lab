package pkg

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sns"
)

// AwsSendEmail service is use aws sdk send email from particular domain name.
func AwsSendEmail(sess client.ConfigProvider, recipient string) {
	const CharSet = "UTF-8"
	// Create SES service client
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data: aws.String("<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
						"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
						"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String("This email was sent with Amazon SES using the AWS SDK for Go."),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String("Amazon SES Test (AWS SDK for Go)"),
			},
		},
		Source: aws.String("support@example.com"),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return
	}

	fmt.Println("Email Sent to address: " + recipient)
	fmt.Println(result)
}

// AwsSendNotification service is use aws sdk send notification to firebase.
func AwsSendNotification(sess client.ConfigProvider) {
	const (
		topicArn       = "arn:aws:sns:ap-southeast-1::"
		applicationARN = "arn:aws:sns:ap-southeast-1::app/GCM/abcd"
		token          = "d_XLROUQRh"
		userData       = ""
	)

	svc := sns.New(sess)

	params := &sns.CreatePlatformEndpointInput{
		PlatformApplicationArn: aws.String(applicationARN),
		Token:                  aws.String(token),
		Attributes: map[string]*string{
			"Token":          aws.String(token),
			"CustomUserData": aws.String(userData),
			"Enabled":        aws.String("true"),
		},
		CustomUserData: aws.String(userData),
	}

	endPoints, _ := svc.CreatePlatformEndpoint(params)
	svc.Subscribe(&sns.SubscribeInput{
		Endpoint:              endPoints.EndpointArn,
		Protocol:              aws.String("application"),
		ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
		TopicArn:              aws.String(topicArn),
	})

	publishResult, err := svc.Publish(&sns.PublishInput{
		Message:          aws.String(`{"default":"Sample fallback message","GCM":"{\"notification\":{\"body\":\"Sample message for Android endpoints\",\"title\":\"TitleTest\" }}"}`),
		MessageStructure: aws.String("json"),
		TopicArn:         aws.String(topicArn),
	})

	fmt.Println(err)
	fmt.Println(*publishResult.MessageId)
}

func DeleteRegisteredEndpoint(sess client.ConfigProvider, endpointArn string) (err error) {
	deleteEndpointInput := &sns.DeleteEndpointInput{
		EndpointArn: aws.String(endpointArn),
	}

	svc := sns.New(sess)
	if _, err = svc.DeleteEndpoint(deleteEndpointInput); err != nil {
		fmt.Println(deleteEndpointInput)
	}

	fmt.Println(err)
	return
}

// AwsSendSms service is use aws sdk send sms.
func AwsSendSms(sess client.ConfigProvider, phoneNumber string) {
	svc := sns.New(sess)
	result, err := svc.Publish(&sns.PublishInput{
		Message:     aws.String("Amazon SMS Test (AWS SDK for Go)"),
		PhoneNumber: aws.String(phoneNumber),
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result.MessageId)
}
