# zero-svr

使用 [go-zero](https://github.com/zeromicro/go-zero) 框架编写的企业微信应用接收消息回调API

后台使用时填写的接收消息服务器URL为

    https://your.host.name/zero/receive

请将配置中的`your.host.name`替换成真实域名地址

## 注意事项
- 服务器部署协议`http` or `https`
- 前端加`nginx`或其它代理到端口`8888`;端口号可在`etc/zero-api.yaml`里配置
- 注意修改`etc/zero-api.yaml`里的 `CorpId`、`AppToken`、`EncodingAesKey`