package apolloadapter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/contrib/config/apollo/v2"
)

// Config is the configuration object for apollo client.
type AdapterApolloCustom struct {
	Client *apollo.Client
}

// 重写 Data(ctx)，自动解析 content
func (a *AdapterApolloCustom) Data(ctx context.Context) map[string]interface{} {
	data := a.Data(ctx)

	if content, ok := data["content"].(string); ok {
		var result map[string]interface{}
		if err := json.Unmarshal([]byte(content), &result); err != nil {
			panic(fmt.Sprintf("Apollo 配置解析失败: %v", err))
		}
		return result
	}
	return data
}
