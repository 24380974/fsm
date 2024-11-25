package maps

import (
	"errors"
	"unsafe"

	"github.com/cilium/ebpf"
	"golang.org/x/sys/unix"

	"github.com/flomesh-io/fsm/pkg/sidecar/v2/xnet/bpf"
	"github.com/flomesh-io/fsm/pkg/sidecar/v2/xnet/fs"
)

func AddAclEntry(aclKey *AclKey, aclVal *AclVal) error {
	pinnedFile := fs.GetPinningFile(bpf.FSM_MAP_NAME_ACL)
	if aclMap, err := ebpf.LoadPinnedMap(pinnedFile, &ebpf.LoadPinOptions{}); err == nil {
		defer aclMap.Close()
		return aclMap.Update(unsafe.Pointer(aclKey), unsafe.Pointer(aclVal), ebpf.UpdateAny)
	} else {
		return err
	}
}

func DelAclEntry(aclKey *AclKey) error {
	pinnedFile := fs.GetPinningFile(bpf.FSM_MAP_NAME_ACL)
	if aclMap, err := ebpf.LoadPinnedMap(pinnedFile, &ebpf.LoadPinOptions{}); err == nil {
		defer aclMap.Close()
		err = aclMap.Delete(unsafe.Pointer(aclKey))
		if errors.Is(err, unix.ENOENT) {
			return nil
		}
		return err
	} else {
		return err
	}
}

func GetAclEntries() map[AclKey]AclVal {
	items := make(map[AclKey]AclVal)
	pinnedFile := fs.GetPinningFile(bpf.FSM_MAP_NAME_ACL)
	aclMap, mapErr := ebpf.LoadPinnedMap(pinnedFile, &ebpf.LoadPinOptions{})
	if mapErr != nil {
		log.Fatal().Err(mapErr).Msgf("failed to load ebpf map: %s", pinnedFile)
	}
	defer aclMap.Close()
	aclKey := new(AclKey)
	aclVal := new(AclVal)
	it := aclMap.Iterate()
	for it.Next(unsafe.Pointer(aclKey), unsafe.Pointer(aclVal)) {
		items[*aclKey] = *aclVal
	}
	return items
}
