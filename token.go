// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package main

import (
	"errors"
	"math/big"
)

var token = map[string]interface{}{
	"code":           "PCI",
	"decimal":        "8",
	"max_supply":     "394100000000000000",
	"initial_supply": "394100000000000000",
	"fee":            "transfer=0.001,10000000000;pay=0.002", // 0.001 = 0.1%, 10000000000 = 100PCI, 0.002 = 0.2%
	// "target_address": "PCI010A7733130C356643AE671FFA962E9D34AD9A595EF6294D5E" // testnet
	"target_address": "PCI01639F0B0A494B6040CE8B5B0DC4C56ACA7E78F6BAB02271AF", // mainnet
	"wrap_bridge":    wrapBridge,
	// optional fields
	"description": "Payprotocol",
	"website":     "https://payprotocol.io/",
}

var wrapBridge = map[string]string{
	// "{{ext_code}}":"{{wrap_address}};{{ext_chain}}"
	// eg. "wpci": "PCI01AE18973BDC22C8834620E86B2EFF77988F9D340D840409D0;ETH"
	// "wpci": "PCI01AE18973BDC22C8834620E86B2EFF77988F9D340D840409D0;ETH",	// testnet
	"wpci": "PCI0147A0FFEA1F2A1FA5D08967D297D29CA495CEC34E5BDF11F3;ETH", // mainnet
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
