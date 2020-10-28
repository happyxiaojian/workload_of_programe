package config

import (
	"github.com/spf13/viper"
)

// 主应用配置
type Application struct {
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Port          string
	Name          string
	Mode          string
	Debug 		  bool
	DemoMsg       string
	UploadDir	  string

	CollectionPriHappyUrl string
	CollectionPriSpeedUrl string
	CollectionPriLuckyUrl string
	CollectionPriTopSpeedUrl string
	CollectionPriUGUrl string
	PackShell string
	CPURL	[]string

	ChatRoomApiAddr string

	// game api
	AllWallets string
	RefreshBalance string
	P2mWallet string
	All2mWallet string
	CheckOrder string
	LogPath string

	AppName string
	HttpPort int
	RunMode string
	AutoRender bool
	CopyRequestBody bool
	EnableDocs bool
	EnableAdmin bool
	AdminAddr string
	AdminPort int

}
func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		ReadTimeout:   cfg.GetInt("readTimeout"),
		WriterTimeout: cfg.GetInt("writerTimeout"),
		Host:          cfg.GetString("host"),
		Port:          cfg.GetString("port"),
		Name:          cfg.GetString("name"),
		Mode:          cfg.GetString("mode"),
		Debug:         cfg.GetBool("debug"),
		DemoMsg:       cfg.GetString("demoMsg"),

		UploadDir:      			cfg.GetString("uploadDir"),
		PackShell:					cfg.GetString("packShell"),

		CollectionPriHappyUrl:      cfg.GetString("collectionPriHappyUrl"),
		CollectionPriSpeedUrl:      cfg.GetString("collectionPriSpeedUrl"),
		CollectionPriLuckyUrl:      cfg.GetString("collectionPriLuckyUrl"),
		CollectionPriTopSpeedUrl:	cfg.GetString("collectionPriTopSpeedUrl"),
		CollectionPriUGUrl:			cfg.GetString("collectionPriUGUrl"),

		ChatRoomApiAddr:      		cfg.GetString("chatRoomApiAddr"),
		CPURL:	cfg.GetStringSlice("CPURL"),

		// game api
		AllWallets: 				cfg.GetString("all_wallets"),
		RefreshBalance: 		    cfg.GetString("refresh_balance"),
		P2mWallet: 					cfg.GetString("p_2_m_wallet"),
		All2mWallet: 				cfg.GetString("all_2_m_wallet"),
		CheckOrder: 				cfg.GetString("check_order"),
		LogPath: 					cfg.GetString("logPath"),

		// beego配置
		AppName: 					cfg.GetString("appName"),
		HttpPort: 		    		cfg.GetInt("httpPort"),
		RunMode: 					cfg.GetString("runMode"),
		AutoRender: 				cfg.GetBool("autoRender"),
		CopyRequestBody: 			cfg.GetBool("copyRequestBody"),
		EnableDocs: 				cfg.GetBool("enableDocs"),
		EnableAdmin: 				cfg.GetBool("enableAdmin"),
		AdminAddr: 					cfg.GetString("adminAddr"),
		AdminPort: 					cfg.GetInt("adminPort"),
	}
}
var ApplicationConfig *Application


// 前台配置
type Frontend struct {
	Port string
}
func InitFrontend(cfg *viper.Viper) *Frontend {
	return &Frontend{
		Port: cfg.GetString("port"),
	}
}
var FrontendConfig = new(Frontend)


// 后台配置
type Backend struct {
	Port string
}

func InitBackend(cfg *viper.Viper) *Backend {
	return &Backend{
		Port: cfg.GetString("port"),
	}
}

var BackendConfig = new(Backend)

// 定时任务
type Jobs struct {
	Port string
}

func InitJobs(cfg *viper.Viper) *Jobs {
	return &Jobs{
		Port: cfg.GetString("Port"),
	}
}

var JobsConfig = new(Jobs)

// 私彩配置
type PrivateGame struct {
	FrontendPort string
	BackendPort string
}

func InitPrivateGame(cfg *viper.Viper) *PrivateGame {
	return &PrivateGame{
		FrontendPort: 	cfg.GetString("FrontendPort"),
		BackendPort: 	cfg.GetString("BackendPort"),
	}
}

var PrivateGameConfig = new(PrivateGame)

// saas 软件即服务
type Saas struct {
	RedisAddr string
	RedisPasswd string
	RedisDbnum int
	RedisProto string
	RedisKeypairs string
	RabbitmqQmqp string
	OosToken string
}
func InitSaas(cfg *viper.Viper) *Saas {
	return &Saas{
		RedisAddr: 		cfg.GetString("redisAddr"),
		RedisPasswd: 	cfg.GetString("redisPasswd"),
		RedisDbnum: 	cfg.GetInt("redisDbnum"),
		RedisProto: 	cfg.GetString("redisProto"),
		RedisKeypairs: 	cfg.GetString("redisKeypairs"),
		RabbitmqQmqp: 	cfg.GetString("rabbitmqQmqp"),
		OosToken: 		cfg.GetString("oosToken"),
	}
}
var SaasConfig = new(Saas)