## WEB-noVNC-免登录代理服务

### 功能介绍:
参照websockify实现的TCP转WebSocket代理服务。并在代理转发的过程中拦截注入了VNC的密码,使
VNC的TCP连接地址、端口和VNC密码可以通过后端注册代理的时候进行传入,对前端无感。用于实现较为
安全的，可根据当前session登录用户的可访问资源范围，进行远程连接控制的方式，而无需暴露VNC密
码。将VNC的远程链接安全风险转嫁给用户认证部分。由于本服务和NVC控制的远程桌面(一般是虚拟化环
境)通常部署在服务器网络内，可通过网络隔离的方式，配合本服务在进行代理时的权限业务逻辑处理(需
自行实现)，解决生产用户在WEB端浏览器上访问NVC资源时的授信和密码的业务问题。

### 依赖软件:
- noVNC 1.4.0 (内部集成)

### 支持的后端内部认证方式:
- VNC Authentication
- None

### 依赖包:
- github.com/gin-gonic/gin v1.9.0
- github.com/gofrs/uuid v4.4.0
- github.com/gorilla/websocket v1.5.0
- github.com/bytedance/sonic v1.8.0
- github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311
- github.com/gin-contrib/sse v0.1.0
- github.com/go-playground/locales v0.14.1
- github.com/go-playground/universal-translator v0.18.1
- github.com/go-playground/validator/v10 v10.11.2
- github.com/goccy/go-json v0.10.0
- github.com/json-iterator/go v1.1.12
- github.com/klauspost/cpuid/v2 v2.0.9
- github.com/leodido/go-urn v1.2.1
- github.com/mattn/go-isatty v0.0.17
- github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421
- github.com/modern-go/reflect2 v1.0.2
- github.com/pelletier/go-toml/v2 v2.0.6
- github.com/twitchyliquid64/golang-asm v0.15.1
- github.com/ugorji/go/codec v1.2.9
- golang.org/x/arch v0.0.0-20210923205945-b76863e36670
- golang.org/x/crypto v0.5.0
- golang.org/x/net v0.7.0
- golang.org/x/sys v0.5.0
- golang.org/x/text v0.7.0
- google.golang.org/protobuf v1.28.1
- gopkg.in/yaml.v3 v3.0.1