// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package nodeconfigs

type NodeIPAddr struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	IP        string `json:"ip"`
	IsOn      bool   `json:"isOn"`
	IsUp      bool   `json:"isUp"`
	CanAccess bool   `json:"canAccess"`
}
