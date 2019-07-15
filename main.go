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

type healthcheck_response struct {
    Condition Condition `json:"condition"`
    consumerData consumerData `json:"consumerData"`
    versionData versionData `json:"versionData"`
}

type Condition struct {
    health string `json:"health"`
    reason string `json:"reason"`
}

type consumerData struct {
    consumerData []map[string]interface{} `json:"consumerData"`
}

type versionData struct {
    commitAuthor string `json:"commitAuthor"`
    commitCommitter string `json:"commitCommitter"`
    description string `json:"description"`
    gitBranch string `json:"gitBranch"`
    gitCommitHash string `json:"gitCommitHash"`
    homepage string `json:"homepage"`
    version string `json:"version"`
    workingDirectoryState string `json:"workingDirectoryState"`
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
    tap_metrics := healthcheck_response{}
    jsonErr := json.Unmarshal(body, &tap_metrics)
    if jsonErr != nil {
        log.Fatal(jsonErr)
    }

    fmt.Println(tap_metrics.Condition.health)
    fmt.Println(tap_metrics)
}

func main() {
	// http.Handle("/metrics", promhttp.Handler())
	// http.ListenAndServe(":80", nil)
	getMetrics()
}
