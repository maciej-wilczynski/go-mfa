package main

import (
	"log"
	"os"

	"github.com/maciej-wilczynski/go-mfa/argparser"
	awsclient "github.com/maciej-wilczynski/go-mfa/aws-client"
)

func main() {
	argparser, err := argparser.New(os.Args)
	if err != nil {
		log.Printf("Unable to parse arguments: %s", err)
		os.Exit(1)
	}

	targetProfile, err := awsclient.LoadProfile(awsclient.DefaultConfigFile, argparser.TargetProfile)
	if err != nil {
		log.Printf("Unable to parse config file: %s", err)
		os.Exit(1)
	}

	client, err := awsclient.New()
	if err != nil {
		log.Printf("Unable to setup AWS client: %s", err)
		os.Exit(1)
	}

	err = client.AssumeToRole(targetProfile, argparser.Pin)
	if err != nil {
		log.Printf("Error assuming role: %s", err)
		os.Exit(1)
	}

	client.Export()
}
