package vultr

import (
	"fmt"
	"net/http"
	"strings"
)

type APIName int

const (
	APINameServerReinstall APIName = iota
	APINameServerHalt
	APINameServerStart
	APINameServerReboot
	APINameServerDestroy
	APINameServerCreate
)

type APITask struct {
	SubID  string
	MainIP string
	Label  string
	Error  error
}

func (vc *Client) PostSubID(apiName APIName, subID string) error {
	var api string
	switch apiName {
	case APINameServerReinstall:
		api = APIServerReinstall
	case APINameServerHalt:
		api = APIServerHalt
	case APINameServerStart:
		api = APIServerStart
	case APINameServerReboot:
		api = APIServerReboot
	case APINameServerDestroy:
		api = APIServerDestroy
	}

	data := fmt.Sprintf("SUBID=%s", subID)
	req, err := http.NewRequest(http.MethodPost, api, strings.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set(HeaderAPIKey, vc.APIKey)
	req.Header.Set(HeaderContentType, "application/x-www-form-urlencoded")

	resp, err := vc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %d", resp.StatusCode)
	}

	return nil
}

func (vc *Client) PostAllSubIDs(apiName APIName) ([]ServerBatchResult, error) {
	serverInfo, err := vc.ServerList()
	if err != nil {
		return nil, err
	}

	ch := make(chan ServerBatchResult, len(serverInfo))
	results := make([]ServerBatchResult, 0, len(serverInfo))
	for subID, info := range serverInfo {
		task := APITask{
			SubID:  subID,
			MainIP: info.MainIP,
			Label:  info.Label,
		}

		go vc.doAPITask(apiName, task, ch)
	}

	for i := 0; i < len(serverInfo); i++ {
		res := <-ch
		results = append(results, res)
	}

	return results, nil
}

func (vc *Client) doAPITask(apiName APIName, task APITask, ch chan ServerBatchResult) {
	var err error
	switch apiName {
	case APINameServerReinstall:
		err = vc.ServerReinstall(task.SubID)
	case APINameServerHalt:
		err = vc.ServerHalt(task.SubID)
	case APINameServerStart:
		err = vc.ServerStart(task.SubID)
	case APINameServerReboot:
		err = vc.ServerReboot(task.SubID)
	case APINameServerDestroy:
		err = vc.ServerDestroy(task.SubID)
	}

	res := ServerBatchResult{
		SubID:  task.SubID,
		MainIP: task.MainIP,
		Label:  task.Label,
		Error:  err,
	}

	ch <- res
}
