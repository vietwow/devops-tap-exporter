package main

import (
    // "fmt"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "os"
    "os/signal"
    "syscall"
    "log"
    "time"
    "net/http"

    "github.com/vietwow/devops-tap-exporter/pkg/healthcheck"
)

func getMetrics() {

    // url := fmt.Sprintf(baseURL+"/%s/todos", s.Username)

    client := healthcheck.NewClient()

    // res, _ := client.DoHck("http://13.112.47.182:8086/health")
    // gauge0.Set(float64(res.ConsumerData["13.112.47.182:8186"].ConnectionCount))

    res, _ := client.DoHck("http://172.31.19.76:8086/health")
    gauge1.Set(float64(res.ConsumerData["172.31.19.76:443"].ConnectionCount))
    gauge2.Set(float64(res.ConsumerData["172.31.19.76:6502"].ConnectionCount))

    res, _ = client.DoHck("http://172.31.23.27:8086/health")
    gauge3.Set(float64(res.ConsumerData["172.31.23.27:443"].ConnectionCount))
    gauge4.Set(float64(res.ConsumerData["172.31.23.27:6502"].ConnectionCount))

    res, _ = client.DoHck("http://172.31.25.57:8086/health")
    gauge5.Set(float64(res.ConsumerData["172.31.25.57:443"].ConnectionCount))
    gauge6.Set(float64(res.ConsumerData["172.31.25.57:6502"].ConnectionCount))

    res, _ = client.DoHck("http://172.31.16.71:8086/health")
    gauge7.Set(float64(res.ConsumerData["172.31.16.71:443"].ConnectionCount))
    gauge8.Set(float64(res.ConsumerData["172.31.16.71:6502"].ConnectionCount))
}

var (
    // counter = prometheus.NewCounter(
    //    prometheus.CounterOpts{
    //       Namespace: "golang",
    //       Name:      "my_counter",
    //       Help:      "This is my counter",
    //    })

    // gauge0 = prometheus.NewGauge(
    //     prometheus.GaugeOpts{
    //         Namespace: "golang",
    //         Name:      "node1a",
    //         Help:      "This is my gauge",
    //         ConstLabels: prometheus.Labels{
    //             "node":   "13.112.47.182:8086",
    //         },
    //     })

    gauge1 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "tap",
            Name:      "ws_connections",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.19.76:443",
            },
        })

    gauge2 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "tap",
            Name:      "ws_connections",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.19.76:6502",
            },
        })

    gauge3 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "tap",
            Name:      "ws_connections",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.23.27:443",
            },
        })

    gauge4 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "tap",
            Name:      "ws_connections",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.23.27:6502",
            },
        })

    gauge5 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "tap",
            Name:      "ws_connections",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.25.57:443",
            },
        })

    gauge6 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "tap",
            Name:      "ws_connections",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.25.57:6502",
            },
        })

    gauge7 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "tap",
            Name:      "ws_connections",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.16.71:443",
            },
        })

    gauge8 = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "tap",
            Name:      "ws_connections",
            Help:      "This is my gauge",
            ConstLabels: prometheus.Labels{
                "node":   "172.31.16.71:6502",
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
    // prometheus.MustRegister(gauge0)
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

    // fmt.Println("====== Start =======")
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
