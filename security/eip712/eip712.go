package eip712

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_pool_registry"
)

// MessageHash calculates the hash for EIP-712 message.
func MessageHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

type ScannerNodeRegistration contract_scanner_pool_registry.ScannerPoolRegistryCoreScannerNodeRegistration

// SignScannerRegistration encodes registration data using EIP712
// typed structured data encoding rules and signs.
func SignScannerRegistration(
	scannerKey *ecdsa.PrivateKey, verifyingContract common.Address, reg *ScannerNodeRegistration,
) ([]byte, error) {
	data := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "version",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
				{
					Name: "verifyingContract",
					Type: "address",
				},
			},
			"ScannerNodeRegistration": {
				{
					Name: "scanner",
					Type: "address",
				},
				{
					Name: "scannerPoolId",
					Type: "uint256",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
				{
					Name: "metadata",
					Type: "string",
				},
				{
					Name: "timestamp",
					Type: "uint256",
				},
			},
		},
		Domain: apitypes.TypedDataDomain{
			Name:              "ScannerPoolRegistry",
			Version:           "1",
			ChainId:           (*math.HexOrDecimal256)(reg.ChainId),
			VerifyingContract: verifyingContract.Hex(),
		},
		PrimaryType: "ScannerNodeRegistration",
		Message: apitypes.TypedDataMessage{
			"scanner":       reg.Scanner.Hex(),
			"scannerPoolId": (*hexutil.Big)(reg.ScannerPoolId).String(),
			"chainId":       (*hexutil.Big)(reg.ChainId).String(),
			"metadata":      reg.Metadata,
			"timestamp":     (*hexutil.Big)(reg.Timestamp).String(),
		},
	}

	hash, err := hashTypedData(data)
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Sign(hash, scannerKey)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

func hashTypedData(data *apitypes.TypedData) ([]byte, error) {
	separator, err := data.HashStruct("EIP712Domain", data.Domain.Map())
	if err != nil {
		return nil, err
	}
	hash, err := data.HashStruct(data.PrimaryType, data.Message)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256([]byte(fmt.Sprintf("\x19\x01%s%s", string(separator), string(hash)))), nil
}