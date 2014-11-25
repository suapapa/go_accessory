package accessory

import (
	"errors"

	"github.com/kylelemons/gousb/usb"
)

type Context struct {
	ctx    *usb.Context
	device *usb.Device

	Protocol uint16
}

func NewContext() *Context {
	return &Context{
		ctx: usb.NewContext(),
	}
}

func (c *Context) Close() {
	if c.device != nil {
		c.device.Close()
	}
	c.ctx.Close()
}

func (c *Context) switchToAccessoryMode(manufacturer, model, description, version, uri, serial string) error {
	if c.device != nil {
		return errors.New("accessory: already has device")
	}

	// list up candidates
	devs, err := c.ctx.ListDevices(func(desc *usb.Descriptor) bool {
		switch desc.Vendor {
		case 0x18D1, 0x22B8, 0x04E8:
			return true
		}
		return false
	})
	if err != nil {
		return err
	}

	for _, d := range devs {
		defer d.Close()

		v, err := getProtocol(d)
		if err != nil || v < 1 {
			continue
		}
		c.Protocol = v

		sendString(d, STRING_MANUFACTURER, manufacturer)
		sendString(d, STRING_MODEL, model)
		sendString(d, STRING_DESCRIPTION, description)
		sendString(d, STRING_VERSION, version)
		sendString(d, STRING_URI, uri)
		sendString(d, STRING_SERIAL, serial)

		err = start(d)
		if err != nil {
			continue
		}
	}

	devs, err = c.ctx.ListDevices(func(desc *usb.Descriptor) bool {
		if desc.Vendor == USB_VENDOR_ID {
			switch desc.Product {
			case USB_PRODUCT_ID, USB_ADB_PRODUCT_ID:
				return true
			}
		}
		return false
	})

	if err != nil || len(devs) == 0 {
		return errors.New("accessory: failed to switch to accessory")
	}

	if len(devs) > 1 {
		for _, d := range devs {
			defer d.Close()
		}
		return errors.New("accessory: more then one accessory")
	}

	c.device = devs[0]

	return nil
}

func listAccessoryDevice(desc *usb.Descriptor) bool {
	if desc.Vendor == USB_VENDOR_ID {
		switch desc.Product {
		case USB_PRODUCT_ID, USB_ADB_PRODUCT_ID:
			return true
		}
	}
	return false
}
