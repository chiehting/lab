# Nginx

## 前置

把要測試的配置改為default.conf,例如

```
mv ./config/conf.d/regex-path.cnf ./config/conf.d/default.conf
```

## 實驗

#### regex-path.cnf

測試路徑使用正規表示式,並顯示正規表示條件的變數.

```bash
➜ curl 'http://localhost/user/index.html?arg1=1&arg2=a'
var1: user, var2: index.html, arges: ?arg1=1&arg2=a

➜ curl 'http://localhost/developer/?type=foo'
var1:developer, var2:, arges: ?type=foo
```

#### cors.cnf

測試Cross-Origin Resource Sharing的配置.


```bash
➜ echo '127.0.0.1 cors.com'|sudo tee -a /etc/hosts
➜ open http://localhost/cors/client.html
```