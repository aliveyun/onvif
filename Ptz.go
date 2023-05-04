package onvif

import (
	//"context"
	//"fmt"
	//"log"
	"fmt"
	"net/http"

	goonvif "github.com/use-go/onvif"
	//"github.com/use-go/onvif/device"
	//sdk "github.com/use-go/onvif/sdk/device"
	//"github.com/use-go/onvif/sdk/ptz"
	//"github.com/use-go/onvif/xsd/onvif"
	//"io/ioutil"
	//"github.com/use-go/onvif/ptz"
	//"github.com/beevik/etree"
	//"github.com/use-go/onvif/xsd"
	//"github.com/use-go/onvif/media"
)

type DeviceParams struct {
	Xaddr      string
	Username   string
	Password   string
	HttpClient *http.Client
}

type Device struct {
	dev    *goonvif.Device
}

func NewDevice(params DeviceParams) (*Device, error) {
	dev1 := &Device{}

	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      params.Xaddr,
		Username:   params.Username,
		Password:   params.Password,
		HttpClient: new(http.Client),
	})
	dev1.dev=dev
   fmt.Println(dev ,err)
	return dev1, nil
}
