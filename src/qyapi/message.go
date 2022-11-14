package qyapi

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wezhai/kubesphere-webhook-proxy-go/config"
	"github.com/wezhai/kubesphere-webhook-proxy-go/logger"
)

func SendMarkdownMessage(c *fiber.Ctx, content string) (err error) {
	key := c.Query("key")
	if key != "" && len(key) != 36 {
		err = errors.New("key不符合规范，请传入正确的key")
		logger.Error(fmt.Sprintln(err))
		return err
	}

	url := fmt.Sprintf("%v?key=%v&debug=1", config.Config.RobotBaseUrl, key)
	payload := MarkdownMessage{
		Msgtype: "markdown",
	}
	payload.Markdown.Content = content
	jsonstr, err := json.Marshal(payload)
	if err != nil {
		err = errors.New("解析数据失败")
		return err
	}
	_, err = request(url, jsonstr)
	return err
}
