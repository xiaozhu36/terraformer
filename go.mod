module github.com/GoogleCloudPlatform/terraformer

go 1.15

require (
	cloud.google.com/go v0.60.0
	cloud.google.com/go/logging v1.0.0
	cloud.google.com/go/storage v1.8.0
	github.com/Azure/azure-sdk-for-go v42.0.0+incompatible
	github.com/Azure/azure-storage-blob-go v0.10.0
	github.com/Azure/go-autorest/autorest v0.10.0
	github.com/OctopusDeploy/go-octopusdeploy v1.6.0
	github.com/Sirupsen/logrus v0.0.0-20181010200618-458213699411 // indirect
	github.com/alibabacloud-go/tea-rpc v1.1.6
	github.com/alibabacloud-go/tea-utils v1.3.4
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.517
	github.com/aliyun/aliyun-datahub-sdk-go v0.0.0-20180929121038-c1c85baca7c0
	github.com/aliyun/aliyun-log-go-sdk v0.1.13
	github.com/aliyun/aliyun-mns-go-sdk v0.0.0-20191205082232-b251b9d95415
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190103054945-8205d1f41e70
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible
	github.com/aliyun/credentials-go v1.1.2
	github.com/aliyun/fc-go-sdk v0.0.0-20200922070653-c9c4d8e5539a
	github.com/aws/aws-sdk-go v1.30.19
	github.com/aws/aws-sdk-go-v2 v0.22.0
	github.com/cloudflare/cloudflare-go v0.11.7
	github.com/ddelnano/terraform-provider-mikrotik v0.0.0-20200501162830-a217572b326c
	github.com/denverdino/aliyungo v0.0.0-20200327235253-d59c209c7e93
	github.com/digitalocean/godo v1.35.1
	github.com/dollarshaveclub/new-relic-synthetics-go v0.0.0-20170605224734-4dc3dd6ae884
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/fastly/go-fastly v1.15.0
	github.com/gogap/errors v0.0.0-20200228125012-531a6449b28c // indirect
	github.com/gogap/stack v0.0.0-20150131034635-fef68dddd4f8 // indirect
	github.com/google/go-github/v25 v25.1.3
	github.com/gophercloud/gophercloud v0.10.0
	github.com/hashicorp/go-azure-helpers v0.10.0
	github.com/hashicorp/go-hclog v0.12.2
	github.com/hashicorp/go-plugin v1.3.0
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/terraform v0.12.29
	github.com/heroku/heroku-go/v5 v5.1.0
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/jmespath/go-jmespath v0.3.0
	github.com/jonboydell/logzio_client v1.2.0
	github.com/labd/commercetools-go-sdk v0.0.0-20200309143931-ca72e918a79d
	github.com/linode/linodego v0.15.0
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-20200506151941-509881368409
	github.com/ns1/ns1-go v2.4.0+incompatible
	github.com/paultyng/go-newrelic/v4 v4.10.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/valyala/fasthttp v1.16.0 // indirect
	github.com/vultr/govultr v0.4.0
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0
	github.com/yandex-cloud/go-genproto v0.0.0-20200722140432-762fe965ce77
	github.com/yandex-cloud/go-sdk v0.0.0-20200722140627-2194e5077f13
	github.com/zclconf/go-cty v1.4.0
	github.com/zorkian/go-datadog-api v2.29.0+incompatible
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/text v0.3.2
	gonum.org/v1/gonum v0.7.0
	google.golang.org/api v0.28.0
	google.golang.org/genproto v0.0.0-20200626011028-ee7919e894b5
	gopkg.in/jarcoal/httpmock.v1 v1.0.0-00010101000000-000000000000 // indirect
	k8s.io/apimachinery v0.17.5
	k8s.io/client-go v0.17.5
	k8s.io/utils v0.0.0-20191218082557-f07c713de883 // indirect
)

replace gopkg.in/jarcoal/httpmock.v1 => github.com/jarcoal/httpmock v1.0.5

replace github.com/Sirupsen/logrus v0.0.0-20181010200618-458213699411 => github.com/sirupsen/logrus v0.0.0-20181010200618-458213699411
