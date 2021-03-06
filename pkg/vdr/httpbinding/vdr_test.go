/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package httpbinding

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVDR_Close(t *testing.T) {
	t.Run("test success", func(t *testing.T) {
		v, err := New("/did:example:334455")
		require.NoError(t, err)
		require.NoError(t, v.Close())
	})
}

func TestVDR_Build(t *testing.T) {
	t.Run("test HTTP Binding VDR build with default opts", func(t *testing.T) {
		resolver, err := New("/did:example:334455")
		require.NoError(t, err)
		require.NotNil(t, resolver)

		result, err := resolver.Build(nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "build not supported in http binding vdr")
		require.Nil(t, result)
	})
}
