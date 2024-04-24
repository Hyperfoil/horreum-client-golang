package main

import (
	"context"
	"log"
	"strconv"

	"github.com/hyperfoil/horreum-client-golang/pkg/horreum"
	"github.com/hyperfoil/horreum-client-golang/pkg/raw_client/api"
	"github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

var (
	ctx                            = context.Background()
	username                       = "user"
	password                       = "secret"
	cleanupData                    = true
	acmeBenchmarkSchema            = models.NewSchema()
	acmeHorreumSchema              = models.NewSchema()
	acmeBenchmarkSchemaTransformer = models.NewTransformer()
	acmeExtractor1                 = models.NewExtractor()
	acmeExtractor2                 = models.NewExtractor()
	roadrunnerTest                 = models.NewTest()
	roadrunnerRun                  = models.NewRun()
)

// Of returns a pointer to the provided literal/const input
func Of[E any](e E) *E {
	return &e
}

// This will setup all data variables
func init() {
	log.Printf("setting up data")

	// Benchmark Schema
	acmeBenchmarkSchema.SetName(Of("ACME Benchmark Schema"))
	acmeBenchmarkSchema.SetDescription(Of("Data produced by benchmarking tool"))
	acmeBenchmarkSchema.SetOwner(Of("dev-team"))
	acmeBenchmarkSchema.SetAccess(Of(models.PUBLIC_PROTECTEDTYPE_ACCESS))
	acmeBenchmarkSchema.SetUri(Of("urn:acme:benchmark:0.1"))

	// Horreum Schema
	acmeHorreumSchema.SetName(Of("ACME Horreum Schema"))
	acmeHorreumSchema.SetDescription(Of("Used in Datasets"))
	acmeHorreumSchema.SetOwner(Of("dev-team"))
	acmeHorreumSchema.SetAccess(Of(models.PUBLIC_PROTECTEDTYPE_ACCESS))
	acmeHorreumSchema.SetUri(Of("urn:acme:horreum:0.1"))

	// Transformer
	acmeBenchmarkSchemaTransformer.SetName(Of("Acme Transformer"))
	acmeBenchmarkSchemaTransformer.SetDescription(Of("Transformer for converting complex runs into individual datasets"))
	acmeBenchmarkSchemaTransformer.SetOwner(Of("dev-team"))
	acmeBenchmarkSchemaTransformer.SetAccess(Of(models.PUBLIC_PROTECTEDTYPE_ACCESS))
	acmeBenchmarkSchemaTransformer.SetTargetSchemaUri(Of("urn:acme:horreum:0.1"))
	acmeBenchmarkSchemaTransformer.SetFunction(Of("({results, hash}) => results.map(r => ({ ...r, hash }))"))

	acmeExtractor1.SetName(Of("hash"))
	acmeExtractor1.SetJsonpath(Of("$.buildHash"))
	acmeExtractor1.SetIsarray(Of(false))

	acmeExtractor2.SetName(Of("results"))
	acmeExtractor2.SetJsonpath(Of("$.results"))
	acmeExtractor2.SetIsarray(Of(false))

	acmeBenchmarkSchemaTransformer.SetExtractors([]models.Extractorable{acmeExtractor1, acmeExtractor2})

	// Roadrunner Test
	roadrunnerTest.SetName(Of("Roadrunner Test"))
	roadrunnerTest.SetDescription(Of("acme.com benchmark"))
	roadrunnerTest.SetOwner(Of("dev-team"))
	roadrunnerTest.SetAccess(Of(models.PUBLIC_PROTECTEDTYPE_ACCESS))
	roadrunnerTest.SetFingerprintLabels([]string{"benchmark_test"})

	// Roadrunner Run
	roadrunnerRun.SetDescription(Of("acme.com benchmark"))
	roadrunnerRun.SetOwner(Of("dev-team"))
	roadrunnerRun.SetAccess(Of(models.PUBLIC_RUN_ACCESS))
	roadrunnerRun.SetStart(Of(int64(1669388931000)))
	roadrunnerRun.SetStop(Of(int64(1669388932000)))
	roadrunnerRun.SetData(Of("{\n    \"$schema\": \"urn:acme:benchmark:0.1\",\n    \"something\": \"This gets lost by the transformer\",\n    \"buildHash\": \"defec8eddeadbeafcafebabeb16b00b5\",\n    \"results\": [\n      {\n        \"test\": \"Foo\",\n        \"requests\": 123,\n        \"duration\": 10\n      },\n      {\n        \"test\": \"Bar\",\n        \"requests\": 456,\n        \"duration\": 20\n      }\n    ]\n}"))
}

func createSchema(client *horreum.HorreumClient, schema models.Schemaable) int32 {
	log.Printf("creating schema for %s", *schema.GetName())

	schemaId, err := client.RawClient.Api().Schema().Post(ctx, schema, nil)
	if err != nil {
		log.Fatalf("unable to create new schema: %s", err.Error())
	}

	return *schemaId
}

func createSchemaTransformers(client *horreum.HorreumClient, schemaId int32, transformer models.Transformerable) int32 {
	transformerId, err := client.RawClient.Api().Schema().ByIdIdInteger(schemaId).Transformers().Post(ctx, transformer, nil)
	if err != nil {
		log.Fatalf("unable to create schema transformers for %d: %s", schemaId, err.Error())
	}

	return *transformerId
}

func createTest(client *horreum.HorreumClient, test models.Testable) models.Testable {
	createdTest, err := client.RawClient.Api().Test().Post(ctx, test, nil)
	if err != nil {
		log.Fatalf("unable to create new test: %s", err.Error())
	}

	return createdTest
}

func uploadRun(client *horreum.HorreumClient, testId int32, run models.Runable) {
	client.RawClient.Api().Run().Test().Post(ctx, run, &api.RunTestRequestBuilderPostRequestConfiguration{
		QueryParameters: &api.RunTestRequestBuilderPostQueryParameters{
			Test: Of(strconv.Itoa(int(testId))),
		},
	})
}

func setupRoadrunnerTest(client *horreum.HorreumClient) {
	log.Print("creating roadrunner test")

	benchmarkSchemaId := createSchema(client, acmeBenchmarkSchema)
	_ = createSchema(client, acmeHorreumSchema)
	createSchemaTransformers(client, benchmarkSchemaId, acmeBenchmarkSchemaTransformer)
	test := createTest(client, roadrunnerTest)
	uploadRun(client, *test.GetId(), roadrunnerRun)
}

// delete all records that this examples will create
func deleteAll(client *horreum.HorreumClient) {
	log.Printf("cleaning up tests")
	testQueryRes, err := client.RawClient.Api().Test().Get(ctx, nil)
	if err != nil {
		log.Fatalf("unable to query all tests: %s", err.Error())
	}
	for _, t := range testQueryRes.GetTests() {
		err = client.RawClient.Api().Test().ByIdInteger(*t.GetId()).Delete(ctx, nil)
		if err != nil {
			log.Fatalf("unable to delete test %d: %s", *t.GetId(), err.Error())
		}
	}

	log.Printf("cleaning up schemas")
	schemaQueryRes, err := client.RawClient.Api().Schema().Get(ctx, nil)
	if err != nil {
		log.Fatalf("unable to query all schemas: %s", err.Error())
	}
	for _, s := range schemaQueryRes.GetSchemas() {
		err = client.RawClient.Api().Schema().ByIdIdInteger(*s.GetId()).Delete(ctx, nil)
		if err != nil {
			log.Fatalf("unable to delete schema %d: %s", *s.GetId(), err.Error())
		}
	}
}

func main() {
	client, err := horreum.NewHorreumClient("http://localhost:8080", &horreum.HorreumCredentials{
		Username: &username,
		Password: &password,
	}, nil)
	if err != nil {
		log.Fatalf("error creating Horreum client: %s", err.Error())
	}

	clientVersion := client.Version()
	serverVersion, err := client.RawClient.Api().Config().Version().Get(ctx, nil)
	if err != nil {
		log.Fatalf("error fetching Horreum server version: %s", err.Error())
	}

	log.Printf("Running client %s against server %s", clientVersion, *serverVersion.GetVersion())

	if cleanupData {
		deleteAll(client)
	}

	setupRoadrunnerTest(client)

	// assert data into Horreum server
	if getSchemas, _ := client.RawClient.Api().Schema().Get(ctx, nil); *getSchemas.GetCount() != int64(2) {
		log.Fatalf("expected 2 schemas in the server")
	}
	if getTests, _ := client.RawClient.Api().Test().Get(ctx, nil); *getTests.GetCount() != int64(1) {
		log.Fatalf("expected 1 test in the server")
	}
	if getRuns, _ := client.RawClient.Api().Run().List().Get(ctx, nil); *getRuns.GetTotal() != int64(1) {
		log.Fatalf("expected 1 run in the server")
	}
}
