package app_test

import (
	"fmt"
	"strconv"
	"testing"
  )

const (
	errCode     = uint32(19) // checking the completion of the work of the worker
	successCode = uint32(0)
    )

// If an error occursm, so....
func CommonArgs(net *network.Network) []string {
	return []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastAsync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String()),
	}
