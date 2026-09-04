// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"oa.98ent.com/p9/platform-base/pkg/i18n"
)

type Config struct {
	rest.RestConf
	I18n i18n.Config // i18n 国际化配置
}
