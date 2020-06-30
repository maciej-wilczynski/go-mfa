package awsclient

import (
	"fmt"

	"gopkg.in/ini.v1"
)

const (
	DefaultConfigFile = "/Users/maciej.wilczynski/.aws/config"
)

type Profile struct {
	Name          string
	RoleARN       string `ini:"role_arn"`
	MFASerial     string `ini:"mfa_serial"`
	Duration      int    `ini:"duration_seconds"`
	SourceProfile string `ini:"source_profile"`
	Region        string `ini:"region"`
	Credential    Credential
}

type Credential struct {
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
}

func LoadProfile(filename, profile string) (p Profile, err error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return p, err
	}

	p.Name = profile

	sectionName := fmt.Sprintf("profile %s", profile)
	return p, cfg.Section(sectionName).MapTo(&p)

}
