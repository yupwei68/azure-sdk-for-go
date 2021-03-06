package dnsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns"
	"github.com/Azure/go-autorest/autorest"
)

// RecordSetsClientAPI contains the set of methods on the RecordSetsClient type.
type RecordSetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, recordType dns.RecordType, relativeRecordSetName string, parameters dns.RecordSet, ifMatch string, ifNoneMatch string) (result dns.RecordSet, err error)
	Delete(ctx context.Context, resourceGroupName string, zoneName string, recordType dns.RecordType, relativeRecordSetName string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, zoneName string, recordType dns.RecordType, relativeRecordSetName string) (result dns.RecordSet, err error)
	List(ctx context.Context, resourceGroupName string, zoneName string, recordType dns.RecordType, top string, filter string) (result dns.RecordSetListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, zoneName string, recordType dns.RecordType, top string, filter string) (result dns.RecordSetListResultIterator, err error)
	ListAll(ctx context.Context, resourceGroupName string, zoneName string, top string, filter string) (result dns.RecordSetListResultPage, err error)
	ListAllComplete(ctx context.Context, resourceGroupName string, zoneName string, top string, filter string) (result dns.RecordSetListResultIterator, err error)
}

var _ RecordSetsClientAPI = (*dns.RecordSetsClient)(nil)

// ZonesClientAPI contains the set of methods on the ZonesClient type.
type ZonesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, parameters dns.Zone, ifMatch string, ifNoneMatch string) (result dns.Zone, err error)
	Delete(ctx context.Context, resourceGroupName string, zoneName string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, zoneName string) (result dns.Zone, err error)
	ListZonesInResourceGroup(ctx context.Context, resourceGroupName string, top string, filter string) (result dns.ZoneListResultPage, err error)
	ListZonesInResourceGroupComplete(ctx context.Context, resourceGroupName string, top string, filter string) (result dns.ZoneListResultIterator, err error)
	ListZonesInSubscription(ctx context.Context, top string, filter string) (result dns.ZoneListResultPage, err error)
	ListZonesInSubscriptionComplete(ctx context.Context, top string, filter string) (result dns.ZoneListResultIterator, err error)
}

var _ ZonesClientAPI = (*dns.ZonesClient)(nil)
