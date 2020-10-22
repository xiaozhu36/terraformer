package alicloud

import (
	"github.com/GoogleCloudPlatform/terraformer/providers/alicloud/connectivity"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dcdn"
)

type DCDNGenerator struct {
	AliCloudService
}

func (g *DCDNGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}
	allDomains, err := initDCDNDomains(client)
	if err != nil {
		return err
	}
	for _, domain := range allDomains {
		resource := resourceFromDCDNDomain(domain)
		g.Resources = append(g.Resources, resource)
	}
	return nil
}
func initDCDNDomains(client *connectivity.AliyunClient) ([]dcdn.PageData, error) {
	allDomains := make([]dcdn.PageData, 0)
	remaining := 1
	pageNumber := 1
	pageSize := 10
	for remaining > 0 {
		raw, err := client.WithDcdnClient(func(dcdnClient *dcdn.Client) (interface{}, error) {
			request := dcdn.CreateDescribeDcdnUserDomainsRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return dcdnClient.DescribeDcdnUserDomains(request)
		})
		if err != nil {
			return nil, err
		}
		response := raw.(*dcdn.DescribeDcdnUserDomainsResponse)
		allDomains = append(allDomains, response.Domains.PageData...)
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}
	return allDomains, nil
}
func resourceFromDCDNDomain(domain dcdn.PageData) terraformutils.Resource {
	return terraformutils.NewResource(
		domain.DomainName,
		domain.DomainName,
		"alicloud_dcdn_domain",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}
