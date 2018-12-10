// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package main

import (
	"errors"
	"math/big"
)

var token = map[string]string{
	"code":           "PCL",
	"decimal":        "2",
	"max_supply":     "10000000000000", // 100 billion
	"initial_supply": "200000000000",   // 2 million
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
