package lib

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/rs/xid"
	"github.com/workflow-interoperability/manufacturer/types"

	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/worker"
)

// FailJob fail a job
func FailJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())
	client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send()
}

// BlockchainTransaction contact blockchain
func BlockchainTransaction(url, body string) error {
	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(resbody))
	}
	return nil
}

// GetPIIS return piis
func GetPIIS(url string) (types.PIIS, error) {
	var ret types.PIIS
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ret, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	resbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}
	if resp.StatusCode != 200 {
		return ret, errors.New(string(resbody))
	}
	err = json.Unmarshal(resbody, &ret)
	return ret, err
}

// GetIM return im
func GetIM(url string) (types.IM, error) {
	var ret types.IM
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ret, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	resbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}
	if resp.StatusCode != 200 {
		return ret, errors.New(string(resbody))
	}
	err = json.Unmarshal(resbody, &ret)
	return ret, err
}

// GenerateXID generate unique id
func GenerateXID() string {
	return xid.New().String()
}
