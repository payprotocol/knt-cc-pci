// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package main

import (
	"errors"
	"math/big"
)

var token = map[string]string{
	"code":           "PCI",
	"decimal":        "8",
	"max_supply":     "3941000000000000000",
	"initial_supply": "394100000000000000",
	"fee":            "transfer=0.1,10000;pay=-0.2",
	// unnecessary fields
	"description": "Payprotocol",
	"website":     "https://payprotocol.io/",
}

func mintableAmount(totalSupply *big.Int, balance *big.Int, amount *big.Int) (*big.Int, error) {
	if nil == amount || amount.Sign() <= 0 {
		return nil, errors.New("invalid amount to mint")
	}
	return amount, nil
}

func burnableAmount(totalSupply *big.Int, balance *big.Int, amount *big.Int) (*big.Int, error) {
	if nil == amount || amount.Sign() <= 0 {
		return nil, errors.New("invalid amount to burn")
	}
	return amount, nil
}
