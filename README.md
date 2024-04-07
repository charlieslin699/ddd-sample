# ddd-sample
domain-driven design(領域驅動設計) CQRS 架構範例

## 系統架構
```
ddd-sample
├─ application                     # CQRS
│  ├─ command
│  ├─ eventhandler
│  └─ query
├─ cmd                             # 程式進入點 (main)
│  ├─ apiserver
│  │  ├─ system
│  │  └─ swagger
│  └─ schedule
├─ config                          # 設定檔目錄
├─ infra                           # infrastructure 統整外部資料連線
│  ├─ api
│  ├─ db
│  ├─ rabbitmq
│  └─ redis
├─ internal                        # 領域目錄
│  ├─ auth
│  │  ├─ adapter                   # 介接層
│  │  ├─ aggregate                 # 聚合
│  │  ├─ entity                    # 領域核心
│  │  ├─ event                     # 領域事件
│  │  ├─ repository                # 資料溝通介面
│  │  └─ valueobject               # 領域物件
│  ├─ core
│  │  ├─ adapter                   # base adapter 內含事件發佈邏輯
│  │  ├─ aggregate                 # base aggregate 內含事件收集邏輯
│  │  ├─ event                     # base event 領域事件封裝
│  │  └─ repository                # base repository
│  └─ lang
│     ├─ adapter
│     └─ repository
├─ mocks                           # mocks
├─ pkg                             # 工具集
│  ├─ config                       # 設定檔相關邏輯
│  ├─ env                          # 環境變數相關邏輯
│  ├─ errorcode                    # error code集合
│  ├─ httpserver                   # http服務相關邏輯
│  ├─ localtime                    # 時間相關邏輯
│  ├─ log                          # log相關邏輯
│  ├─ otp                          # OTP相關邏輯
│  ├─ random                       # 隨機數相關邏輯
│  ├─ token                        # token產生器
│  ├─ uid                          # uid產生器
│  └─ util                         # 常用方法集合
└─ userinterface                   # 系統訪問入口
   ├─ amqp
   ├─ api
   │  ├─ auth
   │  │  ├─ errorhandler
   │  │  ├─ middleware
   │  │  ├─ model
   │  │  ├─ panichandler
   │  │  ├─ restful
   │  │  └─ router.go
   │  └─ common
   ├─ crontab
   └─ grpc

```