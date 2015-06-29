package brain

import (
	"net"
)

type Instance struct {
	Ip      net.IP
	Enabled bool
}

type Instances struct {
	Pool []Instance
}
