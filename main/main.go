package main

import (
	"context"
	"encoding/json"
	"fmt"
	openapiclient "github.com/zhuqw12/slurmrestapi"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	cfg := openapiclient.NewConfiguration()
	cfg.HTTPClient = &http.Client{Timeout: time.Second * 3600}
	cfg.Scheme = "http"
	cfg.Host = "172.32.1.24:6820"
	cfg.DefaultHeader = map[string]string{
		"X-SLURM-USER-NAME":  "root",
		"X-SLURM-USER-TOKEN": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjIwMDAwODQwODgsImlhdCI6MTY4NDcyNDA4OCwic3VuIjoicm9vdCJ9.3NoFEsowIPSvbjTXDJlN_fEIcIKh8VRGwOHb5LHd6o4"}

	client := openapiclient.NewAPIClient(cfg)
	//v0038JobSubmission(client)

	v0038GetJobs(client)
	v0038GetJobID(client, "89")
	v0038GetNodes(client)
	v0038GetNodeName(client, "phy0023")
	v0038GetPartitions(client)
	v0038GetPartitionName(client, "compute")
	v0038GetReservations(client)

	v0038GetDBAccounts(client)
	v0038GetDBAccountName(client, "root")
	v0038GetDBAssociations(client)
	v0038GetDBClusters(client)
	v0038GetDBClusterByName(client, "zcycluster")
	v0038GetDBConfig(client)
	v0038GetDBJobs(client)
	v0038GetDBJobID(client, "89")
	v0038GetDBQos(client)
	v0038GetDBQosByName(client, "normal")
	v0038GetDBTres(client)
	v0038GetDBUsers(client)
	v0038GetDBUserName(client, "wqtest")
	v0038GetDBWckeys(client)
}

func v0038JobSubmission(client *openapiclient.APIClient) {
	fmt.Println("===================================v0038JobSubmission Jobs=====================================")
	v0038JobProperties := openapiclient.NewV0038JobProperties(map[string]interface{}{"USER": "u000000"})
	v0038JobProperties.SetQos("normal")
	v0038JobProperties.SetTimeLimit(5)
	v0038JobProperties.SetAccount("root")
	v0038JobProperties.SetNodes([]int32{2, 2})
	v0038JobSubmission := *openapiclient.NewV0038JobSubmission("\"echo 'Hello, world'\"") // V0038JobSubmission | submit new job
	v0038JobSubmission.Job = v0038JobProperties
	job := client.SlurmAPI.SlurmV0038SubmitJob(context.Background())
	job = job.V0038JobSubmission(v0038JobSubmission)
	//resp, r, err := client.SlurmAPI.SlurmV0038SubmitJobExecute(job)
	resp, r, err := job.V0038JobSubmission(v0038JobSubmission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038SubmitJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038SubmitJob`: V0038JobSubmissionResponse
	fmt.Printf("Job %d - %s\n", resp.GetJobId(), resp.GetJobSubmitUserMsg())
}

func v0038GetJobs(client *openapiclient.APIClient) {
	fmt.Println("===================================GET Jobs=====================================")
	jreq := client.SlurmAPI.SlurmV0038GetJobs(context.Background())
	jobs, resp, err := client.SlurmAPI.SlurmV0038GetJobsExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, job := range jobs.GetJobs() {
		//fmt.Println(job)
		marshal, err := json.Marshal(job)
		if err == nil {
			fmt.Printf("Job %d - %s [%s]\n", job.GetJobId(), job.GetJobState(), marshal)
		}
	}
}

func v0038GetJobID(client *openapiclient.APIClient, id string) {
	fmt.Println("===================================GET Job ID=====================================")
	jreq := client.SlurmAPI.SlurmV0038GetJob(context.Background(), id)
	jobs, resp, err := client.SlurmAPI.SlurmV0038GetJobExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, job := range jobs.GetJobs() {
		//fmt.Println(job)
		marshal, err := json.Marshal(job)
		if err == nil {
			fmt.Printf("Job %d - %s [%s]\n", job.GetJobId(), job.GetJobState(), marshal)
		}
	}
}

func v0038GetDBJobs(client *openapiclient.APIClient) {
	fmt.Println("===================================GET DB Jobs=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetJobs(context.Background())
	jreq = jreq.StartTime("2023-06-01").EndTime("2023-06-02")

	jobs, resp, err := client.SlurmAPI.SlurmdbV0038GetJobsExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, job := range jobs.GetJobs() {
		//fmt.Println(job)
		marshal, err := json.Marshal(job)
		if err == nil {
			fmt.Printf("Job %d - %s\n", job.GetJobId(), marshal)
		}
	}
}

func v0038GetDBJobID(client *openapiclient.APIClient, id string) {
	fmt.Println("===================================GET DB Job ID=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetJob(context.Background(), id)
	jobs, resp, err := client.SlurmAPI.SlurmdbV0038GetJobExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, job := range jobs.GetJobs() {
		//fmt.Println(job)
		marshal, err := json.Marshal(job)
		if err == nil {
			fmt.Printf("Job %d - %s\n", job.GetJobId(), marshal)
		}
	}
}

func v0038GetNodes(client *openapiclient.APIClient) {
	fmt.Println("===================================GET Nodes=====================================")
	jreq := client.SlurmAPI.SlurmV0038GetNodes(context.Background())
	nodes, resp, err := client.SlurmAPI.SlurmV0038GetNodesExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, node := range nodes.GetNodes() {
		//fmt.Println(job)
		marshal, err := json.Marshal(node)
		if err == nil {
			fmt.Printf("Node %s [%s]\n", node.GetName(), marshal)
		}
	}
}

func v0038GetNodeName(client *openapiclient.APIClient, name string) {
	fmt.Println("===================================GET Node Name=====================================")
	jreq := client.SlurmAPI.SlurmV0038GetNode(context.Background(), name)
	nodes, resp, err := client.SlurmAPI.SlurmV0038GetNodeExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	for _, node := range nodes.GetNodes() {
		//fmt.Println(job)
		marshal, err := json.Marshal(node)
		if err == nil {
			fmt.Printf("Node %s [%s]\n", node.GetName(), marshal)
		}
	}
}

func v0038GetPartitions(client *openapiclient.APIClient) {
	fmt.Println("===================================GET Partitions=====================================")
	jreq := client.SlurmAPI.SlurmV0038GetPartitions(context.Background())
	partitions, resp, err := client.SlurmAPI.SlurmV0038GetPartitionsExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, partition := range partitions.GetPartitions() {
		//fmt.Println(job)
		marshal, err := json.Marshal(partition)
		if err == nil {
			fmt.Printf("Node %s [%s]\n", partition.GetName(), marshal)
		}
	}
}
func v0038GetReservations(client *openapiclient.APIClient) {
	fmt.Println("===================================GET Reservations=====================================")
	jreq := client.SlurmAPI.SlurmV0038GetReservations(context.Background())
	reservations, resp, err := client.SlurmAPI.SlurmV0038GetReservationsExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, reservation := range reservations.GetReservations() {
		//fmt.Println(job)
		marshal, err := json.Marshal(reservation)
		if err == nil {
			fmt.Printf("Node %s [%s]\n", reservation.GetName(), marshal)
		}
	}
}

func v0038GetPartitionName(client *openapiclient.APIClient, name string) {
	fmt.Println("===================================GET Partition Name=====================================")
	jreq := client.SlurmAPI.SlurmV0038GetPartition(context.Background(), name)
	partitions, resp, err := client.SlurmAPI.SlurmV0038GetPartitionExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, partition := range partitions.GetPartitions() {
		//fmt.Println(job)
		marshal, err := json.Marshal(partition)
		if err == nil {
			fmt.Printf("Node %s [%s]\n", partition.GetName(), marshal)
		}
	}
}

func v0038GetDBAccounts(client *openapiclient.APIClient) {
	fmt.Println("===================================GET DBAccounts=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetAccounts(context.Background())
	accounts, resp, err := client.SlurmAPI.SlurmdbV0038GetAccountsExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, account := range accounts.GetAccounts() {
		//fmt.Println(job)
		marshal, err := json.Marshal(account)
		if err == nil {
			fmt.Printf("Account %s [%s]\n", account.GetName(), marshal)
		}
	}
}

func v0038GetDBAccountName(client *openapiclient.APIClient, name string) {
	fmt.Println("===================================GET DB AccountName=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetAccount(context.Background(), name)
	accounts, resp, err := client.SlurmAPI.SlurmdbV0038GetAccountExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, account := range accounts.GetAccounts() {
		//fmt.Println(job)
		marshal, err := json.Marshal(account)
		if err == nil {
			fmt.Printf("Account %s [%s]\n", account.GetName(), marshal)
		}
	}
}

func v0038GetDBUsers(client *openapiclient.APIClient) {
	fmt.Println("===================================GET DBUsers=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetUsers(context.Background())
	users, resp, err := client.SlurmAPI.SlurmdbV0038GetUsersExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, user := range users.GetUsers() {
		//fmt.Println(job)
		marshal, err := json.Marshal(user)
		if err == nil {
			fmt.Printf("user %s [%s]\n", user.GetName(), marshal)
		}
	}
}

func v0038GetDBUserName(client *openapiclient.APIClient, name string) {
	fmt.Println("===================================GET DB UserName=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetUser(context.Background(), name)
	users, resp, err := client.SlurmAPI.SlurmdbV0038GetUserExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, user := range users.GetUsers() {
		//fmt.Println(job)
		marshal, err := json.Marshal(user)
		if err == nil {
			fmt.Printf("user %s [%s]\n", user.GetName(), marshal)
		}
	}
}

func v0038GetDBQos(client *openapiclient.APIClient) {
	fmt.Println("===================================GET Qos=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetQos(context.Background())
	qoss, resp, err := client.SlurmAPI.SlurmdbV0038GetQosExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, qos := range qoss.GetQos() {
		//fmt.Println(job)
		marshal, err := json.Marshal(qos)
		if err == nil {
			fmt.Printf("qos %s [%s]\n", qos.GetName(), marshal)
		}
	}
}

func v0038GetDBQosByName(client *openapiclient.APIClient, name string) {
	fmt.Println("===================================GET Qos Name=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetSingleQos(context.Background(), name)
	qoss, resp, err := client.SlurmAPI.SlurmdbV0038GetSingleQosExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, qos := range qoss.GetQos() {
		//fmt.Println(job)
		marshal, err := json.Marshal(qos)
		if err == nil {
			fmt.Printf("qos %s [%s]\n", qos.GetName(), marshal)
		}
	}
}

func v0038GetDBClusters(client *openapiclient.APIClient) {
	fmt.Println("===================================GET DB Clusters=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetClusters(context.Background())
	clusters, resp, err := client.SlurmAPI.SlurmdbV0038GetClustersExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, cluster := range clusters.GetClusters() {
		//fmt.Println(job)
		marshal, err := json.Marshal(cluster)
		if err == nil {
			fmt.Printf("cluster - %s [%s]\n", cluster.GetName(), marshal)
		}
	}
}
func v0038GetDBAssociations(client *openapiclient.APIClient) {
	fmt.Println("===================================GET DB Associations=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetAssociations(context.Background())
	associations, resp, err := client.SlurmAPI.SlurmdbV0038GetAssociationsExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, association := range associations.GetAssociations() {
		//fmt.Println(job)
		marshal, err := json.Marshal(association)
		if err == nil {
			fmt.Printf("associations - %s [%s]\n", association.GetAccount(), marshal)
		}
	}
}

func v0038GetDBTres(client *openapiclient.APIClient) {
	fmt.Println("===================================GET DB Tres=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetTres(context.Background())
	tres, resp, err := client.SlurmAPI.SlurmdbV0038GetTresExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, tre := range tres.GetTres() {
		//fmt.Println(job)
		marshal, err := json.Marshal(tre)
		if err == nil {
			fmt.Printf("associations - %d [%s]\n", tre.GetId(), marshal)
		}
	}
}
func v0038GetDBWckeys(client *openapiclient.APIClient) {
	fmt.Println("===================================GET DB Wckeys=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetWckeys(context.Background())
	wckeys, resp, err := client.SlurmAPI.SlurmdbV0038GetWckeysExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, wckey := range wckeys.GetWckeys() {
		//fmt.Println(job)
		marshal, err := json.Marshal(wckey)
		if err == nil {
			fmt.Printf("wckey - %d [%s]\n", wckey.GetId(), marshal)
		}
	}
}

func v0038GetDBClusterByName(client *openapiclient.APIClient, name string) {
	fmt.Println("===================================GET DB Cluster Name=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetCluster(context.Background(), name)
	clusters, resp, err := client.SlurmAPI.SlurmdbV0038GetClusterExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, cluster := range clusters.GetClusters() {
		//fmt.Println(job)
		marshal, err := json.Marshal(cluster)
		if err == nil {
			fmt.Printf("cluster - %s [%s]\n", cluster.GetName(), marshal)
		}
	}
}
func v0038GetDBConfig(client *openapiclient.APIClient) {
	fmt.Println("===================================GET DB Config=====================================")
	jreq := client.SlurmAPI.SlurmdbV0038GetConfig(context.Background())
	config, resp, err := client.SlurmAPI.SlurmdbV0038GetConfigExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	marshal, err := json.Marshal(config)
	if err == nil {
		fmt.Printf("config [%s]\n", marshal)
	}
}
