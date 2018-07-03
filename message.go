package gdax

import (
	"encoding/json"
)

type Message struct {
	Type          string           `json:"type"`
	ProductId     string           `json:"product_id,omitempty"`
	ProductIds    []string         `json:"product_ids,omitempty"`
	TradeId       int              `json:"trade_id,number,omitempty"`
	OrderId       string           `json:"order_id,omitempty"`
	Sequence      int64            `json:"sequence,number,omitempty"`
	MakerOrderId  string           `json:"maker_order_id,omitempty"`
	TakerOrderId  string           `json:"taker_order_id,omitempty"`
	Time          Time             `json:"time,string,omitempty"`
	RemainingSize string           `json:"remaining_size,omitempty"`
	NewSize       string           `json:"new_size,omitempty"`
	OldSize       string           `json:"old_size,omitempty"`
	Size          string           `json:"size,omitempty"`
	Price         string           `json:"price,omitempty"`
	Side          string           `json:"side,omitempty"`
	Reason        string           `json:"reason,omitempty"`
	OrderType     string           `json:"order_type,omitempty"`
	Funds         string           `json:"funds,omitempty"`
	NewFunds      string           `json:"new_funds,omitempty"`
	OldFunds      string           `json:"old_funds,omitempty"`
	Message       string           `json:"message,omitempty"`
	Bids          []SnapshotEntry  `json:"bids,omitempty"`
	Asks          []SnapshotEntry  `json:"asks,omitempty"`
	Changes       []SnapshotChange `json:"changes,omitempty"`
	LastSize      string           `json:"last_size,omitempty"`
	BestBid       string           `json:"best_bid,omitempty"`
	BestAsk       string           `json:"best_ask,omitempty"`
	Channels      []MessageChannel `json:"channels,omitempty"`
	UserId        string           `json:"user_id,omitempty"`
	ProfileId     string           `json:"profile_id,omitempty"`

	// Fields added for ticker message support.
	Open24h   string `json:"open_24h,omitempty"`
	High24h   string `json:"high_24h,omitempty"`
	Low24h    string `json:"low_24h,omitempty"`
	Volume24h string `json:"volume_24h,omitempty"`
	Volume30d string `json:"volume_30d,omitempty"`

	// Fields added for heartbeat message support.
	LastTradeId int `json:"last_trade_id,number,omitempty"`
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
