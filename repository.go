package hive_go

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/proto"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"net/http"
	"time"
)

type SecretsRepository interface {
	GetSecret(id uuid.UUID) *Secret
}

type Repository struct {
	client  *http.Client
	hiveAPI string
	cache   *cache.Cache
}

func InitRepository(hiveAPI string) *Repository {
	return &Repository{client: &http.Client{}, hiveAPI: hiveAPI, cache: cache.New(time.Minute*10, time.Minute*5)}
}

func (repository Repository) getFullURL(path string) string {
	return fmt.Sprintf("%s/%s", repository.hiveAPI, path)
}

func (repository Repository) getSecretsURL(id uuid.UUID) string {
	return repository.getFullURL(fmt.Sprintf("api/v1/secrets/%s", id.String()))
}

func (repository Repository) getSecret(id uuid.UUID) *Secret {
	if x, found := repository.cache.Get(id.String()); found {
		if secret, ok := x.(*Secret); ok {
			return secret
		}
	}

	return nil
}

func (repository Repository) SetSecret(secret *Secret) {
	repository.cache.Set(secret.Id.String(), secret, time.Hour*24)
}

func (repository Repository) GetSecret(id uuid.UUID) *Secret {
	secret := repository.getSecret(id)
	if secret != nil {
		return secret
	}

	url := repository.getSecretsURL(id)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil
	}
	request.Header.Add("Content-Type", "application/octet-stream")
	resp, err := repository.client.Do(request)
	if err != nil {
		return nil
	}
	if resp.StatusCode != 200 {
		return nil
	}
	var response GetSecretResponseV1
	body, err := ioutil.ReadAll(resp.Body)
	err = proto.Unmarshal(body, &response)
	if err != nil {
		return nil
	}

	secretData := response.Data

	uuidID, err := uuid.FromBytes(secretData.Id)
	if err != nil {
		return nil
	}
	uuidValue, err := uuid.FromBytes(secretData.Value)
	if err != nil {
		return nil
	}

	secret = &Secret{
		Id:      uuidID,
		Created: secretData.Created,
		Value:   uuidValue,
	}

	repository.SetSecret(secret)

	return secret
}
