package s3

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	localConfig "github.com/chiragthapa777/expense-tracker-api/internal/config"
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
)

type S3 struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	bucketName    string
}

var (
	s3Instance *S3
	once       sync.Once
)

// GetLogger returns the singleton logger instance
func GetS3() *S3 {
	once.Do(func() {
		localConfig := localConfig.GetConfig()
		log := logger.GetLogger()
		var bucketName = localConfig.R2_BUCKET_NAME
		var accountId = localConfig.R2_ACCOUNT_ID
		var accessKeyId = localConfig.R2_ACCESS_KEY_ID
		var accessKeySecret = localConfig.R2_SECRET_ACCESS_KEY

		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
			config.WithRegion("auto"),
		)

		if err != nil {
			log.Error(err.Error())
			panic(err.Error())
		}

		client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))
		})
		presignClient := s3.NewPresignClient(client)

		s3Instance = &S3{
			client:        client,
			presignClient: presignClient,
			bucketName:    bucketName,
		}
	})
	return s3Instance
}

func (s S3) UploadFile(ctx context.Context, objectKey string, ContentType *string, fileHeader multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	_, err = s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: ContentType,
	})

	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "EntityTooLarge" {
			log.Printf("Error while uploading object to %s. The object is too large.\n"+
				"To upload objects larger than 5GB, use the S3 console (160GB max)\n"+
				"or the multipart upload API (5TB max).", s.bucketName)
		} else {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				fileHeader.Filename, s.bucketName, objectKey, err)
		}
		return err
	} else {
		err = s3.NewObjectExistsWaiter(s.client).Wait(
			ctx, &s3.HeadObjectInput{Bucket: aws.String(s.bucketName), Key: aws.String(objectKey)}, time.Minute)
		if err != nil {
			log.Printf("Failed attempt to wait for object %s to exist.\n", objectKey)
		}
	}

	return nil
}

func (s S3) DeleteObject(ctx context.Context, key string, versionId string, bypassGovernance bool) (bool, error) {
	deleted := false
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	}
	if versionId != "" {
		input.VersionId = aws.String(versionId)
	}
	if bypassGovernance {
		input.BypassGovernanceRetention = aws.Bool(true)
	}
	_, err := s.client.DeleteObject(ctx, input)
	if err != nil {
		var noKey *types.NoSuchKey
		var apiErr *smithy.GenericAPIError
		if errors.As(err, &noKey) {
			log.Printf("Object %s does not exist in %s.\n", key, s.bucketName)
			err = noKey
		} else if errors.As(err, &apiErr) {
			switch apiErr.ErrorCode() {
			case "AccessDenied":
				log.Printf("Access denied: cannot delete object %s from %s.\n", key, s.bucketName)
				err = nil
			case "InvalidArgument":
				if bypassGovernance {
					log.Printf("You cannot specify bypass governance on a bucket without lock enabled.")
					err = nil
				}
			}
		}
		return false, err
	}
	err = s3.NewObjectNotExistsWaiter(s.client).Wait(
		ctx, &s3.HeadObjectInput{Bucket: aws.String(s.bucketName), Key: aws.String(key)}, time.Minute)
	if err != nil {
		log.Printf("Failed attempt to wait for object %s in bucket %s to be deleted.\n", key, s.bucketName)
	} else {
		deleted = true
	}

	return deleted, err
}

func (s S3) GetObject(
	ctx context.Context, objectKey string, lifetimeSecs int64) (*v4.PresignedHTTPRequest, error) {
	request, err := s.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
	})
	if err != nil {
		log.Printf("Couldn't get a presigned request to get %v:%v. Here's why: %v\n",
			s.bucketName, objectKey, err)
		return nil, err
	}
	return request, err
}
