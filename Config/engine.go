package Config

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"techku/Constant"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func GetEnvironment(env string) Config {
	_, filename, _, _ := runtime.Caller(1)
	envPath := path.Join(path.Dir(filename), Constant.ENVIRONMENT_PATH+env+".yml")
	_, err := os.Stat(envPath)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	content, err := ioutil.ReadFile(envPath)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	var config envFile = content
	return config
}

func (e envFile) LoadConfig() *ConfigSetting {

	var config Environment

	err := yaml.Unmarshal([]byte(string(e)), &config)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	if !config.App.Debug {
		log.SetOutput(ioutil.Discard)
	}
	log.Println("Environment Config load successfully!")
	return &ConfigSetting{&config, nil, &config.App, &config}
}

func (e *Environment) BuildConnection() DbConInterface {
	var connectionPool connectionPool = &database{}
	var dbCon DbCon
	for i := 0; i < len(e.Databases); i++ {
		connectionPool = &e.Databases[i]
		switch e.Databases[i].Engine {
		case Constant.POSTGRES:
			con := sql.DB{}
			log.Println("ENGINE " + Constant.POSTGRES + " start....")
			connectionPool.PostgresConnectionPool(&con)
			dbCon.setSqlConnection(DbSqlConfigName(e.Databases[i].Connection), &con)
		}
	}
	return &dbCon
}

func (app *app) Run(route *gin.Engine) {
	//run with https or http
	if app.Service == "https" {
		app.runWithHttps(route)
	}
	app.runWithHttp(route)
}
func (app *app) runWithHttp(route *gin.Engine) {

	log.Println("Http Service running ....")
	address := app.Host + ":" + app.Port
	err := route.Run(address)
	if err != nil {
		//panic error
		panic(err)
	}
}

func (app *app) runWithHttps(route *gin.Engine) {

	log.Println("Https Service running ....")
	address := app.Host + ":" + app.Port
	// Setting Directory Path Key and Pem File
	_, filename, _, _ := runtime.Caller(1)
	filepathKey := path.Join(path.Dir(filename), "../Infrastructures/certificate/"+app.Pem_key)
	filepathCert := path.Join(path.Dir(filename), "../Infrastructures/certificate/"+app.Certificate)
	// Setting Listen TLS
	err := route.RunTLS(address, filepathCert, filepathKey)
	if err != nil {
		log.Println(err.Error())
		//panic error
		panic(err)
	}
}
