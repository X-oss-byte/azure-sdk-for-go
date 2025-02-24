//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armqumulo_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/liftrqumulo/armqumulo"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9918780b4dc4bdc111cf3facc11561904d609ad7/specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_ListBySubscription_MaximumSet_Gen.json
func ExampleFileSystemsClient_NewListBySubscriptionPager_fileSystemsListBySubscriptionMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSystemsClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.FileSystemResourceListResult = armqumulo.FileSystemResourceListResult{
		// 	Value: []*armqumulo.FileSystemResource{
		// 		{
		// 			Name: to.Ptr("bii"),
		// 			Type: to.Ptr("qtvxrqwpoistduq"),
		// 			ID: to.Ptr("tvelgpobdtazrweunifqzaxkgjauyx"),
		// 			SystemData: &armqumulo.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T04:40:17.991Z"); return t}()),
		// 				CreatedBy: to.Ptr("mtdhqooysjhueaojwpmvophkgntl"),
		// 				CreatedByType: to.Ptr(armqumulo.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T04:40:17.991Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("jcywglomzuamsxltnoegdrkzlscxl"),
		// 				LastModifiedByType: to.Ptr(armqumulo.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("przdlsmlzsszphnixq"),
		// 			Tags: map[string]*string{
		// 				"key6565": to.Ptr("cgdhmupta"),
		// 			},
		// 			Identity: &armqumulo.ManagedServiceIdentity{
		// 				Type: to.Ptr(armqumulo.ManagedServiceIdentityTypeNone),
		// 				PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				UserAssignedIdentities: map[string]*armqumulo.UserAssignedIdentity{
		// 					"key4522": &armqumulo.UserAssignedIdentity{
		// 						ClientID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 						PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armqumulo.FileSystemResourceProperties{
		// 				AvailabilityZone: to.Ptr("maseyqhlnhoiwbabcqabtedbjpip"),
		// 				ClusterLoginURL: to.Ptr("jjqhgevy"),
		// 				DelegatedSubnetID: to.Ptr("neqctctqdmjezfgt"),
		// 				InitialCapacity: to.Ptr[int32](9),
		// 				MarketplaceDetails: &armqumulo.MarketplaceDetails{
		// 					MarketplaceSubscriptionID: to.Ptr("ujrcqvxfnhxxheoth"),
		// 					MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 					OfferID: to.Ptr("eiyhbmpwgezcmzrrfoiskuxlcvwojf"),
		// 					PlanID: to.Ptr("x"),
		// 					PublisherID: to.Ptr("wfmokfdjbwpjhz"),
		// 				},
		// 				PrivateIPs: []*string{
		// 					to.Ptr("kslguxrwbwkrj")},
		// 					ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
		// 					StorageSKU: to.Ptr(armqumulo.StorageSKUStandard),
		// 					UserDetails: &armqumulo.UserDetails{
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9918780b4dc4bdc111cf3facc11561904d609ad7/specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_ListBySubscription_MinimumSet_Gen.json
func ExampleFileSystemsClient_NewListBySubscriptionPager_fileSystemsListBySubscriptionMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSystemsClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.FileSystemResourceListResult = armqumulo.FileSystemResourceListResult{
		// 	Value: []*armqumulo.FileSystemResource{
		// 		{
		// 			Name: to.Ptr("aaaaa"),
		// 			ID: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 			Location: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			Properties: &armqumulo.FileSystemResourceProperties{
		// 				DelegatedSubnetID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				InitialCapacity: to.Ptr[int32](9),
		// 				MarketplaceDetails: &armqumulo.MarketplaceDetails{
		// 					MarketplaceSubscriptionID: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 					MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 					OfferID: to.Ptr("aaaaaaaaa"),
		// 					PlanID: to.Ptr("aaaaaaa"),
		// 					PublisherID: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 				},
		// 				ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
		// 				StorageSKU: to.Ptr(armqumulo.StorageSKUStandard),
		// 				UserDetails: &armqumulo.UserDetails{
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9918780b4dc4bdc111cf3facc11561904d609ad7/specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_ListByResourceGroup_MaximumSet_Gen.json
func ExampleFileSystemsClient_NewListByResourceGroupPager_fileSystemsListByResourceGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSystemsClient().NewListByResourceGroupPager("rgQumulo", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.FileSystemResourceListResult = armqumulo.FileSystemResourceListResult{
		// 	Value: []*armqumulo.FileSystemResource{
		// 		{
		// 			Name: to.Ptr("bii"),
		// 			Type: to.Ptr("qtvxrqwpoistduq"),
		// 			ID: to.Ptr("tvelgpobdtazrweunifqzaxkgjauyx"),
		// 			SystemData: &armqumulo.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T04:40:17.991Z"); return t}()),
		// 				CreatedBy: to.Ptr("mtdhqooysjhueaojwpmvophkgntl"),
		// 				CreatedByType: to.Ptr(armqumulo.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T04:40:17.991Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("jcywglomzuamsxltnoegdrkzlscxl"),
		// 				LastModifiedByType: to.Ptr(armqumulo.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("przdlsmlzsszphnixq"),
		// 			Tags: map[string]*string{
		// 				"key6565": to.Ptr("cgdhmupta"),
		// 			},
		// 			Identity: &armqumulo.ManagedServiceIdentity{
		// 				Type: to.Ptr(armqumulo.ManagedServiceIdentityTypeNone),
		// 				PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				UserAssignedIdentities: map[string]*armqumulo.UserAssignedIdentity{
		// 					"key4522": &armqumulo.UserAssignedIdentity{
		// 						ClientID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 						PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armqumulo.FileSystemResourceProperties{
		// 				AvailabilityZone: to.Ptr("maseyqhlnhoiwbabcqabtedbjpip"),
		// 				ClusterLoginURL: to.Ptr("jjqhgevy"),
		// 				DelegatedSubnetID: to.Ptr("neqctctqdmjezfgt"),
		// 				InitialCapacity: to.Ptr[int32](9),
		// 				MarketplaceDetails: &armqumulo.MarketplaceDetails{
		// 					MarketplaceSubscriptionID: to.Ptr("ujrcqvxfnhxxheoth"),
		// 					MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 					OfferID: to.Ptr("eiyhbmpwgezcmzrrfoiskuxlcvwojf"),
		// 					PlanID: to.Ptr("x"),
		// 					PublisherID: to.Ptr("wfmokfdjbwpjhz"),
		// 				},
		// 				PrivateIPs: []*string{
		// 					to.Ptr("kslguxrwbwkrj")},
		// 					ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
		// 					StorageSKU: to.Ptr(armqumulo.StorageSKUStandard),
		// 					UserDetails: &armqumulo.UserDetails{
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9918780b4dc4bdc111cf3facc11561904d609ad7/specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_ListByResourceGroup_MinimumSet_Gen.json
func ExampleFileSystemsClient_NewListByResourceGroupPager_fileSystemsListByResourceGroupMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSystemsClient().NewListByResourceGroupPager("rgQumulo", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.FileSystemResourceListResult = armqumulo.FileSystemResourceListResult{
		// 	Value: []*armqumulo.FileSystemResource{
		// 		{
		// 			Name: to.Ptr("aaaaa"),
		// 			ID: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 			Location: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			Properties: &armqumulo.FileSystemResourceProperties{
		// 				DelegatedSubnetID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				InitialCapacity: to.Ptr[int32](9),
		// 				MarketplaceDetails: &armqumulo.MarketplaceDetails{
		// 					MarketplaceSubscriptionID: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 					MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 					OfferID: to.Ptr("aaaaaaaaa"),
		// 					PlanID: to.Ptr("aaaaaaa"),
		// 					PublisherID: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 				},
		// 				ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
		// 				StorageSKU: to.Ptr(armqumulo.StorageSKUStandard),
		// 				UserDetails: &armqumulo.UserDetails{
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9918780b4dc4bdc111cf3facc11561904d609ad7/specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_Get_MaximumSet_Gen.json
func ExampleFileSystemsClient_Get_fileSystemsGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSystemsClient().Get(ctx, "rgQumulo", "nauwwbfoqehgbhdsmkewoboyxeqg", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileSystemResource = armqumulo.FileSystemResource{
	// 	Name: to.Ptr("bii"),
	// 	Type: to.Ptr("qtvxrqwpoistduq"),
	// 	ID: to.Ptr("tvelgpobdtazrweunifqzaxkgjauyx"),
	// 	SystemData: &armqumulo.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T04:40:17.991Z"); return t}()),
	// 		CreatedBy: to.Ptr("mtdhqooysjhueaojwpmvophkgntl"),
	// 		CreatedByType: to.Ptr(armqumulo.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T04:40:17.991Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("jcywglomzuamsxltnoegdrkzlscxl"),
	// 		LastModifiedByType: to.Ptr(armqumulo.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("przdlsmlzsszphnixq"),
	// 	Tags: map[string]*string{
	// 		"key6565": to.Ptr("cgdhmupta"),
	// 	},
	// 	Identity: &armqumulo.ManagedServiceIdentity{
	// 		Type: to.Ptr(armqumulo.ManagedServiceIdentityTypeNone),
	// 		PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		UserAssignedIdentities: map[string]*armqumulo.UserAssignedIdentity{
	// 			"key4522": &armqumulo.UserAssignedIdentity{
	// 				ClientID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 				PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armqumulo.FileSystemResourceProperties{
	// 		AvailabilityZone: to.Ptr("maseyqhlnhoiwbabcqabtedbjpip"),
	// 		ClusterLoginURL: to.Ptr("jjqhgevy"),
	// 		DelegatedSubnetID: to.Ptr("neqctctqdmjezfgt"),
	// 		InitialCapacity: to.Ptr[int32](9),
	// 		MarketplaceDetails: &armqumulo.MarketplaceDetails{
	// 			MarketplaceSubscriptionID: to.Ptr("ujrcqvxfnhxxheoth"),
	// 			MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 			OfferID: to.Ptr("eiyhbmpwgezcmzrrfoiskuxlcvwojf"),
	// 			PlanID: to.Ptr("x"),
	// 			PublisherID: to.Ptr("wfmokfdjbwpjhz"),
	// 		},
	// 		PrivateIPs: []*string{
	// 			to.Ptr("kslguxrwbwkrj")},
	// 			ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
	// 			StorageSKU: to.Ptr(armqumulo.StorageSKUStandard),
	// 			UserDetails: &armqumulo.UserDetails{
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9918780b4dc4bdc111cf3facc11561904d609ad7/specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_Get_MinimumSet_Gen.json
func ExampleFileSystemsClient_Get_fileSystemsGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSystemsClient().Get(ctx, "rgQumulo", "aaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileSystemResource = armqumulo.FileSystemResource{
	// 	Name: to.Ptr("aaaaa"),
	// 	ID: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 	Location: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	Properties: &armqumulo.FileSystemResourceProperties{
	// 		DelegatedSubnetID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		InitialCapacity: to.Ptr[int32](9),
	// 		MarketplaceDetails: &armqumulo.MarketplaceDetails{
	// 			MarketplaceSubscriptionID: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 			MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 			OfferID: to.Ptr("aaaaaaaaa"),
	// 			PlanID: to.Ptr("aaaaaaa"),
	// 			PublisherID: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 		ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
	// 		StorageSKU: to.Ptr(armqumulo.StorageSKUStandard),
	// 		UserDetails: &armqumulo.UserDetails{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9918780b4dc4bdc111cf3facc11561904d609ad7/specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_CreateOrUpdate_MaximumSet_Gen.json
func ExampleFileSystemsClient_BeginCreateOrUpdate_fileSystemsCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileSystemsClient().BeginCreateOrUpdate(ctx, "rgQumulo", "nauwwbfoqehgbhdsmkewoboyxeqg", armqumulo.FileSystemResource{
		Location: to.Ptr("przdlsmlzsszphnixq"),
		Tags: map[string]*string{
			"key6565": to.Ptr("cgdhmupta"),
		},
		Identity: &armqumulo.ManagedServiceIdentity{
			Type: to.Ptr(armqumulo.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armqumulo.UserAssignedIdentity{
				"key4522": {},
			},
		},
		Properties: &armqumulo.FileSystemResourceProperties{
			AdminPassword:     to.Ptr("ekceujoecaashtjlsgcymnrdozk"),
			AvailabilityZone:  to.Ptr("maseyqhlnhoiwbabcqabtedbjpip"),
			ClusterLoginURL:   to.Ptr("jjqhgevy"),
			DelegatedSubnetID: to.Ptr("neqctctqdmjezfgt"),
			InitialCapacity:   to.Ptr[int32](9),
			MarketplaceDetails: &armqumulo.MarketplaceDetails{
				MarketplaceSubscriptionID:     to.Ptr("ujrcqvxfnhxxheoth"),
				MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
				OfferID:                       to.Ptr("eiyhbmpwgezcmzrrfoiskuxlcvwojf"),
				PlanID:                        to.Ptr("x"),
				PublisherID:                   to.Ptr("wfmokfdjbwpjhz"),
			},
			PrivateIPs: []*string{
				to.Ptr("kslguxrwbwkrj")},
			ProvisioningState: to.Ptr(armqumulo.ProvisioningStateAccepted),
			StorageSKU:        to.Ptr(armqumulo.StorageSKUStandard),
			UserDetails: &armqumulo.UserDetails{
				Email: to.Ptr("viptslwulnpaupfljvnjeq"),
			},
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
	// res.FileSystemResource = armqumulo.FileSystemResource{
	// 	Name: to.Ptr("bii"),
	// 	Type: to.Ptr("qtvxrqwpoistduq"),
	// 	ID: to.Ptr("tvelgpobdtazrweunifqzaxkgjauyx"),
	// 	SystemData: &armqumulo.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T04:40:17.991Z"); return t}()),
	// 		CreatedBy: to.Ptr("mtdhqooysjhueaojwpmvophkgntl"),
	// 		CreatedByType: to.Ptr(armqumulo.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T04:40:17.991Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("jcywglomzuamsxltnoegdrkzlscxl"),
	// 		LastModifiedByType: to.Ptr(armqumulo.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("przdlsmlzsszphnixq"),
	// 	Tags: map[string]*string{
	// 		"key6565": to.Ptr("cgdhmupta"),
	// 	},
	// 	Identity: &armqumulo.ManagedServiceIdentity{
	// 		Type: to.Ptr(armqumulo.ManagedServiceIdentityTypeNone),
	// 		PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		UserAssignedIdentities: map[string]*armqumulo.UserAssignedIdentity{
	// 			"key4522": &armqumulo.UserAssignedIdentity{
	// 				ClientID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 				PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armqumulo.FileSystemResourceProperties{
	// 		AvailabilityZone: to.Ptr("maseyqhlnhoiwbabcqabtedbjpip"),
	// 		ClusterLoginURL: to.Ptr("jjqhgevy"),
	// 		DelegatedSubnetID: to.Ptr("neqctctqdmjezfgt"),
	// 		InitialCapacity: to.Ptr[int32](9),
	// 		MarketplaceDetails: &armqumulo.MarketplaceDetails{
	// 			MarketplaceSubscriptionID: to.Ptr("ujrcqvxfnhxxheoth"),
	// 			MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 			OfferID: to.Ptr("eiyhbmpwgezcmzrrfoiskuxlcvwojf"),
	// 			PlanID: to.Ptr("x"),
	// 			PublisherID: to.Ptr("wfmokfdjbwpjhz"),
	// 		},
	// 		PrivateIPs: []*string{
	// 			to.Ptr("kslguxrwbwkrj")},
	// 			ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
	// 			StorageSKU: to.Ptr(armqumulo.StorageSKUStandard),
	// 			UserDetails: &armqumulo.UserDetails{
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9918780b4dc4bdc111cf3facc11561904d609ad7/specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_CreateOrUpdate_MinimumSet_Gen.json
func ExampleFileSystemsClient_BeginCreateOrUpdate_fileSystemsCreateOrUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileSystemsClient().BeginCreateOrUpdate(ctx, "rgopenapi", "aaaaaaaa", armqumulo.FileSystemResource{
		Location: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		Properties: &armqumulo.FileSystemResourceProperties{
			AdminPassword:     to.Ptr("ekceujoecaashtjlsgcymnrdozk"),
			DelegatedSubnetID: to.Ptr("aaaaaaaaaa"),
			InitialCapacity:   to.Ptr[int32](9),
			MarketplaceDetails: &armqumulo.MarketplaceDetails{
				MarketplaceSubscriptionID:     to.Ptr("aaaaaaaaaaaaa"),
				MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
				OfferID:                       to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
				PlanID:                        to.Ptr("aaaaaa"),
				PublisherID:                   to.Ptr("aa"),
			},
			ProvisioningState: to.Ptr(armqumulo.ProvisioningStateAccepted),
			StorageSKU:        to.Ptr(armqumulo.StorageSKUStandard),
			UserDetails: &armqumulo.UserDetails{
				Email: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
			},
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
	// res.FileSystemResource = armqumulo.FileSystemResource{
	// 	ID: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 	Location: to.Ptr("aaaaaaaaa"),
	// 	Properties: &armqumulo.FileSystemResourceProperties{
	// 		DelegatedSubnetID: to.Ptr("aaaaaaaaaa"),
	// 		InitialCapacity: to.Ptr[int32](9),
	// 		MarketplaceDetails: &armqumulo.MarketplaceDetails{
	// 			MarketplaceSubscriptionID: to.Ptr("aaaaaaaaaaaaa"),
	// 			MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 			OfferID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			PlanID: to.Ptr("aaaaaa"),
	// 			PublisherID: to.Ptr("aa"),
	// 		},
	// 		ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
	// 		StorageSKU: to.Ptr(armqumulo.StorageSKUStandard),
	// 		UserDetails: &armqumulo.UserDetails{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9918780b4dc4bdc111cf3facc11561904d609ad7/specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_Update_MaximumSet_Gen.json
func ExampleFileSystemsClient_Update_fileSystemsUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSystemsClient().Update(ctx, "rgQumulo", "nauwwbfoqehgbhdsmkewoboyxeqg", armqumulo.FileSystemResourceUpdate{
		Identity: &armqumulo.ManagedServiceIdentity{
			Type: to.Ptr(armqumulo.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armqumulo.UserAssignedIdentity{
				"key4522": {},
			},
		},
		Properties: &armqumulo.FileSystemResourceUpdateProperties{
			ClusterLoginURL:   to.Ptr("adabmuthwrbjshzfbo"),
			DelegatedSubnetID: to.Ptr("vjfirtaljehawmflyfianw"),
			MarketplaceDetails: &armqumulo.MarketplaceDetails{
				MarketplaceSubscriptionID:     to.Ptr("ujrcqvxfnhxxheoth"),
				MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
				OfferID:                       to.Ptr("eiyhbmpwgezcmzrrfoiskuxlcvwojf"),
				PlanID:                        to.Ptr("x"),
				PublisherID:                   to.Ptr("wfmokfdjbwpjhz"),
			},
			PrivateIPs: []*string{
				to.Ptr("eugjqbaoucgjsopzfrq")},
			UserDetails: &armqumulo.UserDetails{
				Email: to.Ptr("viptslwulnpaupfljvnjeq"),
			},
		},
		Tags: map[string]*string{
			"key7534": to.Ptr("jsgqvqbagquvxowbrkanyhzvo"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileSystemResource = armqumulo.FileSystemResource{
	// 	Name: to.Ptr("bii"),
	// 	Type: to.Ptr("qtvxrqwpoistduq"),
	// 	ID: to.Ptr("tvelgpobdtazrweunifqzaxkgjauyx"),
	// 	SystemData: &armqumulo.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T04:40:17.991Z"); return t}()),
	// 		CreatedBy: to.Ptr("mtdhqooysjhueaojwpmvophkgntl"),
	// 		CreatedByType: to.Ptr(armqumulo.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T04:40:17.991Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("jcywglomzuamsxltnoegdrkzlscxl"),
	// 		LastModifiedByType: to.Ptr(armqumulo.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("przdlsmlzsszphnixq"),
	// 	Tags: map[string]*string{
	// 		"key6565": to.Ptr("cgdhmupta"),
	// 	},
	// 	Identity: &armqumulo.ManagedServiceIdentity{
	// 		Type: to.Ptr(armqumulo.ManagedServiceIdentityTypeNone),
	// 		PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		UserAssignedIdentities: map[string]*armqumulo.UserAssignedIdentity{
	// 			"key4522": &armqumulo.UserAssignedIdentity{
	// 				ClientID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 				PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armqumulo.FileSystemResourceProperties{
	// 		AvailabilityZone: to.Ptr("maseyqhlnhoiwbabcqabtedbjpip"),
	// 		ClusterLoginURL: to.Ptr("jjqhgevy"),
	// 		DelegatedSubnetID: to.Ptr("neqctctqdmjezfgt"),
	// 		InitialCapacity: to.Ptr[int32](9),
	// 		MarketplaceDetails: &armqumulo.MarketplaceDetails{
	// 			MarketplaceSubscriptionID: to.Ptr("ujrcqvxfnhxxheoth"),
	// 			MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 			OfferID: to.Ptr("eiyhbmpwgezcmzrrfoiskuxlcvwojf"),
	// 			PlanID: to.Ptr("x"),
	// 			PublisherID: to.Ptr("wfmokfdjbwpjhz"),
	// 		},
	// 		PrivateIPs: []*string{
	// 			to.Ptr("eugjqbaoucgjsopzfrq")},
	// 			ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
	// 			StorageSKU: to.Ptr(armqumulo.StorageSKUStandard),
	// 			UserDetails: &armqumulo.UserDetails{
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9918780b4dc4bdc111cf3facc11561904d609ad7/specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_Update_MinimumSet_Gen.json
func ExampleFileSystemsClient_Update_fileSystemsUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSystemsClient().Update(ctx, "rgQumulo", "aaaaaaaaaaaaaaaaa", armqumulo.FileSystemResourceUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileSystemResource = armqumulo.FileSystemResource{
	// 	Name: to.Ptr("aaaaa"),
	// 	Location: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	Properties: &armqumulo.FileSystemResourceProperties{
	// 		DelegatedSubnetID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		InitialCapacity: to.Ptr[int32](9),
	// 		MarketplaceDetails: &armqumulo.MarketplaceDetails{
	// 			MarketplaceSubscriptionID: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 			MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 			OfferID: to.Ptr("aaaaaaaaa"),
	// 			PlanID: to.Ptr("aaaaaaa"),
	// 			PublisherID: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 		ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
	// 		StorageSKU: to.Ptr(armqumulo.StorageSKUStandard),
	// 		UserDetails: &armqumulo.UserDetails{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9918780b4dc4bdc111cf3facc11561904d609ad7/specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_Delete_MaximumSet_Gen.json
func ExampleFileSystemsClient_BeginDelete_fileSystemsDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileSystemsClient().BeginDelete(ctx, "rgQumulo", "nauwwbfoqehgbhdsmkewoboyxeqg", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9918780b4dc4bdc111cf3facc11561904d609ad7/specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_Delete_MinimumSet_Gen.json
func ExampleFileSystemsClient_BeginDelete_fileSystemsDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileSystemsClient().BeginDelete(ctx, "rgQumulo", "nauwwbfoqehgbhdsmkewoboyxeqg", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
