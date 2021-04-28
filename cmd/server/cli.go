package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type CardanoCli struct {
	byron   Byron
	shelley Shelley
}

func (cli *CardanoCli) run(
	topologyFilePath string,
	databasePath string,
	socketPath string,
	configFilePath string,
	port string,
) {
	args := []string{
		"run",
		"--topology", topologyFilePath,
		"--database-path", databasePath,
		"--socket-path", socketPath,
		"--port", port,
		"--config", configFilePath,
	}
	cmd := exec.Command("cardano-node", args...)

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
}

type Byron struct {
	byronNode ByronNode
}

type ByronNode struct {
}

type Shelley struct {
	genesis      Genesis
	devops       Devops
	system       System
	block        Block
	query        Query
	stakePool    StakePool
	shelleynode  ShelleyNode
	transaction  Transaction
	stakeAddress StakeAddress
	address      Address
}

type Genesis struct{}

func (genesis *Genesis) Create(dir string) {
	args := []string{
		"shelley",
		"genesis",
		"create",
		"--genesis-dir", dir,
	}
	_, err := exec.Command("cardano-cli", args...).Output()
	if err != nil {
		log.Print(err)
	}
}

func (genesis *Genesis) KeyGenGenesis(verificationKeyFile string, signingKeyFile string) {
	args := []string{
		"shelley",
		"genesis",
		"key-gen-genesis",
		"--verification-key-file", verificationKeyFile,
		"--signing-key-file", signingKeyFile,
	}
	_, err := exec.Command("cardano-cli", args...).Output()
	if err != nil {
		log.Print(err)
	}
}

type Devops struct{}
type System struct{}
type Block struct{}
type Query struct{}

type QueryTipResponse struct {
	Epoch int
	Hash  string
	Slot  int
	Block int
}

func (query *Query) Tip() (QueryTipResponse, error) {
	args := []string{
		"query",
		"tip",
		"--mainnet",
	}
	out, err := exec.Command("cardano-cli", args...).Output()
	if err != nil {
		log.Print(err)
	}
	var queryTipResponse QueryTipResponse
	json.Unmarshal([]byte(out), &queryTipResponse)

	return queryTipResponse, nil
}

type ProtocolParameter struct {
	TxFeePerByte       int
	MinUTxOValue       int
	StakePoolDeposit   int
	Decentralization   int
	PoolRetireMaxEpoch int
	ExtraPraosEntropy  int
	StakePoolTargetNum int
	MaxBlockBodySize   int
	MaxTxSize          int
	TreasuryCut        string
	MinPoolCost        int
	MaxBlockHeaderSize int
	ProtocolVersion    struct {
		Minor int
		Major int
	}
	TxFeeFixed          int
	StakeAddressDeposit int
	MonetaryExpansion   string
	PoolPledgeInfluence string
}

func (query *Query) ProtocolParameters() (ProtocolParameter, error) {
	args := []string{
		"query",
		"protocol-parameters",
		"--out-file", "protocol.json",
	}
	out, err := exec.Command("cardano-cli", args...).Output()
	if err != nil {
		log.Print(err)
	}
	var protocolParameter ProtocolParameter
	json.Unmarshal([]byte(out), &protocolParameter)

	return protocolParameter, nil
}

type StakePool struct{}
type ShelleyNode struct{}
type Transaction struct{}
type StakeAddress struct{}

func (stakeAddress *StakeAddress) keyGen(
	verificationKeyFilePath string,
	signingKeyFilePath string,
) {
	args := []string{
		"stake-address",
		"key-gen",
		"--verification-key-file", "stake.vkey",
		"--signing-key-file", "stake.skey",
	}
	_, err := exec.Command("cardano-cli", args...).Output()
	if err != nil {
		log.Print(err)
	}
}

func (stakeAddress *StakeAddress) build(
	stakeVerificationKeyFilePath string,
	outFileOutputPath string,
) {
	args := []string{
		"stake-address",
		"build",
		"--verification-key-file", stakeVerificationKeyFilePath,
		"--out-file", outFileOutputPath,
	}
	_, err := exec.Command("cardano-cli", args...).Output()
	if err != nil {
		log.Print(err)
	}
}

type Address struct{}

func (address *Address) keyGen(
	verificationKeyFileOutputPath string,
	signingKeyFileOutputPath string,
) {
	args := []string{
		"address",
		"key-gen",
		"--verification-key-file", verificationKeyFileOutputPath,
		"--signing-key-file", signingKeyFileOutputPath,
	}
	_, err := exec.Command("cardano-cli", args...).Output()
	if err != nil {
		log.Print(err)
	}
}

func (address *Address) build(
	paymentVerificationKeyFilePath string,
	stakeVerificationKeyFilePath string,
	outFileOutputPath string,
) {
	args := []string{
		"address",
		"build",
		"--payment-verification-key-file", paymentVerificationKeyFilePath,
		"--stake-verification-key-file", stakeVerificationKeyFilePath,
		"--out-file", outFileOutputPath,
		"--mainnet",
	}
	_, err := exec.Command("cardano-cli", args...).Output()
	if err != nil {
		log.Print(err)
	}
}
