package hive_go

import (
	"context"
	"net/http"
	"strings"
)

type Backend interface {
	GetUser(ctx context.Context, token string) (int, IAuthenticationBackendUser)
}

type Client struct {
	backends                    map[string]Backend
	authorizationHeader         string
	authenticatedUserContextKey string
	subscription                *Subscription
	repository                  *Repository
}

type Config struct {
	AuthorizationHeader         string
	Backends                    map[string]Backend
	HiveAPI                     string
	SubscriptionTopic           string
	ServiceName                 string
	NSQLookupAddress            string
	AuthenticatedUserContextKey string
}

func InitHiveClient(config *Config) *Client {
	authorizationHeader := config.AuthorizationHeader
	if authorizationHeader == "" {
		authorizationHeader = "Authorization"
	}

	authenticatedUserContextKey := config.AuthenticatedUserContextKey
	if authenticatedUserContextKey == "" {
		authenticatedUserContextKey = "authenticatedUser"
	}

	repository := InitRepository(config.HiveAPI)

	backends := config.Backends
	if backends == nil {
		backends = map[string]Backend{
			"Bearer": InitJWTBackend(repository),
		}
	}

	var subscription *Subscription

	if config.ServiceName != "" && config.NSQLookupAddress != "" {
		subscription = InitHiveSubscription(config.ServiceName, config.SubscriptionTopic, config.NSQLookupAddress, repository)
	}

	return &Client{
		authorizationHeader:         authorizationHeader,
		backends:                    backends,
		authenticatedUserContextKey: authenticatedUserContextKey,
		subscription:                subscription,
		repository:                  repository,
	}
}

func (client *Client) GetToken(authorizationHeader string) (string, string) {

	if authorizationHeader == "" {
		return "", ""
	}

	parts := strings.Split(authorizationHeader, " ")

	if len(parts) < 2 {
		return "", ""
	}

	return parts[0], parts[1]
}

func (client *Client) AddBackend(key string, backend Backend) {
	client.backends[key] = backend
}

func (client *Client) RemoveBackend(key string) Backend {
	backend := client.GetBackend(key)
	delete(client.backends, key)
	return backend
}

func (client *Client) GetBackend(key string) Backend {
	return client.backends[key]
}

func (client *Client) GetRepository() Repository {
	return *client.repository
}

func (client *Client) Login(ctx context.Context, r *http.Request) (int, IAuthenticationBackendUser) {

	authorizationHeader := r.Header.Get(client.authorizationHeader)
	tokenType, token := client.GetToken(authorizationHeader)
	if tokenType == "" || token == "" {
		return Ok, nil
	}

	backend := client.backends[tokenType]
	if backend == nil {
		return BackendNotFound, nil
	}

	status, user := backend.GetUser(ctx, token)
	return status, user
}

func (client *Client) Stop() {
	if client.subscription != nil {
		client.subscription.Stop()
	}
}

func (client *Client) GetUserFromContext(ctx context.Context) IAuthenticationBackendUser {
	user, _ := ctx.Value(client.authenticatedUserContextKey).(IAuthenticationBackendUser)
	return user
}

func (client *Client) SetUserToContext(ctx context.Context, user IAuthenticationBackendUser) context.Context {
	return context.WithValue(ctx, client.authenticatedUserContextKey, user)
}
