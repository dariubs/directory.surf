package services

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var S3Client *s3.Client

func InitR2() {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           os.Getenv("CLOUDFLARE_R2_ENDPOINT"),
			SigningRegion: "auto",
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("auto"),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("CLOUDFLARE_R2_ACCESS_KEY"),
			os.Getenv("CLOUDFLARE_R2_SECRET_KEY"),
			"",
		)),
	)

	if err != nil {
		panic("unable to load R2 config: " + err.Error())
	}

	S3Client = s3.NewFromConfig(cfg)
}

func GetPresignedUploadURL(filename string) (string, error) {
	psClient := s3.NewPresignClient(S3Client)

	req, err := psClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("CLOUDFLARE_R2_BUCKET_NAME")),
		Key:         aws.String(filename),
		ACL:         s3types.ObjectCannedACLPublicRead,
		ContentType: aws.String("image/jpeg"),
	}, s3.WithPresignExpires(15*time.Minute))

	if err != nil {
		return "", err
	}

	return req.URL, nil
}
