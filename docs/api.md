# 广告中台 API（MVP）

统一响应：`{ code, message, data }`，`code===0` 成功。

## 公开

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/healthz` | 探活 |

## Admin（`X-Admin-Token`）

### 调用方

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/admin/clients` | 列表 |
| PUT | `/admin/clients` | 同步凭证（manage 开站写入） |
| POST | `/admin/clients/{app_key}/disable` | 停用 |

### 广告位

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/admin/slots` | 列表 |
| POST | `/admin/slots` | 创建（code 唯一） |
| GET | `/admin/slots/{id}` | 详情 |
| PUT | `/admin/slots/{id}` | 更新 |
| DELETE | `/admin/slots/{id}` | 停用 |

### 素材

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/admin/creatives` | 列表 |
| POST | `/admin/creatives` | 创建（返回 16 位 id） |
| GET | `/admin/creatives/{id}` | 详情 |
| PUT | `/admin/creatives/{id}` | 更新 |
| DELETE | `/admin/creatives/{id}` | 下架 |

`media_url` 必填；可选 `storage_object_id` 关联统一存储对象。

### 投放

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/admin/campaigns` | 列表 |
| POST | `/admin/campaigns` | 创建 |
| GET | `/admin/campaigns/{id}` | 详情 |
| PUT | `/admin/campaigns/{id}` | 更新 |
| POST | `/admin/campaigns/{id}/status` | 启停 `{status:0\|1}` |

`site_code` 空=全站；非空=仅该站。`start_at`/`end_at` 可空。

## Open（`X-App-Key` / `X-App-Secret`）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/open/ads?slot_code=&limit=` | 按广告位拉取本站可展示广告 |
| POST | `/open/events` | 上报 `{event_type:impression\|click,...}` |

拉取规则：广告位启用 + 投放启用 + 素材就绪 +（全站或匹配 site_code）+ 在排期内；按 priority/weight 排序。
