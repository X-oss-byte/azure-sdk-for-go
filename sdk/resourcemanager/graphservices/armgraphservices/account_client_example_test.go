//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armgraphservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/graphservices/armgraphservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/graphservicesprod/resource-manager/Microsoft.GraphServices/preview/2022-09-22-preview/examples/Accounts_Get.json
func ExampleAccountClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armgraphservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountClient().Get(ctx, "testResourceGroupGRAM", "11111111-aaaa-1111-bbbb-111111111111", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccountResource = armgraphservices.AccountResource{
	// 	Name: to.Ptr("11111111-aaaa-1111-bbbb-111111111111"),
	// 	Type: to.Ptr("Microsoft.GraphServices/accounts"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testResourceGroupGRAM/providers/Microsoft.GraphServices/accounts/11111111-aaaa-1111-bbbb-111111111111"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armgraphservices.AccountResourceProperties{
	// 		AppID: to.Ptr("11111111-aaaa-1111-bbbb-111111111111"),
	// 		ProvisioningState: to.Ptr(armgraphservices.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/graphservicesprod/resource-manager/Microsoft.GraphServices/preview/2022-09-22-preview/examples/Accounts_Create.json
func ExampleAccountClient_BeginCreateAndUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armgraphservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountClient().BeginCreateAndUpdate(ctx, "testResourceGroupGRAM", "11111111-aaaa-1111-bbbb-1111111111111", armgraphservices.AccountResource{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
		Properties: &armgraphservices.AccountResourceProperties{
			AppID: to.Ptr("11111111-aaaa-1111-bbbb-111111111111"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccountResource = armgraphservices.AccountResource{
	// 	Name: to.Ptr("11111111-aaaa-1111-bbbb-111111111111"),
	// 	Type: to.Ptr("Microsoft.GraphServices/accounts"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testResourceGroupGRAM/providers/Microsoft.GraphServices/accounts/11111111-aaaa-1111-bbbb-111111111111"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armgraphservices.AccountResourceProperties{
	// 		AppID: to.Ptr("11111111-aaaa-1111-bbbb-111111111111"),
	// 		BillingPlanID: to.Ptr("11111111-aaaa-1111-bbbb-111111111111"),
	// 		ProvisioningState: to.Ptr(armgraphservices.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/graphservicesprod/resource-manager/Microsoft.GraphServices/preview/2022-09-22-preview/examples/Accounts_Update.json
func ExampleAccountClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armgraphservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountClient().Update(ctx, "testResourceGroupGRAM", "11111111-aaaa-1111-bbbb-111111111111", armgraphservices.AccountPatchResource{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccountResource = armgraphservices.AccountResource{
	// 	Name: to.Ptr("11111111-aaaa-1111-bbbb-111111111111"),
	// 	Type: to.Ptr("Microsoft.GraphServices/accounts"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testResourceGroupGRAM/providers/Microsoft.GraphServices/accounts/11111111-aaaa-1111-bbbb-111111111111"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armgraphservices.AccountResourceProperties{
	// 		AppID: to.Ptr("11111111-aaaa-1111-bbbb-111111111111"),
	// 		BillingPlanID: to.Ptr("11111111-aaaa-1111-bbbb-111111111111"),
	// 		ProvisioningState: to.Ptr(armgraphservices.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/graphservicesprod/resource-manager/Microsoft.GraphServices/preview/2022-09-22-preview/examples/Accounts_Delete.json
func ExampleAccountClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armgraphservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAccountClient().Delete(ctx, "testResourceGroupGRAM", "11111111-aaaa-1111-bbbb-111111111111", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
