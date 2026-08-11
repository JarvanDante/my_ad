package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/JarvanDante/my_ad/internal/boot"
	"github.com/JarvanDante/my_ad/internal/dao"
	"github.com/JarvanDante/my_ad/internal/modules/campaign"
	campaignlogic "github.com/JarvanDante/my_ad/internal/modules/campaign/logic"
	"github.com/JarvanDante/my_ad/internal/modules/client"
	"github.com/JarvanDante/my_ad/internal/modules/creative"
	creativelogic "github.com/JarvanDante/my_ad/internal/modules/creative/logic"
	"github.com/JarvanDante/my_ad/internal/modules/health"
	"github.com/JarvanDante/my_ad/internal/modules/open"
	"github.com/JarvanDante/my_ad/internal/modules/slot"
	slotlogic "github.com/JarvanDante/my_ad/internal/modules/slot/logic"
	"github.com/JarvanDante/my_ad/internal/shared/middleware"
)

// Main 广告中台 API(:8006)。
var Main = gcmd.Command{
	Name:  "adapi",
	Brief: "广告中台(PaaS) API",
	Func: func(ctx context.Context, parser *gcmd.Parser) error {
		boot.InitNacosConfig(ctx)
		slotSvc := slotlogic.New(dao.NewSlotRepo())
		creativeSvc := creativelogic.New(dao.NewCreativeRepo())
		campaignSvc := campaignlogic.New(dao.NewCampaignRepo(), dao.NewSlotRepo(), dao.NewCreativeRepo())
		clientRepo := dao.NewClientRepo()

		s := g.Server()
		s.Use(middleware.CORS, ghttp.MiddlewareHandlerResponse)
		s.BindStatusHandler(404, middleware.NotFound)

		health.Register(s.Group("/"))

		s.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.AdminToken)
			slot.RegisterAdmin(group, slotSvc)
			creative.RegisterAdmin(group, creativeSvc)
			campaign.RegisterAdmin(group, campaignSvc)
			client.RegisterAdmin(group, clientRepo)
		})

		s.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.AppKey)
			open.Register(group, campaignSvc)
		})

		s.Run()
		return nil
	},
}
