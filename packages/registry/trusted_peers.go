// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package registry

import (
	"fmt"
	"os"
	"path"

	"github.com/iotaledger/hive.go/core/generics/lo"
	"github.com/iotaledger/hive.go/core/generics/onchangemap"
	"github.com/iotaledger/hive.go/core/ioutils"
	"github.com/iotaledger/hive.go/core/kvstore"
	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/iotaledger/wasp/packages/peering"
)

type jsonTrustedPeers struct {
	TrustedPeers []*peering.TrustedPeer `json:"trustedPeers"`
}

type TrustedPeersRegistryImpl struct {
	onChangeMap *onchangemap.OnChangeMap[string, *peering.ComparablePubKey, *peering.TrustedPeer]

	filePath string
}

var _ TrustedPeersRegistryProvider = &TrustedPeersRegistryImpl{}

// NewTrustedPeersRegistryImpl creates new instance of the trusted peers registry implementation.
func NewTrustedPeersRegistryImpl(filePath string) (*TrustedPeersRegistryImpl, error) {
	registry := &TrustedPeersRegistryImpl{
		filePath: filePath,
	}

	registry.onChangeMap = onchangemap.NewOnChangeMap(
		onchangemap.WithChangedCallback[string, *peering.ComparablePubKey](registry.writeTrustedPeersJSON),
	)

	// load TrustedPeers on startup
	if err := registry.loadTrustedPeersJSON(); err != nil {
		return nil, fmt.Errorf("unable to read TrustedPeers configuration (%s): %s", filePath, err)
	}

	registry.onChangeMap.CallbacksEnabled(true)

	return registry, nil
}

func (p *TrustedPeersRegistryImpl) loadTrustedPeersJSON() error {
	if p.filePath == "" {
		// do not load entries if no path is given
		return nil
	}

	tmpTrustedPeers := &jsonTrustedPeers{}
	if err := ioutils.ReadJSONFromFile(p.filePath, tmpTrustedPeers); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("unable to unmarshal json file: %w", err)
	}

	for _, trustedPeer := range tmpTrustedPeers.TrustedPeers {
		if _, err := p.TrustPeer(trustedPeer.PubKey(), trustedPeer.NetID); err != nil {
			return fmt.Errorf("unable to add trusted peer (%s): %s", p.filePath, err)
		}
	}

	return nil
}

func (p *TrustedPeersRegistryImpl) writeTrustedPeersJSON(trustedPeers []*peering.TrustedPeer) error {
	if p.filePath == "" {
		// do not store entries if no path is given
		return nil
	}

	if err := os.MkdirAll(path.Dir(p.filePath), 0o770); err != nil {
		return fmt.Errorf("unable to create folder \"%s\": %w", path.Dir(p.filePath), err)
	}

	if err := ioutils.WriteJSONToFile(p.filePath, &jsonTrustedPeers{TrustedPeers: trustedPeers}, 0o600); err != nil {
		return fmt.Errorf("unable to marshal json file: %w", err)
	}

	return nil
}

func (p *TrustedPeersRegistryImpl) IsTrustedPeer(pubKey *cryptolib.PublicKey) error {
	_, err := p.onChangeMap.Get(peering.NewComparablePubKey(pubKey))
	if err != nil {
		return kvstore.ErrKeyNotFound
	}

	return nil
}

func (p *TrustedPeersRegistryImpl) TrustPeer(pubKey *cryptolib.PublicKey, netID string) (*peering.TrustedPeer, error) {
	trustedPeer := peering.NewTrustedPeer(pubKey, netID)
	if err := p.onChangeMap.Add(trustedPeer); err != nil {
		// already exists, modify the existing
		return p.onChangeMap.Modify(peering.NewComparablePubKey(pubKey), func(item *peering.TrustedPeer) bool {
			*item = *peering.NewTrustedPeer(pubKey, netID)
			return true
		})
	}

	return trustedPeer, nil
}

func (p *TrustedPeersRegistryImpl) DistrustPeer(pubKey *cryptolib.PublicKey) (*peering.TrustedPeer, error) {
	addr := peering.NewComparablePubKey(pubKey)

	trustedPeer, err := p.onChangeMap.Get(addr)
	if err != nil {
		return nil, nil
	}

	if err := p.onChangeMap.Delete(addr); err != nil {
		return nil, nil
	}

	return trustedPeer, nil
}

func (p *TrustedPeersRegistryImpl) TrustedPeers() ([]*peering.TrustedPeer, error) {
	return lo.Values(p.onChangeMap.All()), nil
}