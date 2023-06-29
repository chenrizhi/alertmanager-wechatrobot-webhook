package transformer

import (
	"bytes"
	"fmt"
	"os"

	"github.com/chenrizhi/alertmanager-wechatrobot-webhook/model"
)

// TransformToMarkdown transform alertmanager notification to WeChat markdown message
func TransformToMarkdown(notification model.Notification) (markdown *model.WeChatMarkdown, robotURL string, err error) {

	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["wechatRobot"]

	var buffer bytes.Buffer

	clusterEnv := os.Getenv("CLUSTER")
	if clusterEnv != "" {
		clusterEnv = fmt.Sprintf("[%s] ", clusterEnv)
	}
	buffer.WriteString(fmt.Sprintf("### %s当前状态: %s \n", clusterEnv, status))

	for _, alert := range notification.Alerts {
		labels := alert.Labels
		buffer.WriteString(fmt.Sprintf("> 告警级别: %s\n", labels["severity"]))
		buffer.WriteString(fmt.Sprintf("> 告警类型: %s\n", labels["alertname"]))
		buffer.WriteString(fmt.Sprintf("> 故障主机: %s\n", labels["instance"]))

		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("> 告警主题: %s\n", annotations["summary"]))
		buffer.WriteString(fmt.Sprintf("> 告警详情: %s\n", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("> 告警时间: %s\n", alert.StartsAt.Local().Format("2006-01-02 15:04:05")))
		if status == "resolved" {
			buffer.WriteString(fmt.Sprintf("> 恢复时间: %s\n", alert.EndsAt.Local().Format("2006-01-02 15:04:05")))
		}
		buffer.WriteString("\n")
	}

	markdown = &model.WeChatMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content: buffer.String(),
		},
	}

	return
}
