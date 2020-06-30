package awsclient

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
)

type AWSClient struct {
	sts stsiface.STSAPI

	sessionCredentials Credential
}

func New() (AWSClient, error) {
	mySession := session.Must(session.NewSession())

	return AWSClient{
		sts: sts.New(mySession),
	}, nil
}
