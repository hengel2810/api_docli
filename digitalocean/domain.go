package digitalocean

import (
	"github.com/digitalocean/godo"
	"github.com/digitalocean/godo/context"
	"golang.org/x/oauth2"
	"errors"
	"fmt"
	"os"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func CreateSubdomain(subdomain string) error {
	tokenSource := &TokenSource{
		AccessToken: os.Getenv("DO_TOKEN"),
	}
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)
	createRequest := &godo.DomainRecordEditRequest{
		Type: "A",
		Name: subdomain,
		Data: "46.101.222.225",
		TTL: 3600,
	}
	domainRecord, _, err := client.Domains.CreateRecord(context.TODO(), "valas.cloud", createRequest)
	if err != nil {
		return errors.New("error creating domain record")
	}
	fmt.Println(domainRecord)
	return nil
}
