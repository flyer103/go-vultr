package vultr

import (
	"encoding/json"
	"fmt"
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
	Availability   []int  `json:"availability,omitempty"`
}

// Retrieve a list of all active regions. Note that just because a region is listed here,
// does not mean that there is room for new servers.
func (vc *Client) RegionsList(needAvai bool) (map[string]RegionInfo, error) {
	api := APIRegionsList
	if needAvai {
		api = fmt.Sprintf("%s?availability=yes", api)
	}

	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}

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

// Retrieve a list of the VPSPLANIDs currently available in this location.
// If your account has special plans available, you will need to pass your API key in
// order to see them. For all other accounts, the API key is not required.
func (vc *Client) RegionsAvailability(dcID string, needAPIKey bool) ([]int, error) {
	if dcID == "" {
		return nil, ErrNoDCID
	}
	api := fmt.Sprintf("%s?DCID=%s", APIRegionsAvailability, dcID)
	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}
	if needAPIKey {
		req.Header.Set(HeaderAPIKey, vc.APIKey)
	}

	resp, err := vc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dcIDs := []int{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&dcIDs); err != nil {
		return nil, err
	}

	return dcIDs, nil
}
