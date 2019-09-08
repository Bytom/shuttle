package swap

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDeployHTLCContract(t *testing.T) {
	account := HTLCAccount{
		AccountID: "10CJPO1HG0A02",
		Password:  "12345",
		TxFee:     uint64(100000000),
	}
	contractArgs := HTLCContractArgs{
		SenderPublicKey:    "3e5d7d52d334964eef173021ef6a04dc0807ac8c41700fe718f5a80c2109f79e",
		RecipientPublicKey: "6ea28f3f1389efd6a731de070fb38ab69dc93dae6c73b6524bac901b662f601d",
		BlockHeight:        uint64(950),
		Hash:               "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
	}
	contractValue := AssetAmount{
		Asset:  "bae7e17bb8f5d0cfbfd87a92f3204da082d388d4c9b10e8dcd36b3d0a18ceb3a",
		Amount: uint64(20000000000),
	}
	contractUTXOID, err := DeployHTLCContract(account, contractValue, contractArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("contractUTXOID:", contractUTXOID)
}

func TestBuildUnlockHTLCContractTransaction(t *testing.T) {
	account := HTLCAccount{
		AccountID: "10CKAD3000A02",
		Password:  "12345",
		Receiver:  "00140fdee108543d305308097019ceb5aec3da60ec66",
		TxFee:     uint64(100000000),
	}
	contractUTXOID := "e5b3b14c03eaab17c21cc23b925309bd7b8f8ed85b3fd078e0170498f5e069c8"
	contractValue := AssetAmount{
		Asset:  "bae7e17bb8f5d0cfbfd87a92f3204da082d388d4c9b10e8dcd36b3d0a18ceb3a",
		Amount: uint64(20000000000),
	}
	buildTxResp, err := buildUnlockHTLCContractTransaction(account, contractUTXOID, contractValue)
	if err != nil {
		fmt.Println(err)
	}
	signingInst, err := json.Marshal(buildTxResp.SigningInstructions[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("raw transaction:", buildTxResp.RawTransaction)
	fmt.Println("signingInst:", string(signingInst))
	contractControlProgram, signData, err := decodeRawTransaction(buildTxResp.RawTransaction, contractValue)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("contractControlProgram:", contractControlProgram)
	fmt.Println("signData:", signData)

	// preimage := "68656c6c6f" // b'hello'.hex()
	// recipientSig := ""
	// signedTransaction, err := signUnlockHTLCContractTransaction(account, preimage, recipientSig, *buildTxResp)

}

func TestListAddresses(t *testing.T) {
	accountID := "10CJPO1HG0A02"
	addressInfos, err := listAddresses(accountID)
	if err != nil {
		fmt.Println(err)
	}
	controlProgram := "00145b0a81adc5c2d68a9967082a09c96e82d62aa058"
	for _, addressInfo := range addressInfos {
		if addressInfo.ControlProgram == controlProgram {
			fmt.Println("address:", addressInfo.Address)
		}
	}
}

func TestSignMessage(t *testing.T) {
	address := "sm1q828d7re2wp20kgx4zyrw4e049k4v0enwdadq40"
	message := "719b521e9f341b1a07edf3805f8a5c5f9de453b61eb6e60f14a4bf94fa2bf6bc"
	password := "12345"
	sig, err := signMessage(address, message, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("signature:", sig)
}

func TestCallHTLCContract(t *testing.T) {
	account := HTLCAccount{
		AccountID: "10CKAD3000A02",
		Password:  "12345",
		Receiver:  "00140fdee108543d305308097019ceb5aec3da60ec66",
		TxFee:     uint64(100000000),
	}
	contractUTXOID := "68b6f62e826d219fba4997794e1399014a2a093184ec01fcf9be9cd4bae892cd"
	preimage := "68656c6c6f"
	contractArgs := HTLCContractArgs{
		SenderPublicKey:    "3e5d7d52d334964eef173021ef6a04dc0807ac8c41700fe718f5a80c2109f79e",
		RecipientPublicKey: "198787c8380ed1ba6fec1f81bb68c17c16432c4bc646effe0a5fae4f1b528f16",
		BlockHeight:        uint64(950),
		Hash:               "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
	}
	contractValue := AssetAmount{
		Asset:  "bae7e17bb8f5d0cfbfd87a92f3204da082d388d4c9b10e8dcd36b3d0a18ceb3a",
		Amount: uint64(20000000000),
	}
	txID, err := CallHTLCContract(account, contractUTXOID, preimage, contractArgs, contractValue)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("txID:", txID)
}
