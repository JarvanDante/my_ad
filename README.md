# my_ad · 广告中台 (PaaS)

平台级广告位 / 素材 / 投放；子站用 `APPKEY/APPSECRET` 按广告位拉取可展示内容。素材文件建议走统一存储 `my_storage`，本服务只存 `media_url` / `storage_object_id` 引用。

## 端口

| 环境 | 地址 |
|------|------|
| 进程内 | `:8006` |
| Docker 宿主机 | `8016:8006` |

## 鉴权

- Admin：`X-Admin-Token`
- Open：`X-App-Key` + `X-App-Secret`

## 快速启动

```bash
docker exec postgres psql -U postgres -c "CREATE DATABASE my_ad;"
make migrate
make build && ./bin/adapi
# 或
make docker-up
```

Swagger：`http://127.0.0.1:8006/swagger/`

## 接口概览

见 `docs/api.md`。
