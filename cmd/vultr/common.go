package vultr

import (
	"encoding/json"
	"fmt"

	"github.com/flyer103/go-vultr/pkg/apikeyctrl"
	pvultr "github.com/flyer103/go-vultr/pkg/vultr"
)

func GetAPIKey() (string, error) {
	ctrl, err := apikeyctrl.New()
	if err != nil {
		return "", err
	}

	apikey, err := ctrl.Get()
	if err != nil {
		return "", err
	}

	return apikey, nil
}

func NewVultrClient() (*pvultr.Client, error) {
	apikey, err := GetAPIKey()
	if err != nil {
		return nil, err
	}

	cfg := &pvultr.Config{
		APIKey: apikey,

		DialTimeout:           3,
		ResponseHeaderTimeout: 30,
	}
	client, err := pvultr.New(cfg)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func PrettyJsonString(data interface{}) (string, error) {
	dataBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}

	return string(dataBytes), nil
}

func CmdDo(apiName pvultr.APIName, isAll bool, args []string) error {
	if !isAll && (len(args) == 0 || args[0] == "") {
		return ErrNoSUBID
	}

	client, err := NewVultrClient()
	if err != nil {
		return err
	}

	if isAll {
		var res []pvultr.ServerBatchResult
		var err error

		switch apiName {
		case pvultr.APINameServerReinstall:
			res, err = client.ServerReinstallAll()
		case pvultr.APINameServerHalt:
			res, err = client.ServerHaltAll()
		case pvultr.APINameServerStart:
			res, err = client.ServerStartAll()
		case pvultr.APINameServerReboot:
			res, err = client.ServerRebootAll()
		}
		if err != nil {
			return err
		}

		data, err := PrettyJsonString(res)
		if err != nil {
			return err
		}
		fmt.Println(data)

		return nil
	}

	switch apiName {
	case pvultr.APINameServerReinstall:
		return client.ServerReinstall(args[0])
	case pvultr.APINameServerHalt:
		return client.ServerHalt(args[0])
	case pvultr.APINameServerStart:
		return client.ServerStart(args[0])
	case pvultr.APINameServerReboot:
		return client.ServerReboot(args[0])
	default:
		return nil
	}
}
