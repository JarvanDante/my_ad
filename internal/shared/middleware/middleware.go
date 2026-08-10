package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/JarvanDante/my_ad/internal/dao"
	"github.com/JarvanDante/my_ad/internal/shared/authz"
)

func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func AdminToken(r *ghttp.Request) {
	want := g.Cfg().MustGet(r.Context(), "security.admin_token").String()
	got := r.Header.Get("X-Admin-Token")
	if want == "" || got == "" || got != want {
		r.Response.WriteStatus(http.StatusUnauthorized)
		r.Response.WriteJsonExit(g.Map{"code": 401, "message": "invalid admin token", "data": nil})
		return
	}
	r.Middleware.Next()
}

func AppKey(r *ghttp.Request) {
	key := r.Header.Get("X-App-Key")
	secret := r.Header.Get("X-App-Secret")
	if key == "" || secret == "" {
		r.Response.WriteStatus(http.StatusUnauthorized)
		r.Response.WriteJsonExit(g.Map{"code": 401, "message": "missing app credentials", "data": nil})
		return
	}
	cli, err := dao.NewClientRepo().FindActive(r.Context(), key)
	if err != nil || cli == nil || !authz.MatchSecret(secret, cli.AppSecret, cli.SecretHashed == 1) {
		r.Response.WriteStatus(http.StatusUnauthorized)
		r.Response.WriteJsonExit(g.Map{"code": 401, "message": "invalid app credentials", "data": nil})
		return
	}
	if cli.SecretHashed == 0 {
		_ = dao.NewClientRepo().Upsert(r.Context(), key, secret, cli.SiteCode, cli.Remark, 1)
	}
	r.SetCtxVar("app_key", key)
	r.SetCtxVar("site_code", cli.SiteCode)
	r.Middleware.Next()
}

func NotFound(r *ghttp.Request) {
	r.Response.WriteStatus(http.StatusNotFound)
	r.Response.WriteJson(g.Map{"code": 404, "message": "not found", "data": nil})
}
