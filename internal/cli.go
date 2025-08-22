package cli

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	bitbucket "github.com/gfleury/go-bitbucket-v1"
	gitHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/sirupsen/logrus"
)

type BitbucketCLI struct {
	cloneCredentials gitHttp.BasicAuth
	restURL          *url.URL
	client           *bitbucket.APIClient
	logger           *logrus.Logger
	httpClient       *http.Client
	auth             Authenticator
}

func (b *BitbucketCLI) SetLogger(logger *logrus.Logger) {
	if logger == nil {
		// We don't set nil loggers
		return
	}
	b.logger = logger
}

func NewCLI(auth Authenticator, restURL string) (*BitbucketCLI, error) {
	mURL, err := url.Parse(restURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx = auth.GetContext(ctx)
	c := bitbucket.NewAPIClient(ctx, bitbucket.NewConfiguration(
		strings.TrimRight(mURL.String(), "/"), // https://git.example.com/rest/ -> https://git.example.com/rest
	))
	logger := logrus.New()

	return &BitbucketCLI{
		cloneCredentials: auth.GetCloneCredentials(),
		restURL:          mURL,
		auth:             auth,
		client:           c,
		logger:           logger,
		httpClient:       http.DefaultClient,
	}, nil
}
