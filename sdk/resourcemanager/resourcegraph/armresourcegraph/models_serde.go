//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcegraph

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Column.
func (c Column) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Column.
func (c *Column) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &c.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DateTimeInterval.
func (d DateTimeInterval) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeRFC3339(objectMap, "end", d.End)
	populateTimeRFC3339(objectMap, "start", d.Start)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DateTimeInterval.
func (d *DateTimeInterval) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "end":
			err = unpopulateTimeRFC3339(val, "End", &d.End)
			delete(rawMsg, key)
		case "start":
			err = unpopulateTimeRFC3339(val, "Start", &d.Start)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Error.
func (e Error) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Error.
func (e *Error) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &e.Details)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetails.
func (e ErrorDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	if e.AdditionalProperties != nil {
		for key, val := range e.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetails.
func (e *ErrorDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		default:
			if e.AdditionalProperties == nil {
				e.AdditionalProperties = map[string]any{}
			}
			if val != nil {
				var aux any
				err = json.Unmarshal(val, &aux)
				e.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &e.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Facet.
func (f Facet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "expression", f.Expression)
	objectMap["resultType"] = f.ResultType
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Facet.
func (f *Facet) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "expression":
			err = unpopulate(val, "Expression", &f.Expression)
			delete(rawMsg, key)
		case "resultType":
			err = unpopulate(val, "ResultType", &f.ResultType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FacetError.
func (f FacetError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "errors", f.Errors)
	populate(objectMap, "expression", f.Expression)
	objectMap["resultType"] = "FacetError"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FacetError.
func (f *FacetError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "errors":
			err = unpopulate(val, "Errors", &f.Errors)
			delete(rawMsg, key)
		case "expression":
			err = unpopulate(val, "Expression", &f.Expression)
			delete(rawMsg, key)
		case "resultType":
			err = unpopulate(val, "ResultType", &f.ResultType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FacetRequest.
func (f FacetRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "expression", f.Expression)
	populate(objectMap, "options", f.Options)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FacetRequest.
func (f *FacetRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "expression":
			err = unpopulate(val, "Expression", &f.Expression)
			delete(rawMsg, key)
		case "options":
			err = unpopulate(val, "Options", &f.Options)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FacetRequestOptions.
func (f FacetRequestOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "filter", f.Filter)
	populate(objectMap, "sortBy", f.SortBy)
	populate(objectMap, "sortOrder", f.SortOrder)
	populate(objectMap, "$top", f.Top)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FacetRequestOptions.
func (f *FacetRequestOptions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "filter":
			err = unpopulate(val, "Filter", &f.Filter)
			delete(rawMsg, key)
		case "sortBy":
			err = unpopulate(val, "SortBy", &f.SortBy)
			delete(rawMsg, key)
		case "sortOrder":
			err = unpopulate(val, "SortOrder", &f.SortOrder)
			delete(rawMsg, key)
		case "$top":
			err = unpopulate(val, "Top", &f.Top)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FacetResult.
func (f FacetResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "count", f.Count)
	populateAny(objectMap, "data", f.Data)
	populate(objectMap, "expression", f.Expression)
	objectMap["resultType"] = "FacetResult"
	populate(objectMap, "totalRecords", f.TotalRecords)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FacetResult.
func (f *FacetResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "count":
			err = unpopulate(val, "Count", &f.Count)
			delete(rawMsg, key)
		case "data":
			err = unpopulate(val, "Data", &f.Data)
			delete(rawMsg, key)
		case "expression":
			err = unpopulate(val, "Expression", &f.Expression)
			delete(rawMsg, key)
		case "resultType":
			err = unpopulate(val, "ResultType", &f.ResultType)
			delete(rawMsg, key)
		case "totalRecords":
			err = unpopulate(val, "TotalRecords", &f.TotalRecords)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "origin", o.Origin)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Operation.
func (o *Operation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "display":
			err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "origin":
			err = unpopulate(val, "Origin", &o.Origin)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", o.Description)
	populate(objectMap, "operation", o.Operation)
	populate(objectMap, "provider", o.Provider)
	populate(objectMap, "resource", o.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationDisplay.
func (o *OperationDisplay) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &o.Description)
			delete(rawMsg, key)
		case "operation":
			err = unpopulate(val, "Operation", &o.Operation)
			delete(rawMsg, key)
		case "provider":
			err = unpopulate(val, "Provider", &o.Provider)
			delete(rawMsg, key)
		case "resource":
			err = unpopulate(val, "Resource", &o.Resource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationListResult.
func (o *OperationListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &o.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type QueryRequest.
func (q QueryRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "facets", q.Facets)
	populate(objectMap, "managementGroups", q.ManagementGroups)
	populate(objectMap, "options", q.Options)
	populate(objectMap, "query", q.Query)
	populate(objectMap, "subscriptions", q.Subscriptions)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type QueryRequest.
func (q *QueryRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", q, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "facets":
			err = unpopulate(val, "Facets", &q.Facets)
			delete(rawMsg, key)
		case "managementGroups":
			err = unpopulate(val, "ManagementGroups", &q.ManagementGroups)
			delete(rawMsg, key)
		case "options":
			err = unpopulate(val, "Options", &q.Options)
			delete(rawMsg, key)
		case "query":
			err = unpopulate(val, "Query", &q.Query)
			delete(rawMsg, key)
		case "subscriptions":
			err = unpopulate(val, "Subscriptions", &q.Subscriptions)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", q, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type QueryRequestOptions.
func (q QueryRequestOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "allowPartialScopes", q.AllowPartialScopes)
	populate(objectMap, "authorizationScopeFilter", q.AuthorizationScopeFilter)
	populate(objectMap, "resultFormat", q.ResultFormat)
	populate(objectMap, "$skip", q.Skip)
	populate(objectMap, "$skipToken", q.SkipToken)
	populate(objectMap, "$top", q.Top)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type QueryRequestOptions.
func (q *QueryRequestOptions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", q, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "allowPartialScopes":
			err = unpopulate(val, "AllowPartialScopes", &q.AllowPartialScopes)
			delete(rawMsg, key)
		case "authorizationScopeFilter":
			err = unpopulate(val, "AuthorizationScopeFilter", &q.AuthorizationScopeFilter)
			delete(rawMsg, key)
		case "resultFormat":
			err = unpopulate(val, "ResultFormat", &q.ResultFormat)
			delete(rawMsg, key)
		case "$skip":
			err = unpopulate(val, "Skip", &q.Skip)
			delete(rawMsg, key)
		case "$skipToken":
			err = unpopulate(val, "SkipToken", &q.SkipToken)
			delete(rawMsg, key)
		case "$top":
			err = unpopulate(val, "Top", &q.Top)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", q, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type QueryResponse.
func (q QueryResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "count", q.Count)
	populateAny(objectMap, "data", q.Data)
	populate(objectMap, "facets", q.Facets)
	populate(objectMap, "resultTruncated", q.ResultTruncated)
	populate(objectMap, "$skipToken", q.SkipToken)
	populate(objectMap, "totalRecords", q.TotalRecords)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type QueryResponse.
func (q *QueryResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", q, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "count":
			err = unpopulate(val, "Count", &q.Count)
			delete(rawMsg, key)
		case "data":
			err = unpopulate(val, "Data", &q.Data)
			delete(rawMsg, key)
		case "facets":
			q.Facets, err = unmarshalFacetClassificationArray(val)
			delete(rawMsg, key)
		case "resultTruncated":
			err = unpopulate(val, "ResultTruncated", &q.ResultTruncated)
			delete(rawMsg, key)
		case "$skipToken":
			err = unpopulate(val, "SkipToken", &q.SkipToken)
			delete(rawMsg, key)
		case "totalRecords":
			err = unpopulate(val, "TotalRecords", &q.TotalRecords)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", q, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourcesHistoryRequest.
func (r ResourcesHistoryRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "managementGroups", r.ManagementGroups)
	populate(objectMap, "options", r.Options)
	populate(objectMap, "query", r.Query)
	populate(objectMap, "subscriptions", r.Subscriptions)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourcesHistoryRequest.
func (r *ResourcesHistoryRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "managementGroups":
			err = unpopulate(val, "ManagementGroups", &r.ManagementGroups)
			delete(rawMsg, key)
		case "options":
			err = unpopulate(val, "Options", &r.Options)
			delete(rawMsg, key)
		case "query":
			err = unpopulate(val, "Query", &r.Query)
			delete(rawMsg, key)
		case "subscriptions":
			err = unpopulate(val, "Subscriptions", &r.Subscriptions)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourcesHistoryRequestOptions.
func (r ResourcesHistoryRequestOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "interval", r.Interval)
	populate(objectMap, "resultFormat", r.ResultFormat)
	populate(objectMap, "$skip", r.Skip)
	populate(objectMap, "$skipToken", r.SkipToken)
	populate(objectMap, "$top", r.Top)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourcesHistoryRequestOptions.
func (r *ResourcesHistoryRequestOptions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "interval":
			err = unpopulate(val, "Interval", &r.Interval)
			delete(rawMsg, key)
		case "resultFormat":
			err = unpopulate(val, "ResultFormat", &r.ResultFormat)
			delete(rawMsg, key)
		case "$skip":
			err = unpopulate(val, "Skip", &r.Skip)
			delete(rawMsg, key)
		case "$skipToken":
			err = unpopulate(val, "SkipToken", &r.SkipToken)
			delete(rawMsg, key)
		case "$top":
			err = unpopulate(val, "Top", &r.Top)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Table.
func (t Table) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "columns", t.Columns)
	populate(objectMap, "rows", t.Rows)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Table.
func (t *Table) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "columns":
			err = unpopulate(val, "Columns", &t.Columns)
			delete(rawMsg, key)
		case "rows":
			err = unpopulate(val, "Rows", &t.Rows)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateAny(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
