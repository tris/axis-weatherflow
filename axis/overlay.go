package axis

import (
	"fmt"

	"github.com/godbus/dbus"
)

type Overlay struct {
	properties OverlayProperties
	path       dbus.ObjectPath
	conn       *dbus.Conn
}

type OverlayProperties struct {
	Text           *string
	TextColor      *string
	TextOLColor    *string
	TextBGColor    *string
	FontSize       *string
	ScrollSpeed    *int
	Position       *string
	PtPosition     *string
	ZoomInterval   *string
	Indicator      *string
	IndicatorSize  *int
	IndicatorColor *string
	IndicatorBG    *string
	IndicatorOL    *string
	OverlayPath    *string
	Size           *string
	Visible        *bool
	AnchorPoint    *string
	Reference      *string
	ZIndex         *int
	AlwaysOn       *bool
	Legacy         *bool
	OverlayId      *int
	Owner          *string
	Camera         *string
	OverlayType    *string
	Stream         *int
	OverlayFormat  *string
	DestroyOnClose *bool
	Scalable       *bool
	TextLength     *int
	NbrOfBuffers   *int
	SwapMethod     *int
}

func NewOverlay(props OverlayProperties) (*Overlay, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, fmt.Errorf("error connecting to system bus: %v", err)
	}

	obj := conn.Object("com.axis.Overlay2", "/com/axis/Overlay2")

	data := structToArrayOfPairs(props)

	var path dbus.ObjectPath
	err = obj.Call("com.axis.Overlay2.Factory.CreateTextOverlay", 0, data).Store(&path)
	if err != nil {
		return nil, fmt.Errorf("error calling CreateTextOverlay: %v", err)
	}

	overlay := &Overlay{
		properties: props,
		path: path,
		conn: conn,
	}

	return overlay, nil
}

func (o *Overlay) Set(props OverlayProperties) error {
	obj := o.conn.Object("com.axis.Overlay2", o.path)

	data := structToArrayOfPairs(props)

	call := obj.Call("com.axis.Overlay2.Overlay.SetProperties", 0, data)
	if call.Err != nil {
		return fmt.Errorf("error calling SetProperties: %v", call.Err)
	}

	o.properties = props
	return nil
}

func (o *Overlay) Remove() error {
	obj := o.conn.Object("com.axis.Overlay2", "/com/axis/Overlay2")

	call := obj.Call("com.axis.Overlay2.Factory.RemoveOverlay", 0, o.path)
	if call.Err != nil {
		return fmt.Errorf("error calling RemoveOverlay: %v", call.Err)
	}

	o.conn.Close()
	o = nil

	return nil
}
