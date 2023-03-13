package Services

import (
	"flag"
	"github.com/gin-gonic/gin"
	"techku/Config"
	"techku/Constant"
	"techku/Controller"
	"techku/Controller/Dto"
	"techku/Modules"
	"techku/Modules/User"
	"techku/Repository"
	"techku/Routes"
	"techku/Routes/middleware"
	"techku/Services/jwt"
)

var AppEnv = flag.String("env", "", "define environment")
var routesConfig *gin.Engine

func init() {
	flag.Parse()
	gin.SetMode(gin.DebugMode)
	routesConfig = gin.Default()
	if *AppEnv == "" {
		*AppEnv = Constant.Localhost
	}
	if *AppEnv == Constant.Production {
		gin.SetMode(gin.ReleaseMode)
		routesConfig = gin.New()
	}
}

// AppInitialization Application Engine initialization
func AppInitialization() {
	newConfig := Config.GetEnvironment(*AppEnv).LoadConfig()
	connection := newConfig.Database.BuildConnection()
	service := serviceInit(newConfig.Environment)

	utilities := Dto.Utilities{
		Jwt: service.Jwt,
		Modules: Modules.Modules{
			UserModule: User.NewModules(Repository.InitRepo(connection)),
		},
	}
	newConfig.Routes = &Routes.Routes{
		Gin:        routesConfig,
		Controller: Controller.InitControllerApi(utilities),
		Middleware: &middleware.Middleware{TokenHeader: newConfig.Environment.ApiToken},
	}

	newConfig.Routes.SetCors()
	routes := newConfig.Routes.CollectRoutes()

	newConfig.HttpEngine.Run(routes)

}

//Initialization Application Services and Dependency
func serviceInit(Env *Config.Environment) service {
	serv := service{
		Jwt: jwt.JwtStruct{Config: &Env.Jwt},
	}
	return serv
}
