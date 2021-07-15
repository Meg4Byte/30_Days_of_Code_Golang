package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func GetIamConfigure() aws.Config {
	//cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("neil"))
	cfg, err := config.LoadDefaultConfig(context.TODO())
	checkError(err)
	return cfg
}

func ListAllBuckets(cfg aws.Config) {
	s3svc := s3.NewFromConfig(cfg)
	bl, err := s3svc.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	checkError(err)

	for _, name := range bl.Buckets {
		if *name.Name == "cf-templates-10g05dtt4n6x7-us-east-1" {
			fmt.Printf("%v\n", *name.Name)
		}
	}
}

func stsGetCallerIdentity(cfg aws.Config) string {
	svc := sts.NewFromConfig(cfg)

	v, err := svc.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	checkError(err)
	return string(*v.Arn)
}

func main() {
	defConfg := GetIamConfigure()
	fmt.Println("Let's List Buckets")
	ListAllBuckets(defConfg)
	fmt.Println("Hello I am IAM: ")
	fmt.Printf("%v\n", stsGetCallerIdentity(defConfg))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
}

// example output
// Let's List Buckets
// cf-templates-10g05dtt4n6x7-us-east-1
// Hello I am IAM:
// arn:aws:iam::807885433112:user/neil
