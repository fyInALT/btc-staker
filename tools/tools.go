//go:build tools
// +build tools

package btcstaker

import (
	_ "github.com/btcsuite/btcd"
	_ "github.com/btcsuite/btcwallet"
	_ "github.com/babylonchain/babylon/cmd/babylond"
)
