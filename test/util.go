// util.go
package test

import (
	"os"

	cli "github.com/swisscom/bitbucket-cli/internal"
)

func MustGetCLI() *cli.BitbucketCLI {
	c, err := cli.NewCLI(
		&cli.BasicAuth{
			Username: os.Getenv("BITBUCKET_USERNAME"),
			Password: os.Getenv("BITBUCKET_PASSWORD"),
		},
		os.Getenv("BITBUCKET_URL"),
	)
	if err != nil {
		panic(err)
	}
	return c
}
