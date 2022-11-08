package config

import (
	"github.com/zeromicro/go-zero/core/logx"
)

type Config struct {
	Postgres struct {
		DataSource string
	}
	ChainConfig struct {
		NetworkRPCSysConfigName string
		StartL1BlockHeight      int64
		ConfirmBlocksCount      uint64
		MaxHandledBlocksCount   int64
		KeptHistoryBlocksCount  int64 // KeptHistoryBlocksCount define the count of blocks to keep in table, old blocks will be cleaned
	}
	LogConf logx.LogConf
}

func (c Config) Validate() {
	if c.ChainConfig.StartL1BlockHeight <= 0 || c.ChainConfig.MaxHandledBlocksCount <= 0 || c.ChainConfig.KeptHistoryBlocksCount <= 0 {
		logx.Error("invalid chain config")
		panic("invalid chain config")
	}
}
