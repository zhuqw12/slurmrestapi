# \SlurmAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmV0038CancelJob**](SlurmAPI.md#SlurmV0038CancelJob) | **Delete** /slurm/v0.0.38/job/{job_id} | cancel or signal job
[**SlurmV0038Diag**](SlurmAPI.md#SlurmV0038Diag) | **Get** /slurm/v0.0.38/diag | get diagnostics
[**SlurmV0038GetJob**](SlurmAPI.md#SlurmV0038GetJob) | **Get** /slurm/v0.0.38/job/{job_id} | get job info
[**SlurmV0038GetJobs**](SlurmAPI.md#SlurmV0038GetJobs) | **Get** /slurm/v0.0.38/jobs | get list of jobs
[**SlurmV0038GetNode**](SlurmAPI.md#SlurmV0038GetNode) | **Get** /slurm/v0.0.38/node/{node_name} | get node info
[**SlurmV0038GetNodes**](SlurmAPI.md#SlurmV0038GetNodes) | **Get** /slurm/v0.0.38/nodes | get all node info
[**SlurmV0038GetPartition**](SlurmAPI.md#SlurmV0038GetPartition) | **Get** /slurm/v0.0.38/partition/{partition_name} | get partition info
[**SlurmV0038GetPartitions**](SlurmAPI.md#SlurmV0038GetPartitions) | **Get** /slurm/v0.0.38/partitions | get all partition info
[**SlurmV0038GetReservation**](SlurmAPI.md#SlurmV0038GetReservation) | **Get** /slurm/v0.0.38/reservation/{reservation_name} | get reservation info
[**SlurmV0038GetReservations**](SlurmAPI.md#SlurmV0038GetReservations) | **Get** /slurm/v0.0.38/reservations | get all reservation info
[**SlurmV0038Ping**](SlurmAPI.md#SlurmV0038Ping) | **Get** /slurm/v0.0.38/ping | ping test
[**SlurmV0038SlurmctldGetLicenses**](SlurmAPI.md#SlurmV0038SlurmctldGetLicenses) | **Get** /slurm/v0.0.38/licenses | get all Slurm tracked license info
[**SlurmV0038SubmitJob**](SlurmAPI.md#SlurmV0038SubmitJob) | **Post** /slurm/v0.0.38/job/submit | submit new job
[**SlurmV0038UpdateJob**](SlurmAPI.md#SlurmV0038UpdateJob) | **Post** /slurm/v0.0.38/job/{job_id} | update job
[**SlurmdbV0038AddClusters**](SlurmAPI.md#SlurmdbV0038AddClusters) | **Post** /slurmdb/v0.0.38/clusters | Add clusters
[**SlurmdbV0038AddWckeys**](SlurmAPI.md#SlurmdbV0038AddWckeys) | **Post** /slurmdb/v0.0.38/wckeys | Add wckeys
[**SlurmdbV0038DeleteAccount**](SlurmAPI.md#SlurmdbV0038DeleteAccount) | **Delete** /slurmdb/v0.0.38/account/{account_name} | Delete account
[**SlurmdbV0038DeleteAssociation**](SlurmAPI.md#SlurmdbV0038DeleteAssociation) | **Delete** /slurmdb/v0.0.38/association | Delete association
[**SlurmdbV0038DeleteAssociations**](SlurmAPI.md#SlurmdbV0038DeleteAssociations) | **Delete** /slurmdb/v0.0.38/associations | Delete associations
[**SlurmdbV0038DeleteCluster**](SlurmAPI.md#SlurmdbV0038DeleteCluster) | **Delete** /slurmdb/v0.0.38/cluster/{cluster_name} | Delete cluster
[**SlurmdbV0038DeleteQos**](SlurmAPI.md#SlurmdbV0038DeleteQos) | **Delete** /slurmdb/v0.0.38/qos/{qos_name} | Delete QOS
[**SlurmdbV0038DeleteUser**](SlurmAPI.md#SlurmdbV0038DeleteUser) | **Delete** /slurmdb/v0.0.38/user/{user_name} | Delete user
[**SlurmdbV0038DeleteWckey**](SlurmAPI.md#SlurmdbV0038DeleteWckey) | **Delete** /slurmdb/v0.0.38/wckey/{wckey} | Delete wckey
[**SlurmdbV0038Diag**](SlurmAPI.md#SlurmdbV0038Diag) | **Get** /slurmdb/v0.0.38/diag | Get slurmdb diagnostics
[**SlurmdbV0038GetAccount**](SlurmAPI.md#SlurmdbV0038GetAccount) | **Get** /slurmdb/v0.0.38/account/{account_name} | Get account info
[**SlurmdbV0038GetAccounts**](SlurmAPI.md#SlurmdbV0038GetAccounts) | **Get** /slurmdb/v0.0.38/accounts | Get account list
[**SlurmdbV0038GetAssociation**](SlurmAPI.md#SlurmdbV0038GetAssociation) | **Get** /slurmdb/v0.0.38/association | Get association info
[**SlurmdbV0038GetAssociations**](SlurmAPI.md#SlurmdbV0038GetAssociations) | **Get** /slurmdb/v0.0.38/associations | Get association list
[**SlurmdbV0038GetCluster**](SlurmAPI.md#SlurmdbV0038GetCluster) | **Get** /slurmdb/v0.0.38/cluster/{cluster_name} | Get cluster info
[**SlurmdbV0038GetClusters**](SlurmAPI.md#SlurmdbV0038GetClusters) | **Get** /slurmdb/v0.0.38/clusters | Get cluster list
[**SlurmdbV0038GetConfig**](SlurmAPI.md#SlurmdbV0038GetConfig) | **Get** /slurmdb/v0.0.38/config | Dump all configuration information
[**SlurmdbV0038GetJob**](SlurmAPI.md#SlurmdbV0038GetJob) | **Get** /slurmdb/v0.0.38/job/{job_id} | Get job info
[**SlurmdbV0038GetJobs**](SlurmAPI.md#SlurmdbV0038GetJobs) | **Get** /slurmdb/v0.0.38/jobs | Get job list
[**SlurmdbV0038GetQos**](SlurmAPI.md#SlurmdbV0038GetQos) | **Get** /slurmdb/v0.0.38/qos | Get QOS list
[**SlurmdbV0038GetSingleQos**](SlurmAPI.md#SlurmdbV0038GetSingleQos) | **Get** /slurmdb/v0.0.38/qos/{qos_name} | Get QOS info
[**SlurmdbV0038GetTres**](SlurmAPI.md#SlurmdbV0038GetTres) | **Get** /slurmdb/v0.0.38/tres | Get TRES info
[**SlurmdbV0038GetUser**](SlurmAPI.md#SlurmdbV0038GetUser) | **Get** /slurmdb/v0.0.38/user/{user_name} | Get user info
[**SlurmdbV0038GetUsers**](SlurmAPI.md#SlurmdbV0038GetUsers) | **Get** /slurmdb/v0.0.38/users | Get user list
[**SlurmdbV0038GetWckey**](SlurmAPI.md#SlurmdbV0038GetWckey) | **Get** /slurmdb/v0.0.38/wckey/{wckey} | Get wckey info
[**SlurmdbV0038GetWckeys**](SlurmAPI.md#SlurmdbV0038GetWckeys) | **Get** /slurmdb/v0.0.38/wckeys | Get wckey list
[**SlurmdbV0038SetConfig**](SlurmAPI.md#SlurmdbV0038SetConfig) | **Post** /slurmdb/v0.0.38/config | Load all configuration information
[**SlurmdbV0038UpdateAccount**](SlurmAPI.md#SlurmdbV0038UpdateAccount) | **Post** /slurmdb/v0.0.38/accounts | Update accounts
[**SlurmdbV0038UpdateAssociations**](SlurmAPI.md#SlurmdbV0038UpdateAssociations) | **Post** /slurmdb/v0.0.38/associations | Set associations info
[**SlurmdbV0038UpdateQos**](SlurmAPI.md#SlurmdbV0038UpdateQos) | **Post** /slurmdb/v0.0.38/qos | Set QOS info
[**SlurmdbV0038UpdateTres**](SlurmAPI.md#SlurmdbV0038UpdateTres) | **Post** /slurmdb/v0.0.38/tres | Set TRES info
[**SlurmdbV0038UpdateUsers**](SlurmAPI.md#SlurmdbV0038UpdateUsers) | **Post** /slurmdb/v0.0.38/users | Update user



## SlurmV0038CancelJob

> SlurmV0038CancelJob(ctx, jobId).Signal(signal).Execute()

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
    signal := openapiclient.v0.0.38_signal("HUP") // V0038Signal | signal to send to job (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlurmAPI.SlurmV0038CancelJob(context.Background(), jobId).Signal(signal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038CancelJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038CancelJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | [**V0038Signal**](V0038Signal.md) | signal to send to job | 

### Return type

 (empty response body)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038Diag

> V0038Diag SlurmV0038Diag(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmV0038Diag(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038Diag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0038Diag`: V0038Diag
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038Diag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038DiagRequest struct via the builder pattern


### Return type

[**V0038Diag**](V0038Diag.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetJob

> V0038JobsResponse SlurmV0038GetJob(ctx, jobId).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmV0038GetJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0038GetJob`: V0038JobsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm JobID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0038JobsResponse**](V0038JobsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetJobs

> V0038JobsResponse SlurmV0038GetJobs(ctx).UpdateTime(updateTime).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmV0038GetJobs(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0038GetJobs`: V0038JobsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0038JobsResponse**](V0038JobsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetNode

> V0038NodesResponse SlurmV0038GetNode(ctx, nodeName).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmV0038GetNode(context.Background(), nodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0038GetNode`: V0038NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Slurm Node Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0038NodesResponse**](V0038NodesResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetNodes

> V0038NodesResponse SlurmV0038GetNodes(ctx).UpdateTime(updateTime).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmV0038GetNodes(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0038GetNodes`: V0038NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0038NodesResponse**](V0038NodesResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetPartition

> V0038PartitionsResponse SlurmV0038GetPartition(ctx, partitionName).UpdateTime(updateTime).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmV0038GetPartition(context.Background(), partitionName).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetPartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0038GetPartition`: V0038PartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Slurm Partition Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. | 

### Return type

[**V0038PartitionsResponse**](V0038PartitionsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetPartitions

> V0038PartitionsResponse SlurmV0038GetPartitions(ctx).UpdateTime(updateTime).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmV0038GetPartitions(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetPartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0038GetPartitions`: V0038PartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0038PartitionsResponse**](V0038PartitionsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetReservation

> V0038ReservationsResponse SlurmV0038GetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmV0038GetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0038GetReservation`: V0038ReservationsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Slurm Reservation Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if no reservation (not limited to reservation in URL) changed since update_time. | 

### Return type

[**V0038ReservationsResponse**](V0038ReservationsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetReservations

> V0038ReservationsResponse SlurmV0038GetReservations(ctx).UpdateTime(updateTime).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmV0038GetReservations(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0038GetReservations`: V0038ReservationsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0038ReservationsResponse**](V0038ReservationsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038Ping

> V0038Pings SlurmV0038Ping(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmV0038Ping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038Ping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0038Ping`: V0038Pings
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038PingRequest struct via the builder pattern


### Return type

[**V0038Pings**](V0038Pings.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038SlurmctldGetLicenses

> V0038Licenses SlurmV0038SlurmctldGetLicenses(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmV0038SlurmctldGetLicenses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038SlurmctldGetLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0038SlurmctldGetLicenses`: V0038Licenses
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038SlurmctldGetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038SlurmctldGetLicensesRequest struct via the builder pattern


### Return type

[**V0038Licenses**](V0038Licenses.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038SubmitJob

> V0038JobSubmissionResponse SlurmV0038SubmitJob(ctx).V0038JobSubmission(v0038JobSubmission).Execute()

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
    v0038JobSubmission := *openapiclient.NewV0038JobSubmission("Script_example") // V0038JobSubmission | submit new job

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmV0038SubmitJob(context.Background()).V0038JobSubmission(v0038JobSubmission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038SubmitJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmV0038SubmitJob`: V0038JobSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038SubmitJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038SubmitJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0038JobSubmission** | [**V0038JobSubmission**](V0038JobSubmission.md) | submit new job | 

### Return type

[**V0038JobSubmissionResponse**](V0038JobSubmissionResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038UpdateJob

> SlurmV0038UpdateJob(ctx, jobId).V0038JobProperties(v0038JobProperties).Execute()

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
    v0038JobProperties := *openapiclient.NewV0038JobProperties(map[string]interface{}(123)) // V0038JobProperties | update job

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlurmAPI.SlurmV0038UpdateJob(context.Background(), jobId).V0038JobProperties(v0038JobProperties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038UpdateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038UpdateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0038JobProperties** | [**V0038JobProperties**](V0038JobProperties.md) | update job | 

### Return type

 (empty response body)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038AddClusters

> Dbv0038ResponseClusterAdd SlurmdbV0038AddClusters(ctx).Dbv0038ClustersProperties(dbv0038ClustersProperties).Execute()

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
    dbv0038ClustersProperties := *openapiclient.NewDbv0038ClustersProperties() // Dbv0038ClustersProperties | Add or update clusters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038AddClusters(context.Background()).Dbv0038ClustersProperties(dbv0038ClustersProperties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038AddClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038AddClusters`: Dbv0038ResponseClusterAdd
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038AddClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038AddClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0038ClustersProperties** | [**Dbv0038ClustersProperties**](Dbv0038ClustersProperties.md) | Add or update clusters | 

### Return type

[**Dbv0038ResponseClusterAdd**](Dbv0038ResponseClusterAdd.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038AddWckeys

> Dbv0038ResponseWckeyAdd SlurmdbV0038AddWckeys(ctx).Dbv0038WckeyInfo(dbv0038WckeyInfo).Execute()

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
    dbv0038WckeyInfo := *openapiclient.NewDbv0038WckeyInfo() // Dbv0038WckeyInfo | add wckeys (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038AddWckeys(context.Background()).Dbv0038WckeyInfo(dbv0038WckeyInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038AddWckeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038AddWckeys`: Dbv0038ResponseWckeyAdd
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038AddWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038AddWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0038WckeyInfo** | [**Dbv0038WckeyInfo**](Dbv0038WckeyInfo.md) | add wckeys | 

### Return type

[**Dbv0038ResponseWckeyAdd**](Dbv0038ResponseWckeyAdd.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038DeleteAccount

> Dbv0038ResponseAccountDelete SlurmdbV0038DeleteAccount(ctx, accountName).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038DeleteAccount(context.Background(), accountName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038DeleteAccount`: Dbv0038ResponseAccountDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Slurm Account Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0038ResponseAccountDelete**](Dbv0038ResponseAccountDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038DeleteAssociation

> Dbv0038ResponseAssociationsDelete SlurmdbV0038DeleteAssociation(ctx).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038DeleteAssociation(context.Background()).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038DeleteAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038DeleteAssociation`: Dbv0038ResponseAssociationsDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038DeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038DeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Cluster name | 
 **account** | **string** | Account name | 
 **user** | **string** | User name | 
 **partition** | **string** | Partition Name | 

### Return type

[**Dbv0038ResponseAssociationsDelete**](Dbv0038ResponseAssociationsDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038DeleteAssociations

> Dbv0038ResponseAssociationsDelete SlurmdbV0038DeleteAssociations(ctx).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038DeleteAssociations(context.Background()).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038DeleteAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038DeleteAssociations`: Dbv0038ResponseAssociationsDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038DeleteAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038DeleteAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Cluster name | 
 **account** | **string** | Account name | 
 **user** | **string** | User name | 
 **partition** | **string** | Partition Name | 

### Return type

[**Dbv0038ResponseAssociationsDelete**](Dbv0038ResponseAssociationsDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038DeleteCluster

> Dbv0038ResponseClusterDelete SlurmdbV0038DeleteCluster(ctx, clusterName).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038DeleteCluster(context.Background(), clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038DeleteCluster`: Dbv0038ResponseClusterDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Slurm cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038DeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0038ResponseClusterDelete**](Dbv0038ResponseClusterDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038DeleteQos

> Dbv0038ResponseQosDelete SlurmdbV0038DeleteQos(ctx, qosName).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038DeleteQos(context.Background(), qosName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038DeleteQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038DeleteQos`: Dbv0038ResponseQosDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038DeleteQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qosName** | **string** | Slurm QOS Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038DeleteQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0038ResponseQosDelete**](Dbv0038ResponseQosDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038DeleteUser

> Dbv0038ResponseUserDelete SlurmdbV0038DeleteUser(ctx, userName).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038DeleteUser(context.Background(), userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038DeleteUser`: Dbv0038ResponseUserDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | Slurm User Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0038ResponseUserDelete**](Dbv0038ResponseUserDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038DeleteWckey

> Dbv0038ResponseWckeyDelete SlurmdbV0038DeleteWckey(ctx, wckey).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038DeleteWckey(context.Background(), wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038DeleteWckey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038DeleteWckey`: Dbv0038ResponseWckeyDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038DeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wckey** | **string** | Slurm wckey name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038DeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0038ResponseWckeyDelete**](Dbv0038ResponseWckeyDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038Diag

> Dbv0038Diag SlurmdbV0038Diag(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038Diag(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038Diag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038Diag`: Dbv0038Diag
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038Diag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038DiagRequest struct via the builder pattern


### Return type

[**Dbv0038Diag**](Dbv0038Diag.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetAccount

> Dbv0038AccountInfo SlurmdbV0038GetAccount(ctx, accountName).WithDeleted(withDeleted).Execute()

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
    withDeleted := true // bool | Include deleted accounts. False by default. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetAccount(context.Background(), accountName).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetAccount`: Dbv0038AccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Slurm Account Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **bool** | Include deleted accounts. False by default. | 

### Return type

[**Dbv0038AccountInfo**](Dbv0038AccountInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetAccounts

> Dbv0038AccountInfo SlurmdbV0038GetAccounts(ctx).WithDeleted(withDeleted).Execute()

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
    withDeleted := true // bool | Include deleted accounts. False by default. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetAccounts(context.Background()).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetAccounts`: Dbv0038AccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withDeleted** | **bool** | Include deleted accounts. False by default. | 

### Return type

[**Dbv0038AccountInfo**](Dbv0038AccountInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetAssociation

> Dbv0038AssociationsInfo SlurmdbV0038GetAssociation(ctx).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetAssociation(context.Background()).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetAssociation`: Dbv0038AssociationsInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Cluster name | 
 **account** | **string** | Account name | 
 **user** | **string** | User name | 
 **partition** | **string** | Partition Name | 

### Return type

[**Dbv0038AssociationsInfo**](Dbv0038AssociationsInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetAssociations

> Dbv0038AssociationsInfo SlurmdbV0038GetAssociations(ctx).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetAssociations(context.Background()).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetAssociations`: Dbv0038AssociationsInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Cluster name | 
 **account** | **string** | Account name | 
 **user** | **string** | User name | 
 **partition** | **string** | Partition Name | 

### Return type

[**Dbv0038AssociationsInfo**](Dbv0038AssociationsInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetCluster

> Dbv0038ClusterInfo SlurmdbV0038GetCluster(ctx, clusterName).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetCluster(context.Background(), clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetCluster`: Dbv0038ClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Slurm cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0038ClusterInfo**](Dbv0038ClusterInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetClusters

> Dbv0038ClusterInfo SlurmdbV0038GetClusters(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetClusters`: Dbv0038ClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetClustersRequest struct via the builder pattern


### Return type

[**Dbv0038ClusterInfo**](Dbv0038ClusterInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetConfig

> Dbv0038ConfigInfo SlurmdbV0038GetConfig(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetConfig`: Dbv0038ConfigInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetConfigRequest struct via the builder pattern


### Return type

[**Dbv0038ConfigInfo**](Dbv0038ConfigInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetJob

> Dbv0038JobInfo SlurmdbV0038GetJob(ctx, jobId).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetJob`: Dbv0038JobInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm JobID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0038JobInfo**](Dbv0038JobInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetJobs

> Dbv0038JobInfo SlurmdbV0038GetJobs(ctx).SubmitTime(submitTime).StartTime(startTime).EndTime(endTime).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SkipSteps(skipSteps).DisableWaitForResult(disableWaitForResult).ExitCode(exitCode).Format(format).Group(group).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).State(state).Step(step).Node(node).Wckey(wckey).Execute()

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
    submitTime := "submitTime_example" // string | Filter by submission time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] (optional)
    startTime := "startTime_example" // string | Filter by start time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] (optional)
    endTime := "endTime_example" // string | Filter by end time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] (optional)
    account := "account_example" // string | Comma delimited list of accounts to match (optional)
    association := "association_example" // string | Comma delimited list of associations to match (optional)
    cluster := "cluster_example" // string | Comma delimited list of cluster to match (optional)
    constraints := "constraints_example" // string | Comma delimited list of constraints to match (optional)
    cpusMax := "cpusMax_example" // string | Number of CPUs high range (optional)
    cpusMin := "cpusMin_example" // string | Number of CPUs low range (optional)
    skipSteps := true // bool | Report job step information (optional)
    disableWaitForResult := true // bool | Disable waiting for result from slurmdbd (optional)
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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetJobs(context.Background()).SubmitTime(submitTime).StartTime(startTime).EndTime(endTime).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SkipSteps(skipSteps).DisableWaitForResult(disableWaitForResult).ExitCode(exitCode).Format(format).Group(group).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).State(state).Step(step).Node(node).Wckey(wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetJobs`: Dbv0038JobInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitTime** | **string** | Filter by submission time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] | 
 **startTime** | **string** | Filter by start time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] | 
 **endTime** | **string** | Filter by end time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] | 
 **account** | **string** | Comma delimited list of accounts to match | 
 **association** | **string** | Comma delimited list of associations to match | 
 **cluster** | **string** | Comma delimited list of cluster to match | 
 **constraints** | **string** | Comma delimited list of constraints to match | 
 **cpusMax** | **string** | Number of CPUs high range | 
 **cpusMin** | **string** | Number of CPUs low range | 
 **skipSteps** | **bool** | Report job step information | 
 **disableWaitForResult** | **bool** | Disable waiting for result from slurmdbd | 
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

[**Dbv0038JobInfo**](Dbv0038JobInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetQos

> Dbv0038QosInfo SlurmdbV0038GetQos(ctx).WithDeleted(withDeleted).Execute()

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
    withDeleted := true // bool | Include deleted QOSs. False by default. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetQos(context.Background()).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetQos`: Dbv0038QosInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withDeleted** | **bool** | Include deleted QOSs. False by default. | 

### Return type

[**Dbv0038QosInfo**](Dbv0038QosInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetSingleQos

> Dbv0038QosInfo SlurmdbV0038GetSingleQos(ctx, qosName).WithDeleted(withDeleted).Execute()

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
    withDeleted := true // bool | Include deleted QOSs. False by default. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetSingleQos(context.Background(), qosName).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetSingleQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetSingleQos`: Dbv0038QosInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qosName** | **string** | Slurm QOS Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **bool** | Include deleted QOSs. False by default. | 

### Return type

[**Dbv0038QosInfo**](Dbv0038QosInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetTres

> Dbv0038TresInfo SlurmdbV0038GetTres(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetTres(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetTres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetTres`: Dbv0038TresInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetTresRequest struct via the builder pattern


### Return type

[**Dbv0038TresInfo**](Dbv0038TresInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetUser

> Dbv0038UserInfo SlurmdbV0038GetUser(ctx, userName).WithDeleted(withDeleted).Execute()

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
    withDeleted := true // bool | Include deleted users. False by default. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetUser(context.Background(), userName).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetUser`: Dbv0038UserInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | Slurm User Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **bool** | Include deleted users. False by default. | 

### Return type

[**Dbv0038UserInfo**](Dbv0038UserInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetUsers

> Dbv0038UserInfo SlurmdbV0038GetUsers(ctx).WithDeleted(withDeleted).Execute()

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
    withDeleted := true // bool | Include deleted users. False by default. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetUsers(context.Background()).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetUsers`: Dbv0038UserInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withDeleted** | **bool** | Include deleted users. False by default. | 

### Return type

[**Dbv0038UserInfo**](Dbv0038UserInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetWckey

> Dbv0038WckeyInfo SlurmdbV0038GetWckey(ctx, wckey).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetWckey(context.Background(), wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetWckey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetWckey`: Dbv0038WckeyInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wckey** | **string** | Slurm wckey name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0038WckeyInfo**](Dbv0038WckeyInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038GetWckeys

> Dbv0038WckeyInfo SlurmdbV0038GetWckeys(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038GetWckeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038GetWckeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038GetWckeys`: Dbv0038WckeyInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038GetWckeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038GetWckeysRequest struct via the builder pattern


### Return type

[**Dbv0038WckeyInfo**](Dbv0038WckeyInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038SetConfig

> Dbv0038ConfigResponse SlurmdbV0038SetConfig(ctx).Dbv0038SetConfig(dbv0038SetConfig).Execute()

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
    dbv0038SetConfig := *openapiclient.NewDbv0038SetConfig() // Dbv0038SetConfig | Add or update config (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038SetConfig(context.Background()).Dbv0038SetConfig(dbv0038SetConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038SetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038SetConfig`: Dbv0038ConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038SetConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038SetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0038SetConfig** | [**Dbv0038SetConfig**](Dbv0038SetConfig.md) | Add or update config | 

### Return type

[**Dbv0038ConfigResponse**](Dbv0038ConfigResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038UpdateAccount

> Dbv0038AccountResponse SlurmdbV0038UpdateAccount(ctx).Dbv0038UpdateAccount(dbv0038UpdateAccount).Execute()

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
    dbv0038UpdateAccount := *openapiclient.NewDbv0038UpdateAccount() // Dbv0038UpdateAccount | update/create accounts

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038UpdateAccount(context.Background()).Dbv0038UpdateAccount(dbv0038UpdateAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038UpdateAccount`: Dbv0038AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038UpdateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038UpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0038UpdateAccount** | [**Dbv0038UpdateAccount**](Dbv0038UpdateAccount.md) | update/create accounts | 

### Return type

[**Dbv0038AccountResponse**](Dbv0038AccountResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038UpdateAssociations

> Dbv0038ResponseAssociations SlurmdbV0038UpdateAssociations(ctx).Dbv0038AssociationsInfo(dbv0038AssociationsInfo).Execute()

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
    dbv0038AssociationsInfo := *openapiclient.NewDbv0038AssociationsInfo() // Dbv0038AssociationsInfo | Add or update associations

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038UpdateAssociations(context.Background()).Dbv0038AssociationsInfo(dbv0038AssociationsInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038UpdateAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038UpdateAssociations`: Dbv0038ResponseAssociations
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038UpdateAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038UpdateAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0038AssociationsInfo** | [**Dbv0038AssociationsInfo**](Dbv0038AssociationsInfo.md) | Add or update associations | 

### Return type

[**Dbv0038ResponseAssociations**](Dbv0038ResponseAssociations.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038UpdateQos

> Dbv0038ResponseQos SlurmdbV0038UpdateQos(ctx).Dbv0038UpdateQos(dbv0038UpdateQos).Execute()

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
    dbv0038UpdateQos := *openapiclient.NewDbv0038UpdateQos() // Dbv0038UpdateQos | Add or update QOSs

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038UpdateQos(context.Background()).Dbv0038UpdateQos(dbv0038UpdateQos).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038UpdateQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038UpdateQos`: Dbv0038ResponseQos
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038UpdateQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038UpdateQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0038UpdateQos** | [**Dbv0038UpdateQos**](Dbv0038UpdateQos.md) | Add or update QOSs | 

### Return type

[**Dbv0038ResponseQos**](Dbv0038ResponseQos.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038UpdateTres

> Dbv0038ResponseTres SlurmdbV0038UpdateTres(ctx).Dbv0038TresUpdate(dbv0038TresUpdate).Execute()

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
    dbv0038TresUpdate := *openapiclient.NewDbv0038TresUpdate() // Dbv0038TresUpdate | Add or Update TRES

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038UpdateTres(context.Background()).Dbv0038TresUpdate(dbv0038TresUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038UpdateTres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038UpdateTres`: Dbv0038ResponseTres
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038UpdateTres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038UpdateTresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0038TresUpdate** | [**Dbv0038TresUpdate**](Dbv0038TresUpdate.md) | Add or Update TRES | 

### Return type

[**Dbv0038ResponseTres**](Dbv0038ResponseTres.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0038UpdateUsers

> Dbv0038ResponseUserUpdate SlurmdbV0038UpdateUsers(ctx).Dbv0038UpdateUsers(dbv0038UpdateUsers).Execute()

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
    dbv0038UpdateUsers := *openapiclient.NewDbv0038UpdateUsers() // Dbv0038UpdateUsers | add or update user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbV0038UpdateUsers(context.Background()).Dbv0038UpdateUsers(dbv0038UpdateUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbV0038UpdateUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbV0038UpdateUsers`: Dbv0038ResponseUserUpdate
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbV0038UpdateUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0038UpdateUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbv0038UpdateUsers** | [**Dbv0038UpdateUsers**](Dbv0038UpdateUsers.md) | add or update user | 

### Return type

[**Dbv0038ResponseUserUpdate**](Dbv0038ResponseUserUpdate.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

