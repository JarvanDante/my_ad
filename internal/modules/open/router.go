package open

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/JarvanDante/my_ad/internal/modules/campaign/logic"
	"github.com/JarvanDante/my_ad/internal/modules/open/controller"
)

func Register(group *ghttp.RouterGroup, svc *logic.Service) {
	ctrl := controller.New(svc)
	group.Bind(ctrl.Fetch, ctrl.Event)
}
