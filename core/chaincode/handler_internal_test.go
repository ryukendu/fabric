/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chaincode

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/common/sysccprovider"
	"github.com/stretchr/testify/assert"
)

func TestGetChaincodeInstance(t *testing.T) {
	tests := []struct {
		ccName   string
		expected *sysccprovider.ChaincodeInstance
	}{
		{"name", &sysccprovider.ChaincodeInstance{ChaincodeName: "name"}},
		{"name:version", &sysccprovider.ChaincodeInstance{ChaincodeName: "name", ChaincodeVersion: "version"}},
		{"name/chain-id", &sysccprovider.ChaincodeInstance{ChaincodeName: "name", ChainID: "chain-id"}},
		{"name:version/chain-id", &sysccprovider.ChaincodeInstance{ChaincodeName: "name", ChaincodeVersion: "version", ChainID: "chain-id"}},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ci := getChaincodeInstance(tc.ccName)
			assert.Equal(t, tc.expected, ci)
		})
	}
}
