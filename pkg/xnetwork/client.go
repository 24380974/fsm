package xnetwork

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	xnetv1alpha1 "github.com/flomesh-io/fsm/pkg/apis/xnetwork/v1alpha1"

	"github.com/flomesh-io/fsm/pkg/announcements"
	"github.com/flomesh-io/fsm/pkg/k8s"
	"github.com/flomesh-io/fsm/pkg/k8s/informers"
	"github.com/flomesh-io/fsm/pkg/messaging"
)

// NewXNetworkController returns a xnetwork.Controller interface related to functionality provided by the resources in the xnetwork.flomesh.io API group
func NewXNetworkController(informerCollection *informers.InformerCollection, kubeClient kubernetes.Interface, kubeController k8s.Controller, msgBroker *messaging.Broker) Controller {
	client := &Client{
		informers:      informerCollection,
		kubeClient:     kubeClient,
		kubeController: kubeController,
	}

	shouldObserve := func(obj interface{}) bool {
		object, ok := obj.(metav1.Object)
		if !ok {
			return false
		}
		return kubeController.IsMonitoredNamespace(object.GetNamespace())
	}

	xAccessControlEventTypes := k8s.EventTypes{
		Add:    announcements.XAccessControlAdded,
		Update: announcements.XAccessControlUpdated,
		Delete: announcements.XAccessControlDeleted,
	}
	client.informers.AddEventHandler(informers.InformerKeyXNetworkAccessControl, k8s.GetEventHandlerFuncs(shouldObserve, xAccessControlEventTypes, msgBroker))

	podEventTypes := k8s.EventTypes{
		Add:    announcements.PodAdded,
		Update: announcements.PodUpdated,
		Delete: announcements.PodDeleted,
	}
	client.informers.AddEventHandler(informers.InformerKeyPod, k8s.GetEventHandlerFuncs(shouldObserve, podEventTypes, msgBroker))
	return client
}

// GetAccessControls lists AccessControls
func (c *Client) GetAccessControls() []*xnetv1alpha1.AccessControl {
	var accessControls []*xnetv1alpha1.AccessControl
	for _, accessControlIface := range c.informers.List(informers.InformerKeyXNetworkAccessControl) {
		accessControl := accessControlIface.(*xnetv1alpha1.AccessControl)
		accessControls = append(accessControls, accessControl)
	}
	return accessControls
}
