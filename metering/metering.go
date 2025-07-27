package metering

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// The well-known metering service tenant ID for SaaS and Azure Applications with managed application plan type is 20e940b3-4c77-4b0b-9a53-9e16a1b010a7.
// See https://learn.microsoft.com/en-us/partner-center/marketplace-offers/marketplace-metering-service-authentication#using-the-microsoft-entra-security-token
const WellKnownScopeSaaSMetering = "20e940b3-4c77-4b0b-9a53-9e16a1b010a7/.default"

// API endpoint for Azure Marketplace SaaS publisher operations
const endpoint = "https://marketplaceapi.microsoft.com/api"

// NewMeteringClient creates a new instance of OperationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMeteringClient(credential azcore.TokenCredential, options *policy.ClientOptions) (*OperationsClient, error) {
	popts := runtime.PipelineOptions{
		PerCall: []policy.Policy{runtime.NewBearerTokenPolicy(credential, []string{WellKnownScopeSaaSMetering}, nil)}}
	cl, err := azcore.NewClient(moduleName, moduleVersion, popts, options)
	if err != nil {
		return nil, err
	}
	soc := &OperationsClient{internal: cl, endpoint: endpoint}
	return soc, nil
}
