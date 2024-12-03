package curator_discovery

import (
	"sync"

	"github.com/flomesh-io/fsm/cmd/zk/kv"
	"github.com/flomesh-io/fsm/cmd/zk/zookeeper"
	"github.com/flomesh-io/fsm/pkg/logger"
)

var (
	log = logger.New("curator_discovery")
)

type ServiceDiscovery struct {
	client   *kv.ZookeeperClient
	mutex    *sync.Mutex
	basePath string
	services *sync.Map
	listener *zookeeper.ZkEventListener
}
