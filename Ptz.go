package onvif

import (
	//"context"
	//"fmt"
	//"log"
	//"fmt"
	"errors"
	"fmt"
	"net/http"

	goonvif "github.com/use-go/onvif"

	//"github.com/use-go/onvif/device"
	//sdk "github.com/use-go/onvif/sdk/device"
	//"github.com/use-go/onvif/sdk/ptz"
	"io/ioutil"

	vif "github.com/use-go/onvif/xsd/onvif"

	"github.com/use-go/onvif/ptz"
	"github.com/beevik/etree"
	"github.com/use-go/onvif/media"
	"github.com/use-go/onvif/xsd"
)

type DeviceParams struct {
	Xaddr      string
	Username   string
	Password   string
	HttpClient *http.Client
}

type Device struct {
	Dev    *goonvif.Device
	onvifTokens map[int]string
	videoTokens  map[int]string
}
//UP, DOWN, LEFT, RIGHT, UP_LEFT, DOWN_LEFT, UP_RIGHT, DOWN_RIGHT, STOP

// 设备类型
type PtzType int32

const (
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
	
)



func NewDevice(params DeviceParams) (*Device, error) {
	dev := &Device{
		onvifTokens:make(map[int]string),
		videoTokens:make(map[int]string),
	}
	err:=errors.New("")
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
	fmt.Println("2222222222",dev)
	return dev, err
}
//https://www.freesion.com/article/65451238750/
func (dev *Device) PtzUp() error{


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
				X:  0,
				Y: 0.2,
			},
			Zoom: vif.Vector1D{
				X: 0,
			},
		},
		//Timeout: "PT10S",
	}

	res, err := dev.Dev.CallMethod(ptzRelReq)
	fmt.Println("11111111111",ptzRelReq ,res)
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

func (dev *Device) ControlPTZ( control_type int32,  control bool,speed float64) error{


	ptzRelReq := ptz.RelativeMove{
		ProfileToken: vif.ReferenceToken(dev.onvifTokens[0]),
		Translation: vif.PTZVector{
			PanTilt: vif.Vector2D{
				X: 0.0,
				Y: 0.0,
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
	}

	switch PtzType(control_type) {
	case PTZ_UP:
		ptzRelReq.Translation.PanTilt.Y=-speed
		//ptzRelReq.Speed.PanTilt.Y=speed
	case PTZ_DOWN:
		ptzRelReq.Translation.PanTilt.Y=speed
		//ptzRelReq.Speed.PanTilt.Y=-speed
	case PTZ_LEFT:
		ptzRelReq.Translation.PanTilt.X=speed
	case PTZ_RIGHT:
		ptzRelReq.Translation.PanTilt.X=-speed
	case PTZ_LEFTUP:
		ptzRelReq.Translation.PanTilt.X = speed;
        ptzRelReq.Translation.PanTilt.Y = -speed;
    case PTZ_LEFTDOWN:
        ptzRelReq.Translation.PanTilt.X = speed;
        ptzRelReq.Translation.PanTilt.Y = speed;
	case PTZ_RIGHTUP:
		ptzRelReq.Translation.PanTilt.X = -speed;
		ptzRelReq.Translation.PanTilt.Y = -speed;
    case PTZ_RIGHTDOWN:
		ptzRelReq.Translation.PanTilt.X = -speed;
		ptzRelReq.Translation.PanTilt.Y = speed;
	case PTZ_ZOOMIN:
		ptzRelReq.Translation.Zoom.X=speed
	case PTZ_ZOOMOUT:
        ptzRelReq.Translation.Zoom.X = -speed;
	}
	res, err := dev.Dev.CallMethod(ptzRelReq)
	fmt.Println("11111111111",ptzRelReq ,res,err)

	return nil
}