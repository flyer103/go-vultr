package vultr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ServerV6Network struct {
	V6MainIP      string `json:"v6_main_ip"`
	V6NetworkSize string `json:"v6_network_size"`
	V6Network     string `json:"v6_network"`
}

type ServerInfo struct {
	SubID              string            `json:"SUBID"`
	OS                 string            `json:"os"`
	Ram                string            `json:"ram"`
	Disk               string            `json:"disk"`
	MainIP             string            `json:"main_ip"`
	VCPUCount          string            `json:"vcpu_count"`
	Location           string            `json:"location"`
	DCID               string            `json::"DCID"`
	DefaultPasswd      string            `json:"default_password"`
	DataCreated        string            `json:"date_created"`
	PendingCharges     string            `json:"pending_charges"`
	Status             string            `json:"status"`
	CostPerMonth       string            `json:"cost_per_month"`
	CurrentBandwidthGB float64           `json:"current_bandwidth_gb"`
	AllowedBandwidthGB string            `json:"allowed_bandwidth_gb"`
	NetMaskV4          string            `json:"netmask_v4"`
	GatewayV4          string            `json:"gateway_v4"`
	PowerStatus        string            `json:"power_status"`
	ServerState        string            `json:"server_state"`
	VPSPLanID          string            `json:"VPSPLANID"`
	V6MainIP           string            `json:"v6_main_ip"`
	V6NetworkSize      string            `json:"v6_network_size"`
	V6Network          string            `json:"v6_network"`
	V6Networks         []ServerV6Network `json:"v6_networks"`

	Label           string `json:"label"`
	InternalIP      string `json:"internal_ip"`
	KVMURL          string `json:"kvm_url"`
	AutoBackups     string `json:"auto_backups"`
	Tag             string `json:"tag"`
	OSID            string `json:"OSID"`
	APPID           string `json:"APPID"`
	FirewallGroupID string `json:"FIREWALLGROUPID"`
}

type ServerReinstallResult struct {
	SubID  string
	MainIP string
	Label  string
	Error  error
}

func (vc *Client) ServerList() (map[string]ServerInfo, error) {
	req, err := http.NewRequest(http.MethodGet, APIServerList, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(HeaderAPIKey, vc.APIKey)

	resp, err := vc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	serverInfo := map[string]ServerInfo{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&serverInfo)

	return serverInfo, err
}

func (vc *Client) ServerReinstall(subID string) error {
	data := fmt.Sprintf("SUBID=%s", subID)
	req, err := http.NewRequest(http.MethodPost, APIServerReInstall, strings.NewReader(data))
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

func (vc *Client) ServerReinstallAll() ([]ServerReinstallResult, error) {
	serverInfo, err := vc.ServerList()
	if err != nil {
		return nil, err
	}

	results := []ServerReinstallResult{}
	for subID, info := range serverInfo {
		err := vc.ServerReinstall(subID)
		res := ServerReinstallResult{
			SubID:  subID,
			MainIP: info.MainIP,
			Label:  info.Label,
			Error:  err,
		}
		results = append(results, res)
	}

	return results, nil
}
