package carrierdb

import (
	"github.com/datumtechs/datum-network-carrier/carrierdb/iface"
	"github.com/datumtechs/datum-network-carrier/types"
)

type CarrierDB interface {
	iface.LocalStoreCarrierDB
	iface.MetadataCarrierDB
	iface.MetadataAuthorityCarrierDB
	iface.ResourceCarrierDB
	iface.IdentityCarrierDB
	iface.TaskCarrierDB
	InsertData(blocks types.Blocks) (int, error)
	Stop()
}