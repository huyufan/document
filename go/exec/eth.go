package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// 1. 生成随机助记词（Mnemonic）
	entropy, err := bip39.NewEntropy(128) // 128-bit 生成 12 个单词
	if err != nil {
		log.Fatalf("生成 Entropy 失败: %v", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatalf("生成助记词失败: %v", err)
	}
	fmt.Println("助记词:", mnemonic)

	// 2. 生成种子（Seed）
	seed := bip39.NewSeed(mnemonic, "")

	// 3. 使用 HD 钱包库
	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatalf("创建 HD 钱包失败: %v", err)
	}

	// 4. 派生路径（BIP44 规范）
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatalf("派生账户失败: %v", err)
	}

	// 5. 获取私钥
	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		log.Fatalf("获取私钥失败: %v", err)
	}
	privateKeyHex := hex.EncodeToString(crypto.FromECDSA(privateKey))
	fmt.Println("私钥:", privateKeyHex)

	// 6. 获取公钥
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	publicKeyHex := hex.EncodeToString(crypto.FromECDSAPub(publicKey))
	fmt.Println("公钥:", publicKeyHex)

	// 7. 生成以太坊地址
	address := crypto.PubkeyToAddress(*publicKey).Hex()
	fmt.Println("以太坊地址:", address)
}
