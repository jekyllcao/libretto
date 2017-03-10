package commitmentplans

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// GroupClient is the these APIs allow end users to operate on Azure Machine
// Learning Commitment Plans resources and their child Commitment Association
// resources. They support CRUD operations for commitment plans, get and list
// operations for commitment associations, moving commitment associations
// between commitment plans, and retrieving commitment plan usage history.
type GroupClient struct {
	ManagementClient
}

// NewGroupClient creates an instance of the GroupClient client.
func NewGroupClient(subscriptionID string) GroupClient {
	return NewGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGroupClientWithBaseURI creates an instance of the GroupClient client.
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return GroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create a new Azure ML commitment plan resource or updates an
// existing one.
//
// createOrUpdatePayload is the payload to create or update the Azure ML
// commitment plan. resourceGroupName is the resource group name.
// commitmentPlanName is the Azure ML commitment plan name.
func (client GroupClient) CreateOrUpdate(createOrUpdatePayload CommitmentPlan, resourceGroupName string, commitmentPlanName string) (result CommitmentPlan, err error) {
	req, err := client.CreateOrUpdatePreparer(createOrUpdatePayload, resourceGroupName, commitmentPlanName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client GroupClient) CreateOrUpdatePreparer(createOrUpdatePayload CommitmentPlan, resourceGroupName string, commitmentPlanName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"commitmentPlanName": autorest.Encode("path", commitmentPlanName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/commitmentPlans/{commitmentPlanName}", pathParameters),
		autorest.WithJSON(createOrUpdatePayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client GroupClient) CreateOrUpdateResponder(resp *http.Response) (result CommitmentPlan, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get retrieve an Azure ML commitment plan by its subscription, resource group
// and name.
//
// resourceGroupName is the resource group name. commitmentPlanName is the
// Azure ML commitment plan name.
func (client GroupClient) Get(resourceGroupName string, commitmentPlanName string) (result CommitmentPlan, err error) {
	req, err := client.GetPreparer(resourceGroupName, commitmentPlanName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client GroupClient) GetPreparer(resourceGroupName string, commitmentPlanName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"commitmentPlanName": autorest.Encode("path", commitmentPlanName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/commitmentPlans/{commitmentPlanName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GroupClient) GetResponder(resp *http.Response) (result CommitmentPlan, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieve all Azure ML commitment plans in a subscription.
//
// skipToken is continuation token for pagination.
func (client GroupClient) List(skipToken string) (result ListResult, err error) {
	req, err := client.ListPreparer(skipToken)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client GroupClient) ListPreparer(skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearning/commitmentPlans", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client GroupClient) ListResponder(resp *http.Response) (result ListResult, err error) {
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
func (client GroupClient) ListNextResults(lastResults ListResult) (result ListResult, err error) {
	req, err := lastResults.ListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListInResourceGroup retrieve all Azure ML commitment plans in a resource
// group.
//
// resourceGroupName is the resource group name. skipToken is continuation
// token for pagination.
func (client GroupClient) ListInResourceGroup(resourceGroupName string, skipToken string) (result ListResult, err error) {
	req, err := client.ListInResourceGroupPreparer(resourceGroupName, skipToken)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "ListInResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListInResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "ListInResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListInResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "ListInResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListInResourceGroupPreparer prepares the ListInResourceGroup request.
func (client GroupClient) ListInResourceGroupPreparer(resourceGroupName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/commitmentPlans", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListInResourceGroupSender sends the ListInResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) ListInResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListInResourceGroupResponder handles the response to the ListInResourceGroup request. The method always
// closes the http.Response Body.
func (client GroupClient) ListInResourceGroupResponder(resp *http.Response) (result ListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListInResourceGroupNextResults retrieves the next set of results, if any.
func (client GroupClient) ListInResourceGroupNextResults(lastResults ListResult) (result ListResult, err error) {
	req, err := lastResults.ListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "ListInResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListInResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "ListInResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListInResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "ListInResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// Patch patch an existing Azure ML commitment plan resource.
//
// patchPayload is the payload to use to patch the Azure ML commitment plan.
// Only tags and SKU may be modified on an existing commitment plan.
// resourceGroupName is the resource group name. commitmentPlanName is the
// Azure ML commitment plan name.
func (client GroupClient) Patch(patchPayload PatchPayload, resourceGroupName string, commitmentPlanName string) (result CommitmentPlan, err error) {
	req, err := client.PatchPreparer(patchPayload, resourceGroupName, commitmentPlanName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "Patch", nil, "Failure preparing request")
	}

	resp, err := client.PatchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "Patch", resp, "Failure sending request")
	}

	result, err = client.PatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "Patch", resp, "Failure responding to request")
	}

	return
}

// PatchPreparer prepares the Patch request.
func (client GroupClient) PatchPreparer(patchPayload PatchPayload, resourceGroupName string, commitmentPlanName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"commitmentPlanName": autorest.Encode("path", commitmentPlanName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/commitmentPlans/{commitmentPlanName}", pathParameters),
		autorest.WithJSON(patchPayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) PatchSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client GroupClient) PatchResponder(resp *http.Response) (result CommitmentPlan, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Remove remove an existing Azure ML commitment plan.
//
// resourceGroupName is the resource group name. commitmentPlanName is the
// Azure ML commitment plan name.
func (client GroupClient) Remove(resourceGroupName string, commitmentPlanName string) (result autorest.Response, err error) {
	req, err := client.RemovePreparer(resourceGroupName, commitmentPlanName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "Remove", nil, "Failure preparing request")
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "Remove", resp, "Failure sending request")
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.GroupClient", "Remove", resp, "Failure responding to request")
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client GroupClient) RemovePreparer(resourceGroupName string, commitmentPlanName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"commitmentPlanName": autorest.Encode("path", commitmentPlanName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/commitmentPlans/{commitmentPlanName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client GroupClient) RemoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
