package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v9/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v9/modules/core/23-commitment/types/v2"
	host "github.com/cosmos/ibc-go/v9/modules/core/24-host"
	ibcerrors "github.com/cosmos/ibc-go/v9/modules/core/errors"
)

var (
	_ sdk.Msg              = (*MsgProvideCounterparty)(nil)
	_ sdk.HasValidateBasic = (*MsgProvideCounterparty)(nil)

	_ sdk.Msg              = (*MsgCreateChannel)(nil)
	_ sdk.HasValidateBasic = (*MsgCreateChannel)(nil)
)

// NewMsgProvideCounterparty creates a new MsgProvideCounterparty instance
func NewMsgProvideCounterparty(channelID, counterpartyChannelID string, signer string) *MsgProvideCounterparty {
	return &MsgProvideCounterparty{
		Signer:                signer,
		ChannelId:             channelID,
		CounterpartyChannelId: counterpartyChannelID,
	}
}

// ValidateBasic performs basic checks on a MsgProvideCounterparty.
func (msg *MsgProvideCounterparty) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return errorsmod.Wrapf(ibcerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}

	if err := host.ChannelIdentifierValidator(msg.ChannelId); err != nil {
		return err
	}

	if err := host.ChannelIdentifierValidator(msg.CounterpartyChannelId); err != nil {
		return err
	}

	return nil
}

// NewMsgCreateChannel creates a new MsgCreateChannel instance
func NewMsgCreateChannel(clientID string, merklePathPrefix commitmenttypes.MerklePath, signer string) *MsgCreateChannel {
	return &MsgCreateChannel{
		Signer:           signer,
		ClientId:         clientID,
		MerklePathPrefix: merklePathPrefix,
	}
}

// ValidateBasic performs basic checks on a MsgCreateChannel.
func (msg *MsgCreateChannel) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return errorsmod.Wrapf(ibcerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}

	if err := host.ClientIdentifierValidator(msg.ClientId); err != nil {
		return err
	}

	if err := msg.MerklePathPrefix.ValidateAsPrefix(); err != nil {
		return err
	}

	return nil
}

// NewMsgSendPacket creates a new MsgSendPacket instance.
func NewMsgSendPacket(sourceChannel string, timeoutTimestamp uint64, signer string, packetData ...PacketData) *MsgSendPacket {
	return &MsgSendPacket{
		SourceChannel:    sourceChannel,
		TimeoutTimestamp: timeoutTimestamp,
		PacketData:       packetData,
		Signer:           signer,
	}
}

// NewMsgRecvPacket creates a new MsgRecvPacket instance.
func NewMsgRecvPacket(packet Packet, proofCommitment []byte, proofHeight clienttypes.Height, signer string) *MsgRecvPacket {
	return &MsgRecvPacket{
		Packet:          packet,
		ProofCommitment: proofCommitment,
		ProofHeight:     proofHeight,
		Signer:          signer,
	}
}