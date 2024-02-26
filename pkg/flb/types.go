// Package flb contains controller logic for the flb
package flb

import "github.com/flomesh-io/fsm/pkg/logger"

var (
	log = logger.New("flb-utilities")
)

type TLSSecretMode string

const (
	TLSSecretModeLocal  TLSSecretMode = "local"
	TLSSecretModeRemote TLSSecretMode = "remote"
)
