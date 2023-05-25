package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	slurmrest "github.com/zhuqw12/slurmrestapi"
)

func main() {
	cfg := slurmrest.NewConfiguration()
	cfg.HTTPClient = &http.Client{Timeout: time.Second * 3600}
	cfg.Scheme = "http"
	cfg.Host = "172.32.1.24:6820"
	cfg.DefaultHeader = map[string]string{
		"X-SLURM-USER-NAME":  "root",
		"X-SLURM-USER-TOKEN": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjIwMDAwODQwODgsImlhdCI6MTY4NDcyNDA4OCwic3VuIjoicm9vdCJ9.3NoFEsowIPSvbjTXDJlN_fEIcIKh8VRGwOHb5LHd6o4"}

	client := slurmrest.NewAPIClient(cfg)

	jreq := client.SlurmAPI.SlurmV0039GetJobs(context.Background())
	jobs, resp, err := client.SlurmAPI.SlurmV0039GetJobsExecute(jreq)
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	//fmt.Println(jobs)
	for _, job := range jobs.GetJobs() {
		//fmt.Println(job)
		fmt.Printf("Job %s - %s\n", job.GetJobId(), job.GetJobState())
	}
}
