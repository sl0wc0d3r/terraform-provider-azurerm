package artifacts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// SparkJobDefinitionClient is the client for the SparkJobDefinition methods of the Artifacts service.
type SparkJobDefinitionClient struct {
	BaseClient
}

// NewSparkJobDefinitionClient creates an instance of the SparkJobDefinitionClient client.
func NewSparkJobDefinitionClient(endpoint string) SparkJobDefinitionClient {
	return SparkJobDefinitionClient{New(endpoint)}
}

// CreateOrUpdateSparkJobDefinition creates or updates a Spark Job Definition.
// Parameters:
// sparkJobDefinitionName - the spark job definition name.
// sparkJobDefinition - spark Job Definition resource definition.
// ifMatch - eTag of the Spark Job Definition entity.  Should only be specified for update, for which it should
// match existing entity or can be * for unconditional update.
func (client SparkJobDefinitionClient) CreateOrUpdateSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, ifMatch string) (result SparkJobDefinitionCreateOrUpdateSparkJobDefinitionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SparkJobDefinitionClient.CreateOrUpdateSparkJobDefinition")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: sparkJobDefinition,
			Constraints: []validation.Constraint{{Target: "sparkJobDefinition.Properties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "sparkJobDefinition.Properties.TargetBigDataPool", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "sparkJobDefinition.Properties.TargetBigDataPool.Type", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "sparkJobDefinition.Properties.TargetBigDataPool.ReferenceName", Name: validation.Null, Rule: true, Chain: nil},
					}},
					{Target: "sparkJobDefinition.Properties.JobProperties", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "sparkJobDefinition.Properties.JobProperties.File", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "sparkJobDefinition.Properties.JobProperties.DriverMemory", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "sparkJobDefinition.Properties.JobProperties.DriverCores", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "sparkJobDefinition.Properties.JobProperties.ExecutorMemory", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "sparkJobDefinition.Properties.JobProperties.ExecutorCores", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "sparkJobDefinition.Properties.JobProperties.NumExecutors", Name: validation.Null, Rule: true, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("artifacts.SparkJobDefinitionClient", "CreateOrUpdateSparkJobDefinition", err.Error())
	}

	req, err := client.CreateOrUpdateSparkJobDefinitionPreparer(ctx, sparkJobDefinitionName, sparkJobDefinition, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "CreateOrUpdateSparkJobDefinition", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSparkJobDefinitionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "CreateOrUpdateSparkJobDefinition", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateSparkJobDefinitionPreparer prepares the CreateOrUpdateSparkJobDefinition request.
func (client SparkJobDefinitionClient) CreateOrUpdateSparkJobDefinitionPreparer(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, ifMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"sparkJobDefinitionName": autorest.Encode("path", sparkJobDefinitionName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/sparkJobDefinitions/{sparkJobDefinitionName}", pathParameters),
		autorest.WithJSON(sparkJobDefinition),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSparkJobDefinitionSender sends the CreateOrUpdateSparkJobDefinition request. The method will close the
// http.Response Body if it receives an error.
func (client SparkJobDefinitionClient) CreateOrUpdateSparkJobDefinitionSender(req *http.Request) (future SparkJobDefinitionCreateOrUpdateSparkJobDefinitionFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateSparkJobDefinitionResponder handles the response to the CreateOrUpdateSparkJobDefinition request. The method always
// closes the http.Response Body.
func (client SparkJobDefinitionClient) CreateOrUpdateSparkJobDefinitionResponder(resp *http.Response) (result SparkJobDefinitionResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DebugSparkJobDefinition debug the spark job definition.
// Parameters:
// sparkJobDefinitionAzureResource - spark Job Definition resource definition.
func (client SparkJobDefinitionClient) DebugSparkJobDefinition(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource) (result SparkJobDefinitionDebugSparkJobDefinitionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SparkJobDefinitionClient.DebugSparkJobDefinition")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: sparkJobDefinitionAzureResource,
			Constraints: []validation.Constraint{{Target: "sparkJobDefinitionAzureResource.Properties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "sparkJobDefinitionAzureResource.Properties.TargetBigDataPool", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "sparkJobDefinitionAzureResource.Properties.TargetBigDataPool.Type", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "sparkJobDefinitionAzureResource.Properties.TargetBigDataPool.ReferenceName", Name: validation.Null, Rule: true, Chain: nil},
					}},
					{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties.File", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties.DriverMemory", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties.DriverCores", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties.ExecutorMemory", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties.ExecutorCores", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "sparkJobDefinitionAzureResource.Properties.JobProperties.NumExecutors", Name: validation.Null, Rule: true, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("artifacts.SparkJobDefinitionClient", "DebugSparkJobDefinition", err.Error())
	}

	req, err := client.DebugSparkJobDefinitionPreparer(ctx, sparkJobDefinitionAzureResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "DebugSparkJobDefinition", nil, "Failure preparing request")
		return
	}

	result, err = client.DebugSparkJobDefinitionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "DebugSparkJobDefinition", nil, "Failure sending request")
		return
	}

	return
}

// DebugSparkJobDefinitionPreparer prepares the DebugSparkJobDefinition request.
func (client SparkJobDefinitionClient) DebugSparkJobDefinitionPreparer(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"endpoint": client.Endpoint,
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPath("/debugSparkJobDefinition"),
		autorest.WithJSON(sparkJobDefinitionAzureResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DebugSparkJobDefinitionSender sends the DebugSparkJobDefinition request. The method will close the
// http.Response Body if it receives an error.
func (client SparkJobDefinitionClient) DebugSparkJobDefinitionSender(req *http.Request) (future SparkJobDefinitionDebugSparkJobDefinitionFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DebugSparkJobDefinitionResponder handles the response to the DebugSparkJobDefinition request. The method always
// closes the http.Response Body.
func (client SparkJobDefinitionClient) DebugSparkJobDefinitionResponder(resp *http.Response) (result SparkBatchJob, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteSparkJobDefinition deletes a Spark Job Definition.
// Parameters:
// sparkJobDefinitionName - the spark job definition name.
func (client SparkJobDefinitionClient) DeleteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string) (result SparkJobDefinitionDeleteSparkJobDefinitionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SparkJobDefinitionClient.DeleteSparkJobDefinition")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteSparkJobDefinitionPreparer(ctx, sparkJobDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "DeleteSparkJobDefinition", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSparkJobDefinitionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "DeleteSparkJobDefinition", nil, "Failure sending request")
		return
	}

	return
}

// DeleteSparkJobDefinitionPreparer prepares the DeleteSparkJobDefinition request.
func (client SparkJobDefinitionClient) DeleteSparkJobDefinitionPreparer(ctx context.Context, sparkJobDefinitionName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"sparkJobDefinitionName": autorest.Encode("path", sparkJobDefinitionName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/sparkJobDefinitions/{sparkJobDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSparkJobDefinitionSender sends the DeleteSparkJobDefinition request. The method will close the
// http.Response Body if it receives an error.
func (client SparkJobDefinitionClient) DeleteSparkJobDefinitionSender(req *http.Request) (future SparkJobDefinitionDeleteSparkJobDefinitionFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteSparkJobDefinitionResponder handles the response to the DeleteSparkJobDefinition request. The method always
// closes the http.Response Body.
func (client SparkJobDefinitionClient) DeleteSparkJobDefinitionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ExecuteSparkJobDefinition executes the spark job definition.
// Parameters:
// sparkJobDefinitionName - the spark job definition name.
func (client SparkJobDefinitionClient) ExecuteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string) (result SparkJobDefinitionExecuteSparkJobDefinitionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SparkJobDefinitionClient.ExecuteSparkJobDefinition")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ExecuteSparkJobDefinitionPreparer(ctx, sparkJobDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "ExecuteSparkJobDefinition", nil, "Failure preparing request")
		return
	}

	result, err = client.ExecuteSparkJobDefinitionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "ExecuteSparkJobDefinition", nil, "Failure sending request")
		return
	}

	return
}

// ExecuteSparkJobDefinitionPreparer prepares the ExecuteSparkJobDefinition request.
func (client SparkJobDefinitionClient) ExecuteSparkJobDefinitionPreparer(ctx context.Context, sparkJobDefinitionName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"sparkJobDefinitionName": autorest.Encode("path", sparkJobDefinitionName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/sparkJobDefinitions/{sparkJobDefinitionName}/execute", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteSparkJobDefinitionSender sends the ExecuteSparkJobDefinition request. The method will close the
// http.Response Body if it receives an error.
func (client SparkJobDefinitionClient) ExecuteSparkJobDefinitionSender(req *http.Request) (future SparkJobDefinitionExecuteSparkJobDefinitionFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// ExecuteSparkJobDefinitionResponder handles the response to the ExecuteSparkJobDefinition request. The method always
// closes the http.Response Body.
func (client SparkJobDefinitionClient) ExecuteSparkJobDefinitionResponder(resp *http.Response) (result SparkBatchJob, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSparkJobDefinition gets a Spark Job Definition.
// Parameters:
// sparkJobDefinitionName - the spark job definition name.
// ifNoneMatch - eTag of the Spark Job Definition entity. Should only be specified for get. If the ETag matches
// the existing entity tag, or if * was provided, then no content will be returned.
func (client SparkJobDefinitionClient) GetSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, ifNoneMatch string) (result SparkJobDefinitionResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SparkJobDefinitionClient.GetSparkJobDefinition")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSparkJobDefinitionPreparer(ctx, sparkJobDefinitionName, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "GetSparkJobDefinition", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSparkJobDefinitionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "GetSparkJobDefinition", resp, "Failure sending request")
		return
	}

	result, err = client.GetSparkJobDefinitionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "GetSparkJobDefinition", resp, "Failure responding to request")
		return
	}

	return
}

// GetSparkJobDefinitionPreparer prepares the GetSparkJobDefinition request.
func (client SparkJobDefinitionClient) GetSparkJobDefinitionPreparer(ctx context.Context, sparkJobDefinitionName string, ifNoneMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"sparkJobDefinitionName": autorest.Encode("path", sparkJobDefinitionName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/sparkJobDefinitions/{sparkJobDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSparkJobDefinitionSender sends the GetSparkJobDefinition request. The method will close the
// http.Response Body if it receives an error.
func (client SparkJobDefinitionClient) GetSparkJobDefinitionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetSparkJobDefinitionResponder handles the response to the GetSparkJobDefinition request. The method always
// closes the http.Response Body.
func (client SparkJobDefinitionClient) GetSparkJobDefinitionResponder(resp *http.Response) (result SparkJobDefinitionResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotModified),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSparkJobDefinitionsByWorkspace lists spark job definitions.
func (client SparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspace(ctx context.Context) (result SparkJobDefinitionsListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SparkJobDefinitionClient.GetSparkJobDefinitionsByWorkspace")
		defer func() {
			sc := -1
			if result.sjdlr.Response.Response != nil {
				sc = result.sjdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getSparkJobDefinitionsByWorkspaceNextResults
	req, err := client.GetSparkJobDefinitionsByWorkspacePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "GetSparkJobDefinitionsByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSparkJobDefinitionsByWorkspaceSender(req)
	if err != nil {
		result.sjdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "GetSparkJobDefinitionsByWorkspace", resp, "Failure sending request")
		return
	}

	result.sjdlr, err = client.GetSparkJobDefinitionsByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "GetSparkJobDefinitionsByWorkspace", resp, "Failure responding to request")
		return
	}
	if result.sjdlr.hasNextLink() && result.sjdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetSparkJobDefinitionsByWorkspacePreparer prepares the GetSparkJobDefinitionsByWorkspace request.
func (client SparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspacePreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"endpoint": client.Endpoint,
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPath("/sparkJobDefinitions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSparkJobDefinitionsByWorkspaceSender sends the GetSparkJobDefinitionsByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client SparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetSparkJobDefinitionsByWorkspaceResponder handles the response to the GetSparkJobDefinitionsByWorkspace request. The method always
// closes the http.Response Body.
func (client SparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspaceResponder(resp *http.Response) (result SparkJobDefinitionsListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getSparkJobDefinitionsByWorkspaceNextResults retrieves the next set of results, if any.
func (client SparkJobDefinitionClient) getSparkJobDefinitionsByWorkspaceNextResults(ctx context.Context, lastResults SparkJobDefinitionsListResponse) (result SparkJobDefinitionsListResponse, err error) {
	req, err := lastResults.sparkJobDefinitionsListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "getSparkJobDefinitionsByWorkspaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetSparkJobDefinitionsByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "getSparkJobDefinitionsByWorkspaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetSparkJobDefinitionsByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "getSparkJobDefinitionsByWorkspaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetSparkJobDefinitionsByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client SparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspaceComplete(ctx context.Context) (result SparkJobDefinitionsListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SparkJobDefinitionClient.GetSparkJobDefinitionsByWorkspace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetSparkJobDefinitionsByWorkspace(ctx)
	return
}

// RenameSparkJobDefinition renames a sparkJobDefinition.
// Parameters:
// sparkJobDefinitionName - the spark job definition name.
// request - proposed new name.
func (client SparkJobDefinitionClient) RenameSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, request RenameRequest) (result SparkJobDefinitionRenameSparkJobDefinitionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SparkJobDefinitionClient.RenameSparkJobDefinition")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: request,
			Constraints: []validation.Constraint{{Target: "request.NewName", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "request.NewName", Name: validation.MaxLength, Rule: 260, Chain: nil},
					{Target: "request.NewName", Name: validation.MinLength, Rule: 1, Chain: nil},
					{Target: "request.NewName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("artifacts.SparkJobDefinitionClient", "RenameSparkJobDefinition", err.Error())
	}

	req, err := client.RenameSparkJobDefinitionPreparer(ctx, sparkJobDefinitionName, request)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "RenameSparkJobDefinition", nil, "Failure preparing request")
		return
	}

	result, err = client.RenameSparkJobDefinitionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.SparkJobDefinitionClient", "RenameSparkJobDefinition", nil, "Failure sending request")
		return
	}

	return
}

// RenameSparkJobDefinitionPreparer prepares the RenameSparkJobDefinition request.
func (client SparkJobDefinitionClient) RenameSparkJobDefinitionPreparer(ctx context.Context, sparkJobDefinitionName string, request RenameRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"sparkJobDefinitionName": autorest.Encode("path", sparkJobDefinitionName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/sparkJobDefinitions/{sparkJobDefinitionName}/rename", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RenameSparkJobDefinitionSender sends the RenameSparkJobDefinition request. The method will close the
// http.Response Body if it receives an error.
func (client SparkJobDefinitionClient) RenameSparkJobDefinitionSender(req *http.Request) (future SparkJobDefinitionRenameSparkJobDefinitionFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// RenameSparkJobDefinitionResponder handles the response to the RenameSparkJobDefinition request. The method always
// closes the http.Response Body.
func (client SparkJobDefinitionClient) RenameSparkJobDefinitionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
