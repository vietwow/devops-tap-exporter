package main

import (
    // "fmt"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "encoding/json"
    "time"
    "log"
    "io/ioutil"
    "os"
    "os/signal"
    "syscall"

    "github.com/vietwow/devops-tap-exporter/pkg/healthcheck"
)

// type ConsumerDataBlock struct {
//   ConnectionCount int `json:"connectionCount"`
//   ConnectionLimit int `json:"connectionLimit"`
//   ConnectionLoad float64 `json:"connectionLoad"`
//   ConnectionsRemaining int `json:"connectionsRemaining"`
// }

// type Healthcheck_response struct {
//     Condition Condition `json:"condition"`
//     // ConsumerData map[string]map[string]interface{} `json:"consumerData"`
//     ConsumerData map[string]ConsumerDataBlock `json:"consumerData"`
//     VersionData VersionData `json:"versionData"`
// }

// type Condition struct {
//     Health string `json:"health"`
//     Reason string `json:"reason"`
// }

// type VersionData struct {
//     CommitAuthor string `json:"commitAuthor"`
//     CommitCommitter string `json:"commitCommitter"`
//     Description string `json:"description"`
//     GitBranch string `json:"gitBranch"`
//     GitCommitHash string `json:"gitCommitHash"`
//     Homepage string `json:"homepage"`
//     Version string `json:"version"`
//     WorkingDirectoryState string `json:"workingDirectoryState"`
// }


func getMetrics() {
    // fmt.Println("Inside function getMetrics...")

    // url := fmt.Sprintf(baseURL+"/%s/todos", s.Username)

    client := healthcheck.NewClient()

    res, _ := client.doHck("http://172.31.19.76:8086/health")
    gauge1.Set(float64(res.ConsumerData["172.31.19.76:8186"].ConnectionCount))

    res, _ := client.doHck("http://172.31.23.27:8086/health")
    gauge2.Set(float64(res.ConsumerData["172.31.23.27:8186"].ConnectionCount))

    res, _ := client.doHck("http://172.31.25.57:8086/health")
    gauge3.Set(float64(res.ConsumerData["172.31.25.57:8186"].ConnectionCount))

    res, _ := client.doHck("http://172.31.16.71:8086/health")
    gauge4.Set(float64(res.ConsumerData["172.31.16.71:8186"].ConnectionCount))

    // tClient := http.Client{
    //     Timeout: time.Second * 2, // Maximum of 2 secs
    // }

    // req, err := http.NewRequest(http.MethodGet, url, nil)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // // req.Header.Set("User-Agent", "spacecount-tutorial")

    // res, getErr := tClient.Do(req)
    // if getErr != nil {
    //     log.Fatal(getErr)
    // }

    // body, readErr := ioutil.ReadAll(res.Body)
    // if readErr != nil {
    //     log.Fatal(readErr)
    // }

    // // parse
    // tap_metrics := Healthcheck_response{}
    // jsonErr := json.Unmarshal(body, &tap_metrics)
    // if jsonErr != nil {
    //     log.Fatal(jsonErr)
    // }


    

    // fmt.Println(tap_metrics.Condition.Health)
    // fmt.Println(tap_metrics.ConsumerData)

    // empty interface
    // for node, v := range tap_metrics.ConsumerData {
    //     // fmt.Println(node, v["connectionCount"])
    //     if node == "172.31.19.76:443" {
    //         // fmt.Println("add gauge1")
    //         gauge1.Set(float64(v["connectionCount"]))

    //     } else if node == "172.31.19.76:6502" {
    //         // fmt.Println("add gauge2")
    //         gauge2.Set(float64(v["connectionCount"]))
    //     }
    // }

    // fmt.Print(tap_metrics.ConsumerData["172.31.19.76:443"].ConnectionCount)
    gauge1.Set(float64(tap_metrics1.ConsumerData["172.31.19.76:8186"].ConnectionCount))
    gauge2.Set(float64(tap_metrics2.ConsumerData["172.31.23.27:8186"].ConnectionCount))
    gauge3.Set(float64(tap_metrics3.ConsumerData["172.31.25.57:8186"].ConnectionCount))
    gauge4.Set(float64(tap_metrics4.ConsumerData["172.31.16.71:8186"].ConnectionCount))
    gauge1.Set(float64(tap_metrics5.ConsumerData["172.31.19.76:8186"].ConnectionCount))
    gauge2.Set(float64(tap_metrics6.ConsumerData["172.31.23.27:8186"].ConnectionCount))
    gauge3.Set(float64(tap_metrics7.ConsumerData["172.31.25.57:8186"].ConnectionCount))
    gauge4.Set(float64(tap_metrics8.ConsumerData["172.31.16.71:8186"].ConnectionCount))
}

var (
    // counter = prometheus.NewCounter(
    //    prometheus.CounterOpts{
    //       Namespace: "golang",
    //       Name:      "my_counter",
    //       Help:      "This is my counter",
    //    })

    gauge1 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "golang",
            Name:      "172.31.19.76:8186",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.19.76:8186",
            },
        })

    gauge2 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "golang",
            Name:      "172.31.19.76:443",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.19.76:443",
            },
        })

    gauge3 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "golang",
            Name:      "172.31.23.27:8186",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.23.27:8186",
            },
        })

    gauge4 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "golang",
            Name:      "172.31.23.27:443",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.23.27:443",
            },
        })

    gauge5 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "golang",
            Name:      "172.31.25.57:8186",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.25.57:8186",
            },
        })

    gauge6 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "golang",
            Name:      "172.31.25.57:443",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.25.57:443",
            },
        })

    gauge7 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "golang",
            Name:      "172.31.16.71:8186",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.16.71:8186",
            },
        })

    gauge8 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "golang",
            Name:      "172.31.16.71:443",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.16.71:443",
            },
        })

    // histogram = prometheus.NewHistogram(
    //    prometheus.HistogramOpts{
    //       Namespace: "golang",
    //       Name:      "my_histogram",
    //       Help:      "This is my histogram",
    //    })

    // summary = prometheus.NewSummary(
    //    prometheus.SummaryOpts{
    //       Namespace: "golang",
    //       Name:      "my_summary",
    //       Help:      "This is my summary",
    //    })
)

func init() {
    // prometheus.MustRegister(counter)
    prometheus.MustRegister(gauge1)
    prometheus.MustRegister(gauge2)
    prometheus.MustRegister(gauge3)
    prometheus.MustRegister(gauge4)
    prometheus.MustRegister(gauge5)
    prometheus.MustRegister(gauge6)
    prometheus.MustRegister(gauge7)
    prometheus.MustRegister(gauge8)
    // prometheus.MustRegister(histogram)
    // prometheus.MustRegister(summary)
}
func main() {
    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGTERM)
    go func() {
        log.Printf("SIGTERM received: %v. Exiting...", <-signalChan)
        os.Exit(0)
    }()

    go func() {
        for {
            // counter.Add(rand.Float64() * 5)
            //gauge.Add(rand.Float64()*15 - 5)
            getMetrics()
            // histogram.Observe(rand.Float64() * 10)
            // summary.Observe(rand.Float64() * 10)

            time.Sleep(time.Second)
        }
    }()

    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":9114", nil))
}
