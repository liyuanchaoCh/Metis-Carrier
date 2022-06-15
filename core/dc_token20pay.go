package core

import (
	"github.com/datumtechs/datum-network-carrier/carrierdb/rawdb"
	"github.com/datumtechs/datum-network-carrier/types"
)

func (dc *DataCenter) StoreOrgWallet(sysWallet *types.OrgWallet) error {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.StoreOrgWallet(dc.db, sysWallet)
}

// QueryOrgWallet does not return ErrNotFound if the organization wallet not found.
func (dc *DataCenter) QueryOrgWallet() (*types.OrgWallet, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryOrgWallet(dc.db)
}
