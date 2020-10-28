package config

import "github.com/spf13/viper"

// 数据库配置
type Database struct {
	MaxIdleConns			int
	MaxOpenConns    		int
	Driver                	string
	Source                	string
	User     				[]string
	Merchant 				[]string
	Lottery  				[]string
	Pay      				[]string
	Chat 					[]string
	Stat     				[]string
	Bet      				[]string
	Private					[]string
}

func InitDatabase(cfg *viper.Viper) *Database {

	db := &Database{
		MaxIdleConns : cfg.GetInt("MaxIdleConns"),
		MaxOpenConns : cfg.GetInt("MaxOpenConns"),
		Driver: cfg.GetString("driver"),
		Source: cfg.GetString("source"),
		User: cfg.GetStringSlice("user"),
		Merchant: cfg.GetStringSlice("merchant"),
		Lottery: cfg.GetStringSlice("lottery"),
		Pay: cfg.GetStringSlice("pay"),
		Chat: cfg.GetStringSlice("chat"),
		Stat: cfg.GetStringSlice("stat"),
		Bet: cfg.GetStringSlice("bet"),
		Private: cfg.GetStringSlice("private"),
	}
	return db
}

var DatabaseConfig *Database
