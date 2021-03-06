package account

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// Client is the creates an Azure Data Lake Analytics account management
// client.
type Client struct {
	ManagementClient
}

// NewClient creates an instance of the Client client.
func NewClient(subscriptionID string) Client {
	return NewClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClientWithBaseURI creates an instance of the Client client.
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return Client{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates the specified Data Lake Analytics account. This supplies the
// user with computation services for Data Lake Analytics workloads This
// method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and
// any outstanding HTTP requests.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account.the account will be associated with. name is
// the name of the Data Lake Analytics account to create. parameters is
// parameters supplied to the create Data Lake Analytics account operation.
func (client Client) Create(resourceGroupName string, name string, parameters DataLakeAnalyticsAccount, cancel <-chan struct{}) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.DataLakeAnalyticsAccountProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.DataLakeAnalyticsAccountProperties.DefaultDataLakeStoreAccount", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.MaxDegreeOfParallelism", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.DataLakeAnalyticsAccountProperties.MaxDegreeOfParallelism", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.QueryStoreRetention", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.DataLakeAnalyticsAccountProperties.QueryStoreRetention", Name: validation.InclusiveMaximum, Rule: 180, Chain: nil},
							{Target: "parameters.DataLakeAnalyticsAccountProperties.QueryStoreRetention", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
						}},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.MaxJobCount", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.DataLakeAnalyticsAccountProperties.MaxJobCount", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.DataLakeStoreAccounts", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.ProvisioningState", Name: validation.ReadOnly, Rule: true, Chain: nil},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.State", Name: validation.ReadOnly, Rule: true, Chain: nil},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.SystemMaxDegreeOfParallelism", Name: validation.ReadOnly, Rule: true, Chain: nil},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.SystemMaxJobCount", Name: validation.ReadOnly, Rule: true, Chain: nil},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.CreationTime", Name: validation.ReadOnly, Rule: true, Chain: nil},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.LastModifiedTime", Name: validation.ReadOnly, Rule: true, Chain: nil},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.Endpoint", Name: validation.ReadOnly, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "account.Client", "Create")
	}

	req, err := client.CreatePreparer(resourceGroupName, name, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.Client", "Create", nil, "Failure preparing request")
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "account.Client", "Create", resp, "Failure sending request")
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.Client", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client Client) CreatePreparer(resourceGroupName string, name string, parameters DataLakeAnalyticsAccount, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{name}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client Client) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete begins the delete delete process for the Data Lake Analytics account
// object specified by the account name. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The
// channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. accountName is the name of the Data Lake
// Analytics account to delete
func (client Client) Delete(resourceGroupName string, accountName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, accountName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.Client", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "account.Client", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.Client", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client Client) DeletePreparer(resourceGroupName string, accountName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client Client) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client Client) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets details of the specified Data Lake Analytics account.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. accountName is the name of the Data Lake
// Analytics account to retrieve.
func (client Client) Get(resourceGroupName string, accountName string) (result DataLakeAnalyticsAccount, err error) {
	req, err := client.GetPreparer(resourceGroupName, accountName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.Client", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.Client", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.Client", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client Client) GetPreparer(resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client Client) GetResponder(resp *http.Response) (result DataLakeAnalyticsAccount, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the first page of Data Lake Analytics accounts, if any, within
// the current subscription. This includes a link to the next page, if any.
//
// filter is oData filter. Optional. top is the number of items to return.
// Optional. skip is the number of items to skip over before returning
// elements. Optional. selectParameter is oData Select statement. Limits the
// properties on each entry to just those requested, e.g.
// Categories?$select=CategoryName,Description. Optional. orderby is orderBy
// clause. One or more comma-separated expressions with an optional "asc"
// (the default) or "desc" depending on the order you'd like the values
// sorted, e.g. Categories?$orderby=CategoryName desc. Optional. count is the
// Boolean value of true or false to request a count of the matching
// resources included with the resources in the response, e.g.
// Categories?$count=true. Optional.
func (client Client) List(filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result DataLakeAnalyticsAccountListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "account.Client", "List")
	}

	req, err := client.ListPreparer(filter, top, skip, selectParameter, orderby, count)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.Client", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.Client", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.Client", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client Client) ListPreparer(filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if count != nil {
		queryParameters["$count"] = autorest.Encode("query", *count)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeAnalytics/accounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client Client) ListResponder(resp *http.Response) (result DataLakeAnalyticsAccountListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client Client) ListNextResults(lastResults DataLakeAnalyticsAccountListResult) (result DataLakeAnalyticsAccountListResult, err error) {
	req, err := lastResults.DataLakeAnalyticsAccountListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.Client", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.Client", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.Client", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListByResourceGroup gets the first page of Data Lake Analytics accounts, if
// any, within a specific resource group. This includes a link to the next
// page, if any.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. filter is oData filter. Optional. top is the
// number of items to return. Optional. skip is the number of items to skip
// over before returning elements. Optional. selectParameter is oData Select
// statement. Limits the properties on each entry to just those requested,
// e.g. Categories?$select=CategoryName,Description. Optional. orderby is
// orderBy clause. One or more comma-separated expressions with an optional
// "asc" (the default) or "desc" depending on the order you'd like the values
// sorted, e.g. Categories?$orderby=CategoryName desc. Optional. count is the
// Boolean value of true or false to request a count of the matching
// resources included with the resources in the response, e.g.
// Categories?$count=true. Optional.
func (client Client) ListByResourceGroup(resourceGroupName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result DataLakeAnalyticsAccountListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "account.Client", "ListByResourceGroup")
	}

	req, err := client.ListByResourceGroupPreparer(resourceGroupName, filter, top, skip, selectParameter, orderby, count)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.Client", "ListByResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.Client", "ListByResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.Client", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client Client) ListByResourceGroupPreparer(resourceGroupName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if count != nil {
		queryParameters["$count"] = autorest.Encode("query", *count)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client Client) ListByResourceGroupResponder(resp *http.Response) (result DataLakeAnalyticsAccountListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client Client) ListByResourceGroupNextResults(lastResults DataLakeAnalyticsAccountListResult) (result DataLakeAnalyticsAccountListResult, err error) {
	req, err := lastResults.DataLakeAnalyticsAccountListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.Client", "ListByResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.Client", "ListByResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.Client", "ListByResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// Update updates the Data Lake Analytics account object specified by the
// accountName with the contents of the account object. This method may poll
// for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. name is the name of the Data Lake Analytics
// account to update. parameters is parameters supplied to the update Data
// Lake Analytics account operation.
func (client Client) Update(resourceGroupName string, name string, parameters *DataLakeAnalyticsAccountUpdateParameters, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, name, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.Client", "Update", nil, "Failure preparing request")
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "account.Client", "Update", resp, "Failure sending request")
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.Client", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client Client) UpdatePreparer(resourceGroupName string, name string, parameters *DataLakeAnalyticsAccountUpdateParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client Client) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client Client) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}
