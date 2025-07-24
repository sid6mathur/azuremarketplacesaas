package metering

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func TestMeteringClient(t *testing.T) {
	// If you have an App Registration that's already registered in the MPN portal for SaaS fulfillment+metering,
	// configure it's identity via the following environment variables:
	// AZURE_TENANT_ID, AZURE_CLIENT_ID, AZURE_CLIENT_SECRET
	cred, err := azidentity.NewEnvironmentCredential(nil)
	if err != nil {
		t.Logf("failed to obtain a credential: %v", err)
		return
	}
	// We request a token for scope AzureMarketplaceWellKnownTenantID + "/.default"
	// since AZURE_TENANT_ID above is already allow-listed in the MPN portal's SaaS offer configuration.
	mc, err := NewMeteringClient(cred, []string{AzureMarketplaceSaaSMeteringWellKnownTenantID + "/.default"}, nil)
	if err != nil {
		t.Logf("failed to create an Azure Marketplace metering client: %v", err)
		return
	}
	ctx := context.Background()
	// Use the filter fields in OperationsClientGetUsageEventOptions to narrow the scope of your query
	events, err := mc.GetUsageEvent(ctx, startOfYearUTC(time.Now()), &OperationsClientGetUsageEventOptions{})
	if err != nil {
		t.Logf("failed to fetch previously-posted usage events: %v", err)
		return
	}
	for _, event := range events.GetUsageEventArray {
		t.Logf("Previously-posted Metering event: Date %s, Usage ResourceId = %s, dim = %s, units = %f on Plan %s", (*event.UsageDate).UTC().Format("2006-01-02"), *event.UsageResourceID, *event.Dimension, *event.ProcessedQuantity, *event.PlanID)
	}
}

func TestMeteringPostEvent(t *testing.T) {
	cred, err := azidentity.NewEnvironmentCredential(nil)
	if err != nil {
		t.Logf("failed to obtain a credential: %v", err)
		return
	}
	mc, err := NewMeteringClient(cred, []string{AzureMarketplaceSaaSMeteringWellKnownTenantID + "/.default"}, nil)
	if err != nil {
		t.Logf("failed to create an Azure Marketplace metering client: %v", err)
		return
	}
	ctx := context.Background()
	// List of usage events to post, one at a time
	// See event schema at https://learn.microsoft.com/en-us/partner-center/marketplace-offers/marketplace-metering-service-apis#metered-billing-single-usage-event
	units := 22.0
	list := []UsageEvent{
		{EffectiveStartTime: to.Ptr(time.Date(2025, 02, 28, 0, 0, 0, 0, time.UTC)), // TODO: Replace with the start time of the usage report
			PlanID:     to.Ptr("smallbiz_01"),                    // TODO: Replace this with the Plan ID from your offer (tier Id)
			Quantity:   &units,                                   // TODO: number of units to be reported; this is usually (consumed units - what's included the base price of the SaaS)
			ResourceID: to.Ptr("pppppp-xxxx-yyyy-zzzz-qqqqqqqq"), // TODO: This is the SaaS Subscription ID
			Dimension:  to.Ptr("api_calls_1k")},                  // TODO: The dimension - here this is a custom defined unit of measure; also visible in the plan setup in the MPN portal
	}
	for _, ue := range list {
		if pr, err := mc.PostUsageEvent(ctx, ue, nil); err != nil {
			t.Logf("failed to post usage event: %v", err)
		} else {
			t.Logf("[%s] Posted Metering event OK: Status = %s, Usage EventId = %s, Qty = %f", *pr.ResourceID, *pr.Status, *pr.UsageEventID, *pr.Quantity)
		}
	}
}

// startOfYearUTC returns the first moment of the year in UTC
func startOfYearUTC(date time.Time) time.Time {
	utc := date.UTC()
	return time.Date(utc.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
}
