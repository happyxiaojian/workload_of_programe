package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// 数据库配置项
var cfgDatabase *viper.Viper

// 应用配置项
var cfgApplication *viper.Viper
var cfgFrontend *viper.Viper
var cfgBackend *viper.Viper
var cfgJobs *viper.Viper
var cfgPrivateGame *viper.Viper
var cfgSaas *viper.Viper

// Token配置项
var cfgJwt *viper.Viper

// Log配置项
var cfgLogger *viper.Viper

// Ssl配置项 非必须
var cfgSsl *viper.Viper

// 代码生成配置项 非必须
var cfgGen *viper.Viper

//载入配置文件 路径 文件类型
func Setup(path string, suffix string) {
	fmt.Println("---- local Setup config path", path) // config/setting_common_prod.yml
	viper.SetConfigFile(path)
	viper.SetConfigType(suffix) // "yaml"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	// application
	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("No found settings.application in the configuration")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	// frontend 前台
	cfgFrontend = viper.Sub("settings.frontend")
	if cfgFrontend == nil {
		panic("No found settings.frontend in the configuration")
	}
	FrontendConfig = InitFrontend(cfgFrontend)

	// backend 后台
	cfgBackend = viper.Sub("settings.backend")
	if cfgBackend == nil {
		panic("No found settings.backend in the configuration")
	}
	BackendConfig = InitBackend(cfgBackend)

	// jobs 定时任务
	cfgJobs = viper.Sub("settings.jobs")
	if cfgJobs == nil {
		panic("No found settings.jobs in the configuration")
	}
	JobsConfig = InitJobs(cfgJobs)

	// 私有彩种配置
	cfgPrivateGame = viper.Sub("settings.private")
	if cfgPrivateGame == nil {
		panic("No found settings.private in the configuration")
	}
	PrivateGameConfig = InitPrivateGame(cfgPrivateGame)

	cfgLogger = viper.Sub("settings.logger")
	if cfgLogger == nil {
		panic("No found settings.logger in the configuration")
	}
	LoggerConfig = InitLog(cfgLogger)

	// database 数据库
	cfgDatabase = viper.Sub("settings.database")
	if cfgDatabase == nil {
		panic("No found settings.database in the configuration")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	// saas 云服务配置
	cfgSaas = viper.Sub("settings.saas")
	if cfgSaas == nil {
		panic("No found settings.saas in the configuration")
	}
	SaasConfig = InitSaas(cfgSaas)

	//cfgSsl = viper.Sub("settings.ssl")
	//if cfgSsl == nil {
	//	// Ssl不是系统强制要求的配置，默认可以不用配置，将设置为关闭状态
	//	fmt.Println("warning config not found settings.ssl in the configuration")
	//	SslConfig = new(Ssl)
	//	SslConfig.Enable = false
	//} else {
	//	SslConfig = InitSsl(cfgSsl)
	//}

	//cfgGen = viper.Sub("settings.gen")
	//if cfgGen == nil {
	//	panic("No found settings.gen")
	//}
	//GenConfig = InitGen(cfgGen)

}

// 获取配置文件绝对路径
func GetConfigAbsolutePath(fileName string, path string) (filePath string, err error){
	var pathFile string
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(dir, path)
	pathFile = s[0] + string(os.PathSeparator) + path + string(os.PathSeparator) + "config" + string(os.PathSeparator) + fileName
	return pathFile, err
}