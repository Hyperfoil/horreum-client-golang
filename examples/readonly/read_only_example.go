package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/hyperfoil/horreum-client-golang/pkg/horreum"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

var (
	ctx      = context.Background()
	baseUrl  = "http://localhost:8080"
	username = "horreum.bootstrap"
	password = "secret"

	// assertions
	enableAssertion      = true
	expectedSchemasCount = 2
	expectedTestsCount   = 1
)

// Of returns a pointer to the provided literal/const input
func Of[E any](e E) *E {
	return &e
}

func main() {
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: time.Second * 100,
	}
	parentTransport, _ := http.DefaultTransport.(*http.Transport)
	parentTransport.ForceAttemptHTTP2 = true
	parentTransport.DisableCompression = true
	parentTransport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	client, err := horreum.NewHorreumClient(baseUrl,
		&horreum.HorreumCredentials{
			Username: &username,
			Password: &password,
		},
		&horreum.ClientConfiguration{
			HttpClient:            httpClient,
			ParentTransport:       parentTransport,
			UseDefaultMiddlewares: true,
			Options:               []abstractions.RequestOption{},
		},
	)
	if err != nil {
		log.Fatalf("error creating Horreum client: %s", err.Error())
	}

	clientVersion := client.Version()
	serverVersion, err := client.RawClient.Api().Config().Version().Get(ctx, nil)
	if err != nil {
		log.Fatalf("error fetching Horreum server version: %s", err.Error())
	}

	log.Printf("Running client %s against server %s", clientVersion, *serverVersion.GetVersion())

	// assert data into Horreum server
	if getSchemas, _ := client.RawClient.Api().Schema().Get(ctx, nil); enableAssertion && *getSchemas.GetCount() != int64(2) {
		log.Fatalf("expected %d schemas in the server", expectedSchemasCount)
	}
	if getTests, _ := client.RawClient.Api().Test().Get(ctx, nil); enableAssertion && *getTests.GetCount() != int64(1) {
		log.Fatalf("expected %d test in the server", expectedTestsCount)
	}
}
