# \SlurmAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmV0039CancelJob**](SlurmAPI.md#SlurmV0039CancelJob) | **Delete** /slurm/v0.0.39/job/{job_id} | cancel or signal job
[**SlurmV0039DeleteNode**](SlurmAPI.md#SlurmV0039DeleteNode) | **Delete** /slurm/v0.0.39/node/{node_name} | delete node
[**SlurmV0039Diag**](SlurmAPI.md#SlurmV0039Diag) | **Get** /slurm/v0.0.39/diag | get diagnostics
[**SlurmV0039GetJob**](SlurmAPI.md#SlurmV0039GetJob) | **Get** /slurm/v0.0.39/job/{job_id} | get job info
[**SlurmV0039GetJobs**](SlurmAPI.md#SlurmV0039GetJobs) | **Get** /slurm/v0.0.39/jobs | get list of jobs
[**SlurmV0039GetNode**](SlurmAPI.md#SlurmV0039GetNode) | **Get** /slurm/v0.0.39/node/{node_name} | get node info
[**SlurmV0039GetNodes**](SlurmAPI.md#SlurmV0039GetNodes) | **Get** /slurm/v0.0.39/nodes | get all node info
[**SlurmV0039GetPartition**](SlurmAPI.md#SlurmV0039GetPartition) | **Get** /slurm/v0.0.39/partition/{partition_name} | get partition info
[**SlurmV0039GetPartitions**](SlurmAPI.md#SlurmV0039GetPartitions) | **Get** /slurm/v0.0.39/partitions | get all partition info
[**SlurmV0039GetReservation**](SlurmAPI.md#SlurmV0039GetReservation) | **Get** /slurm/v0.0.39/reservation/{reservation_name} | get reservation info
[**SlurmV0039GetReservations**](SlurmAPI.md#SlurmV0039GetReservations) | **Get** /slurm/v0.0.39/reservations | get all reservation info
[**SlurmV0039Ping**](SlurmAPI.md#SlurmV0039Ping) | **Get** /slurm/v0.0.39/ping | ping test
[**SlurmV0039SlurmctldGetLicenses**](SlurmAPI.md#SlurmV0039SlurmctldGetLicenses) | **Get** /slurm/v0.0.39/licenses | get all Slurm tracked license info
[**SlurmV0039SubmitJob**](SlurmAPI.md#SlurmV0039SubmitJob) | **Post** /slurm/v0.0.39/job/submit | submit new job
[**SlurmV0039UpdateJob**](SlurmAPI.md#SlurmV0039UpdateJob) | **Post** /slurm/v0.0.39/job/{job_id} | update job
[**SlurmV0039UpdateNode**](SlurmAPI.md#SlurmV0039UpdateNode) | **Post** /slurm/v0.0.39/node/{node_name} | update node properties
[**SlurmdbV0039AddClusters**](SlurmAPI.md#SlurmdbV0039AddClusters) | **Post** /slurmdb/v0.0.39/clusters | Add clusters
[**SlurmdbV0039AddWckeys**](SlurmAPI.md#SlurmdbV0039AddWckeys) | **Post** /slurmdb/v0.0.39/wckeys | Add wckeys
[**SlurmdbV0039DeleteAccount**](SlurmAPI.md#SlurmdbV0039DeleteAccount) | **Delete** /slurmdb/v0.0.39/account/{account_name} | Delete account
[**SlurmdbV0039DeleteAssociation**](SlurmAPI.md#SlurmdbV0039DeleteAssociation) | **Delete** /slurmdb/v0.0.39/association | Delete association
[**SlurmdbV0039DeleteAssociations**](SlurmAPI.md#SlurmdbV0039DeleteAssociations) | **Delete** /slurmdb/v0.0.39/associations | Delete associations
[**SlurmdbV0039DeleteCluster**](SlurmAPI.md#SlurmdbV0039DeleteCluster) | **Delete** /slurmdb/v0.0.39/cluster/{cluster_name} | Delete cluster
[**SlurmdbV0039DeleteQos**](SlurmAPI.md#SlurmdbV0039DeleteQos) | **Delete** /slurmdb/v0.0.39/qos/{qos_name} | Delete QOS
[**SlurmdbV0039DeleteUser**](SlurmAPI.md#SlurmdbV0039DeleteUser) | **Delete** /slurmdb/v0.0.39/user/{user_name} | Delete user
[**SlurmdbV0039DeleteWckey**](SlurmAPI.md#SlurmdbV0039DeleteWckey) | **Delete** /slurmdb/v0.0.39/wckey/{wckey} | Delete wckey
[**SlurmdbV0039Diag**](SlurmAPI.md#SlurmdbV0039Diag) | **Get** /slurmdb/v0.0.39/diag | Get slurmdb diagnostics
[**SlurmdbV0039GetAccount**](SlurmAPI.md#SlurmdbV0039GetAccount) | **Get** /slurmdb/v0.0.39/account/{account_name} | Get account info
[**SlurmdbV0039GetAccounts**](SlurmAPI.md#SlurmdbV0039GetAccounts) | **Get** /slurmdb/v0.0.39/accounts | Get account list
[**SlurmdbV0039GetAssociation**](SlurmAPI.md#SlurmdbV0039GetAssociation) | **Get** /slurmdb/v0.0.39/association | Get association info
[**SlurmdbV0039GetAssociations**](SlurmAPI.md#SlurmdbV0039GetAssociations) | **Get** /slurmdb/v0.0.39/associations | Get association list
[**SlurmdbV0039GetCluster**](SlurmAPI.md#SlurmdbV0039GetCluster) | **Get** /slurmdb/v0.0.39/cluster/{cluster_name} | Get cluster info
[**SlurmdbV0039GetClusters**](SlurmAPI.md#SlurmdbV0039GetClusters) | **Get** /slurmdb/v0.0.39/clusters | Get cluster list
[**SlurmdbV0039GetConfig**](SlurmAPI.md#SlurmdbV0039GetConfig) | **Get** /slurmdb/v0.0.39/config | Dump all configuration information
[**SlurmdbV0039GetJob**](SlurmAPI.md#SlurmdbV0039GetJob) | **Get** /slurmdb/v0.0.39/job/{job_id} | Get job info
[**SlurmdbV0039GetJobs**](SlurmAPI.md#SlurmdbV0039GetJobs) | **Get** /slurmdb/v0.0.39/jobs | Get job list
[**SlurmdbV0039GetQos**](SlurmAPI.md#SlurmdbV0039GetQos) | **Get** /slurmdb/v0.0.39/qos | Get QOS list
[**SlurmdbV0039GetSingleQos**](SlurmAPI.md#SlurmdbV0039GetSingleQos) | **Get** /slurmdb/v0.0.39/qos/{qos_name} | Get QOS info
[**SlurmdbV0039GetTres**](SlurmAPI.md#SlurmdbV0039GetTres) | **Get** /slurmdb/v0.0.39/tres | Get TRES info
[**SlurmdbV0039GetUser**](SlurmAPI.md#SlurmdbV0039GetUser) | **Get** /slurmdb/v0.0.39/user/{user_name} | Get user info
[**SlurmdbV0039GetUsers**](SlurmAPI.md#SlurmdbV0039GetUsers) | **Get** /slurmdb/v0.0.39/users | Get user list
[**SlurmdbV0039GetWckey**](SlurmAPI.md#SlurmdbV0039GetWckey) | **Get** /slurmdb/v0.0.39/wckey/{wckey} | Get wckey info
[**SlurmdbV0039GetWckeys**](SlurmAPI.md#SlurmdbV0039GetWckeys) | **Get** /slurmdb/v0.0.39/wckeys | Get wckey list
[**SlurmdbV0039SetConfig**](SlurmAPI.md#SlurmdbV0039SetConfig) | **Post** /slurmdb/v0.0.39/config | Load all configuration information
[**SlurmdbV0039UpdateAccounts**](SlurmAPI.md#SlurmdbV0039UpdateAccounts) | **Post** /slurmdb/v0.0.39/accounts | Update accounts
[**SlurmdbV0039UpdateAssociations**](SlurmAPI.md#SlurmdbV0039UpdateAssociations) | **Post** /slurmdb/v0.0.39/associations | Set associations info
[**SlurmdbV0039UpdateQos**](SlurmAPI.md#SlurmdbV0039UpdateQos) | **Post** /slurmdb/v0.0.39/qos | Set QOS info
[**SlurmdbV0039UpdateTres**](SlurmAPI.md#SlurmdbV0039UpdateTres) | **Post** /slurmdb/v0.0.39/tres | Set TRES info
[**SlurmdbV0039UpdateUsers**](SlurmAPI.md#SlurmdbV0039UpdateUsers) | **Post** /slurmdb/v0.0.39/users | Update user



## SlurmV0039CancelJob

> Status SlurmV0039CancelJob(ctx, jobId).Signal(signal).Execute()

cancel or signal job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    jobId := "jobId_example" // string | Slurm Job ID
    signal := "signal_example" // string | signal to send to job (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039CancelJob(context.Background(), jobId).Signal(signal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039CancelJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039CancelJob`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039CancelJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039CancelJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | **string** | signal to send to job | 

### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039DeleteNode

> Status SlurmV0039DeleteNode(ctx, nodeName).Execute()

delete node

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    nodeName := "nodeName_example" // string | Slurm Node Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039DeleteNode(context.Background(), nodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039DeleteNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039DeleteNode`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Slurm Node Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039DeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039Diag

> V0039Diag SlurmV0039Diag(ctx).Execute()

get diagnostics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039Diag(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039Diag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039Diag`: V0039Diag
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039Diag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039DiagRequest struct via the builder pattern


### Return type

[**V0039Diag**](V0039Diag.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetJob

> V0039JobsResponse SlurmV0039GetJob(ctx, jobId).Execute()

get job info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    jobId := "jobId_example" // string | Slurm JobID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039GetJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039GetJob`: V0039JobsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm JobID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0039JobsResponse**](V0039JobsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetJobs

> V0039JobsResponse SlurmV0039GetJobs(ctx).UpdateTime(updateTime).Execute()

get list of jobs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039GetJobs(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039GetJobs`: V0039JobsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0039JobsResponse**](V0039JobsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetNode

> V0039NodesResponse SlurmV0039GetNode(ctx, nodeName).Execute()

get node info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    nodeName := "nodeName_example" // string | Slurm Node Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039GetNode(context.Background(), nodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039GetNode`: V0039NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Slurm Node Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0039NodesResponse**](V0039NodesResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetNodes

> V0039NodesResponse SlurmV0039GetNodes(ctx).UpdateTime(updateTime).Execute()

get all node info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039GetNodes(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039GetNodes`: V0039NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0039NodesResponse**](V0039NodesResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetPartition

> V0039PartitionsResponse SlurmV0039GetPartition(ctx, partitionName).UpdateTime(updateTime).Execute()

get partition info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    partitionName := "partitionName_example" // string | Slurm Partition Name
    updateTime := int64(789) // int64 | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039GetPartition(context.Background(), partitionName).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetPartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039GetPartition`: V0039PartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Slurm Partition Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. | 

### Return type

[**V0039PartitionsResponse**](V0039PartitionsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetPartitions

> V0039PartitionsResponse SlurmV0039GetPartitions(ctx).UpdateTime(updateTime).Execute()

get all partition info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039GetPartitions(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetPartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039GetPartitions`: V0039PartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0039PartitionsResponse**](V0039PartitionsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetReservation

> V0039ReservationsResponse SlurmV0039GetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

get reservation info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    reservationName := "reservationName_example" // string | Slurm Reservation Name
    updateTime := int64(789) // int64 | Filter if no reservation (not limited to reservation in URL) changed since update_time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039GetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039GetReservation`: V0039ReservationsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Slurm Reservation Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if no reservation (not limited to reservation in URL) changed since update_time. | 

### Return type

[**V0039ReservationsResponse**](V0039ReservationsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetReservations

> V0039ReservationsResponse SlurmV0039GetReservations(ctx).UpdateTime(updateTime).Execute()

get all reservation info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039GetReservations(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039GetReservations`: V0039ReservationsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0039ReservationsResponse**](V0039ReservationsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039Ping

> V0039Pings SlurmV0039Ping(ctx).Execute()

ping test

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039Ping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039Ping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039Ping`: V0039Pings
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039PingRequest struct via the builder pattern


### Return type

[**V0039Pings**](V0039Pings.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039SlurmctldGetLicenses

> V0039LicensesInfo SlurmV0039SlurmctldGetLicenses(ctx).Execute()

get all Slurm tracked license info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039SlurmctldGetLicenses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039SlurmctldGetLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039SlurmctldGetLicenses`: V0039LicensesInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039SlurmctldGetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039SlurmctldGetLicensesRequest struct via the builder pattern


### Return type

[**V0039LicensesInfo**](V0039LicensesInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039SubmitJob

> V0039JobSubmissionResponse SlurmV0039SubmitJob(ctx).V0039JobSubmission(v0039JobSubmission).Execute()

submit new job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    v0039JobSubmission := *openapiclient.NewV0039JobSubmission() // V0039JobSubmission | submit new job

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039SubmitJob(context.Background()).V0039JobSubmission(v0039JobSubmission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039SubmitJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039SubmitJob`: V0039JobSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039SubmitJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039SubmitJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0039JobSubmission** | [**V0039JobSubmission**](V0039JobSubmission.md) | submit new job | 

### Return type

[**V0039JobSubmissionResponse**](V0039JobSubmissionResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039UpdateJob

> V0039JobUpdateResponse SlurmV0039UpdateJob(ctx, jobId).V0039JobDescMsg(v0039JobDescMsg).Execute()

update job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    jobId := "jobId_example" // string | Slurm Job ID
    v0039JobDescMsg := *openapiclient.NewV0039JobDescMsg() // V0039JobDescMsg | update job

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039UpdateJob(context.Background(), jobId).V0039JobDescMsg(v0039JobDescMsg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039UpdateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039UpdateJob`: V0039JobUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039UpdateJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039UpdateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0039JobDescMsg** | [**V0039JobDescMsg**](V0039JobDescMsg.md) | update job | 

### Return type

[**V0039JobUpdateResponse**](V0039JobUpdateResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039UpdateNode

> Status SlurmV0039UpdateNode(ctx, nodeName).V0039UpdateNodeMsg(v0039UpdateNodeMsg).Execute()

update node properties

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    nodeName := "nodeName_example" // string | Slurm Node Name
    v0039UpdateNodeMsg := *openapiclient.NewV0039UpdateNodeMsg() // V0039UpdateNodeMsg | update node

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0039UpdateNode(context.Background(), nodeName).V0039UpdateNodeMsg(v0039UpdateNodeMsg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039UpdateNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0039UpdateNode`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039UpdateNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Slurm Node Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039UpdateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0039UpdateNodeMsg** | [**V0039UpdateNodeMsg**](V0039UpdateNodeMsg.md) | update node | 

### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039AddClusters

> Status SlurmdbV0039AddClusters(ctx).Dbv0039ClustersInfo(dbv0039ClustersInfo).Execute()

Add clusters

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    dbv0039ClustersInfo := *openapiclient.NewDbv0039ClustersInfo() // Dbv0039ClustersInfo | Add or update clusters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039AddClusters(context.Background()).Dbv0039ClustersInfo(dbv0039ClustersInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039AddClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039AddClusters`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039AddClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039AddClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0039ClustersInfo** | [**Dbv0039ClustersInfo**](Dbv0039ClustersInfo.md) | Add or update clusters | 

### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039AddWckeys

> Status SlurmdbV0039AddWckeys(ctx).Dbv0039WckeyInfo(dbv0039WckeyInfo).Execute()

Add wckeys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    dbv0039WckeyInfo := *openapiclient.NewDbv0039WckeyInfo() // Dbv0039WckeyInfo | add wckeys (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039AddWckeys(context.Background()).Dbv0039WckeyInfo(dbv0039WckeyInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039AddWckeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039AddWckeys`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039AddWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039AddWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0039WckeyInfo** | [**Dbv0039WckeyInfo**](Dbv0039WckeyInfo.md) | add wckeys | 

### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039DeleteAccount

> Status SlurmdbV0039DeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    accountName := "accountName_example" // string | Slurm Account Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039DeleteAccount(context.Background(), accountName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039DeleteAccount`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Slurm Account Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039DeleteAssociation

> Dbv0039ResponseAssociationsDelete SlurmdbV0039DeleteAssociation(ctx).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()

Delete association

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    cluster := "cluster_example" // string | Cluster name (optional)
    account := "account_example" // string | Account name (optional)
    user := "user_example" // string | User name (optional)
    partition := "partition_example" // string | Partition Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039DeleteAssociation(context.Background()).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039DeleteAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039DeleteAssociation`: Dbv0039ResponseAssociationsDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039DeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039DeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Cluster name | 
 **account** | **string** | Account name | 
 **user** | **string** | User name | 
 **partition** | **string** | Partition Name | 

### Return type

[**Dbv0039ResponseAssociationsDelete**](Dbv0039ResponseAssociationsDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039DeleteAssociations

> Dbv0039ResponseAssociationsDelete SlurmdbV0039DeleteAssociations(ctx).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()

Delete associations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    cluster := "cluster_example" // string | Cluster name (optional)
    account := "account_example" // string | Account name (optional)
    user := "user_example" // string | User name (optional)
    partition := "partition_example" // string | Partition Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039DeleteAssociations(context.Background()).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039DeleteAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039DeleteAssociations`: Dbv0039ResponseAssociationsDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039DeleteAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039DeleteAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Cluster name | 
 **account** | **string** | Account name | 
 **user** | **string** | User name | 
 **partition** | **string** | Partition Name | 

### Return type

[**Dbv0039ResponseAssociationsDelete**](Dbv0039ResponseAssociationsDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039DeleteCluster

> Status SlurmdbV0039DeleteCluster(ctx, clusterName).Execute()

Delete cluster

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    clusterName := "clusterName_example" // string | Slurm cluster name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039DeleteCluster(context.Background(), clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039DeleteCluster`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Slurm cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039DeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039DeleteQos

> Status SlurmdbV0039DeleteQos(ctx, qosName).Execute()

Delete QOS

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    qosName := "qosName_example" // string | Slurm QOS Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039DeleteQos(context.Background(), qosName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039DeleteQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039DeleteQos`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039DeleteQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qosName** | **string** | Slurm QOS Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039DeleteQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039DeleteUser

> Status SlurmdbV0039DeleteUser(ctx, userName).Execute()

Delete user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    userName := "userName_example" // string | Slurm User Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039DeleteUser(context.Background(), userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039DeleteUser`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | Slurm User Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039DeleteWckey

> Status SlurmdbV0039DeleteWckey(ctx, wckey).Execute()

Delete wckey

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    wckey := "wckey_example" // string | Slurm wckey name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039DeleteWckey(context.Background(), wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039DeleteWckey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039DeleteWckey`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039DeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wckey** | **string** | Slurm wckey name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039DeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039Diag

> Dbv0039Diag SlurmdbV0039Diag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039Diag(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039Diag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039Diag`: Dbv0039Diag
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039Diag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039DiagRequest struct via the builder pattern


### Return type

[**Dbv0039Diag**](Dbv0039Diag.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetAccount

> Dbv0039AccountInfo SlurmdbV0039GetAccount(ctx, accountName).WithDeleted(withDeleted).Execute()

Get account info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    accountName := "accountName_example" // string | Slurm Account Name
    withDeleted := "withDeleted_example" // string | Include deleted accounts. False by default. (optional) (default to "false")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetAccount(context.Background(), accountName).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetAccount`: Dbv0039AccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Slurm Account Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted accounts. False by default. | [default to &quot;false&quot;]

### Return type

[**Dbv0039AccountInfo**](Dbv0039AccountInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetAccounts

> Dbv0039AccountInfo SlurmdbV0039GetAccounts(ctx).WithDeleted(withDeleted).Execute()

Get account list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    withDeleted := "withDeleted_example" // string | Include deleted accounts. False by default. (optional) (default to "false")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetAccounts(context.Background()).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetAccounts`: Dbv0039AccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withDeleted** | **string** | Include deleted accounts. False by default. | [default to &quot;false&quot;]

### Return type

[**Dbv0039AccountInfo**](Dbv0039AccountInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetAssociation

> Dbv0039AssociationsInfo SlurmdbV0039GetAssociation(ctx).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()

Get association info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    cluster := "cluster_example" // string | Cluster name (optional)
    account := "account_example" // string | Account name (optional)
    user := "user_example" // string | User name (optional)
    partition := "partition_example" // string | Partition Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetAssociation(context.Background()).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetAssociation`: Dbv0039AssociationsInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Cluster name | 
 **account** | **string** | Account name | 
 **user** | **string** | User name | 
 **partition** | **string** | Partition Name | 

### Return type

[**Dbv0039AssociationsInfo**](Dbv0039AssociationsInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetAssociations

> Dbv0039AssociationsInfo SlurmdbV0039GetAssociations(ctx).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()

Get association list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    cluster := "cluster_example" // string | Cluster name (optional)
    account := "account_example" // string | Account name (optional)
    user := "user_example" // string | User name (optional)
    partition := "partition_example" // string | Partition Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetAssociations(context.Background()).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetAssociations`: Dbv0039AssociationsInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Cluster name | 
 **account** | **string** | Account name | 
 **user** | **string** | User name | 
 **partition** | **string** | Partition Name | 

### Return type

[**Dbv0039AssociationsInfo**](Dbv0039AssociationsInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetCluster

> Dbv0039ClustersInfo SlurmdbV0039GetCluster(ctx, clusterName).Execute()

Get cluster info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    clusterName := "clusterName_example" // string | Slurm cluster name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetCluster(context.Background(), clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetCluster`: Dbv0039ClustersInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Slurm cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0039ClustersInfo**](Dbv0039ClustersInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetClusters

> Dbv0039ClustersInfo SlurmdbV0039GetClusters(ctx).Execute()

Get cluster list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetClusters`: Dbv0039ClustersInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetClustersRequest struct via the builder pattern


### Return type

[**Dbv0039ClustersInfo**](Dbv0039ClustersInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetConfig

> Dbv0039ConfigInfo SlurmdbV0039GetConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetConfig`: Dbv0039ConfigInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetConfigRequest struct via the builder pattern


### Return type

[**Dbv0039ConfigInfo**](Dbv0039ConfigInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetJob

> Dbv0039JobInfo SlurmdbV0039GetJob(ctx, jobId).Execute()

Get job info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    jobId := "jobId_example" // string | Slurm JobID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetJob`: Dbv0039JobInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm JobID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0039JobInfo**](Dbv0039JobInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetJobs

> Dbv0039JobInfo SlurmdbV0039GetJobs(ctx).Users(users).SubmitTime(submitTime).StartTime(startTime).EndTime(endTime).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SkipSteps(skipSteps).DisableWaitForResult(disableWaitForResult).ExitCode(exitCode).Format(format).Group(group).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).State(state).Step(step).Node(node).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    users := "users_example" // string | Filter by comma delimited list of user names (optional)
    submitTime := "submitTime_example" // string | Filter by submission time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] (optional)
    startTime := "startTime_example" // string | Filter by start time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] (optional)
    endTime := "endTime_example" // string | Filter by end time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] (optional)
    account := "account_example" // string | Comma delimited list of accounts to match (optional)
    association := "association_example" // string | Comma delimited list of associations to match (optional)
    cluster := "cluster_example" // string | Comma delimited list of cluster to match (optional)
    constraints := "constraints_example" // string | Comma delimited list of constraints to match (optional)
    cpusMax := "cpusMax_example" // string | Number of CPUs high range (optional)
    cpusMin := "cpusMin_example" // string | Number of CPUs low range (optional)
    skipSteps := "skipSteps_example" // string | Report job step information (optional) (default to "false")
    disableWaitForResult := "disableWaitForResult_example" // string | Disable waiting for result from slurmdbd (optional) (default to "false")
    exitCode := "exitCode_example" // string | Exit code of job (optional)
    format := "format_example" // string | Comma delimited list of formats to match (optional)
    group := "group_example" // string | Comma delimited list of groups to match (optional)
    jobName := "jobName_example" // string | Comma delimited list of job names to match (optional)
    nodesMax := "nodesMax_example" // string | Number of nodes high range (optional)
    nodesMin := "nodesMin_example" // string | Number of nodes low range (optional)
    partition := "partition_example" // string | Comma delimited list of partitions to match (optional)
    qos := "qos_example" // string | Comma delimited list of QOS to match (optional)
    reason := "reason_example" // string | Comma delimited list of job reasons to match (optional)
    reservation := "reservation_example" // string | Comma delimited list of reservations to match (optional)
    state := "state_example" // string | Comma delimited list of states to match (optional)
    step := "step_example" // string | Comma delimited list of job steps to match (optional)
    node := "node_example" // string | Comma delimited list of used nodes to match (optional)
    wckey := "wckey_example" // string | Comma delimited list of wckeys to match (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetJobs(context.Background()).Users(users).SubmitTime(submitTime).StartTime(startTime).EndTime(endTime).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SkipSteps(skipSteps).DisableWaitForResult(disableWaitForResult).ExitCode(exitCode).Format(format).Group(group).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).State(state).Step(step).Node(node).Wckey(wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetJobs`: Dbv0039JobInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **users** | **string** | Filter by comma delimited list of user names | 
 **submitTime** | **string** | Filter by submission time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] | 
 **startTime** | **string** | Filter by start time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] | 
 **endTime** | **string** | Filter by end time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] | 
 **account** | **string** | Comma delimited list of accounts to match | 
 **association** | **string** | Comma delimited list of associations to match | 
 **cluster** | **string** | Comma delimited list of cluster to match | 
 **constraints** | **string** | Comma delimited list of constraints to match | 
 **cpusMax** | **string** | Number of CPUs high range | 
 **cpusMin** | **string** | Number of CPUs low range | 
 **skipSteps** | **string** | Report job step information | [default to &quot;false&quot;]
 **disableWaitForResult** | **string** | Disable waiting for result from slurmdbd | [default to &quot;false&quot;]
 **exitCode** | **string** | Exit code of job | 
 **format** | **string** | Comma delimited list of formats to match | 
 **group** | **string** | Comma delimited list of groups to match | 
 **jobName** | **string** | Comma delimited list of job names to match | 
 **nodesMax** | **string** | Number of nodes high range | 
 **nodesMin** | **string** | Number of nodes low range | 
 **partition** | **string** | Comma delimited list of partitions to match | 
 **qos** | **string** | Comma delimited list of QOS to match | 
 **reason** | **string** | Comma delimited list of job reasons to match | 
 **reservation** | **string** | Comma delimited list of reservations to match | 
 **state** | **string** | Comma delimited list of states to match | 
 **step** | **string** | Comma delimited list of job steps to match | 
 **node** | **string** | Comma delimited list of used nodes to match | 
 **wckey** | **string** | Comma delimited list of wckeys to match | 

### Return type

[**Dbv0039JobInfo**](Dbv0039JobInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetQos

> Dbv0039QosInfo SlurmdbV0039GetQos(ctx).WithDeleted(withDeleted).Execute()

Get QOS list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    withDeleted := "withDeleted_example" // string | Include deleted QOSs. False by default. (optional) (default to "false")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetQos(context.Background()).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetQos`: Dbv0039QosInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withDeleted** | **string** | Include deleted QOSs. False by default. | [default to &quot;false&quot;]

### Return type

[**Dbv0039QosInfo**](Dbv0039QosInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetSingleQos

> Dbv0039QosInfo SlurmdbV0039GetSingleQos(ctx, qosName).WithDeleted(withDeleted).Execute()

Get QOS info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    qosName := "qosName_example" // string | Slurm QOS Name
    withDeleted := "withDeleted_example" // string | Include deleted QOSs. False by default. (optional) (default to "false")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetSingleQos(context.Background(), qosName).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetSingleQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetSingleQos`: Dbv0039QosInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qosName** | **string** | Slurm QOS Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted QOSs. False by default. | [default to &quot;false&quot;]

### Return type

[**Dbv0039QosInfo**](Dbv0039QosInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetTres

> Dbv0039TresInfo SlurmdbV0039GetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetTres(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetTres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetTres`: Dbv0039TresInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetTresRequest struct via the builder pattern


### Return type

[**Dbv0039TresInfo**](Dbv0039TresInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetUser

> Dbv0039UserInfo SlurmdbV0039GetUser(ctx, userName).WithDeleted(withDeleted).Execute()

Get user info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    userName := "userName_example" // string | Slurm User Name
    withDeleted := "withDeleted_example" // string | Include deleted users. False by default. (optional) (default to "false")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetUser(context.Background(), userName).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetUser`: Dbv0039UserInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | Slurm User Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted users. False by default. | [default to &quot;false&quot;]

### Return type

[**Dbv0039UserInfo**](Dbv0039UserInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetUsers

> Dbv0039UserInfo SlurmdbV0039GetUsers(ctx).WithDeleted(withDeleted).Execute()

Get user list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    withDeleted := "withDeleted_example" // string | Include deleted users. False by default. (optional) (default to "false")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetUsers(context.Background()).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetUsers`: Dbv0039UserInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withDeleted** | **string** | Include deleted users. False by default. | [default to &quot;false&quot;]

### Return type

[**Dbv0039UserInfo**](Dbv0039UserInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetWckey

> Dbv0039WckeyInfo SlurmdbV0039GetWckey(ctx, wckey).Execute()

Get wckey info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    wckey := "wckey_example" // string | Slurm wckey name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetWckey(context.Background(), wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetWckey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetWckey`: Dbv0039WckeyInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wckey** | **string** | Slurm wckey name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0039WckeyInfo**](Dbv0039WckeyInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039GetWckeys

> Dbv0039WckeyInfo SlurmdbV0039GetWckeys(ctx).Execute()

Get wckey list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039GetWckeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039GetWckeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039GetWckeys`: Dbv0039WckeyInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039GetWckeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039GetWckeysRequest struct via the builder pattern


### Return type

[**Dbv0039WckeyInfo**](Dbv0039WckeyInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039SetConfig

> Status SlurmdbV0039SetConfig(ctx).Dbv0039SetConfig(dbv0039SetConfig).Execute()

Load all configuration information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    dbv0039SetConfig := *openapiclient.NewDbv0039SetConfig() // Dbv0039SetConfig | Add or update config (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039SetConfig(context.Background()).Dbv0039SetConfig(dbv0039SetConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039SetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039SetConfig`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039SetConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039SetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0039SetConfig** | [**Dbv0039SetConfig**](Dbv0039SetConfig.md) | Add or update config | 

### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039UpdateAccounts

> Status SlurmdbV0039UpdateAccounts(ctx).Dbv0039AccountInfo(dbv0039AccountInfo).Execute()

Update accounts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    dbv0039AccountInfo := *openapiclient.NewDbv0039AccountInfo() // Dbv0039AccountInfo | update/create accounts

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039UpdateAccounts(context.Background()).Dbv0039AccountInfo(dbv0039AccountInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039UpdateAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039UpdateAccounts`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039UpdateAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039UpdateAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0039AccountInfo** | [**Dbv0039AccountInfo**](Dbv0039AccountInfo.md) | update/create accounts | 

### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039UpdateAssociations

> Status SlurmdbV0039UpdateAssociations(ctx).Dbv0039AssociationsInfo(dbv0039AssociationsInfo).Execute()

Set associations info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    dbv0039AssociationsInfo := *openapiclient.NewDbv0039AssociationsInfo() // Dbv0039AssociationsInfo | Add or update associations

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039UpdateAssociations(context.Background()).Dbv0039AssociationsInfo(dbv0039AssociationsInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039UpdateAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039UpdateAssociations`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039UpdateAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039UpdateAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0039AssociationsInfo** | [**Dbv0039AssociationsInfo**](Dbv0039AssociationsInfo.md) | Add or update associations | 

### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039UpdateQos

> Status SlurmdbV0039UpdateQos(ctx).Dbv0039UpdateQos(dbv0039UpdateQos).Execute()

Set QOS info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    dbv0039UpdateQos := *openapiclient.NewDbv0039UpdateQos() // Dbv0039UpdateQos | Add or update QOSs

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039UpdateQos(context.Background()).Dbv0039UpdateQos(dbv0039UpdateQos).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039UpdateQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039UpdateQos`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039UpdateQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039UpdateQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0039UpdateQos** | [**Dbv0039UpdateQos**](Dbv0039UpdateQos.md) | Add or update QOSs | 

### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039UpdateTres

> Status SlurmdbV0039UpdateTres(ctx).Dbv0039TresUpdate(dbv0039TresUpdate).Execute()

Set TRES info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    dbv0039TresUpdate := *openapiclient.NewDbv0039TresUpdate() // Dbv0039TresUpdate | Add or Update TRES

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039UpdateTres(context.Background()).Dbv0039TresUpdate(dbv0039TresUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039UpdateTres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039UpdateTres`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039UpdateTres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039UpdateTresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0039TresUpdate** | [**Dbv0039TresUpdate**](Dbv0039TresUpdate.md) | Add or Update TRES | 

### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0039UpdateUsers

> Status SlurmdbV0039UpdateUsers(ctx).Dbv0039UpdateUsers(dbv0039UpdateUsers).Execute()

Update user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/zhuqw12/slurmrestapi"
)

func main() {
    dbv0039UpdateUsers := *openapiclient.NewDbv0039UpdateUsers() // Dbv0039UpdateUsers | add or update user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0039UpdateUsers(context.Background()).Dbv0039UpdateUsers(dbv0039UpdateUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0039UpdateUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0039UpdateUsers`: Status
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0039UpdateUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0039UpdateUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0039UpdateUsers** | [**Dbv0039UpdateUsers**](Dbv0039UpdateUsers.md) | add or update user | 

### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

