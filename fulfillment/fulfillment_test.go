package fulfillment

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func TestFulfillment(t *testing.T) {
	// If you have an App Registration that's already registered in the MPN portal for SaaS fulfillment+metering,
	// configure it's identity via the following environment variables:
	// AZURE_TENANT_ID, AZURE_CLIENT_ID, AZURE_CLIENT_SECRET
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
}
