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
)

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

	results := make([]ServerBatchResult, 0, len(serverInfo))
	for subID, info := range serverInfo {
		var err error
		switch apiName {
		case APINameServerReinstall:
			err = vc.ServerReinstall(subID)
		case APINameServerHalt:
			err = vc.ServerHalt(subID)
		case APINameServerStart:
			err = vc.ServerStart(subID)
		case APINameServerReboot:
			err = vc.ServerReboot(subID)
		}
		res := ServerBatchResult{
			SubID:  subID,
			MainIP: info.MainIP,
			Label:  info.Label,
			Error:  err,
		}
		results = append(results, res)
	}

	return results, nil
}
