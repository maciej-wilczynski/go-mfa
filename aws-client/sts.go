package awsclient

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
)

func (c *AWSClient) Export() {
	log.Printf("Exporting key %s", c.sessionCredentials.AccessKeyId)

	os.Setenv("AWS_ACCESS_KEY_ID", c.sessionCredentials.AccessKeyId)
	os.Setenv("AWS_SECRET_ACCESS_KEY", c.sessionCredentials.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", c.sessionCredentials.SessionToken)
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
