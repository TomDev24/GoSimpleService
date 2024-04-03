package model

import (
	"time"
)

type Order struct {
	CustomerID  string    `json:"customer_id"`
	DateCreated time.Time `json:"date_created"`
	Delivery    struct {
		Address string `json:"address"`
		City    string `json:"city"`
		Email   string `json:"email"`
		Name    string `json:"name" fake:"{firstname}"`
		Phone   string `json:"phone"`
		Region  string `json:"region"`
		Zip     string `json:"zip"`
	} `json:"delivery"`
	DeliveryService   string `json:"delivery_service"`
	Entry             string `json:"entry"`
	InternalSignature string `json:"internal_signature"`
	Items             []struct {
		Brand       string `json:"brand"`
		ChrtID      int    `json:"chrt_id"`
		Name        string `json:"name"`
		NmID        int    `json:"nm_id"`
		Price       int    `json:"price"`
		Rid         string `json:"rid"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		Status      int    `json:"status"`
		TotalPrice  int    `json:"total_price"`
		TrackNumber string `json:"track_number"`
	} `json:"items"`
	Locale   string `json:"locale"`
	OofShard string `json:"oof_shard"`
	OrderUid string `json:"order_uid"`
	Payment  struct {
		Amount       int    `json:"amount"`
		Bank         string `json:"bank"`
		Currency     string `json:"currency"`
		CustomFee    int    `json:"custom_fee"`
		DeliveryCost int    `json:"delivery_cost"`
		GoodsTotal   int    `json:"goods_total"`
		PaymentDt    int    `json:"payment_dt"`
		Provider     string `json:"provider"`
		RequestID    string `json:"request_id"`
		Transaction  string `json:"transaction"`
	} `json:"payment"`
	Shardkey    string `json:"shardkey"`
	SmID        int    `json:"sm_id"`
	TrackNumber string `json:"track_number" fake:"{number}"`
}
