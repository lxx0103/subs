package cmd

import (
	"subs/api/v1/auth"
	"subs/api/v1/subscription"
	"subs/api/v1/user"
	"subs/core/cache"
	"subs/core/config"
	"subs/core/database"
	"subs/core/log"
	"subs/core/router"
)

func Run(args []string) {
	config.LoadConfig(args[1])
	log.ConfigLogger()
	cache.ConfigCache()
	database.ConfigMysql()
	r := router.InitRouter()
	router.InitPublicRouter(r, auth.Routers)
	router.InitAuthRouter(r, auth.AuthRouter, user.AuthRouter, subscription.AuthRouter) //, setting.AuthRouter)
	router.RunServer(r)
}
