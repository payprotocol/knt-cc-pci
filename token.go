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
	"fee":            "transfer=0.001,10000000000;pay=0.002",
	"target_address": "PCI01639F0B0A494B6040CE8B5B0DC4C56ACA7E78F6BAB02271AF",
	// unnecessary fields
	"description": "Payprotocol",
	"website":     "https://payprotocol.io/",
	//wrap
	"wrap_bridge": wrapBridge,
}

var wrapBridge = map[string]string{
	// "{{ext_code}}":"{{wrap_address}};{{ext_chain}};{{fee}}"
	// eg. "wpci": "PCI0119043794E2DCB197059D2304E6AABB8501B32B8C6F116CB0;ETH;10000"
	// "wpci": "PCI01B6AC0313B49F37934225D53720E71C2B63872C9AC2BAAABF;ETH;0",	// testnet
	"wpci": "PCI0147A0FFEA1F2A1FA5D08967D297D29CA495CEC34E5BDF11F3;ETH;0", // mainnet
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
