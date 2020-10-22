package alicloud

import (
	"github.com/GoogleCloudPlatform/terraformer/providers/alicloud/connectivity"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type AlidnsGenerator struct {
	AliCloudService
}

func (g *AlidnsGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}
	allDomains, err := initAlidnsDomains(client)
	if err != nil {
		return err
	}
	for _, domain := range allDomains {
		resource := resourceFromAlidnsDomain(domain)
		g.Resources = append(g.Resources, resource)
	}
	allDomainGroups, err := initAlidnsDomainGroups(client)
	if err != nil {
		return err
	}
	for _, domainGroup := range allDomainGroups {
		resource := resourceFromAlidnsDomainGroup(domainGroup)
		g.Resources = append(g.Resources, resource)
	}
	allRecords, err := initAlidnsRecords(client, allDomains)
	if err != nil {
		return err
	}
	for _, record := range allRecords {
		resource := resourceFromAlidnsRecord(record)
		g.Resources = append(g.Resources, resource)
	}
	allInstances, err := initAlidnsInstances(client)
	if err != nil {
		return err
	}
	for _, instance := range allInstances {
		resource := resourceFromAlidnsInstance(instance)
		g.Resources = append(g.Resources, resource)
	}
	return nil
}
func initAlidnsDomains(client *connectivity.AliyunClient) ([]alidns.Domain, error) {
	allDomains := make([]alidns.Domain, 0)
	remaining := 1
	pageNumber := 1
	pageSize := 10
	for remaining > 0 {
		raw, err := client.WithAlidnsClient(func(alidnsClient *alidns.Client) (interface{}, error) {
			request := alidns.CreateDescribeDomainsRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return alidnsClient.DescribeDomains(request)
		})
		if err != nil {
			return nil, err
		}
		response := raw.(*alidns.DescribeDomainsResponse)
		allDomains = append(allDomains, response.Domains.Domain...)
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}
	return allDomains, nil
}
func initAlidnsDomainGroups(client *connectivity.AliyunClient) ([]alidns.DomainGroup, error) {
	allDomainGroups := make([]alidns.DomainGroup, 0)
	remaining := 1
	pageNumber := 1
	pageSize := 10
	for remaining > 0 {
		raw, err := client.WithAlidnsClient(func(alidnsClient *alidns.Client) (interface{}, error) {
			request := alidns.CreateDescribeDomainGroupsRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return alidnsClient.DescribeDomainGroups(request)
		})
		if err != nil {
			return nil, err
		}
		response := raw.(*alidns.DescribeDomainGroupsResponse)
		allDomainGroups = append(allDomainGroups, response.DomainGroups.DomainGroup...)
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}
	return allDomainGroups, nil
}
func initAlidnsRecords(client *connectivity.AliyunClient, allDomains []alidns.Domain) ([]alidns.Record, error) {
	allRecords := make([]alidns.Record, 0)
	for _, domain := range allDomains {
		remaining := 1
		pageNumber := 1
		pageSize := 10
		for remaining > 0 {
			raw, err := client.WithAlidnsClient(func(alidnsClient *alidns.Client) (interface{}, error) {
				request := alidns.CreateDescribeDomainRecordsRequest()
				request.RegionId = client.RegionId
				request.DomainName = domain.DomainName
				request.PageSize = requests.NewInteger(pageSize)
				request.PageNumber = requests.NewInteger(pageNumber)
				return alidnsClient.DescribeDomainRecords(request)
			})
			if err != nil {
				return nil, err
			}
			response := raw.(*alidns.DescribeDomainRecordsResponse)
			allRecords = append(allRecords, response.DomainRecords.Record...)
			remaining = int(response.TotalCount) - pageNumber*pageSize
			pageNumber++
		}
	}
	return allRecords, nil
}
func initAlidnsInstances(client *connectivity.AliyunClient) ([]alidns.DnsProduct, error) {
	allInstances := make([]alidns.DnsProduct, 0)
	remaining := 1
	pageNumber := 1
	pageSize := 10
	for remaining > 0 {
		raw, err := client.WithAlidnsClient(func(alidnsClient *alidns.Client) (interface{}, error) {
			request := alidns.CreateDescribeDnsProductInstancesRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return alidnsClient.DescribeDnsProductInstances(request)
		})
		if err != nil {
			return nil, err
		}
		response := raw.(*alidns.DescribeDnsProductInstancesResponse)
		allInstances = append(allInstances, response.DnsProducts.DnsProduct...)
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}
	return allInstances, nil
}
func resourceFromAlidnsDomain(domain alidns.Domain) terraformutils.Resource {
	return terraformutils.NewResource(
		domain.DomainName,
		domain.DomainName,
		"alicloud_alidns_domain",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}
func resourceFromAlidnsDomainGroup(domainGroup alidns.DomainGroup) terraformutils.Resource {
	return terraformutils.NewResource(
		domainGroup.GroupId,
		domainGroup.GroupId,
		"alicloud_alidns_domain_group",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}
func resourceFromAlidnsRecord(record alidns.Record) terraformutils.Resource {
	return terraformutils.NewResource(
		record.RecordId,
		record.RecordId,
		"alicloud_alidns_record",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}
func resourceFromAlidnsInstance(instance alidns.DnsProduct) terraformutils.Resource {
	return terraformutils.NewResource(
		instance.InstanceId,
		instance.InstanceId,
		"alicloud_alidns_instance",
		"alicloud",
		map[string]string{
			"payment_type": "Subscription",
		},
		[]string{},
		map[string]interface{}{},
	)
}
