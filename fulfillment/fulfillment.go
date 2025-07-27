package fulfillment

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// The well-known default Entra ID scope for SaaS fulfillment APIs as described here: https://learn.microsoft.com/en-us/partner-center/marketplace-offers/pc-saas-registration#request-body
const WellKnownScopeSaaSFulfillment = "20e940b3-4c77-4b0b-9a53-9e16a1b010a7/.default"

// API endpoint for Azure Marketplace SaaS publisher operations
const endpoint = "https://marketplaceapi.microsoft.com/api"

// NewFulfillmentClient creates a new instance of a SaaS fulfillment client and an API operation status (tracking) client. Usually you only need the former.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFulfillmentClient(credential azcore.TokenCredential, options *policy.ClientOptions) (*OperationsClient, error) {
	popts := runtime.PipelineOptions{
		PerCall: []policy.Policy{runtime.NewBearerTokenPolicy(credential, []string{WellKnownScopeSaaSFulfillment}, nil)}}
	cl, err := azcore.NewClient(moduleName, moduleVersion, popts, options)
	if err != nil {
		return nil, err
	}
	return &OperationsClient{internal: cl, endpoint: endpoint}, nil
}

func NewSubscriptionOperationsClient(credential azcore.TokenCredential, options *policy.ClientOptions) (*SubscriptionOperationsClient, error) {
	popts := runtime.PipelineOptions{
		PerCall: []policy.Policy{runtime.NewBearerTokenPolicy(credential, []string{WellKnownScopeSaaSFulfillment}, nil)}}
	cl, err := azcore.NewClient(moduleName, moduleVersion, popts, options)
	if err != nil {
		return nil, err
	}
	return &SubscriptionOperationsClient{internal: cl, endpoint: endpoint}, nil
}
