package ps3

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"os"
)

type Ps3_s3 struct {
	bucket      string
	s3          *awss3.Client
	s3_uploader *manager.Uploader
}

func (p *Ps3_s3) Put(key string, data []byte) error {
	_, err := p.s3_uploader.Upload(context.TODO(), &awss3.PutObjectInput{
		Bucket: aws.String(p.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		return err
	}
	return nil
}

func (p *Ps3_s3) Get(key string) ([]byte, error) {
	output, err := p.s3.GetObject(context.TODO(), &awss3.GetObjectInput{
		Bucket: aws.String(p.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return io.ReadAll(output.Body)
}

func Ps3_s3_new(region string, bucket string) (*Ps3_s3, error) {
	panic_msg_env := "env variable not set: "
	if os.Getenv(aws_access_key_id) == "" {
		panic(panic_msg_env + aws_access_key_id)
	}
	if os.Getenv(aws_secret_access_key) == "" {
		panic(panic_msg_env + aws_secret_access_key)
	}

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, err
	}

	s3 := awss3.NewFromConfig(cfg)
	s3_uploader := manager.NewUploader(s3)

	return &Ps3_s3{
		bucket:      bucket,
		s3:          s3,
		s3_uploader: s3_uploader,
	}, nil
}
