package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	baseAddress = "localhost:9001"
	tinyAddress = "localhost:9002"
)

func main() {
	base, err := newMinioClient(baseAddress)
	if err != nil {
		log.Printf("base: %v", err)
		return
	}

	tiny, err := newMinioClient(tinyAddress)
	if err != nil {
		log.Printf("tiny: %v", err)
		return
	}

	content, err := os.ReadFile("tomsawyer.txt")
	if err != nil {
		log.Printf("content: %v", err)
		return
	}

	bi, err := putObject(base, content)
	if err != nil {
		log.Printf("base put: %v", err)
		return
	}

	println("base: ", bi.Size)

	ti, err := putObject(tiny, content)
	if err != nil {
		log.Printf("tiny put: %v", err)
		return
	}

	println("tiny: ", ti.Size)
}

const (
	minioAccess = "accessKey"
	minioSecret = "secretKey"
)

func newMinioClient(address string) (*minio.Client, error) {
	c, err := minio.New(address, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccess, minioSecret, ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	err = createBucketIfNotExists(c)
	if err != nil {
		return nil, fmt.Errorf("bucket: %v", err)
	}

	return c, nil
}

const bucketName = "mybucket"

func createBucketIfNotExists(c *minio.Client) error {
	found, err := c.BucketExists(context.Background(), bucketName)
	if err != nil {
		return fmt.Errorf("exists: %v", err)
	}

	if !found {
		err = c.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("make: %v", err)
		}
	}

	return nil
}

func putObject(c *minio.Client, content []byte) (minio.UploadInfo, error) {
	const filename = "file.txt"

	size := int64(len(content))
	r := bytes.NewReader(content)

	return c.PutObject(context.Background(), bucketName, filename, r, size, minio.PutObjectOptions{})
}
