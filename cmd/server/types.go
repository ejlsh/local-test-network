package main

import "time"

// CONFIG FILES

type ConfigFile struct {
	Applicationname                   string        `json:"ApplicationName"`
	Applicationversion                int           `json:"ApplicationVersion"`
	Byrongenesisfile                  string        `json:"ByronGenesisFile"`
	Byrongenesishash                  string        `json:"ByronGenesisHash"`
	LastknownblockversionAlt          int           `json:"LastKnownBlockVersion-Alt"`
	LastknownblockversionMajor        int           `json:"LastKnownBlockVersion-Major"`
	LastknownblockversionMinor        int           `json:"LastKnownBlockVersion-Minor"`
	Maxknownmajorprotocolversion      int           `json:"MaxKnownMajorProtocolVersion"`
	Protocol                          string        `json:"Protocol"`
	Requiresnetworkmagic              string        `json:"RequiresNetworkMagic"`
	Shelleygenesisfile                string        `json:"ShelleyGenesisFile"`
	Shelleygenesishash                string        `json:"ShelleyGenesisHash"`
	Traceblockfetchclient             bool          `json:"TraceBlockFetchClient"`
	Traceblockfetchdecisions          bool          `json:"TraceBlockFetchDecisions"`
	Traceblockfetchprotocol           bool          `json:"TraceBlockFetchProtocol"`
	Traceblockfetchprotocolserialised bool          `json:"TraceBlockFetchProtocolSerialised"`
	Traceblockfetchserver             bool          `json:"TraceBlockFetchServer"`
	Tracechaindb                      bool          `json:"TraceChainDb"`
	Tracechainsyncblockserver         bool          `json:"TraceChainSyncBlockServer"`
	Tracechainsyncclient              bool          `json:"TraceChainSyncClient"`
	Tracechainsyncheaderserver        bool          `json:"TraceChainSyncHeaderServer"`
	Tracechainsyncprotocol            bool          `json:"TraceChainSyncProtocol"`
	Tracednsresolver                  bool          `json:"TraceDNSResolver"`
	Tracednssubscription              bool          `json:"TraceDNSSubscription"`
	Traceerrorpolicy                  bool          `json:"TraceErrorPolicy"`
	Traceforge                        bool          `json:"TraceForge"`
	Tracehandshake                    bool          `json:"TraceHandshake"`
	Traceipsubscription               bool          `json:"TraceIpSubscription"`
	Tracelocalchainsyncprotocol       bool          `json:"TraceLocalChainSyncProtocol"`
	Tracelocalerrorpolicy             bool          `json:"TraceLocalErrorPolicy"`
	Tracelocalhandshake               bool          `json:"TraceLocalHandshake"`
	Tracelocaltxsubmissionprotocol    bool          `json:"TraceLocalTxSubmissionProtocol"`
	Tracelocaltxsubmissionserver      bool          `json:"TraceLocalTxSubmissionServer"`
	Tracemempool                      bool          `json:"TraceMempool"`
	Tracemux                          bool          `json:"TraceMux"`
	Tracetxinbound                    bool          `json:"TraceTxInbound"`
	Tracetxoutbound                   bool          `json:"TraceTxOutbound"`
	Tracetxsubmissionprotocol         bool          `json:"TraceTxSubmissionProtocol"`
	Tracingverbosity                  string        `json:"TracingVerbosity"`
	Turnonlogmetrics                  bool          `json:"TurnOnLogMetrics"`
	Turnonlogging                     bool          `json:"TurnOnLogging"`
	Defaultbackends                   []string      `json:"defaultBackends"`
	Defaultscribes                    [][]string    `json:"defaultScribes"`
	Hasekg                            int           `json:"hasEKG"`
	Hasprometheus                     []interface{} `json:"hasPrometheus"`
	Minseverity                       string        `json:"minSeverity"`
	Options                           struct {
		Mapbackends struct {
			CardanoNodeMetrics   []string `json:"cardano.node.metrics"`
			CardanoNodeResources []string `json:"cardano.node.resources"`
		} `json:"mapBackends"`
		Mapsubtrace struct {
			CardanoNodeMetrics struct {
				Subtrace string `json:"subtrace"`
			} `json:"cardano.node.metrics"`
		} `json:"mapSubtrace"`
	} `json:"options"`
	Rotation struct {
		Rpkeepfilesnum  int `json:"rpKeepFilesNum"`
		Rploglimitbytes int `json:"rpLogLimitBytes"`
		Rpmaxagehours   int `json:"rpMaxAgeHours"`
	} `json:"rotation"`
	Setupbackends []string `json:"setupBackends"`
	Setupscribes  []struct {
		Scformat   string      `json:"scFormat"`
		Sckind     string      `json:"scKind"`
		Scname     string      `json:"scName"`
		Scrotation interface{} `json:"scRotation"`
	} `json:"setupScribes"`
}

type ShelleyGenesisFile struct {
	Activeslotscoeff float64 `json:"activeSlotsCoeff"`
	Protocolparams   struct {
		Protocolversion struct {
			Minor int `json:"minor"`
			Major int `json:"major"`
		} `json:"protocolVersion"`
		Decentralisationparam int `json:"decentralisationParam"`
		Emax                  int `json:"eMax"`
		Extraentropy          struct {
			Tag string `json:"tag"`
		} `json:"extraEntropy"`
		Maxtxsize          int     `json:"maxTxSize"`
		Maxblockbodysize   int     `json:"maxBlockBodySize"`
		Maxblockheadersize int     `json:"maxBlockHeaderSize"`
		Minfeea            int     `json:"minFeeA"`
		Minfeeb            int     `json:"minFeeB"`
		Minutxovalue       int     `json:"minUTxOValue"`
		Pooldeposit        int     `json:"poolDeposit"`
		Minpoolcost        int     `json:"minPoolCost"`
		Keydeposit         int     `json:"keyDeposit"`
		Nopt               int     `json:"nOpt"`
		Rho                float64 `json:"rho"`
		Tau                float64 `json:"tau"`
		A0                 float64 `json:"a0"`
	} `json:"protocolParams"`
	Gendelegs struct {
		Ad5463153Dc3D24B9Ff133E46136028Bdc1Edbb897F5A7Cf1B37950C struct {
			Delegate string `json:"delegate"`
			Vrf      string `json:"vrf"`
		} `json:"ad5463153dc3d24b9ff133e46136028bdc1edbb897f5a7cf1b37950c"`
		B9547B8A57656539A8D9Bc42C008E38D9C8Bd9C8Adbb1E73Ad529497 struct {
			Delegate string `json:"delegate"`
			Vrf      string `json:"vrf"`
		} `json:"b9547b8a57656539a8d9bc42c008e38d9c8bd9c8adbb1e73ad529497"`
		Six0Baee25Cbc90047E83Fd01E1E57Dc0B06D3D0Cb150D0Ab40Bbfead1 struct {
			Delegate string `json:"delegate"`
			Vrf      string `json:"vrf"`
		} `json:"60baee25cbc90047e83fd01e1e57dc0b06d3d0cb150d0ab40bbfead1"`
		F7B341C14Cd58Fca4195A9B278Cce1Ef402Dc0E06Deb77E543Cd1757 struct {
			Delegate string `json:"delegate"`
			Vrf      string `json:"vrf"`
		} `json:"f7b341c14cd58fca4195a9b278cce1ef402dc0e06deb77e543cd1757"`
		One62F94554Ac8C225383A2248C245659Eda870Eaa82D0Ef25Fc7Dcd82 struct {
			Delegate string `json:"delegate"`
			Vrf      string `json:"vrf"`
		} `json:"162f94554ac8c225383a2248c245659eda870eaa82d0ef25fc7dcd82"`
		Two075A095B3C844A29C24317A94A643Ab8E22D54A3A3A72A420260Af6 struct {
			Delegate string `json:"delegate"`
			Vrf      string `json:"vrf"`
		} `json:"2075a095b3c844a29c24317a94a643ab8e22d54a3a3a72a420260af6"`
		Two68Cfc0B89E910Ead22E0Ade91493D8212F53F3E2164B2E4Bef0819B struct {
			Delegate string `json:"delegate"`
			Vrf      string `json:"vrf"`
		} `json:"268cfc0b89e910ead22e0ade91493d8212f53f3e2164b2e4bef0819b"`
	} `json:"genDelegs"`
	Updatequorum int    `json:"updateQuorum"`
	Networkid    string `json:"networkId"`
	Initialfunds struct {
	} `json:"initialFunds"`
	Maxlovelacesupply int64     `json:"maxLovelaceSupply"`
	Networkmagic      int       `json:"networkMagic"`
	Epochlength       int       `json:"epochLength"`
	Systemstart       time.Time `json:"systemStart"`
	Slotsperkesperiod int       `json:"slotsPerKESPeriod"`
	Slotlength        int       `json:"slotLength"`
	Maxkesevolutions  int       `json:"maxKESEvolutions"`
	Securityparam     int       `json:"securityParam"`
}

type Producer struct {
	Addr    string `json:"addr"`
	Port    int    `json:"port"`
	Valency int    `json:"valency"`
}

type TopologyFile struct {
	Producer `json:"Producers"`
}
