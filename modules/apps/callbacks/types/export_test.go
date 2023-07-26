package types

import (
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
)

/*
	This file is to allow for unexported functions to be accessible to the testing package.
*/

// GetCallbackData is a wrapper around getCallbackData to allow the function to be directly called in tests.
func GetCallbackData(
	packetDataUnmarshaler porttypes.PacketDataUnmarshaler,
	packet ibcexported.PacketI, remainingGas uint64,
	maxGas uint64, callbackKey string,
) (CallbackData, bool, error) {
	return getCallbackData(packetDataUnmarshaler, packet, remainingGas, maxGas, callbackKey)
}

// GetCallbackAddress is a wrapper around getCallbackAddress to allow the function to be directly called in tests.
func GetCallbackAddress(callbackData map[string]interface{}) string {
	return getCallbackAddress(callbackData)
}

// GetUserDefinedGasLimit is a wrapper around getUserDefinedGasLimit to allow the function to be directly called in tests.
func GetUserDefinedGasLimit(callbackData map[string]interface{}) uint64 {
	return getUserDefinedGasLimit(callbackData)
}
