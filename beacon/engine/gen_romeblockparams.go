// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package engine

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ = (*payloadAttributesMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (r RomePayloadAttributes) MarshalJSON() ([]byte, error) {
	type RomePayloadAttributes struct {
		Timestamp             hexutil.Uint64      `json:"timestamp"            gencodec:"required"`
		GasPrice              []uint64            `json:"gasPrices"            gencodec:"required"`
		GasUsed               []uint64            `json:"gasUsed"              gencodec:"required"`
		Random                common.Hash         `json:"prevRandao"            gencodec:"required"`
		SuggestedFeeRecipient common.Address      `json:"suggestedFeeRecipient" gencodec:"required"`
		Withdrawals           []*types.Withdrawal `json:"withdrawals"`
		BeaconRoot            *common.Hash        `json:"parentBeaconBlockRoot"`
		Transactions          []hexutil.Bytes     `json:"transactions,omitempty"  gencodec:"optional"`
		NoTxPool              bool                `json:"noTxPool,omitempty" gencodec:"optional"`
		GasLimit              *hexutil.Uint64     `json:"gasLimit,omitempty" gencodec:"optional"`
	}
	var enc RomePayloadAttributes
	enc.Timestamp = hexutil.Uint64(r.Timestamp)
	enc.GasPrice = r.GasPrice
	enc.GasUsed = r.GasUsed
	enc.Random = r.Random
	enc.SuggestedFeeRecipient = r.SuggestedFeeRecipient
	enc.Withdrawals = r.Withdrawals
	enc.BeaconRoot = r.BeaconRoot
	if r.Transactions != nil {
		enc.Transactions = make([]hexutil.Bytes, len(r.Transactions))
		for k, v := range r.Transactions {
			enc.Transactions[k] = v
		}
	}
	enc.NoTxPool = r.NoTxPool
	enc.GasLimit = (*hexutil.Uint64)(r.GasLimit)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *RomePayloadAttributes) UnmarshalJSON(input []byte) error {
	type RomePayloadAttributes struct {
		Timestamp             *hexutil.Uint64     `json:"timestamp"            gencodec:"required"`
		GasPrice              []uint64            `json:"gasPrices"            gencodec:"required"`
		GasUsed               []uint64            `json:"gasUsed"              gencodec:"required"`
		Random                *common.Hash        `json:"prevRandao"            gencodec:"required"`
		SuggestedFeeRecipient *common.Address     `json:"suggestedFeeRecipient" gencodec:"required"`
		Withdrawals           []*types.Withdrawal `json:"withdrawals"`
		BeaconRoot            *common.Hash        `json:"parentBeaconBlockRoot"`
		Transactions          []hexutil.Bytes     `json:"transactions,omitempty"  gencodec:"optional"`
		NoTxPool              *bool               `json:"noTxPool,omitempty" gencodec:"optional"`
		GasLimit              *hexutil.Uint64     `json:"gasLimit,omitempty" gencodec:"optional"`
	}
	var dec RomePayloadAttributes
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for RomePayloadAttributes")
	}
	r.Timestamp = uint64(*dec.Timestamp)
	if dec.GasPrice == nil {
		return errors.New("missing required field 'gasPrices' for RomePayloadAttributes")
	}
	r.GasPrice = dec.GasPrice
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for RomePayloadAttributes")
	}
	r.GasUsed = dec.GasUsed
	if dec.Random == nil {
		return errors.New("missing required field 'prevRandao' for RomePayloadAttributes")
	}
	r.Random = *dec.Random
	if dec.SuggestedFeeRecipient == nil {
		return errors.New("missing required field 'suggestedFeeRecipient' for RomePayloadAttributes")
	}
	r.SuggestedFeeRecipient = *dec.SuggestedFeeRecipient
	if dec.Withdrawals != nil {
		r.Withdrawals = dec.Withdrawals
	}
	if dec.BeaconRoot != nil {
		r.BeaconRoot = dec.BeaconRoot
	}
	if dec.Transactions != nil {
		r.Transactions = make([][]byte, len(dec.Transactions))
		for k, v := range dec.Transactions {
			r.Transactions[k] = v
		}
	}
	if dec.NoTxPool != nil {
		r.NoTxPool = *dec.NoTxPool
	}
	if dec.GasLimit != nil {
		r.GasLimit = (*uint64)(dec.GasLimit)
	}
	return nil
}
