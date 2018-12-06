// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package main

import (
	"errors"
	"math/big"
)

var token = map[string]string{
	"code":           "PCL",
	"decimal":        "18",
	"max_supply":     "100000000000",
	"initial_supply": "20000000000",
	// unnecessary fields
	"description": "Pay Protocol powered by Key Inside",
	"website":     "",
}

func mintableAmount(totalSupply *big.Int, balance *big.Int, amount *big.Int) (*big.Int, error) {
	if amount.Sign() <= 0 {
		return nil, errors.New("invalid amount to mint")
	}
	return amount, nil
}

func burnableAmount(totalSupply *big.Int, balance *big.Int, amount *big.Int) (*big.Int, error) {
	if amount.Sign() <= 0 {
		return nil, errors.New("invalid amount to burn")
	}
	return amount, nil
}
