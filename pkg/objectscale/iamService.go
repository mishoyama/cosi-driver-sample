package objectscale

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iam"
)

const (
	s3FullPoilicyArn = "urn:osc:iam:::policy/OSCS3FullAccess"
)

type IamService objectScaleService

func (t *IamService) CreateUser(username string) (*iam.User, error) {
	res, err := t.client.iam.CreateUser(&iam.CreateUserInput{UserName: aws.String(username)})
	if err != nil {
		return nil, HandleError(err)
	}
	return res.User, nil
}

func (t *IamService) AttachUserPolicy(username string) error {
	_, err := t.client.iam.AttachUserPolicy(&iam.AttachUserPolicyInput{UserName: aws.String(username), PolicyArn: aws.String(s3FullPoilicyArn)})
	if err != nil {
		return HandleError(err)
	}
	return nil
}

func (t *IamService) CreateAccessKey(username string) (*iam.AccessKey, error) {
	res, err := t.client.iam.CreateAccessKey(&iam.CreateAccessKeyInput{UserName: aws.String(username)})
	if err != nil {
		return nil, HandleError(err)
	}
	return res.AccessKey, nil
}

func (t *IamService) ListUsers() ([]*iam.User, error) {
	res, err := t.client.iam.ListUsersWithContext(context.TODO(), &iam.ListUsersInput{})
	if err != nil {
		return nil, HandleError(err)
	}
	return res.Users, nil
}

func WithHeader(name, value string) request.Option {
	return func(h *request.Request) {
		h.HTTPRequest.Header.Set(name, value)
	}
}
