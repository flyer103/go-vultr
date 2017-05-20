package vultr

import (
	"encoding/json"
	"net/http"
)

type RegionInfo struct {
	DCID           string `json:"DCID"`
	Name           string `json:"name"`
	Country        string `json:"country"`
	Continent      string `json:"continent"`
	State          string `json:"state"`
	DDOSProtection bool   `json:"ddos_protection"`
	BlockStorage   bool   `json:"block_storage"`
	ReginCode      string `json:"region_code"`
}

func (vc *Client) RegionsList() (map[string]RegionInfo, error) {
	req, err := http.NewRequest(http.MethodGet, APIRegionsList, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(HeaderAPIKey, vc.APIKey)

	resp, err := vc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	regionsInfo := map[string]RegionInfo{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&regionsInfo); err != nil {
		return nil, err
	}

	return regionsInfo, nil
}
