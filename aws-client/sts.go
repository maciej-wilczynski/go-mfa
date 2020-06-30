package awsclient

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
)

func (c *AWSClient) Export() {
	fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", c.sessionCredentials.AccessKeyId)
	fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", c.sessionCredentials.SecretAccessKey)
	fmt.Printf("export AWS_SESSION_TOKEN=%s\n", c.sessionCredentials.SessionToken)
}

func (c *AWSClient) AssumeToRole(targetProfile Profile, pin string) error {
	input := &sts.AssumeRoleInput{
		DurationSeconds: aws.Int64(int64(targetProfile.Duration)),
		RoleArn:         aws.String(targetProfile.RoleARN),
		RoleSessionName: aws.String("mfa-session"),
		SerialNumber:    aws.String(targetProfile.MFASerial),
		TokenCode:       aws.String(pin),
	}

	result, err := c.sts.AssumeRole(input)
	if err != nil {
		return err
	}

	c.sessionCredentials = Credential{
		AccessKeyId:     *result.Credentials.AccessKeyId,
		SecretAccessKey: *result.Credentials.SecretAccessKey,
		SessionToken:    *result.Credentials.SessionToken,
	}

	return nil
}
