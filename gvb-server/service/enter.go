package service

import "time1043/gvb-server/service/images_ser"

type ServiceGroup struct {
	ImageService images_ser.ImageService
}

var ServiceApp = new(ServiceGroup)
