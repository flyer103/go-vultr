package vultr

import (
	"encoding/json"
	"net/http"
)

type OSInfo struct {
	OSID    int    `json:"OSID"`
	Name    string `json:"name"`
	Arch    string `json:"arch"`
	Family  string `json:"family"`
	Windows bool   `json:"windows"`
}

// Retrieve a list of available operating systems. If the "windows" flag is true,
// a Windows license will be included with the instance, which will increase the cost.
func (vc *Client) OSList() (map[string]OSInfo, error) {
	req, err := http.NewRequest(http.MethodGet, APIOSList, nil)
	if err != nil {
		return nil, err
	}

	resp, err := vc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	osInfo := map[string]OSInfo{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&osInfo); err != nil {
		return nil, err
	}

	return osInfo, nil
}
