package main

import (
	"encoding/json"
	"io/ioutil"
)

type Actions struct{}

func CreateTopology(topology TopologyFile, fileName string) {
	file, _ := json.MarshalIndent(topology, "", " ")
	_ = ioutil.WriteFile(fileName, file, 0644)
}

func (action *Actions) CreateNetwork() {
	cli := CardanoCli{}
	cli.shelley.genesis.Create("network-config/")
	cli.shelley.genesis.KeyGenKes(
		"network-config/node/kes.vkey",
		"network-config/node/kes.skey",
	)
	cli.shelley.genesis.KeyGenVrf(
		"network-config/node/vrf.vkey",
		"network-config/node/vrf.skey",
	)
	cli.shelley.node.IssueOpCert(
		"network-config/node/kes.vkey",
		"network-config/delegate-keys/delagate1.skey",
		"network-config/delegate-keys/delegate-opcert1.counter",
		"0",
		"network-config/node/cert",
	)

	producer := Producer{
		Addr:    "127.0.0.1",
		Port:    3001,
		Valency: 1,
	}

	topologyFilePath := "network-config/node/topology.json"
	CreateTopology(TopologyFile{producer}, topologyFilePath)
	cli.run(topologyFilePath, "db/", "db/node.socket", "network-config/node/config.json", "3001")
}

func (action *Actions) SendTransaction(from string, to string) error {
	cli := CardanoCli{}
	protocolParameters, err := cli.shelley.query.ProtocolParameters()
	if err != nil {
		return err
	}
}
