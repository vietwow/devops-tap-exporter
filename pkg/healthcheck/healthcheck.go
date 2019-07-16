package healthcheck

import (
    "net/http"
    "log"
)

type ConsumerDataBlock struct {
  ConnectionCount int `json:"connectionCount"`
  ConnectionLimit int `json:"connectionLimit"`
  ConnectionLoad float64 `json:"connectionLoad"`
  ConnectionsRemaining int `json:"connectionsRemaining"`
}

type Healthcheck_response struct {
    Condition Condition `json:"condition"`
    // ConsumerData map[string]map[string]interface{} `json:"consumerData"`
    ConsumerData map[string]ConsumerDataBlock `json:"consumerData"`
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

type Client struct {}
 
func NewClient() *Client {
	return &Client{}
}

func(c *Client) doHck(URL string) (Healthcheck_response, error) {
    tClient := http.Client{
        Timeout: time.Second * 2, // Maximum of 2 secs
    }

    req, err := http.NewRequest(http.MethodGet, URL, nil)
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

    return tap_metrics,nil
}
