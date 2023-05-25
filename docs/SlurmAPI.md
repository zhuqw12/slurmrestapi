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
[**SlurmctldCancelJob**](SlurmAPI.md#SlurmctldCancelJob) | **Delete** /slurm/v0.0.37/job/{job_id} | cancel or signal job
[**SlurmctldDiag**](SlurmAPI.md#SlurmctldDiag) | **Get** /slurm/v0.0.37/diag | get diagnostics
[**SlurmctldGetJob**](SlurmAPI.md#SlurmctldGetJob) | **Get** /slurm/v0.0.37/job/{job_id} | get job info
[**SlurmctldGetJobs**](SlurmAPI.md#SlurmctldGetJobs) | **Get** /slurm/v0.0.37/jobs | get list of jobs
[**SlurmctldGetNode**](SlurmAPI.md#SlurmctldGetNode) | **Get** /slurm/v0.0.37/node/{node_name} | get node info
[**SlurmctldGetNodes**](SlurmAPI.md#SlurmctldGetNodes) | **Get** /slurm/v0.0.37/nodes | get all node info
[**SlurmctldGetPartition**](SlurmAPI.md#SlurmctldGetPartition) | **Get** /slurm/v0.0.37/partition/{partition_name} | get partition info
[**SlurmctldGetPartitions**](SlurmAPI.md#SlurmctldGetPartitions) | **Get** /slurm/v0.0.37/partitions | get all partition info
[**SlurmctldGetReservation**](SlurmAPI.md#SlurmctldGetReservation) | **Get** /slurm/v0.0.37/reservation/{reservation_name} | get reservation info
[**SlurmctldGetReservations**](SlurmAPI.md#SlurmctldGetReservations) | **Get** /slurm/v0.0.37/reservations | get all reservation info
[**SlurmctldPing**](SlurmAPI.md#SlurmctldPing) | **Get** /slurm/v0.0.37/ping | ping test
[**SlurmctldSubmitJob**](SlurmAPI.md#SlurmctldSubmitJob) | **Post** /slurm/v0.0.37/job/submit | submit new job
[**SlurmctldUpdateJob**](SlurmAPI.md#SlurmctldUpdateJob) | **Post** /slurm/v0.0.37/job/{job_id} | update job
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
[**SlurmdbdAddClusters**](SlurmAPI.md#SlurmdbdAddClusters) | **Post** /slurmdb/v0.0.37/clusters | Add clusters
[**SlurmdbdAddWckeys**](SlurmAPI.md#SlurmdbdAddWckeys) | **Post** /slurmdb/v0.0.37/wckeys | Add wckeys
[**SlurmdbdDeleteAccount**](SlurmAPI.md#SlurmdbdDeleteAccount) | **Delete** /slurmdb/v0.0.37/account/{account_name} | Delete account
[**SlurmdbdDeleteAssociation**](SlurmAPI.md#SlurmdbdDeleteAssociation) | **Delete** /slurmdb/v0.0.37/association | Delete association
[**SlurmdbdDeleteCluster**](SlurmAPI.md#SlurmdbdDeleteCluster) | **Delete** /slurmdb/v0.0.37/cluster/{cluster_name} | Delete cluster
[**SlurmdbdDeleteQos**](SlurmAPI.md#SlurmdbdDeleteQos) | **Delete** /slurmdb/v0.0.37/qos/{qos_name} | Delete QOS
[**SlurmdbdDeleteUser**](SlurmAPI.md#SlurmdbdDeleteUser) | **Delete** /slurmdb/v0.0.37/user/{user_name} | Delete user
[**SlurmdbdDeleteWckey**](SlurmAPI.md#SlurmdbdDeleteWckey) | **Delete** /slurmdb/v0.0.37/wckey/{wckey} | Delete wckey
[**SlurmdbdDiag**](SlurmAPI.md#SlurmdbdDiag) | **Get** /slurmdb/v0.0.37/diag | Get slurmdb diagnostics
[**SlurmdbdGetAccount**](SlurmAPI.md#SlurmdbdGetAccount) | **Get** /slurmdb/v0.0.37/account/{account_name} | Get account info
[**SlurmdbdGetAccounts**](SlurmAPI.md#SlurmdbdGetAccounts) | **Get** /slurmdb/v0.0.37/accounts | Get account list
[**SlurmdbdGetAssociation**](SlurmAPI.md#SlurmdbdGetAssociation) | **Get** /slurmdb/v0.0.37/association | Get association info
[**SlurmdbdGetAssociations**](SlurmAPI.md#SlurmdbdGetAssociations) | **Get** /slurmdb/v0.0.37/associations | Get association list
[**SlurmdbdGetCluster**](SlurmAPI.md#SlurmdbdGetCluster) | **Get** /slurmdb/v0.0.37/cluster/{cluster_name} | Get cluster info
[**SlurmdbdGetClusters**](SlurmAPI.md#SlurmdbdGetClusters) | **Get** /slurmdb/v0.0.37/clusters | Get cluster list
[**SlurmdbdGetDbConfig**](SlurmAPI.md#SlurmdbdGetDbConfig) | **Get** /slurmdb/v0.0.37/config | Dump all configuration information
[**SlurmdbdGetJob**](SlurmAPI.md#SlurmdbdGetJob) | **Get** /slurmdb/v0.0.37/job/{job_id} | Get job info
[**SlurmdbdGetJobs**](SlurmAPI.md#SlurmdbdGetJobs) | **Get** /slurmdb/v0.0.37/jobs | Get job list
[**SlurmdbdGetQos**](SlurmAPI.md#SlurmdbdGetQos) | **Get** /slurmdb/v0.0.37/qos | Get QOS list
[**SlurmdbdGetSingleQos**](SlurmAPI.md#SlurmdbdGetSingleQos) | **Get** /slurmdb/v0.0.37/qos/{qos_name} | Get QOS info
[**SlurmdbdGetTres**](SlurmAPI.md#SlurmdbdGetTres) | **Get** /slurmdb/v0.0.37/tres | Get TRES info
[**SlurmdbdGetUser**](SlurmAPI.md#SlurmdbdGetUser) | **Get** /slurmdb/v0.0.37/user/{user_name} | Get user info
[**SlurmdbdGetUsers**](SlurmAPI.md#SlurmdbdGetUsers) | **Get** /slurmdb/v0.0.37/users | Get user list
[**SlurmdbdGetWckey**](SlurmAPI.md#SlurmdbdGetWckey) | **Get** /slurmdb/v0.0.37/wckey/{wckey} | Get wckey info
[**SlurmdbdGetWckeys**](SlurmAPI.md#SlurmdbdGetWckeys) | **Get** /slurmdb/v0.0.37/wckeys | Get wckey list
[**SlurmdbdSetDbConfig**](SlurmAPI.md#SlurmdbdSetDbConfig) | **Post** /slurmdb/v0.0.37/config | Load all configuration information
[**SlurmdbdUpdateAccount**](SlurmAPI.md#SlurmdbdUpdateAccount) | **Post** /slurmdb/v0.0.37/accounts | Update accounts
[**SlurmdbdUpdateAssociations**](SlurmAPI.md#SlurmdbdUpdateAssociations) | **Post** /slurmdb/v0.0.37/associations | Set associations info
[**SlurmdbdUpdateTres**](SlurmAPI.md#SlurmdbdUpdateTres) | **Post** /slurmdb/v0.0.37/tres | Set TRES info
[**SlurmdbdUpdateUsers**](SlurmAPI.md#SlurmdbdUpdateUsers) | **Post** /slurmdb/v0.0.37/users | Update user



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


## SlurmctldCancelJob

> SlurmctldCancelJob(ctx, jobId).Signal(signal).Execute()

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
    signal := openapiclient.v0.0.37_signal("HUP") // V0037Signal | signal to send to job (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlurmAPI.SlurmctldCancelJob(context.Background(), jobId).Signal(signal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldCancelJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSlurmctldCancelJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | [**V0037Signal**](V0037Signal.md) | signal to send to job | 

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


## SlurmctldDiag

> V0037Diag SlurmctldDiag(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmctldDiag(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldDiag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldDiag`: V0037Diag
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldDiagRequest struct via the builder pattern


### Return type

[**V0037Diag**](V0037Diag.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetJob

> V0037JobsResponse SlurmctldGetJob(ctx, jobId).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetJob`: V0037JobsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm JobID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0037JobsResponse**](V0037JobsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetJobs

> V0037JobsResponse SlurmctldGetJobs(ctx).UpdateTime(updateTime).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetJobs(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetJobs`: V0037JobsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0037JobsResponse**](V0037JobsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetNode

> V0037NodesResponse SlurmctldGetNode(ctx, nodeName).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetNode(context.Background(), nodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetNode`: V0037NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Slurm Node Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0037NodesResponse**](V0037NodesResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetNodes

> V0037NodesResponse SlurmctldGetNodes(ctx).UpdateTime(updateTime).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetNodes(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetNodes`: V0037NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0037NodesResponse**](V0037NodesResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetPartition

> V0037PartitionsResponse SlurmctldGetPartition(ctx, partitionName).UpdateTime(updateTime).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetPartition(context.Background(), partitionName).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetPartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetPartition`: V0037PartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Slurm Partition Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. | 

### Return type

[**V0037PartitionsResponse**](V0037PartitionsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetPartitions

> V0037PartitionsResponse SlurmctldGetPartitions(ctx).UpdateTime(updateTime).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetPartitions(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetPartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetPartitions`: V0037PartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0037PartitionsResponse**](V0037PartitionsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetReservation

> V0037ReservationsResponse SlurmctldGetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetReservation`: V0037ReservationsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Slurm Reservation Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if no reservation (not limited to reservation in URL) changed since update_time. | 

### Return type

[**V0037ReservationsResponse**](V0037ReservationsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetReservations

> V0037ReservationsResponse SlurmctldGetReservations(ctx).UpdateTime(updateTime).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetReservations(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetReservations`: V0037ReservationsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0037ReservationsResponse**](V0037ReservationsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldPing

> V0037Pings SlurmctldPing(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmctldPing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldPing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldPing`: V0037Pings
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldPingRequest struct via the builder pattern


### Return type

[**V0037Pings**](V0037Pings.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldSubmitJob

> V0037JobSubmissionResponse SlurmctldSubmitJob(ctx).V0037JobSubmission(v0037JobSubmission).Execute()

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
    v0037JobSubmission := *openapiclient.NewV0037JobSubmission("Script_example") // V0037JobSubmission | submit new job

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmctldSubmitJob(context.Background()).V0037JobSubmission(v0037JobSubmission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldSubmitJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldSubmitJob`: V0037JobSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldSubmitJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldSubmitJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0037JobSubmission** | [**V0037JobSubmission**](V0037JobSubmission.md) | submit new job | 

### Return type

[**V0037JobSubmissionResponse**](V0037JobSubmissionResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldUpdateJob

> SlurmctldUpdateJob(ctx, jobId).V0037JobProperties(v0037JobProperties).Execute()

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
    v0037JobProperties := *openapiclient.NewV0037JobProperties(map[string]interface{}(123)) // V0037JobProperties | update job

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlurmAPI.SlurmctldUpdateJob(context.Background(), jobId).V0037JobProperties(v0037JobProperties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldUpdateJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSlurmctldUpdateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0037JobProperties** | [**V0037JobProperties**](V0037JobProperties.md) | update job | 

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


## SlurmdbdAddClusters

> Dbv0037ResponseClusterAdd SlurmdbdAddClusters(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdAddClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdAddClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdAddClusters`: Dbv0037ResponseClusterAdd
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdAddClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdAddClustersRequest struct via the builder pattern


### Return type

[**Dbv0037ResponseClusterAdd**](Dbv0037ResponseClusterAdd.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdAddWckeys

> Dbv0037ResponseWckeyAdd SlurmdbdAddWckeys(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdAddWckeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdAddWckeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdAddWckeys`: Dbv0037ResponseWckeyAdd
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdAddWckeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdAddWckeysRequest struct via the builder pattern


### Return type

[**Dbv0037ResponseWckeyAdd**](Dbv0037ResponseWckeyAdd.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDeleteAccount

> Dbv0037ResponseAccountDelete SlurmdbdDeleteAccount(ctx, accountName).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdDeleteAccount(context.Background(), accountName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdDeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteAccount`: Dbv0037ResponseAccountDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdDeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Slurm Account Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037ResponseAccountDelete**](Dbv0037ResponseAccountDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDeleteAssociation

> Dbv0037ResponseAssociationDelete SlurmdbdDeleteAssociation(ctx).Account(account).User(user).Cluster(cluster).Partition(partition).Execute()

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
    account := "account_example" // string | Account name
    user := "user_example" // string | User name
    cluster := "cluster_example" // string | Cluster name (optional)
    partition := "partition_example" // string | Partition Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdDeleteAssociation(context.Background()).Account(account).User(user).Cluster(cluster).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdDeleteAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteAssociation`: Dbv0037ResponseAssociationDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdDeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | Account name | 
 **user** | **string** | User name | 
 **cluster** | **string** | Cluster name | 
 **partition** | **string** | Partition Name | 

### Return type

[**Dbv0037ResponseAssociationDelete**](Dbv0037ResponseAssociationDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDeleteCluster

> Dbv0037ResponseClusterDelete SlurmdbdDeleteCluster(ctx, clusterName).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdDeleteCluster(context.Background(), clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdDeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteCluster`: Dbv0037ResponseClusterDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Slurm cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037ResponseClusterDelete**](Dbv0037ResponseClusterDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDeleteQos

> Dbv0037ResponseQosDelete SlurmdbdDeleteQos(ctx, qosName).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdDeleteQos(context.Background(), qosName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdDeleteQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteQos`: Dbv0037ResponseQosDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdDeleteQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qosName** | **string** | Slurm QOS Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037ResponseQosDelete**](Dbv0037ResponseQosDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDeleteUser

> Dbv0037ResponseUserDelete SlurmdbdDeleteUser(ctx, userName).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdDeleteUser(context.Background(), userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdDeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteUser`: Dbv0037ResponseUserDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdDeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | Slurm User Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037ResponseUserDelete**](Dbv0037ResponseUserDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDeleteWckey

> Dbv0037ResponseWckeyDelete SlurmdbdDeleteWckey(ctx, wckey).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdDeleteWckey(context.Background(), wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdDeleteWckey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteWckey`: Dbv0037ResponseWckeyDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdDeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wckey** | **string** | Slurm wckey name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037ResponseWckeyDelete**](Dbv0037ResponseWckeyDelete.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDiag

> Dbv0037Diag SlurmdbdDiag(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdDiag(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdDiag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDiag`: Dbv0037Diag
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDiagRequest struct via the builder pattern


### Return type

[**Dbv0037Diag**](Dbv0037Diag.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetAccount

> Dbv0037AccountInfo SlurmdbdGetAccount(ctx, accountName).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetAccount(context.Background(), accountName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetAccount`: Dbv0037AccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Slurm Account Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037AccountInfo**](Dbv0037AccountInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetAccounts

> Dbv0037AccountInfo SlurmdbdGetAccounts(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetAccounts`: Dbv0037AccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetAccountsRequest struct via the builder pattern


### Return type

[**Dbv0037AccountInfo**](Dbv0037AccountInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetAssociation

> Dbv0037AssociationsInfo SlurmdbdGetAssociation(ctx).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetAssociation(context.Background()).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetAssociation`: Dbv0037AssociationsInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Cluster name | 
 **account** | **string** | Account name | 
 **user** | **string** | User name | 
 **partition** | **string** | Partition Name | 

### Return type

[**Dbv0037AssociationsInfo**](Dbv0037AssociationsInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetAssociations

> Dbv0037AssociationsInfo SlurmdbdGetAssociations(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetAssociations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetAssociations`: Dbv0037AssociationsInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetAssociations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetAssociationsRequest struct via the builder pattern


### Return type

[**Dbv0037AssociationsInfo**](Dbv0037AssociationsInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetCluster

> Dbv0037ClusterInfo SlurmdbdGetCluster(ctx, clusterName).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetCluster(context.Background(), clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetCluster`: Dbv0037ClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Slurm cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037ClusterInfo**](Dbv0037ClusterInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetClusters

> Dbv0037ClusterInfo SlurmdbdGetClusters(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetClusters`: Dbv0037ClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetClustersRequest struct via the builder pattern


### Return type

[**Dbv0037ClusterInfo**](Dbv0037ClusterInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetDbConfig

> Dbv0037ConfigInfo SlurmdbdGetDbConfig(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetDbConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetDbConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetDbConfig`: Dbv0037ConfigInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetDbConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetDbConfigRequest struct via the builder pattern


### Return type

[**Dbv0037ConfigInfo**](Dbv0037ConfigInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetJob

> Dbv0037JobInfo SlurmdbdGetJob(ctx, jobId).Execute()

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
    jobId := int64(789) // int64 | Slurm Job ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetJob`: Dbv0037JobInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **int64** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037JobInfo**](Dbv0037JobInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetJobs

> Dbv0037JobInfo SlurmdbdGetJobs(ctx).SubmitTime(submitTime).StartTime(startTime).EndTime(endTime).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SkipSteps(skipSteps).DisableWaitForResult(disableWaitForResult).ExitCode(exitCode).Format(format).Group(group).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).State(state).Step(step).Node(node).Wckey(wckey).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetJobs(context.Background()).SubmitTime(submitTime).StartTime(startTime).EndTime(endTime).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SkipSteps(skipSteps).DisableWaitForResult(disableWaitForResult).ExitCode(exitCode).Format(format).Group(group).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).State(state).Step(step).Node(node).Wckey(wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetJobs`: Dbv0037JobInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetJobsRequest struct via the builder pattern


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

[**Dbv0037JobInfo**](Dbv0037JobInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetQos

> Dbv0037QosInfo SlurmdbdGetQos(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetQos(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetQos`: Dbv0037QosInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetQos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetQosRequest struct via the builder pattern


### Return type

[**Dbv0037QosInfo**](Dbv0037QosInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetSingleQos

> Dbv0037QosInfo SlurmdbdGetSingleQos(ctx, qosName).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetSingleQos(context.Background(), qosName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetSingleQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetSingleQos`: Dbv0037QosInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qosName** | **string** | Slurm QOS Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037QosInfo**](Dbv0037QosInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetTres

> Dbv0037TresInfo SlurmdbdGetTres(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetTres(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetTres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetTres`: Dbv0037TresInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetTresRequest struct via the builder pattern


### Return type

[**Dbv0037TresInfo**](Dbv0037TresInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetUser

> Dbv0037UserInfo SlurmdbdGetUser(ctx, userName).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetUser(context.Background(), userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetUser`: Dbv0037UserInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | Slurm User Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037UserInfo**](Dbv0037UserInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetUsers

> Dbv0037UserInfo SlurmdbdGetUsers(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetUsers`: Dbv0037UserInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetUsersRequest struct via the builder pattern


### Return type

[**Dbv0037UserInfo**](Dbv0037UserInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetWckey

> Dbv0037WckeyInfo SlurmdbdGetWckey(ctx, wckey).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetWckey(context.Background(), wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetWckey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetWckey`: Dbv0037WckeyInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wckey** | **string** | Slurm wckey name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037WckeyInfo**](Dbv0037WckeyInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetWckeys

> Dbv0037WckeyInfo SlurmdbdGetWckeys(ctx).Execute()

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
    resp, r, err := apiClient.SlurmAPI.SlurmdbdGetWckeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdGetWckeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetWckeys`: Dbv0037WckeyInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdGetWckeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetWckeysRequest struct via the builder pattern


### Return type

[**Dbv0037WckeyInfo**](Dbv0037WckeyInfo.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdSetDbConfig

> Dbv0037ConfigResponse SlurmdbdSetDbConfig(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdSetDbConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdSetDbConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdSetDbConfig`: Dbv0037ConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdSetDbConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdSetDbConfigRequest struct via the builder pattern


### Return type

[**Dbv0037ConfigResponse**](Dbv0037ConfigResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdUpdateAccount

> Dbv0037AccountResponse SlurmdbdUpdateAccount(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdUpdateAccount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdUpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdUpdateAccount`: Dbv0037AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdUpdateAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdUpdateAccountRequest struct via the builder pattern


### Return type

[**Dbv0037AccountResponse**](Dbv0037AccountResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdUpdateAssociations

> Dbv0037ResponseAssociations SlurmdbdUpdateAssociations(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdUpdateAssociations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdUpdateAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdUpdateAssociations`: Dbv0037ResponseAssociations
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdUpdateAssociations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdUpdateAssociationsRequest struct via the builder pattern


### Return type

[**Dbv0037ResponseAssociations**](Dbv0037ResponseAssociations.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdUpdateTres

> Dbv0037ResponseTres SlurmdbdUpdateTres(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdUpdateTres(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdUpdateTres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdUpdateTres`: Dbv0037ResponseTres
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdUpdateTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdUpdateTresRequest struct via the builder pattern


### Return type

[**Dbv0037ResponseTres**](Dbv0037ResponseTres.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdUpdateUsers

> Dbv0037ResponseUserUpdate SlurmdbdUpdateUsers(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmdbdUpdateUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmdbdUpdateUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdUpdateUsers`: Dbv0037ResponseUserUpdate
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmdbdUpdateUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdUpdateUsersRequest struct via the builder pattern


### Return type

[**Dbv0037ResponseUserUpdate**](Dbv0037ResponseUserUpdate.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

