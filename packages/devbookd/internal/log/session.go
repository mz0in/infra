package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	mmdsDefaultAddress = "169.254.169.254"
)

var (
	mmdsTokenExpiration = 60 * time.Second
)

type sessionWriter struct{}

type opts struct {
	SessionID     string `json:"sessionID"`
	CodeSnippetID string `json:"codeSnippetID"`
	Address       string `json:"address"`
}

func addOptsToJSON(jsonLogs []byte, opts *opts) ([]byte, error) {
	var parsed map[string]interface{}

	json.Unmarshal(jsonLogs, &parsed)

	parsed["sessionID"] = opts.SessionID
	parsed["codeSnippetID"] = opts.CodeSnippetID

	data, err := json.Marshal(parsed)
	return data, err
}

func newSessionWriter() *sessionWriter {
	return &sessionWriter{}
}

func getMMDSToken(client http.Client) (string, error) {
	fmt.Println("Retrieving MMDS token")

	request, err := http.NewRequest("PUT", "http://"+mmdsDefaultAddress+"/latest/api/token", new(bytes.Buffer))
	if err != nil {
		return "", err
	}

	request.Header["X-metadata-token-ttl-seconds"] = []string{fmt.Sprint(mmdsTokenExpiration.Seconds())}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	fmt.Printf("Reading mmds token response body")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	token := string(body)

	if len(token) == 0 {
		return "", fmt.Errorf("mmds token is an empty string")
	}

	return token, nil
}

func getMMDSOpts(client http.Client, token string) (*opts, error) {
	fmt.Printf("Retrieving MMDS opts")

	request, err := http.NewRequest("GET", "http://"+mmdsDefaultAddress, new(bytes.Buffer))
	if err != nil {
		return nil, err
	}
	request.Header["X-metadata-token"] = []string{token}
	request.Header["Accept"] = []string{"application/json"}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	fmt.Printf("Reading mmds opts response body")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Unmarshalling body to json")
	var opts opts
	err = json.Unmarshal(body, &opts)
	if err != nil {
		return nil, err
	}

	fmt.Printf("MMDS opts body unmarshalled")

	if opts.Address == "" {
		return nil, fmt.Errorf("no 'address' in mmds opts")
	}

	if opts.CodeSnippetID == "" {
		return nil, fmt.Errorf("no 'codeSnippetID' in mmds opts")
	}

	if opts.SessionID == "" {
		return nil, fmt.Errorf("no 'sessionID' in mmds opts")
	}

	return &opts, nil
}

func sendSessionLogs(client http.Client, logs []byte, address string) error {
	fmt.Printf("Sending session logs")

	request, err := http.NewRequest("POST", address, bytes.NewBuffer(logs))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	fmt.Printf("Session logs sent")

	return nil
}

func sendLogs(logs []byte) {
	client := http.Client{
		Timeout: 6 * time.Second,
	}

	mmdsToken, err := getMMDSToken(client)
	if err != nil {
		fmt.Printf("error getting mmds token: %+v", err)
		return
	}

	mmdsOpts, err := getMMDSOpts(client, mmdsToken)
	if err != nil {
		fmt.Printf("error getting session logging options from mmds (token %s): %+v", mmdsToken, err)
		return
	}

	fmt.Printf("Logs identification: %+v", mmdsOpts)

	sessionLogs, err := addOptsToJSON(logs, mmdsOpts)
	if err != nil {
		fmt.Printf("error adding session logging options (%+v) to JSON (%+v) with logs : %+v", mmdsOpts, logs, err)
		return
	}

	err = sendSessionLogs(client, sessionLogs, mmdsOpts.Address)
	if err != nil {
		fmt.Printf("error sending session logs: %+v", err)
		return
	}
}

func (w *sessionWriter) Write(logs []byte) (int, error) {
	go sendLogs(logs)

	return len(logs), nil
}
