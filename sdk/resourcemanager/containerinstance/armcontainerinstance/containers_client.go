//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerinstance

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ContainersClient contains the methods for the Containers group.
// Don't use this type directly, use NewContainersClient() instead.
type ContainersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewContainersClient creates a new instance of ContainersClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainersClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ContainersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Attach - Attach to the output stream of a specific container instance in a specified resource group and container group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
// resourceGroupName - The name of the resource group.
// containerGroupName - The name of the container group.
// containerName - The name of the container instance.
// options - ContainersClientAttachOptions contains the optional parameters for the ContainersClient.Attach method.
func (client *ContainersClient) Attach(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, options *ContainersClientAttachOptions) (ContainersClientAttachResponse, error) {
	req, err := client.attachCreateRequest(ctx, resourceGroupName, containerGroupName, containerName, options)
	if err != nil {
		return ContainersClientAttachResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainersClientAttachResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainersClientAttachResponse{}, runtime.NewResponseError(resp)
	}
	return client.attachHandleResponse(resp)
}

// attachCreateRequest creates the Attach request.
func (client *ContainersClient) attachCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, options *ContainersClientAttachOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/attach"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// attachHandleResponse handles the Attach response.
func (client *ContainersClient) attachHandleResponse(resp *http.Response) (ContainersClientAttachResponse, error) {
	result := ContainersClientAttachResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerAttachResponse); err != nil {
		return ContainersClientAttachResponse{}, err
	}
	return result, nil
}

// ExecuteCommand - Executes a command for a specific container instance in a specified resource group and container group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
// resourceGroupName - The name of the resource group.
// containerGroupName - The name of the container group.
// containerName - The name of the container instance.
// containerExecRequest - The request for the exec command.
// options - ContainersClientExecuteCommandOptions contains the optional parameters for the ContainersClient.ExecuteCommand
// method.
func (client *ContainersClient) ExecuteCommand(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, containerExecRequest ContainerExecRequest, options *ContainersClientExecuteCommandOptions) (ContainersClientExecuteCommandResponse, error) {
	req, err := client.executeCommandCreateRequest(ctx, resourceGroupName, containerGroupName, containerName, containerExecRequest, options)
	if err != nil {
		return ContainersClientExecuteCommandResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainersClientExecuteCommandResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainersClientExecuteCommandResponse{}, runtime.NewResponseError(resp)
	}
	return client.executeCommandHandleResponse(resp)
}

// executeCommandCreateRequest creates the ExecuteCommand request.
func (client *ContainersClient) executeCommandCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, containerExecRequest ContainerExecRequest, options *ContainersClientExecuteCommandOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/exec"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, containerExecRequest)
}

// executeCommandHandleResponse handles the ExecuteCommand response.
func (client *ContainersClient) executeCommandHandleResponse(resp *http.Response) (ContainersClientExecuteCommandResponse, error) {
	result := ContainersClientExecuteCommandResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerExecResponse); err != nil {
		return ContainersClientExecuteCommandResponse{}, err
	}
	return result, nil
}

// ListLogs - Get the logs for a specified container instance in a specified resource group and container group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
// resourceGroupName - The name of the resource group.
// containerGroupName - The name of the container group.
// containerName - The name of the container instance.
// options - ContainersClientListLogsOptions contains the optional parameters for the ContainersClient.ListLogs method.
func (client *ContainersClient) ListLogs(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, options *ContainersClientListLogsOptions) (ContainersClientListLogsResponse, error) {
	req, err := client.listLogsCreateRequest(ctx, resourceGroupName, containerGroupName, containerName, options)
	if err != nil {
		return ContainersClientListLogsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainersClientListLogsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainersClientListLogsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listLogsHandleResponse(resp)
}

// listLogsCreateRequest creates the ListLogs request.
func (client *ContainersClient) listLogsCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, options *ContainersClientListLogsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/logs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	if options != nil && options.Tail != nil {
		reqQP.Set("tail", strconv.FormatInt(int64(*options.Tail), 10))
	}
	if options != nil && options.Timestamps != nil {
		reqQP.Set("timestamps", strconv.FormatBool(*options.Timestamps))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listLogsHandleResponse handles the ListLogs response.
func (client *ContainersClient) listLogsHandleResponse(resp *http.Response) (ContainersClientListLogsResponse, error) {
	result := ContainersClientListLogsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Logs); err != nil {
		return ContainersClientListLogsResponse{}, err
	}
	return result, nil
}
