package alicloud

import (
	"github.com/GoogleCloudPlatform/terraformer/providers/alicloud/connectivity"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/eci"
)

type EciGenerator struct {
	AliCloudService
}

func (g *EciGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}
	allImageCaches, err := initEciImageCaches(client)
	if err != nil {
		return err
	}
	for _, imageCache := range allImageCaches {
		resource := resourceFromEciImageCache(imageCache)
		g.Resources = append(g.Resources, resource)
	}
	return nil
}
func initEciImageCaches(client *connectivity.AliyunClient) ([]eci.DescribeImageCachesImageCache0, error) {
	allImageCaches := make([]eci.DescribeImageCachesImageCache0, 0)
	raw, err := client.WithEciClient(func(eciClient *eci.Client) (interface{}, error) {
		request := eci.CreateDescribeImageCachesRequest()
		request.RegionId = client.RegionId
		return eciClient.DescribeImageCaches(request)
	})
	if err != nil {
		return nil, err
	}
	response := raw.(*eci.DescribeImageCachesResponse)
	allImageCaches = append(allImageCaches, response.ImageCaches...)
	return allImageCaches, nil
}
func resourceFromEciImageCache(imageCache eci.DescribeImageCachesImageCache0) terraformutils.Resource {
	return terraformutils.NewResource(
		imageCache.ImageCacheId,
		imageCache.ImageCacheId,
		"alicloud_eci_image_cache",
		"alicloud",
		map[string]string{
			"security_group_id": "",
			"vswitch_id":        "",
		},
		[]string{
			"security_group_id",
			"vswitch_id",
		},
		map[string]interface{}{},
	)
}
