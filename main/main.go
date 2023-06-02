package main

import (
	"context"
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

	//v0038GetJobs(client)
	v0038GetDBJobs(client)

	//v0038JobSubmission(client)
}

func v0038JobSubmission(client *openapiclient.APIClient) {
	/*
		{
		    "name":"eeeee",
		    "user_name":"wq",
		    "partition":"compute",
		    "qos":"normal",
		    "node_count": 1,
		    "core_count": 1,
		    "gpu_count": 0,
		    "limit_time": 5,
		    "region_id":"test",
		    "command":"hostname",
		    "work_dir":"/root/wq",
		    "standard_output":"job.%j.out",
		    "error_output":"job.%j.err",
		    "description":"qqq"
		}
	*/

	v0038JobSubmission := *openapiclient.NewV0038JobSubmission("hostname") // V0038JobSubmission | submit new job
	//v0038JobSubmission.Job = &v0038JobProperties
	job := client.SlurmAPI.SlurmV0038SubmitJob(context.Background())
	job = job.V0038JobSubmission(v0038JobSubmission)
	resp, r, err := client.SlurmAPI.SlurmV0038SubmitJobExecute(job)
	//resp, r, err := job.V0038JobSubmission(v0038JobSubmission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038SubmitJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038SubmitJob`: V0038JobSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038SubmitJob`: %v\n", resp)
}

func v0038GetJobs(client *openapiclient.APIClient) {
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
		fmt.Printf("Job %d - %s\n", job.GetJobId(), job.GetJobState())
	}

	job := client.SlurmAPI.SlurmV0038SubmitJob(context.Background())
	client.SlurmAPI.SlurmV0038SubmitJobExecute(job)
}

func v0038GetDBJobs(client *openapiclient.APIClient) {
	jreq := client.SlurmAPI.SlurmdbV0038GetJobs(context.Background())
	jobs, resp, err := client.SlurmAPI.SlurmdbV0038GetJobsExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, job := range jobs.GetJobs() {
		//fmt.Println(job)
		fmt.Printf("Job %d - %v\n", job.GetJobId(), job)
	}

	job := client.SlurmAPI.SlurmV0038SubmitJob(context.Background())
	client.SlurmAPI.SlurmV0038SubmitJobExecute(job)
}
