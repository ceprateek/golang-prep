package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type podDetails struct {
	Name    string                 `json:"name"`
	Service string                 `json:"service"`
	State   map[string]interface{} `json:"state"`
}

func run() {
	for {
		result, err := getPodsStatus("")
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		for service, containersStatus := range result {
			for podname, podstatus := range containersStatus {
				log.Println(fmt.Sprintf("service:%s pod-name:%s status:%v", service, podname, podstatus))
			}
		}
		time.Sleep(3 * time.Second)
		log.Print("\n\n\n\n")
	}

}

func getPodsStatus(namespace string) (map[string]map[string]string, error) {
	cmd := "kubectl get pods -n default -o json | /usr/bin/jq '[.items[] | {name:.metadata.name,service:.status.containerStatuses[].name,state:.status.containerStatuses[].state}]'"
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Print(string(out))
	}

	var pods []podDetails
	err = json.Unmarshal(out, &pods)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := make(map[string]map[string]string)
	for _, pod := range pods {
		if _, ok := result[pod.Service]; !ok {
			containersStatus := make(map[string]string)
			result[pod.Service] = containersStatus
		}
		for status := range pod.State {
			containersStatus := result[pod.Service]
			containersStatus[pod.Name] = status
		}
	}

	return result, nil
}

func createBackupTimeCSV(){
	file, err := os.Open("/Users/prategar/Downloads/backup_times.csv")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var totalTime time.Duration
	for scanner.Scan(){
		line := scanner.Text()
		if len(line)>30{
			cols := strings.Split(line, ",")
			tend,err := time.Parse(time.RFC3339, cols[2])
			if err != nil{
				log.Fatal(err)
			}
			tstart,err := time.Parse(time.RFC3339, cols[1])
			if err != nil{
				log.Fatal(err)
			}
			ttaken := tend.Sub(tstart)
			totalTime = totalTime + ttaken
			fmt.Println(fmt.Sprintf("%s,%s,%s,%v", cols[0],cols[1],cols[2],ttaken))
		} else {
			fmt.Println(totalTime)
			fmt.Println(line)
			totalTime = 0
		}
	}
}
