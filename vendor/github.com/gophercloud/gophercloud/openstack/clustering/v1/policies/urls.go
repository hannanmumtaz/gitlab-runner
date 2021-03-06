package policies

import "github.com/gophercloud/gophercloud"

const (
	apiVersion = "v1"
	apiName    = "policies"
)

func policyListURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func policyCreateURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func policyDeleteURL(client *gophercloud.ServiceClient, policyID string) string {
	return client.ServiceURL(apiVersion, apiName, policyID)
}
