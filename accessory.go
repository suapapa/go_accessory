package accessory

import (
	"github.com/kylelemons/gousb/usb"
)

type Context struct {
	ctx    *usb.Context
	device *usb.Device

	protocol int
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

func listAccessoryDevice(desc *usb.Descriptor) bool {
	if desc.Vendor == USB_VENDOR_ID {
		if desc.Product == USB_PRODUCT_ID || desc.Product == USB_ADB_PRODUCT_ID {
			return true
		}
	}
	return false
}
