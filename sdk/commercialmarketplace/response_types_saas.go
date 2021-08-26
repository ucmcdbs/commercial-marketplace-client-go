// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package commercialmarketplace

import "net/http"

// FulfillmentOperationsActivateSubscriptionResponse contains the response from method FulfillmentOperations.ActivateSubscription.
type FulfillmentOperationsActivateSubscriptionResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FulfillmentOperationsDeleteSubscriptionResponse contains the response from method FulfillmentOperations.DeleteSubscription.
type FulfillmentOperationsDeleteSubscriptionResponse struct {
	FulfillmentOperationsDeleteSubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FulfillmentOperationsDeleteSubscriptionResult contains the result from method FulfillmentOperations.DeleteSubscription.
type FulfillmentOperationsDeleteSubscriptionResult struct {
	// OperationLocationURI contains the information returned from the Operation-Location header response.
	OperationLocationURI *string
}

// FulfillmentOperationsGetSubscriptionResponse contains the response from method FulfillmentOperations.GetSubscription.
type FulfillmentOperationsGetSubscriptionResponse struct {
	FulfillmentOperationsGetSubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FulfillmentOperationsGetSubscriptionResult contains the result from method FulfillmentOperations.GetSubscription.
type FulfillmentOperationsGetSubscriptionResult struct {
	Subscription
}

// FulfillmentOperationsListAvailablePlansResponse contains the response from method FulfillmentOperations.ListAvailablePlans.
type FulfillmentOperationsListAvailablePlansResponse struct {
	FulfillmentOperationsListAvailablePlansResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FulfillmentOperationsListAvailablePlansResult contains the result from method FulfillmentOperations.ListAvailablePlans.
type FulfillmentOperationsListAvailablePlansResult struct {
	SubscriptionPlans
}

// FulfillmentOperationsListSubscriptionsResponse contains the response from method FulfillmentOperations.ListSubscriptions.
type FulfillmentOperationsListSubscriptionsResponse struct {
	FulfillmentOperationsListSubscriptionsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FulfillmentOperationsListSubscriptionsResult contains the result from method FulfillmentOperations.ListSubscriptions.
type FulfillmentOperationsListSubscriptionsResult struct {
	SubscriptionsResponse
}

// FulfillmentOperationsResolveResponse contains the response from method FulfillmentOperations.Resolve.
type FulfillmentOperationsResolveResponse struct {
	FulfillmentOperationsResolveResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FulfillmentOperationsResolveResult contains the result from method FulfillmentOperations.Resolve.
type FulfillmentOperationsResolveResult struct {
	ResolvedSubscription
}

// FulfillmentOperationsUpdateSubscriptionResponse contains the response from method FulfillmentOperations.UpdateSubscription.
type FulfillmentOperationsUpdateSubscriptionResponse struct {
	FulfillmentOperationsUpdateSubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FulfillmentOperationsUpdateSubscriptionResult contains the result from method FulfillmentOperations.UpdateSubscription.
type FulfillmentOperationsUpdateSubscriptionResult struct {
	// OperationLocationURI contains the information returned from the Operation-Location header response.
	OperationLocationURI *string
}

// SubscriptionOperationsGetOperationStatusResponse contains the response from method SubscriptionOperations.GetOperationStatus.
type SubscriptionOperationsGetOperationStatusResponse struct {
	SubscriptionOperationsGetOperationStatusResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionOperationsGetOperationStatusResult contains the result from method SubscriptionOperations.GetOperationStatus.
type SubscriptionOperationsGetOperationStatusResult struct {
	Operation
}

// SubscriptionOperationsListOperationsResponse contains the response from method SubscriptionOperations.ListOperations.
type SubscriptionOperationsListOperationsResponse struct {
	SubscriptionOperationsListOperationsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionOperationsListOperationsResult contains the result from method SubscriptionOperations.ListOperations.
type SubscriptionOperationsListOperationsResult struct {
	OperationList
}

// SubscriptionOperationsUpdateOperationStatusResponse contains the response from method SubscriptionOperations.UpdateOperationStatus.
type SubscriptionOperationsUpdateOperationStatusResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

