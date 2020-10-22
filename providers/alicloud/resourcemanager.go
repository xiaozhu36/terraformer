package alicloud

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/providers/alicloud/connectivity"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/resourcemanager"
)

type ResourceManagerGenerator struct {
	AliCloudService
}

func (g *ResourceManagerGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}
	allAccounts, err := initResourceManagerAccounts(client)
	if err != nil {
		fmt.Println("[Error]init ResourceManagerAccount fail")
		return err
	}
	for _, account := range allAccounts {
		resource := resourceFromResourceManagerAccount(account)
		g.Resources = append(g.Resources, resource)
	}
	allFolders, err := initResourceManagerFolders(client)
	if err != nil {
		fmt.Println("[Error]init ResourceManagerFolder fail")
		return err
	}
	for _, folder := range allFolders {
		resource := resourceFromResourceManagerFolder(folder)
		g.Resources = append(g.Resources, resource)
	}
	allHandshakes, err := initResourceManagerHandshakes(client)
	if err != nil {
		fmt.Println("[Error]init ResourceManagerHandshake fail")
		return err
	}
	for _, handshake := range allHandshakes {
		resource := resourceFromResourceManagerHandshake(handshake)
		g.Resources = append(g.Resources, resource)
	}
	allPolicys, err := initResourceManagerPolicys(client)
	if err != nil {
		fmt.Println("[Error]init ResourceManagerPolicy fail")
		return err
	}
	for _, policy := range allPolicys {
		resource := resourceFromResourceManagerPolicy(policy)
		g.Resources = append(g.Resources, resource)
	}
	allPolicyAttachments, err := initResourceManagerPolicyAttachments(client)
	if err != nil {
		fmt.Println("[Error]init ResourceManagerPolicyAttachment fail")
		return err
	}
	for _, policyAttachment := range allPolicyAttachments {
		resource := resourceFromResourceManagerPolicyAttachment(policyAttachment)
		g.Resources = append(g.Resources, resource)
	}
	allPolicyVersions, policyNames, err := initResourceManagerPolicyVersions(client, allPolicys)
	if err != nil {
		fmt.Println("[Error]init ResourceManagerPolicyVersion fail")
		return err
	}
	for i, policyVersion := range allPolicyVersions {
		resource := resourceFromResourceManagerPolicyVersion(policyVersion, policyNames[i])
		g.Resources = append(g.Resources, resource)
	}
	allResourceDirectorys, err := initResourceManagerResourceDirectorys(client)
	if err != nil {
		fmt.Println("[Error]init ResourceManagerResourceDirectory fail")
		return err
	}
	for _, resourceDirectory := range allResourceDirectorys {
		resource := resourceFromResourceManagerResourceDirectory(resourceDirectory)
		g.Resources = append(g.Resources, resource)
	}
	allResourceGroups, err := initResourceManagerResourceGroups(client)
	if err != nil {
		fmt.Println("[Error]init ResourceManagerResourceGroup fail")
		return err
	}
	for _, resourceGroup := range allResourceGroups {
		resource := resourceFromResourceManagerResourceGroup(resourceGroup)
		g.Resources = append(g.Resources, resource)
	}
	allRoles, err := initResourceManagerRoles(client)
	if err != nil {
		fmt.Println("[Error]init ResourceManagerRole fail")
		return err
	}
	for _, role := range allRoles {
		resource := resourceFromResourceManagerRole(role)
		g.Resources = append(g.Resources, resource)
	}
	return nil
}
func initResourceManagerAccounts(client *connectivity.AliyunClient) ([]resourcemanager.Account, error) {
	allAccounts := make([]resourcemanager.Account, 0)
	remaining := 1
	pageNumber := 1
	pageSize := 10
	for remaining > 0 {
		raw, err := client.WithResourcemanagerClient(func(resourcemanagerClient *resourcemanager.Client) (interface{}, error) {
			request := resourcemanager.CreateListAccountsRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return resourcemanagerClient.ListAccounts(request)
		})
		if err != nil {
			return nil, err
		}
		response := raw.(*resourcemanager.ListAccountsResponse)
		allAccounts = append(allAccounts, response.Accounts.Account...)
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}
	return allAccounts, nil
}
func initResourceManagerFolders(client *connectivity.AliyunClient) ([]resourcemanager.Folder, error) {
	allFolders := make([]resourcemanager.Folder, 0)
	remaining := 1
	pageNumber := 1
	pageSize := 10
	for remaining > 0 {
		raw, err := client.WithResourcemanagerClient(func(resourcemanagerClient *resourcemanager.Client) (interface{}, error) {
			request := resourcemanager.CreateListFoldersForParentRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return resourcemanagerClient.ListFoldersForParent(request)
		})
		if err != nil {
			return nil, err
		}
		response := raw.(*resourcemanager.ListFoldersForParentResponse)
		allFolders = append(allFolders, response.Folders.Folder...)
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}
	return allFolders, nil
}
func initResourceManagerHandshakes(client *connectivity.AliyunClient) ([]resourcemanager.Handshake, error) {
	allHandshakes := make([]resourcemanager.Handshake, 0)
	remaining := 1
	pageNumber := 1
	pageSize := 10
	for remaining > 0 {
		raw, err := client.WithResourcemanagerClient(func(resourcemanagerClient *resourcemanager.Client) (interface{}, error) {
			request := resourcemanager.CreateListHandshakesForResourceDirectoryRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return resourcemanagerClient.ListHandshakesForResourceDirectory(request)
		})
		if err != nil {
			return nil, err
		}
		response := raw.(*resourcemanager.ListHandshakesForResourceDirectoryResponse)
		allHandshakes = append(allHandshakes, response.Handshakes.Handshake...)
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}
	return allHandshakes, nil
}
func initResourceManagerPolicys(client *connectivity.AliyunClient) ([]resourcemanager.Policy, error) {
	allPolicys := make([]resourcemanager.Policy, 0)
	remaining := 1
	pageNumber := 1
	pageSize := 10
	for remaining > 0 {
		raw, err := client.WithResourcemanagerClient(func(resourcemanagerClient *resourcemanager.Client) (interface{}, error) {
			request := resourcemanager.CreateListPoliciesRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return resourcemanagerClient.ListPolicies(request)
		})
		if err != nil {
			return nil, err
		}
		response := raw.(*resourcemanager.ListPoliciesResponse)
		allPolicys = append(allPolicys, response.Policies.Policy...)
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}
	return allPolicys, nil
}
func initResourceManagerPolicyAttachments(client *connectivity.AliyunClient) ([]resourcemanager.PolicyAttachment, error) {
	allPolicyAttachments := make([]resourcemanager.PolicyAttachment, 0)
	remaining := 1
	pageNumber := 1
	pageSize := 10
	for remaining > 0 {
		raw, err := client.WithResourcemanagerClient(func(resourcemanagerClient *resourcemanager.Client) (interface{}, error) {
			request := resourcemanager.CreateListPolicyAttachmentsRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return resourcemanagerClient.ListPolicyAttachments(request)
		})
		if err != nil {
			return nil, err
		}
		response := raw.(*resourcemanager.ListPolicyAttachmentsResponse)
		allPolicyAttachments = append(allPolicyAttachments, response.PolicyAttachments.PolicyAttachment...)
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}
	return allPolicyAttachments, nil
}
func initResourceManagerPolicyVersions(client *connectivity.AliyunClient, allPolicys []resourcemanager.Policy) ([]resourcemanager.PolicyVersion, []string, error) {
	allPolicyVersions := make([]resourcemanager.PolicyVersion, 0)
	policyNames := make([]string, 0)
	for _, policy := range allPolicys {
		raw, err := client.WithResourcemanagerClient(func(resourcemanagerClient *resourcemanager.Client) (interface{}, error) {
			request := resourcemanager.CreateListPolicyVersionsRequest()
			request.RegionId = client.RegionId
			request.PolicyType = policy.PolicyType
			request.PolicyName = policy.PolicyName
			return resourcemanagerClient.ListPolicyVersions(request)
		})
		if err != nil {
			return nil, policyNames, err
		}
		response := raw.(*resourcemanager.ListPolicyVersionsResponse)
		for _, v := range response.PolicyVersions.PolicyVersion {
			allPolicyVersions = append(allPolicyVersions, v)
			policyNames = append(policyNames, policy.PolicyName)
		}
	}
	return allPolicyVersions, policyNames, nil
}
func initResourceManagerResourceDirectorys(client *connectivity.AliyunClient) ([]resourcemanager.ResourceDirectory, error) {
	allResourceDirectorys := make([]resourcemanager.ResourceDirectory, 0)
	raw, err := client.WithResourcemanagerClient(func(resourcemanagerClient *resourcemanager.Client) (interface{}, error) {
		request := resourcemanager.CreateGetResourceDirectoryRequest()
		request.RegionId = client.RegionId
		return resourcemanagerClient.GetResourceDirectory(request)
	})
	if err != nil {
		return nil, err
	}
	response := raw.(*resourcemanager.GetResourceDirectoryResponse)
	allResourceDirectorys = append(allResourceDirectorys, response.ResourceDirectory)
	return allResourceDirectorys, nil
}
func initResourceManagerResourceGroups(client *connectivity.AliyunClient) ([]resourcemanager.ResourceGroup, error) {
	allResourceGroups := make([]resourcemanager.ResourceGroup, 0)
	remaining := 1
	pageNumber := 1
	pageSize := 10
	for remaining > 0 {
		raw, err := client.WithResourcemanagerClient(func(resourcemanagerClient *resourcemanager.Client) (interface{}, error) {
			request := resourcemanager.CreateListResourceGroupsRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return resourcemanagerClient.ListResourceGroups(request)
		})
		if err != nil {
			return nil, err
		}
		response := raw.(*resourcemanager.ListResourceGroupsResponse)
		allResourceGroups = append(allResourceGroups, response.ResourceGroups.ResourceGroup...)
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}
	return allResourceGroups, nil
}
func initResourceManagerRoles(client *connectivity.AliyunClient) ([]resourcemanager.Role, error) {
	allRoles := make([]resourcemanager.Role, 0)
	remaining := 1
	pageNumber := 1
	pageSize := 10
	for remaining > 0 {
		raw, err := client.WithResourcemanagerClient(func(resourcemanagerClient *resourcemanager.Client) (interface{}, error) {
			request := resourcemanager.CreateListRolesRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return resourcemanagerClient.ListRoles(request)
		})
		if err != nil {
			return nil, err
		}
		response := raw.(*resourcemanager.ListRolesResponse)
		allRoles = append(allRoles, response.Roles.Role...)
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}
	return allRoles, nil
}
func resourceFromResourceManagerAccount(account resourcemanager.Account) terraformutils.Resource {
	return terraformutils.NewResource(
		account.AccountId,
		account.AccountId,
		"alicloud_resource_manager_account",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}
func resourceFromResourceManagerFolder(folder resourcemanager.Folder) terraformutils.Resource {
	return terraformutils.NewResource(
		folder.FolderId,
		folder.FolderId,
		"alicloud_resource_manager_folder",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}
func resourceFromResourceManagerHandshake(handshake resourcemanager.Handshake) terraformutils.Resource {
	return terraformutils.NewResource(
		handshake.HandshakeId,
		handshake.HandshakeId,
		"alicloud_resource_manager_handshake",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}
func resourceFromResourceManagerPolicy(policy resourcemanager.Policy) terraformutils.Resource {
	return terraformutils.NewResource(
		policy.PolicyName,
		policy.PolicyName,
		"alicloud_resource_manager_policy",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}
func resourceFromResourceManagerPolicyAttachment(policyAttachment resourcemanager.PolicyAttachment) terraformutils.Resource {
	id := strings.Join([]string{policyAttachment.PolicyName, policyAttachment.PolicyType, policyAttachment.PrincipalName, policyAttachment.PrincipalType, policyAttachment.ResourceGroupId}, ":")
	return terraformutils.NewResource(
		id,
		id,
		"alicloud_resource_manager_policy_attachment",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}
func resourceFromResourceManagerPolicyVersion(policyVersion resourcemanager.PolicyVersion, policyName string) terraformutils.Resource {
	id := strings.Join([]string{policyName, policyVersion.VersionId}, ":")
	return terraformutils.NewResource(
		id,
		id,
		"alicloud_resource_manager_policy_version",
		"alicloud",
		map[string]string{
			"policy_name": "",
		},
		[]string{
			"policy_name",
		},
		map[string]interface{}{},
	)
}
func resourceFromResourceManagerResourceDirectory(resourceDirectory resourcemanager.ResourceDirectory) terraformutils.Resource {
	return terraformutils.NewResource(
		resourceDirectory.ResourceDirectoryId,
		resourceDirectory.ResourceDirectoryId,
		"alicloud_resource_manager_resource_directory",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}
func resourceFromResourceManagerResourceGroup(resourceGroup resourcemanager.ResourceGroup) terraformutils.Resource {
	return terraformutils.NewResource(
		resourceGroup.Id,
		resourceGroup.Id,
		"alicloud_resource_manager_resource_group",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}
func resourceFromResourceManagerRole(role resourcemanager.Role) terraformutils.Resource {
	return terraformutils.NewResource(
		role.RoleName,
		role.RoleName,
		"alicloud_resource_manager_role",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}
func (g *ResourceManagerGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "alicloud_resource_manager_policy" {
			delete(r.Item, "default_version")
		}
	}
	return nil
}
