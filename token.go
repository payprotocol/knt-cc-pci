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
	"fee":            "transfer=0.001,10000000000;pay=0.002;wrap=0",
	// "target_address": "PCI01639F0B0A494B6040CE8B5B0DC4C56ACA7E78F6BAB02271AF", //PROD
	// "wpci_address":   "{{wrap_address}}", //PROD must check address
	// "target_address": "PCI015DEE3F19CCDB7A652136A4B48B6BBD5758F30EAE7713DEB8", //local
	// "wpci_address":   "PCI014CAAB7A72208A519FA4A75D915F5265AA00D3200E1FD5A54", //local
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
