package v102

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetAdminAddressFromPlanInfo(t *testing.T) {
	testCases := []struct {
		info          string
		expected      string
		expectedError bool
	}{
		{
			info:          "",
			expected:      "",
			expectedError: true,
		},
		{
			info:          "This proposal aims to add liquid stake possibility for the TAC network. whitelist_admin_address: tac15lvhklny0khnwy7hgrxsxut6t6ku2cgknw79fr",
			expected:      "tac15lvhklny0khnwy7hgrxsxut6t6ku2cgknw79fr",
			expectedError: false,
		},
	}

	sdk.GetConfig().SetBech32PrefixForAccount("tac", "tacpub")

	for _, tc := range testCases {
		addr, err := getAdminAddressFromPlanInfo(tc.info)
		if !tc.expectedError && err != nil {
			t.Error(err)
		}
		require.Equal(t, tc.expected, addr)
	}
}
