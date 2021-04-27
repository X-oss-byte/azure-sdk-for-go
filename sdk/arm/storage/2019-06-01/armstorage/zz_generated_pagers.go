// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// EncryptionScopeListResultPager provides iteration over EncryptionScopeListResult pages.
type EncryptionScopeListResultPager interface {
	azcore.Pager

	// Page returns the current EncryptionScopeListResultResponse.
	PageResponse() EncryptionScopeListResultResponse
}

type encryptionScopeListResultCreateRequest func(context.Context) (*azcore.Request, error)

type encryptionScopeListResultHandleError func(*azcore.Response) error

type encryptionScopeListResultHandleResponse func(*azcore.Response) (EncryptionScopeListResultResponse, error)

type encryptionScopeListResultAdvancePage func(context.Context, EncryptionScopeListResultResponse) (*azcore.Request, error)

type encryptionScopeListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester encryptionScopeListResultCreateRequest
	// callback for handling response errors
	errorer encryptionScopeListResultHandleError
	// callback for handling the HTTP response
	responder encryptionScopeListResultHandleResponse
	// callback for advancing to the next page
	advancer encryptionScopeListResultAdvancePage
	// contains the current response
	current EncryptionScopeListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *encryptionScopeListResultPager) Err() error {
	return p.err
}

func (p *encryptionScopeListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EncryptionScopeListResult.NextLink == nil || len(*p.current.EncryptionScopeListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *encryptionScopeListResultPager) PageResponse() EncryptionScopeListResultResponse {
	return p.current
}

// FileShareItemsPager provides iteration over FileShareItems pages.
type FileShareItemsPager interface {
	azcore.Pager

	// Page returns the current FileShareItemsResponse.
	PageResponse() FileShareItemsResponse
}

type fileShareItemsCreateRequest func(context.Context) (*azcore.Request, error)

type fileShareItemsHandleError func(*azcore.Response) error

type fileShareItemsHandleResponse func(*azcore.Response) (FileShareItemsResponse, error)

type fileShareItemsAdvancePage func(context.Context, FileShareItemsResponse) (*azcore.Request, error)

type fileShareItemsPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester fileShareItemsCreateRequest
	// callback for handling response errors
	errorer fileShareItemsHandleError
	// callback for handling the HTTP response
	responder fileShareItemsHandleResponse
	// callback for advancing to the next page
	advancer fileShareItemsAdvancePage
	// contains the current response
	current FileShareItemsResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *fileShareItemsPager) Err() error {
	return p.err
}

func (p *fileShareItemsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FileShareItems.NextLink == nil || len(*p.current.FileShareItems.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *fileShareItemsPager) PageResponse() FileShareItemsResponse {
	return p.current
}

// ListContainerItemsPager provides iteration over ListContainerItems pages.
type ListContainerItemsPager interface {
	azcore.Pager

	// Page returns the current ListContainerItemsResponse.
	PageResponse() ListContainerItemsResponse
}

type listContainerItemsCreateRequest func(context.Context) (*azcore.Request, error)

type listContainerItemsHandleError func(*azcore.Response) error

type listContainerItemsHandleResponse func(*azcore.Response) (ListContainerItemsResponse, error)

type listContainerItemsAdvancePage func(context.Context, ListContainerItemsResponse) (*azcore.Request, error)

type listContainerItemsPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester listContainerItemsCreateRequest
	// callback for handling response errors
	errorer listContainerItemsHandleError
	// callback for handling the HTTP response
	responder listContainerItemsHandleResponse
	// callback for advancing to the next page
	advancer listContainerItemsAdvancePage
	// contains the current response
	current ListContainerItemsResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *listContainerItemsPager) Err() error {
	return p.err
}

func (p *listContainerItemsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListContainerItems.NextLink == nil || len(*p.current.ListContainerItems.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *listContainerItemsPager) PageResponse() ListContainerItemsResponse {
	return p.current
}

// ListQueueResourcePager provides iteration over ListQueueResource pages.
type ListQueueResourcePager interface {
	azcore.Pager

	// Page returns the current ListQueueResourceResponse.
	PageResponse() ListQueueResourceResponse
}

type listQueueResourceCreateRequest func(context.Context) (*azcore.Request, error)

type listQueueResourceHandleError func(*azcore.Response) error

type listQueueResourceHandleResponse func(*azcore.Response) (ListQueueResourceResponse, error)

type listQueueResourceAdvancePage func(context.Context, ListQueueResourceResponse) (*azcore.Request, error)

type listQueueResourcePager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester listQueueResourceCreateRequest
	// callback for handling response errors
	errorer listQueueResourceHandleError
	// callback for handling the HTTP response
	responder listQueueResourceHandleResponse
	// callback for advancing to the next page
	advancer listQueueResourceAdvancePage
	// contains the current response
	current ListQueueResourceResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *listQueueResourcePager) Err() error {
	return p.err
}

func (p *listQueueResourcePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListQueueResource.NextLink == nil || len(*p.current.ListQueueResource.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *listQueueResourcePager) PageResponse() ListQueueResourceResponse {
	return p.current
}

// ListTableResourcePager provides iteration over ListTableResource pages.
type ListTableResourcePager interface {
	azcore.Pager

	// Page returns the current ListTableResourceResponse.
	PageResponse() ListTableResourceResponse
}

type listTableResourceCreateRequest func(context.Context) (*azcore.Request, error)

type listTableResourceHandleError func(*azcore.Response) error

type listTableResourceHandleResponse func(*azcore.Response) (ListTableResourceResponse, error)

type listTableResourceAdvancePage func(context.Context, ListTableResourceResponse) (*azcore.Request, error)

type listTableResourcePager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester listTableResourceCreateRequest
	// callback for handling response errors
	errorer listTableResourceHandleError
	// callback for handling the HTTP response
	responder listTableResourceHandleResponse
	// callback for advancing to the next page
	advancer listTableResourceAdvancePage
	// contains the current response
	current ListTableResourceResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *listTableResourcePager) Err() error {
	return p.err
}

func (p *listTableResourcePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListTableResource.NextLink == nil || len(*p.current.ListTableResource.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *listTableResourcePager) PageResponse() ListTableResourceResponse {
	return p.current
}

// StorageAccountListResultPager provides iteration over StorageAccountListResult pages.
type StorageAccountListResultPager interface {
	azcore.Pager

	// Page returns the current StorageAccountListResultResponse.
	PageResponse() StorageAccountListResultResponse
}

type storageAccountListResultCreateRequest func(context.Context) (*azcore.Request, error)

type storageAccountListResultHandleError func(*azcore.Response) error

type storageAccountListResultHandleResponse func(*azcore.Response) (StorageAccountListResultResponse, error)

type storageAccountListResultAdvancePage func(context.Context, StorageAccountListResultResponse) (*azcore.Request, error)

type storageAccountListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester storageAccountListResultCreateRequest
	// callback for handling response errors
	errorer storageAccountListResultHandleError
	// callback for handling the HTTP response
	responder storageAccountListResultHandleResponse
	// callback for advancing to the next page
	advancer storageAccountListResultAdvancePage
	// contains the current response
	current StorageAccountListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *storageAccountListResultPager) Err() error {
	return p.err
}

func (p *storageAccountListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.StorageAccountListResult.NextLink == nil || len(*p.current.StorageAccountListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *storageAccountListResultPager) PageResponse() StorageAccountListResultResponse {
	return p.current
}
