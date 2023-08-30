package onvif

import (
	//"context"
	//"fmt"
	//"log"
	//"fmt"
	//"os"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"context"
	"path"
	"regexp"
	goonvif "github.com/use-go/onvif"

	//"github.com/use-go/onvif/device"
	//sdk "github.com/use-go/onvif/sdk/device"
	//"github.com/use-go/onvif/sdk/ptz"
	"io/ioutil"
	"strings"
	vif "github.com/use-go/onvif/xsd/onvif"

	"github.com/beevik/etree"
	"github.com/use-go/onvif/media"
	"github.com/use-go/onvif/ptz"
	"github.com/use-go/onvif/xsd"
	"github.com/use-go/onvif/device"
	sdkdevice "github.com/use-go/onvif/sdk/device"
	discover "github.com/use-go/onvif/ws-discovery"
)

type DeviceParams struct {
	Xaddr      string
	Username   string
	Password   string
	HttpClient *http.Client
}

type Device struct {
	Dev         *goonvif.Device
	onvifTokens map[int]string
	videoTokens map[int]string
	parms DeviceParams
}

//UP, DOWN, LEFT, RIGHT, UP_LEFT, DOWN_LEFT, UP_RIGHT, DOWN_RIGHT, STOP

// 设备类型
type PtzType int32

const (
	PTZ_UP       PtzType = 0
	PTZ_DOWN     PtzType = 1
	PTZ_LEFT     PtzType = 2
	PTZ_LEFTUP   PtzType = 3
	PTZ_LEFTDOWN PtzType = 4

	PTZ_RIGHT     PtzType = 5
	PTZ_RIGHTUP   PtzType = 6
	PTZ_RIGHTDOWN PtzType = 7

	PTZ_ZOOMIN  PtzType = 8
	PTZ_ZOOMOUT PtzType = 9
)

func NewDevice(params DeviceParams) (*Device, error) {
	dev := &Device{
		onvifTokens: make(map[int]string),
		videoTokens: make(map[int]string),
		parms:DeviceParams{
			Xaddr:      params.Xaddr,
		Username:   params.Username,
		Password:   params.Password,
		},

	}
	err := errors.New("")
	dev.Dev, err = goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      params.Xaddr,
		Username:   params.Username,
		Password:   params.Password,
		HttpClient: new(http.Client),
	})

	// 获取profiles 查询onvifTokens
	profiles := media.GetProfiles{}
	profilesRes, err := dev.Dev.CallMethod(profiles)
	if err != nil {
		return dev, err
	}
   fmt.Println("lll",profilesRes)
	b, err := ioutil.ReadAll(profilesRes.Body)
	if err != nil {
		return dev, err
	}

	doc := etree.NewDocument()
	doc.ReadFromBytes(b)
	root := doc.SelectElement("Envelope")
	if root == nil {
		return dev, err
	}
	token := root.FindElements("./Body/GetProfilesResponse/Profiles")
	if token == nil {
		return dev, err
	}

	for k, res := range token {

		_token := res.SelectAttr("token").Value
		dev.onvifTokens[k] = _token

		// video token
		v := res.FindElement("./VideoSourceConfiguration/SourceToken")
		if v != nil {
			dev.videoTokens[k] = v.Text()
		
		}

	}
	fmt.Println("2222222222", dev)
	return dev, err
}

// https://www.freesion.com/article/65451238750/
func (dev *Device) PtzUp() error {

	/*ptzRelReq := ptz.RelativeMove{
		ProfileToken: vif.ReferenceToken(dev.onvifTokens[0]),
		Translation: vif.PTZVector{
			PanTilt: vif.Vector2D{
				X: 0.0,
				Y: -0.1,
				Space: xsd.AnyURI("http://www.onvif.org/ver10/tptz/PanTiltSpaces/TranslationGenericSpace"),
			},
			Zoom: vif.Vector1D{
				X:	0.0,
				Space: xsd.AnyURI("http://www.onvif.org/ver10/tptz/ZoomSpaces/TranslationGenericSpace"),
			},
		},
		Speed: vif.PTZSpeed{
			PanTilt: vif.Vector2D{
				X:	0.0,
				Y:	0.0,
				Space: xsd.AnyURI("http://www.onvif.org/ver10/tptz/PanTiltSpaces/GenericSpeedSpace"),
			},
			Zoom: vif.Vector1D{
				X:	0.0,
				Space: xsd.AnyURI("http://www.onvif.org/ver10/tptz/ZoomSpaces/ZoomGenericSpeedSpace"),
			},
		},
	}*/

	ptzRelReq := ptz.ContinuousMove{
		ProfileToken: vif.ReferenceToken(dev.onvifTokens[0]),
		//ProfileToken: "Profile_1",
		Velocity: vif.PTZSpeed{
			PanTilt: vif.Vector2D{
				X: 0,
				Y: 0.2,
			},
			Zoom: vif.Vector1D{
				X: 0,
			},
		},
		//Timeout: "PT10S",
	}

	res, err := dev.Dev.CallMethod(ptzRelReq)
	fmt.Println("11111111111", ptzRelReq, res)
	return err
}

//https://www.freesion.com/article/65451238750/

/*const (
PTZ_UP    PtzType = 0
	PTZ_DOWN    PtzType = 1
	PTZ_LEFT    PtzType = 2
	PTZ_LEFTUP    PtzType = 3
	PTZ_LEFTDOWN    PtzType = 4

	PTZ_RIGHT    PtzType = 5
	PTZ_RIGHTUP    PtzType = 6
	PTZ_RIGHTDOWN    PtzType = 7

	PTZ_ZOOMIN   PtzType = 8
	PTZ_ZOOMOUT  PtzType = 9

)*/

func (dev *Device) ControlPTZ(control_type int32, control bool, speed float64) error {

	ptzRelReq := ptz.RelativeMove{
		ProfileToken: vif.ReferenceToken(dev.onvifTokens[0]),
		Translation: vif.PTZVector{
			PanTilt: vif.Vector2D{
				X:     0.0,
				Y:     0.0,
				Space: xsd.AnyURI("http://www.onvif.org/ver10/tptz/PanTiltSpaces/TranslationGenericSpace"),
			},
			Zoom: vif.Vector1D{
				X:     0.0,
				Space: xsd.AnyURI("http://www.onvif.org/ver10/tptz/ZoomSpaces/TranslationGenericSpace"),
			},
		},
		Speed: vif.PTZSpeed{
			PanTilt: vif.Vector2D{
				X:     0.0,
				Y:     0.0,
				Space: xsd.AnyURI("http://www.onvif.org/ver10/tptz/PanTiltSpaces/GenericSpeedSpace"),
			},
			Zoom: vif.Vector1D{
				X:     0.0,
				Space: xsd.AnyURI("http://www.onvif.org/ver10/tptz/ZoomSpaces/ZoomGenericSpeedSpace"),
			},
		},
	}

	switch PtzType(control_type) {
	case PTZ_UP:
		ptzRelReq.Translation.PanTilt.Y = -speed
		//ptzRelReq.Speed.PanTilt.Y=speed
	case PTZ_DOWN:
		ptzRelReq.Translation.PanTilt.Y = speed
		//ptzRelReq.Speed.PanTilt.Y=-speed
	case PTZ_LEFT:
		ptzRelReq.Translation.PanTilt.X = speed
	case PTZ_RIGHT:
		ptzRelReq.Translation.PanTilt.X = -speed
	case PTZ_LEFTUP:
		ptzRelReq.Translation.PanTilt.X = speed
		ptzRelReq.Translation.PanTilt.Y = -speed
	case PTZ_LEFTDOWN:
		ptzRelReq.Translation.PanTilt.X = speed
		ptzRelReq.Translation.PanTilt.Y = speed
	case PTZ_RIGHTUP:
		ptzRelReq.Translation.PanTilt.X = -speed
		ptzRelReq.Translation.PanTilt.Y = -speed
	case PTZ_RIGHTDOWN:
		ptzRelReq.Translation.PanTilt.X = -speed
		ptzRelReq.Translation.PanTilt.Y = speed
	case PTZ_ZOOMIN:
		ptzRelReq.Translation.Zoom.X = speed
	case PTZ_ZOOMOUT:
		ptzRelReq.Translation.Zoom.X = -speed
	}
	res, err := dev.Dev.CallMethod(ptzRelReq)
	fmt.Println("11111111111", ptzRelReq, res, err)

	return nil
}

/**
 * @Description: 搜索设备，返回搜索到的设备列表
 * @Author:ZY
 * @time: 2021-03-25 14:23:04
 * @receiver client
 * @return returnInfo
 */
func (dev *Device) SearchDevice() error {

	s, err := goonvif.GetAvailableDevicesAtSpecificEthernetInterface("以太网 2")

	fmt.Printf("444444444444444444%v %v\n", err, s)
	return nil
}
// GetDeviceInformation 读取设备基础信息
// @author: Sen
// @date  : 2023-03-09 16:36:16
func (dev *Device) GetDeviceInformation(ctx context.Context) {
    // 读取设备基础信息
    getDeviceInformation := device.GetDeviceInformation{}
    getDeviceInformationResponse, err := sdkdevice.Call_GetDeviceInformation(ctx, dev.Dev, getDeviceInformation)
    if err != nil {
        panic(err)
    }
    //HardwareId      string //固件ID/设备编号
    //SerialNumber    string //设备序列号
    //FirmwareVersion string //固件版本
    //Model           string //设备类型
    //Manufacturer    string //厂家信息
    fmt.Println("getDeviceInformationResponse:")
    fmt.Println(getDeviceInformationResponse)
}

func (dev *Device) GetStreamUri() (string, error) {

    profiles := media.GetStreamUri{
		ProfileToken: vif.ReferenceToken(dev.onvifTokens[0]),
	}

	profilesRes, err := dev.Dev.CallMethod(profiles)
	if err != nil {
		fmt.Println("GetStreamUri",err)
		return "", err
	}
	b, err := ioutil.ReadAll(profilesRes.Body)
	if err != nil {
		return "", fmt.Errorf("error:%s", err.Error())
	}
	fmt.Println("444",profilesRes)
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(b); err != nil {
		return "", fmt.Errorf("error:%s", err.Error())
	}

	endpoints := doc.Root().FindElements("./Body/GetStreamUriResponse/MediaUri/Uri")
	if len(endpoints) == 0 {
		fmt.Println("ss")
		return "", fmt.Errorf("error:%s", "no media uri")
	}
	fmt.Println("s1111s")
	mediaUri := endpoints[0].Text()
	if !strings.Contains(mediaUri, "rtsp") {
		fmt.Println("mediaUri:", mediaUri)
		return "", fmt.Errorf("error:%s", "media uri is not rtsp")
	}
	if !strings.Contains(mediaUri, "@") && dev.parms.Username != "" {
		//如果返回的rtsp里没有账号密码，则自己拼接
		mediaUri = strings.Replace(mediaUri, "//", fmt.Sprintf("//%s:%s@", dev.parms.Username, dev.parms.Password), 1)
	}
   fmt.Println("rtsp",mediaUri)
   return "", err
}

type Host struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func  DiscoveryDevice(){
	fmt.Println("123")
	var hosts []*Host
	interfaceName :="以太网 2"  //WLAN
	devices, err := discover.SendProbe(interfaceName, nil, []string{"dn:NetworkVideoTransmitter"}, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})
	if err != nil {
		fmt.Println("SendProbe",err)
		return
	}
	fmt.Println("SendProbe",devices)
	for _, j := range devices {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			fmt.Println("devices",devices)
		} else {

			endpoints := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs")
			scopes := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/Scopes")
			fmt.Println("endpoints",endpoints)
			flag := false

			host := &Host{}

			for _, xaddr := range endpoints {
				xaddr := strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2]
				host.URL = xaddr
				fmt.Println("host.xaddr ",xaddr)
			}
			if flag {
				break
			}
			for _, scope := range scopes {
				re := regexp.MustCompile(`onvif:\/\/www\.onvif\.org\/name\/[A-Za-z0-9-]+`)
				match := re.FindStringSubmatch(scope.Text())
				host.Name = path.Base(match[0])
				fmt.Println("host.Name ",host.Name )
			}

			hosts = append(hosts, host)

		}

	}

	bys, _ := json.Marshal(hosts)
	fmt.Println("qw",bys)
	
}

func (dev *Device) GetSnapshotUri() (string, error) {
	fmt.Println("GetSnapshotUri")
    profiles := media.GetSnapshotUri{
		ProfileToken: vif.ReferenceToken(dev.onvifTokens[0]),
	}


	profilesRes, err := dev.Dev.CallMethod(profiles)
	if err != nil {
		fmt.Println("GetStreamUri",err)
		return "", err
	}
	b, err := ioutil.ReadAll(profilesRes.Body)
	if err != nil {
		return "", fmt.Errorf("error:%s", err.Error())
	}
	fmt.Println("444",profilesRes)
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(b); err != nil {
		return "", fmt.Errorf("error:%s", err.Error())
	}

	endpoints := doc.Root().FindElements("./Body/GetSnapshotUriResponse/MediaUri/Uri")
	if len(endpoints) == 0 {
		fmt.Println("ss")
		return "", fmt.Errorf("error:%s", "no media uri")
	}
	mediaUri := endpoints[0].Text()
	if !strings.Contains(mediaUri, "http") {
		fmt.Println("mediaUri:", mediaUri)
		return "", fmt.Errorf("error:%s", "media uri is not rtsp")
	}
	if !strings.Contains(mediaUri, "@") && dev.parms.Username != "" {
		//如果返回的rtsp里没有账号密码，则自己拼接
		mediaUri = strings.Replace(mediaUri, "//", fmt.Sprintf("//%s:%s@", dev.parms.Username, dev.parms.Password), 1)
	}
	fmt.Printf("保存图像失败： %v\n", endpoints,mediaUri)
   return "", err
}