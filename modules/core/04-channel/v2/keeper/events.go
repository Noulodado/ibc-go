package keeper

import (
	"context"

	channeltypesv2 "github.com/cosmos/ibc-go/v9/modules/core/04-channel/v2/types"
)

// EmitSendPacketEvents emits events for the SendPacket handler.
func EmitSendPacketEvents(ctx context.Context, packet channeltypesv2.Packet) {
	// TODO: https://github.com/cosmos/ibc-go/issues/7386
}