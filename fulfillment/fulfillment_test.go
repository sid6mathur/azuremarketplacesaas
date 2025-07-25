package fulfillment

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

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

func TestSubOps(t *testing.T) {
	cred, err := azidentity.NewEnvironmentCredential(nil)
	if err != nil {
		t.Logf("failed to obtain a credential: %v", err)
		return
	}
	fc, err := NewSubscriptionOperationsClient(cred, nil)
	if err != nil {
		t.Logf("failed to create an Azure Marketplace metering client: %v", err)
		return
	}
	subID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subID == "" {
		t.Logf("AZURE_SUBSCRIPTION_ID environment variable is not set")
		return
	}
	resp, err := fc.ListOperations(context.Background(), subID, nil)
	if err != nil {
		t.Logf("failed to get list of ongoing operations associated with subscription %s: %v", subID, err)
		return
	}
	t.Logf("Found %d ongoing operations for subscription %s", len(resp.OperationList.Operations), subID)
}

// Tests webook JSON payload parsing
// https://learn.microsoft.com/en-us/partner-center/marketplace-offers/pc-saas-fulfillment-webhook
func TestSaaSOperationJSONParsing(t *testing.T) {
	jsonPayload := `{
    "id": "<guid>",
    "activityId": "<guid>",
    "publisherId": "XXX",
    "offerId": "YYY",
    "planId": "plan2",
    "quantity": 10,
    "subscriptionId": "<guid>",
    "timeStamp": "2023-02-10T18:48:58.4449937Z",
    "action": "ChangePlan",
    "status": "InProgress"
}`

	var operation SaaSOperation
	err := json.Unmarshal([]byte(jsonPayload), &operation)

	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Test all fields present in the SaaSOperation struct
	if operation.ID == nil || *operation.ID != "<guid>" {
		t.Errorf("Expected ID to be '<guid>', got %v", operation.ID)
	}

	if operation.ActivityID == nil || *operation.ActivityID != "<guid>" {
		t.Errorf("Expected ActivityID to be '<guid>', got %v", operation.ActivityID)
	}

	if operation.PublisherID == nil || *operation.PublisherID != "XXX" {
		t.Errorf("Expected PublisherID to be 'XXX', got %v", operation.PublisherID)
	}

	if operation.OfferID == nil || *operation.OfferID != "YYY" {
		t.Errorf("Expected OfferID to be 'YYY', got %v", operation.OfferID)
	}

	if operation.PlanID == nil || *operation.PlanID != "plan2" {
		t.Errorf("Expected PlanID to be 'plan2', got %v", operation.PlanID)
	}

	if operation.Quantity == nil || *operation.Quantity != 10 {
		t.Errorf("Expected Quantity to be 10, got %v", operation.Quantity)
	}

	if operation.SubscriptionID == nil || *operation.SubscriptionID != "<guid>" {
		t.Errorf("Expected SubscriptionID to be '<guid>', got %v", operation.SubscriptionID)
	}

	// Check timestamp parsing
	expectedTime, _ := time.Parse(time.RFC3339, "2023-02-10T18:48:58.4449937Z")
	if operation.TimeStamp == nil || !operation.TimeStamp.Equal(expectedTime) {
		t.Errorf("Expected TimeStamp to be %v, got %v", expectedTime, operation.TimeStamp)
	}

	t.Logf("All SaaSOperation struct fields parsed correctly")
}
