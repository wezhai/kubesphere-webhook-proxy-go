package api

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wezhai/kubesphere-webhook-proxy-go/config"
	"github.com/wezhai/kubesphere-webhook-proxy-go/logger"
	"github.com/wezhai/kubesphere-webhook-proxy-go/qyapi"
)

func GwStat(c *fiber.Ctx) error {
	stat_msg := "请求成功"
	return c.SendString(stat_msg)
}

func GwWoker(c *fiber.Ctx) error {
	hook := new(Hook)
	if err := c.BodyParser(hook); err != nil {
		errmsg := "解析Json数据失败"
		logger.Error(errmsg)
		return c.SendString(errmsg)
	}
	logger.Debug(string(c.Body()))
	severityMap := map[string]string{
		"critical": "危险告警",
		"error":    "重要告警",
		"warning":  "一般告警",
	}

	if len(hook.Alerts) == 0 {
		errmsg := "未获取到告警详情"
		logger.Error(errmsg)
		return c.SendString(errmsg)
	}

	firstAlert := hook.Alerts[0]
	aliasName := firstAlert.Annotations.AliasName
	alertName := hook.CommonLables.Alertname
	if aliasName != "" {
		alertName = fmt.Sprintf("%v(%v)", alertName, aliasName)
	}
	startAt, _ := ConvertDatetime(firstAlert.StartsAt)
	kind := firstAlert.Annotations.Kind
	if kind == "" {
		msg := "仅支持Node和Deployment类型告警"
		logger.Warn(msg)
		return c.SendString(msg)
	}
	var resources []string
	switch kind {
	case "Deployment":
		for _, v := range hook.Alerts {
			resources = append(resources, strings.Split(v.Labels.Workload, ":")[1])
		}
	case "Node":
		for _, v := range hook.Alerts {
			resources = append(resources, v.Labels.Node)
		}
	}
	var content string
	content = fmt.Sprintf("您有<font color=\"warning\">**%d条**</font>告警需要关注\n", len(hook.Alerts))
	content = fmt.Sprintf("%s>策略名称: %s\n", content, alertName)
	content = fmt.Sprintf("%s>告警级别: %s\n", content, severityMap[hook.CommonLables.Severity])
	content = fmt.Sprintf("%s>资源类型: %s\n", content, firstAlert.Annotations.Kind)
	if hook.CommonLables.Namespace != "" {
		content = fmt.Sprintf("%s>命名空间: %s\n", content, hook.CommonLables.Namespace)
	}
	content = fmt.Sprintf("%s>告警对象: %s\n", content, SpliceSlice(resources, ","))
	content = fmt.Sprintf("%s>告警详情: %s\n", content, firstAlert.Annotations.Summary)
	content = fmt.Sprintf("%s>开始时间: %s\n", content, startAt.Format("2006-01-02 15:04:05"))
	content = fmt.Sprintf("%s请[登录Kubesphere](%s)查看告警详情", content, config.Config.KubesphereUrl)

	err := qyapi.SendMarkdownMessage(c, content)
	if err != nil {
		return c.SendString(fmt.Sprintf("%s", err))
	}
	return c.SendString("请求成功")
}
