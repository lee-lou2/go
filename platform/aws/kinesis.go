package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"os"
)

// PutKinesisRecord Kinesis에 데이터를 저장
func PutKinesisRecord(data []byte, configs ...Configs) error {
	cfg, err := defaultConfig(configs...)
	if err != nil {
		return err
	}

	// Kinesis 클라이언트 생성
	kc := kinesis.NewFromConfig(cfg)

	// Kinesis PutRecordInput 생성
	putReq := &kinesis.PutRecordInput{
		Data:         data,
		PartitionKey: aws.String(os.Getenv("KINESIS_STREAM_NAME")),
		StreamName:   aws.String(os.Getenv("KINESIS_PARTITION_KEY")),
	}

	// 데이터를 Kinesis 스트림에 저장
	_, err = kc.PutRecord(context.TODO(), putReq)
	if err != nil {
		return err
	}
	return nil
}
