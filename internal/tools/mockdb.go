package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"andre":{
		AuthToken: "123ABC",
		Username: "andre",
	},
	"jason": {
		AuthToken: "456DEF",
		Username: "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username: "marie",
	},
} 

var mockCoinDetails = map[string]CoinDetails{
	"andre": {
		Coins: 100,
		Username: "andre",
	},
	"jason": {
		Coins: 200,
		Username: "jason",
	},
	"marie": {
		Coins: 330,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}