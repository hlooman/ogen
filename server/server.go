package server

import (
	"log"
	"path"

	"github.com/olympus-protocol/ogen/chain"
	"github.com/olympus-protocol/ogen/chainrpc"
	"github.com/olympus-protocol/ogen/config"
	"github.com/olympus-protocol/ogen/db/blockdb"
	"github.com/olympus-protocol/ogen/explorer"
	"github.com/olympus-protocol/ogen/gov"
	"github.com/olympus-protocol/ogen/logger"
	"github.com/olympus-protocol/ogen/mempool"
	"github.com/olympus-protocol/ogen/miner"
	"github.com/olympus-protocol/ogen/params"
	"github.com/olympus-protocol/ogen/peers"
	"github.com/olympus-protocol/ogen/primitives"
	"github.com/olympus-protocol/ogen/users"
	"github.com/olympus-protocol/ogen/wallet"
	"github.com/olympus-protocol/ogen/workers"
)

type Server struct {
	log    *logger.Logger
	config *config.Config
	params params.ChainParams

	Chain     *chain.Blockchain
	PeerMan   *peers.PeerMan
	Wallet    *wallet.Wallet
	Miner     *miner.Miner
	GovMan    *gov.GovMan
	WorkerMan *workers.WorkerMan
	UsersMan  *users.UserMan
	RPC       *chainrpc.Wallet
	Gui       bool
}

func (s *Server) Start() {
	if s.config.Wallet {
		err := s.Wallet.Start()
		if err != nil {
			log.Fatalf("unable to start wallet manager: %s", err)
		}
	}
	err := s.Chain.Start()
	if err != nil {
		log.Fatalln("unable to start chain instance")
	}
	err = s.PeerMan.Start()
	if err != nil {
		log.Fatalln("unable to start peer manager")
	}
	chainrpc.ServeRPC(s.RPC, loadRPCConfig(s.config, s.log))
	if s.Miner != nil {
		err = s.Miner.Start()
		if err != nil {
			log.Fatalln("unable to start miner thread")
		}
	}
	switch s.config.Mode {
	case "api":
		err := explorer.LoadApi(s.config, s.Chain, s.PeerMan)
		if err != nil {
			log.Fatal("unable to start api")
		}
	}
}

func (s *Server) Stop() error {
	s.Chain.Stop()
	s.PeerMan.Stop()
	if s.config.Wallet {
		err := s.Wallet.Stop()
		if err != nil {
			return err
		}
	}
	if s.Miner != nil {
		s.Miner.Stop()
	}
	return nil
}

func NewServer(configParams *config.Config, logger *logger.Logger, currParams params.ChainParams, db *blockdb.BlockDB, gui bool, ip primitives.InitializationParameters) (*Server, error) {
	logger.Tracef("loading network parameters for '%v'", params.NetworkNames[configParams.NetworkName])
	coinsMempool := mempool.NewCoinsMempool()
	ch, err := chain.NewBlockchain(loadChainConfig(configParams, logger), currParams, db, ip, coinsMempool)
	if err != nil {
		return nil, err
	}
	w, err := wallet.NewWallet(loadWalletsManConfig(configParams, logger), currParams, ch)
	if err != nil {
		return nil, err
	}
	rpc := chainrpc.NewRPCWallet(w)
	voteMempool := mempool.NewVoteMempool(&currParams)
	peersMan, err := peers.NewPeersMan(loadPeersManConfig(configParams, logger), currParams, ch, voteMempool)
	if err != nil {
		return nil, err
	}
	var min *miner.Miner
	if configParams.MiningEnabled {
		min, err = miner.NewMiner(loadMinerConfig(configParams, logger), currParams, ch, w, peersMan, voteMempool, coinsMempool)
		if err != nil {
			return nil, err
		}
	}
	workersMan := workers.NewWorkersMan(loadWorkersConfig(configParams, logger), currParams)
	govMan := gov.NewGovMan(loadGovConfig(configParams, logger), currParams)
	usersMan := users.NewUsersMan(loadUsersConfig(configParams, logger), currParams)
	s := &Server{
		config: configParams,
		log:    logger,

		Chain:     ch,
		PeerMan:   peersMan,
		Wallet:    w,
		Miner:     min,
		WorkerMan: workersMan,
		GovMan:    govMan,
		UsersMan:  usersMan,
		Gui:       gui,
		RPC:       rpc,
	}
	return s, nil
}

func loadGovConfig(config *config.Config, logger *logger.Logger) gov.Config {
	cfg := gov.Config{
		Log: logger,
	}
	return cfg
}

func loadUsersConfig(config *config.Config, logger *logger.Logger) users.Config {
	cfg := users.Config{
		Log: logger,
	}
	return cfg
}

func loadWorkersConfig(config *config.Config, logger *logger.Logger) workers.Config {
	cfg := workers.Config{
		Log: logger,
	}
	return cfg
}

func loadChainConfig(config *config.Config, logger *logger.Logger) chain.Config {
	cfg := chain.Config{
		Log: logger,
	}
	return cfg
}

func loadMinerConfig(config *config.Config, logger *logger.Logger) miner.Config {
	cfg := miner.Config{
		Log: logger,
	}
	return cfg
}

func loadPeersManConfig(config *config.Config, logger *logger.Logger) peers.Config {
	cfg := peers.Config{
		Log:          logger,
		Listen:       config.Listen,
		ConnectNodes: config.ConnectNodes,
		AddNodes:     config.AddNodes,
		Port:         config.Port,
		MaxPeers:     config.MaxPeers,
		Path:         config.DataFolder,
	}
	return cfg
}

func loadWalletsManConfig(config *config.Config, logger *logger.Logger) wallet.Config {
	cfg := wallet.Config{
		Log:  logger,
		Path: path.Join(config.DataFolder, "wallet"),
	}
	return cfg
}

func loadRPCConfig(config *config.Config, logger *logger.Logger) chainrpc.Config {
	return chainrpc.Config{
		Log:     logger,
		Address: config.RPCAddress,
		Network: "tcp",
	}
}
