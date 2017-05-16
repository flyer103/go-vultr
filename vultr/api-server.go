package vultr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// IPv6 Network info for Server Info
type ServerV6Network struct {
	V6MainIP      string `json:"v6_main_ip"`
	V6NetworkSize string `json:"v6_network_size"`
	V6Network     string `json:"v6_network"`
}

// Detailed server info
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

// Result of reinstalling server.
type ServerReinstallResult struct {
	SubID  string
	MainIP string
	Label  string
	Error  error
}

// List all active or pending virtual machines on the current account.
// The "status" field represents the status of the subscription and will be one of: pending | active | suspended | closed. If the status is "active", you can check "power_status" to determine if the VPS is powered on or not. When status is "active", you may also use "server_state" for a more detailed status of: none | locked | installingbooting | isomounting | ok.
// The API does not provide any way to determine if the initial installation has completed or not. The "v6_network", "v6_main_ip", and "v6_network_size" fields are deprecated in favor of "v6_networks".
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

// Reinstall the operating system on a virtual machine. All data will be permanently lost, but the IP address will remain the same. There is no going back from this call.
// `subID` can be found using `ServerList()`
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

// Reinstall all active or pending virtual machines on the current account.
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
