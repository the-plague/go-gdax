package gdax

import (
	"encoding/json"
)

type Message struct {
	Type string `json:"type" csv:"type"`

	ProductId  string   `json:"product_id,omitempty" csv:"product_id"`
	ProductIds []string `json:"product_ids,omitempty" csv:"-"`

	TradeId  int    `json:"trade_id,number,omitempty" csv:"trade_id"`
	OrderId  string `json:"order_id,omitempty" csv:"order_id"`
	Sequence int64  `json:"sequence,number,omitempty" csv:"sequence"`

	MakerOrderId string `json:"maker_order_id,omitempty" csv:"maker_order_id"`
	TakerOrderId string `json:"taker_order_id,omitempty" csv:"taker_order_id"`

	Time          Time             `json:"time,string,omitempty" csv:"time"`
	RemainingSize string           `json:"remaining_size,omitempty" csv:"remaining_size"`
	NewSize       string           `json:"new_size,omitempty" csv:"new_size"`
	OldSize       string           `json:"old_size,omitempty" csv:"old_size"`
	Size          string           `json:"size,omitempty" csv:"size"`
	Price         string           `json:"price,omitempty" csv:"price"`
	Side          string           `json:"side,omitempty" csv:"side"`
	Reason        string           `json:"reason,omitempty" csv:"reason"`
	OrderType     string           `json:"order_type,omitempty" csv:"order_type"`
	Funds         string           `json:"funds,omitempty" csv:"funds"`
	NewFunds      string           `json:"new_funds,omitempty" csv:"new_funds"`
	OldFunds      string           `json:"old_funds,omitempty" csv:"old_funds"`
	Message       string           `json:"message,omitempty" csv:"message"`
	Bids          []SnapshotEntry  `json:"bids,omitempty" csv:"-"`
	Asks          []SnapshotEntry  `json:"asks,omitempty" csv:"-"`
	Changes       []SnapshotChange `json:"changes,omitempty" csv:"-"`
	LastSize      string           `json:"last_size,omitempty" csv:"last_size"`
	BestBid       string           `json:"best_bid,omitempty" csv:"best_bid"`
	BestAsk       string           `json:"best_ask,omitempty" csv:"best_ask"`
	Channels      []MessageChannel `json:"channels,omitempty" csv:"-"`
	UserId        string           `json:"user_id,omitempty" csv:"user_id"`
	ProfileId     string           `json:"profile_id,omitempty" csv:"profile_id"`

	// Fields added for ticker message support.
	Open24h   string `json:"open_24h,omitempty" csv:"open_24h"`
	High24h   string `json:"high_24h,omitempty" csv:"high_24h"`
	Low24h    string `json:"low_24h,omitempty" csv:"low_24h"`
	Volume24h string `json:"volume_24h,omitempty" csv:"volume_24h"`
	Volume30d string `json:"volume_30d,omitempty" csv:"volume_30d"`

	// Fields added for heartbeat message support.
	LastTradeId int `json:"last_trade_id,number,omitempty" csv:"last_trade_id"`
}

type MessageChannel struct {
	Name       string   `json:"name"`
	ProductIds []string `json:"product_ids"`
}

type SnapshotChange struct {
	Side  string
	Price string
	Size  string
}

type SnapshotEntry struct {
	Price string
	Size  string
}

type SignedMessage struct {
	Message
	Key        string `json:"key"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Signature  string `json:"signature"`
}

func (e *SnapshotEntry) UnmarshalJSON(data []byte) error {
	var entry []string

	if err := json.Unmarshal(data, &entry); err != nil {
		return err
	}

	e.Price = entry[0]
	e.Size = entry[1]

	return nil
}

func (e *SnapshotChange) UnmarshalJSON(data []byte) error {
	var entry []string

	if err := json.Unmarshal(data, &entry); err != nil {
		return err
	}

	e.Side = entry[0]
	e.Price = entry[1]
	e.Size = entry[2]

	return nil
}
