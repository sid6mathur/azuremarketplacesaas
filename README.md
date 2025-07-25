# Go bindings for Azure Marketplace Saas Fulfilment and Metering APIs

This is an *Unofficial* Go SDK for software publishers who wish to integrate their SaaS product with Azure Marketplace's [SaaS Fulfilment](https://learn.microsoft.com/en-us/partner-center/marketplace-offers/pc-saas-fulfillment-subscription-api) and [SaaS Metering APIs](https://learn.microsoft.com/en-us/partner-center/marketplace-offers/marketplace-metering-service-apis). The SDK uses auto-generated code from OpenAPI 3 specification published by Microsoft's Marketplace team at this [repository](https://github.com/microsoft/commercial-marketplace-openapi/).

## Installation

There are separate packages for Fulfilment and Metering APIs. You can install them using `go get`.

```go
import github.com/fastah/azuremarketplacesaas/fulfillment
```

```go
import github.com/fastah/azuremarketplacesaas/metering
```

If your SaaS product doesn't need metering, you can of course skip the metering package.

## Authorizing with Marketplace API endpoints

Use the same Azure AD app registration that you have allow-listed in the Partner Center in your SaaS offer's technical configuration. You will need to provide the coresponding `AZURE_CLIENT_ID`, `AZURE_CLIENT_SECRET`, and `AZURE_TENANT_ID` via the supplied [env.sh](env.sh) environment file.

## Using the Fulfillment package

See the `fulfillement` package's [fulfillment_test.go](fulfillment/fulfillment_test.go) file for example on how to use the API, including authorization.

### Usage - Fulfillment client

```go
cred, err := azidentity.NewEnvironmentCredential(nil)
	if err != nil {
		t.Logf("failed to obtain a credential: %v", err)
		return
	}
	fc, err := NewFulfillmentClient(cred, nil)
	if err != nil {
		t.Logf("failed to create an Azure Marketplace metering client: %v", err)
		return
	}
	pager := fc.NewListSubscriptionsPager(nil)
	t.Logf("List of subscriptions:")
	for pager.More() {
		resp, err := pager.NextPage(context.Background())
		if err != nil {
			t.Errorf("failed to get next page of subscriptions: %v", err)
			break
		}
		for _, s := range resp.Subscriptions {
			fmt.Printf("sub: ID %s in status %s\n", *s.ID, *s.SaasSubscriptionStatus)
		}
	}
```

## Using the Metering endpoint

See the `metering` package's [metering_test.go](metering/metering_test.go) file for example on how to use the API, including authorization.

Once you have set the Azure AD app credentials in the [env.sh](env.sh) file, use the supplied [Makefile](Makefile) target to run the example/test. It simply calls `go test -v` under the hood after setting the environment variables.

```bash
make test-metering
```

### Usage - Metering client 

```go
mc, err := NewMeteringClient(cred, nil)
	if err != nil {
		t.Logf("failed to create an Azure Marketplace metering client: %v", err)
		return
	}
	ctx := context.Background()
	// Example of listing previously posted metering events
	events, err := mc.GetUsageEvent(ctx, startOfYearUTC(time.Now()), &OperationsClientGetUsageEventOptions{})
	if err != nil {
		t.Logf("failed to fetch previously-posted usage events: %v", err)
		return
	}
	for _, event := range events.GetUsageEventArray {
		t.Logf("Previously-posted Metering event: Date %s, Usage ResourceId = %s, dim = %s, units = %f on Plan %s", (*event.UsageDate).UTC().Format("2006-01-02"), *event.UsageResourceID, *event.Dimension, *event.ProcessedQuantity, *event.PlanID)
	}
```

## Maintainers only: Updating the Go client from OpenAPI spec

### Submodule fetch from Microsoft's OpenAPI repo

```bash
git submodule update --init --recursive
```

### Install and Update AutoRest with Go extension

The following [Makefile](Makefile) target will install the `autorest` tool with the Go extension, and update it to the latest version.

```bash
make autorest-go-update-with-reset
```

### Runnng tests

Add your Azure AD app credentials to the [env.sh](env.sh) file, and run the following command to run the tests:

```bash
make test-fulfillment
make test-metering
```

### Tagging a release

```bash
git tag metering/v0.0.1
git tag fulfillment/v0.0.1
```

### References

- [Overview of SaaS subscription lifecycle](https://learn.microsoft.com/en-us/partner-center/marketplace-offers/pc-saas-fulfillment-life-cycle)

- [SaaS Metering API - Marketplace metering service authentication strategies](https://learn.microsoft.com/en-us/partner-center/marketplace/marketplace-metering-service-authentication)

- 