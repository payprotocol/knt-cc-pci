// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package main

import (
	"errors"
	"math/big"
)

var token = map[string]string{
	"code":           "PCI",
	"decimal":        "8",
	"max_supply":     "394100000000000000",
	"initial_supply": "394100000000000000",
	"fee":            "transfer=0.001,10000000000;pay=0.002",
	// "target_address": "PCI01639F0B0A494B6040CE8B5B0DC4C56ACA7E78F6BAB02271AF", //PROD
	// "wpci_address":   "{{wrap_address}}", //PROD must check address
	// "target_address": "PCI01CF0496510A7CD4890A0BF85F2ED0DFFF61C008A40EBAB45C", //local
	// "wpci_address":   "PCI01AA090DB75243A3892E3422AD6EEC5AE314D8399275F325F2", //local
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
