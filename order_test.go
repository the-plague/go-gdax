package gdax

import (
	"errors"
	"testing"
	"time"
)

const (
	orderCycles = 2
)

// Parameters for test orders. Constants at GDAX precision (crude fix to get tests working.)
const (
	limitSize = "1.00000000"
	limitPrice = "1.00000000"
	
	bigLimitSize = "10000.00000000"
	
	marketSize = "2.00000000"
	marketFunds = "10.00"
)

func TestCreateLimitOrders(t *testing.T) {
	client := NewTestClient()

	order := Order{
		Price:     limitPrice,
		Size:      limitSize,
		Side:      "buy",
		ProductId: "BTC-USD",
	}

	savedOrder, err := client.CreateOrder(&order)
	if err != nil {
		t.Error(err)
	}

	if savedOrder.Id == "" {
		t.Error(errors.New("No create id found"))
	}

	props := []string{"Price", "Size", "Side", "ProductId"}
	_, err = CompareProperties(order, savedOrder, props)
	if err != nil {
		t.Error(err)
	}

	if err := client.CancelOrder(savedOrder.Id); err != nil {
		t.Error(err)
	}
}

func TestCreateMarketOrders(t *testing.T) {
	client := NewTestClient()

	order := Order{
		Funds:     marketFunds,
		Size:      marketSize,
		Side:      "buy",
		Type:      "market",
		ProductId: "BTC-USD",
	}

	savedOrder, err := client.CreateOrder(&order)
	if err != nil {
		t.Error(err)
	}

	if savedOrder.Id == "" {
		t.Error(errors.New("No create id found"))
	}

	props := []string{"Price", "Size", "Side", "ProductId"}
	_, err = CompareProperties(order, savedOrder, props)
	if err != nil {
		t.Error(err)
	}
}

func TestCancelOrder(t *testing.T) {
	client := NewTestClient()

	order := Order{
		Price:     "1.00",
		Size:      "1000.00",
		Side:      "buy",
		ProductId: "BTC-USD",
	}

	savedOrder, err := client.CreateOrder(&order)
	if err != nil {
		t.Error(err)
	}

	if err := client.CancelOrder(savedOrder.Id); err != nil {
		t.Error(err)
		t.Error(err)
	}
}

func TestGetOrder(t *testing.T) {
	client := NewTestClient()

	order := Order{
		Price:     "1.00",
		Size:      "1.00",
		Side:      "buy",
		ProductId: "BTC-USD",
	}

	savedOrder, err := client.CreateOrder(&order)
	if err != nil {
		t.Error(err)
	}

	getOrder, err := client.GetOrder(savedOrder.Id)
	if err != nil {
		t.Error(err)
	}

	if getOrder.Id != savedOrder.Id {
		t.Error(errors.New("Order ids do not match"))
	}

	if err := client.CancelOrder(savedOrder.Id); err != nil {
		t.Error(err)
	}
}

func TestListOrders(t *testing.T) {
	client := NewTestClient()
	cursor := client.ListOrders()
	var orders []Order

	for cursor.HasMore {
		if err := cursor.NextPage(&orders); err != nil {
			t.Error(err)
		}

		for _, o := range orders {
			if StructHasZeroValues(o) {
				t.Error(errors.New("Zero value"))
			}
		}
	}

	cursor = client.ListOrders(ListOrdersParams{Status: "open", ProductId: "LTC-EUR"})
	for cursor.HasMore {
		if err := cursor.NextPage(&orders); err != nil {
			t.Error(err)
		}

		for _, o := range orders {
			if StructHasZeroValues(o) {
				t.Error(errors.New("Zero value"))
			}
		}
	}
}

func TestCancelAllOrders(t *testing.T) {
	client := NewTestClient()

	for _, pair := range []string{"BTC-USD", "ETH-USD", "LTC-USD"} {
		
		// Create orders
		for i := 0; i < orderCycles; i++ {
			order := Order{Price: limitPrice,
			       Size: bigLimitSize,
			       Side: "buy", 
			       ProductId: pair}
			
			if _, err := client.CreateOrder(&order); err != nil {
				t.Error(err)
			}
			
			// Wait a second between requests to avoid running into rate limits
			time.Sleep(time.Second)
		}

		// Attempt to cancel all outstanding orders for this currency
		orderIDs, err := client.CancelAllOrders(CancelAllOrdersParams{ProductId: pair})
		if err != nil {
			t.Error(err)
		}

		if len(orderIDs) != orderCycles {
			t.Error("Did not cancel all orders")
		}
	}
}
