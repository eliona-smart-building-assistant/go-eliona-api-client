# \ProjectsAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectById**](ProjectsAPI.md#GetProjectById) | **Get** /projects/{project-id} | Information about a project
[**GetProjectRoleByName**](ProjectsAPI.md#GetProjectRoleByName) | **Get** /project-roles/{role-name} | Information about a project role
[**GetProjectRoles**](ProjectsAPI.md#GetProjectRoles) | **Get** /project-roles | Information about project roles
[**GetProjects**](ProjectsAPI.md#GetProjects) | **Get** /projects | Information about projects
[**PutProject**](ProjectsAPI.md#PutProject) | **Put** /projects | Create or update a project
[**PutProjectRole**](ProjectsAPI.md#PutProjectRole) | **Put** /project-roles | Create or update a project role



## GetProjectById

> Project GetProjectById(ctx, projectId).Expansions(expansions).Execute()

Information about a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
)

func main() {
	projectId := "69133767-7de1-4134-9130-5f5403e1ee98" // string | The id of the project.  The project ID is deprecated due to the removal of projects and is broadly replaced by sites. This filter now uses siteId, which is internally represented as a UUID.
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectById(context.Background(), projectId).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectById`: Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The id of the project.  The project ID is deprecated due to the removal of projects and is broadly replaced by sites. This filter now uses siteId, which is internally represented as a UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**Project**](Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectRoleByName

> ProjectRole GetProjectRoleByName(ctx, roleName).Execute()

Information about a project role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
)

func main() {
	roleName := "System user" // string | The name of the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectRoleByName(context.Background(), roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectRoleByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectRoleByName`: ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectRoleByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | The name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRoleByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectRole**](ProjectRole.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectRoles

> []ProjectRole GetProjectRoles(ctx).Offset(offset).Size(size).Execute()

Information about project roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
)

func main() {
	offset := int64(3) // int64 | Specifies the starting point for pagination by indicating the number of items to skip.  (optional)
	size := int64(10) // int64 | Specifies the number of items per page for pagination.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectRoles(context.Background()).Offset(offset).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectRoles`: []ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Specifies the starting point for pagination by indicating the number of items to skip.  | 
 **size** | **int64** | Specifies the number of items per page for pagination.  | 

### Return type

[**[]ProjectRole**](ProjectRole.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjects

> []Project GetProjects(ctx).Offset(offset).Size(size).Expansions(expansions).Execute()

Information about projects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
)

func main() {
	offset := int64(3) // int64 | Specifies the starting point for pagination by indicating the number of items to skip.  (optional)
	size := int64(10) // int64 | Specifies the number of items per page for pagination.  (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjects(context.Background()).Offset(offset).Size(size).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjects`: []Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Specifies the starting point for pagination by indicating the number of items to skip.  | 
 **size** | **int64** | Specifies the number of items per page for pagination.  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]Project**](Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProject

> Project PutProject(ctx).Project(project).Expansions(expansions).Execute()

Create or update a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
)

func main() {
	project := *openapiclient.NewProject() // Project | 
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PutProject(context.Background()).Project(project).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PutProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProject`: Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PutProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**Project**](Project.md) |  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**Project**](Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectRole

> ProjectRole PutProjectRole(ctx).ProjectRole(projectRole).Execute()

Create or update a project role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
)

func main() {
	projectRole := *openapiclient.NewProjectRole("readonly") // ProjectRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PutProjectRole(context.Background()).ProjectRole(projectRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PutProjectRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectRole`: ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PutProjectRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectRole** | [**ProjectRole**](ProjectRole.md) |  | 

### Return type

[**ProjectRole**](ProjectRole.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

