package v2

import (
	"net"

	"github.com/flomesh-io/fsm/pkg/service"
	"github.com/flomesh-io/fsm/pkg/sidecar/v2/xnet/maps"
	"github.com/flomesh-io/fsm/pkg/sidecar/v2/xnet/util"
)

type xnetworkConfigJob struct {
	done   chan struct{}
	server *Server
}

func (job *xnetworkConfigJob) GetDoneCh() <-chan struct{} {
	return job.done
}

func (job *xnetworkConfigJob) Run() {
	defer close(job.done)
	aclAddrs := make(map[uint32]uint8)
	acls := job.server.xnetworkController.GetAccessControls()
	for _, acl := range acls {
		if len(acl.Spec.Services) > 0 {
			for _, aclSvc := range acl.Spec.Services {
				meshSvc := service.MeshService{Name: aclSvc.Name}
				if len(aclSvc.Namespace) > 0 {
					meshSvc.Namespace = aclSvc.Namespace
				} else {
					meshSvc.Namespace = acl.Namespace
				}
				if k8sSvc := job.server.kubeController.GetService(meshSvc); k8sSvc != nil {
					aclAddr, _ := util.IPv4ToInt(net.ParseIP(k8sSvc.Spec.ClusterIP))
					aclAddrs[aclAddr] = uint8(maps.ACL_TRUSTED)
				}
			}
		}
	}
	job.server.updateAcls(aclAddrs)

	job.server.updateDNSNat()
}

func (job *xnetworkConfigJob) JobName() string {
	return "fsm-xnetwork-config-job"
}
