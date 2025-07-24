package fulfillment

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

const host = "https://marketplaceapi.microsoft.com/api"

// NewFulfillmentClient creates a new instance of a SaaS fulfillment client and an API status tracking client. Usually you only need the former.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFulfillmentClient(credential azcore.TokenCredential, options *policy.ClientOptions) (foc *OperationsClient, soc *SubscriptionOperationsClient, err error) {
	popts := runtime.PipelineOptions{
		PerCall: []policy.Policy{runtime.NewBearerTokenPolicy(credential, []string{}, nil)}}
	cl, err := azcore.NewClient(moduleName, moduleVersion, popts, options)
	if err != nil {
		return nil, nil, err
	}
	return &OperationsClient{internal: cl, endpoint: host},
		&SubscriptionOperationsClient{internal: cl, endpoint: host},
		nil
}
