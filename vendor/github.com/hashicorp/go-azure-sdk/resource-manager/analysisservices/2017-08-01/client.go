package v2017_08_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/analysisservices/2017-08-01/analysisservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/analysisservices/2017-08-01/servers"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	AnalysisServices *analysisservices.AnalysisServicesClient
	Servers          *servers.ServersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	analysisServicesClient := analysisservices.NewAnalysisServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&analysisServicesClient.Client)

	serversClient := servers.NewServersClientWithBaseURI(endpoint)
	configureAuthFunc(&serversClient.Client)

	return Client{
		AnalysisServices: &analysisServicesClient,
		Servers:          &serversClient,
	}
}
