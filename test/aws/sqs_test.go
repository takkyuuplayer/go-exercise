package aws_test

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func TestSqs(t *testing.T) {
	svc := sqs.New(getSession(t))

	// Create a queue
	created, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String("tp-test"),
		Attributes: map[string]*string{
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", created)

	// Get a queue
	queue, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String("tp-test"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if *created.QueueUrl != *queue.QueueUrl {
		t.Error("Fetched queue url must be the same as created one")
	}
	t.Logf("%v", queue)

	// Enqueue
	enqueued, _ := svc.SendMessage(&sqs.SendMessageInput{
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
		MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 12/11/2016."),
		QueueUrl:    queue.QueueUrl,
	})
	t.Logf("%v", enqueued)

	// Dequeue
	dequeued, _ := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queue.QueueUrl,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(60),
	})
	if *enqueued.MD5OfMessageBody != *dequeued.Messages[0].MD5OfBody {
		t.Error("dequeued MessageId must be the same as enqueued one")
	}
	t.Logf("%v", dequeued)

	// Delete the message
	deleted, _ := svc.DeleteMessageBatch(&sqs.DeleteMessageBatchInput{
		QueueUrl: queue.QueueUrl,
		Entries: []*sqs.DeleteMessageBatchRequestEntry{
			{Id: dequeued.Messages[0].MessageId, ReceiptHandle: dequeued.Messages[0].ReceiptHandle},
		},
	})
	t.Logf("%v", deleted)
}

func getSession(t *testing.T) *session.Session {
	t.Helper()

	if endpoint, ok := os.LookupEnv("AWS_ENDPOINT"); ok && endpoint != "" {
		return session.Must(session.NewSession(&aws.Config{
			Endpoint: aws.String(endpoint),
			Region:   aws.String("us-east-1"),
		}))
	}

	return session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1"),
	}))
}
