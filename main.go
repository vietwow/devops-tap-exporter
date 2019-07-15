package main

import (
    "fmt"
    "net/http"
    // "github.com/prometheus/client_golang/prometheus/promhttp"
    "encoding/json"
    "time"
    "log"
    "io/ioutil"
)

// type Healthcheck_response struct {
//     Condition Condition `json:"condition"`
//     ConsumerData ConsumerData `json:"consumerData"`
//     VersionData VersionData `json:"versionData"`
// }

type Healthcheck_response struct {
    Condition Condition `json:"condition"`
    ConsumerData map[string]map[string]interface{} `json:"consumerData"`
    VersionData VersionData `json:"versionData"`
}

type Condition struct {
    Health string `json:"health"`
    Reason string `json:"reason"`
}

type VersionData struct {
    CommitAuthor string `json:"commitAuthor"`
    CommitCommitter string `json:"commitCommitter"`
    Description string `json:"description"`
    GitBranch string `json:"gitBranch"`
    GitCommitHash string `json:"gitCommitHash"`
    Homepage string `json:"homepage"`
    Version string `json:"version"`
    WorkingDirectoryState string `json:"workingDirectoryState"`
}


// type Metrics struct {
//     RedisRequests *prometheus.CounterVec
// }

func getMetrics() {
    fmt.Println("Inside function getMetrics...")

    url := "http://13.112.47.182:8086/health"

    tClient := http.Client{
        Timeout: time.Second * 2, // Maximum of 2 secs
    }

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Fatal(err)
    }

    // req.Header.Set("User-Agent", "spacecount-tutorial")

    res, getErr := tClient.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }

    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }

    // parse
    tap_metrics := Healthcheck_response{}
    jsonErr := json.Unmarshal(body, &tap_metrics)
    if jsonErr != nil {
        log.Fatal(jsonErr)
    }

    // fmt.Println(tap_metrics.Condition.Health)
    // fmt.Println(tap_metrics.ConsumerData)
    for node, v := range tap_metrics.ConsumerData {
        fmt.Println(node, v["connectionCount"])
    }
}

func main() {
	// http.Handle("/metrics", promhttp.Handler())
	// http.ListenAndServe(":80", nil)
	getMetrics()
}
