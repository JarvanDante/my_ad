package slot

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/JarvanDante/my_ad/internal/modules/slot/controller"
	"github.com/JarvanDante/my_ad/internal/modules/slot/logic"
)

func RegisterAdmin(group *ghttp.RouterGroup, svc *logic.Service) {
	ctrl := controller.NewAdmin(svc)
	group.Bind(ctrl.List, ctrl.Create, ctrl.Detail, ctrl.Update, ctrl.Delete)
}
