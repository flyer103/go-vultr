package vultr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PlanInfo struct {
	VPSPlanID          string `json:"VPSPLANID"`
	Name               string `json:"name"`
	VCPUCount          string `json:"vcpu_count"`
	Ram                string `json:"ram"`
	Disk               string `json:"disk"`
	Bandwidth          string `json:"bandwidth"`
	BandwidthGB        string `json:"bandwidth_gb"`
	PricePerMonth      string `json:"price_per_month"`
	PlanType           string `json:"plan_type"`
	Windows            bool   `json:"windows"`
	AvailableLocations []int  `json:"available_locations"`
}

var validPlanTypes = []string{"all", "vc2", "ssd", "vdc2", "dedicated"}

// Retrieve a list of all active plans. Plans that are no longer available will not be shown.
// The "windows" field is no longer in use, and will always be false. Windows licenses will be
// automatically added to any plan as necessary.
// The "deprecated" field indicates that the plan will be going away in the future. New deployments
// of it will still be accepted, but you should begin to transition away from its usage. Typically,
// deprecated plans are available for 30 days after they are deprecated.
// The sandbox ($2.50) plan is not available in the API.
//
// Parameters:
// + type string (optional) The type of plans to return. Possible values: "all", "vc2", "ssd", "vdc2", "dedicated".
func (vc *Client) PlansList(typ string) (map[string]PlanInfo, error) {
	api := APIPlansList
	if typ != "" {
		if !checkPlanType(typ) {
			return nil, ErrInvalidPlanType
		}

		api = fmt.Sprintf("%s?type=%s", api, typ)
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

	plansInfo := map[string]PlanInfo{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&plansInfo); err != nil {
		return nil, err
	}

	return plansInfo, nil
}

func checkPlanType(typ string) bool {
	for _, validType := range validPlanTypes {
		if typ == validType {
			return true
		}
	}

	return false
}
