package digitalocean

import (
	"github.com/digitalocean/godo"
	"github.com/digitalocean/godo/context"
	"golang.org/x/oauth2"
	"errors"
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

var token string
func CreateSubdomain(subdomain string) (int, error) {
	tokenSource := &TokenSource{
		AccessToken: token,
	}
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)
	createRequest := &godo.DomainRecordEditRequest{
		Type: "A",
		Name: subdomain,
		Data: "46.101.222.225",
		TTL: 3600,
	}
	record, response, err := client.Domains.CreateRecord(context.TODO(), "valas.cloud", createRequest)
	if err != nil || response.StatusCode != 201 {
		return 0, errors.New("error creating domain record")
	}
	return record.ID, nil
}

func DeleteSubdomain(recordId int) error {
	tokenSource := &TokenSource{
		AccessToken: token,
	}
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)
	resp, err := client.Domains.DeleteRecord(context.TODO(), "valas.cloud", recordId)
	if err != nil || resp.StatusCode != 204 {
		return errors.New("error deleting domain record")
	}
	return nil
}
