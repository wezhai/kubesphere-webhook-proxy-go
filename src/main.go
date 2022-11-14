package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wezhai/kubesphere-webhook-proxy-go/api"

	"github.com/gofiber/fiber/v2/middleware/logger"
	// logger "github.com/wezhai/kubesphere-webhook-proxy-go/logger"
)

var (
	BuildTime string
	GoVersion string
	GitHead   string
)

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Shanghai",
	}))
	app.Get("/", api.GwStat)
	app.Post("/alert", api.GwWoker)

	ListenAddress := "0.0.0.0:5200"
	app.Listen(ListenAddress)
}

func init() {
	// 加载配置
	// err := config.LoadConfig()
	// if err != nil {
	// 	log.Println(err)
	// }
	buildInfo := fmt.Sprintf("BuildTime: %s\nGoVersion: %s\nGitHead: %s\n", BuildTime, GoVersion, GitHead)
	if BuildTime != "" {
		fmt.Print(buildInfo)
	}
}
