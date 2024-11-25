package maps

import "github.com/flomesh-io/fsm/pkg/logger"

var (
	log = logger.New("fsm-xnet-ebpf-maps")
)

const (
	IPPROTO_TCP L4Proto = 6
	IPPROTO_UDP L4Proto = 17
)

type L4Proto uint8

const (
	ACL_DENY    Acl = 0
	ACL_TRUSTED Acl = 2
)

type Acl uint8

type AclKey struct {
	Addr  [4]uint32
	Port  uint16
	Proto uint8
}

type AclVal struct {
	Acl  uint8
	Flag uint8
	Id   uint16
}
