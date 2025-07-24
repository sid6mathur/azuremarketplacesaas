# Go bindings for Azure Marketplace Saas Fulfilment and Metering APIs

This is an *Unofficial* Go SDK for Azure Marketplace Saas Fulfilment and Metering APIs, based on the OpenAPI 3 specification published by Microsoft's Marketplace team at this [repository](https://github.com/microsoft/commercial-marketplace-openapi/).

## Installation

There are separate packages for Fulfilment and Metering APIs. You can install them using `go get`. We don't have formal releases yet, so you should use the `@main` branch.

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

## Using the Metering endpoint

See the `metering` package's [metering_test.go](metering/metering_test.go) file for example on how to use the API, including authorization.

Once you have set the Azure AD app credentials in the [env.sh](env.sh) file, use the supplied [Makefile](Makefile) target to run the example/test. It simply calls `go test -v` under the hood after setting the environment variables.

```bash
make test-metering
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


### References

- [Marketplace metering service authentication strategies](https://learn.microsoft.com/en-us/partner-center/marketplace/marketplace-metering-service-authentication)
