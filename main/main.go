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
	v0038JobSubmission(client)

	//v0038GetJobs(client)
	v0038GetJobID(client, "149")
	//v0038GetNodes(client)
	//v0038GetNodeName(client, "phy0023")
	//v0038GetPartitions(client)
	//v0038GetPartitionName(client, "compute")

	//v0038GetDBJobs(client)
	//v0038GetDBJobID(client, "89")
	//v0038GetDBAccounts(client)
	//v0038GetDBAccountName(client, "root")
	//v0038GetDBUsers(client)
	//v0038GetDBUserName(client, "wqtest")

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
	fmt.Println("===================================GET DB Job ID=====================================")
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

func v0038GetPartitionName(client *openapiclient.APIClient, name string) {
	fmt.Println("===================================GET DB Job ID=====================================")
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
	fmt.Println("===================================GET Partitions=====================================")
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
	fmt.Println("===================================GET Partitions=====================================")
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
	fmt.Println("===================================GET Partitions=====================================")
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
	fmt.Println("===================================GET Partitions=====================================")
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
