
```
sequenceDiagram

Note over client, login: 登录
client->>gate: loginREQ（登录请求）
gate->>login: loginREQ（rpc消息转发）
login-->>gate: loginACK（rpc消息回复）
gate-->>client: loginACK（登录回复）

Note over client, game: 进入游戏
client->>gate: EnterREQ（进入游戏请求）
gate->>game: EnterREQ（rpc消息转发）
game-->>gate: EnterACK（rpc消息回复）
gate-->>client: EnterACK（进入游戏回复）

Note over client, game: 离开游戏
client->>gate: LeaveREQ（进入游戏请求）
gate->>game: LeaveREQ（rpc消息转发）
game-->>gate: LeaveACK（rpc消息回复）
gate-->>client: LeaveACK（进入游戏回复）

Note over client, xxx: ...

