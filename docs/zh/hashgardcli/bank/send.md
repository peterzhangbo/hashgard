# hashgardcli bank send

## 描述

发送通证到指定地址 

## 使用方式

```
hashgardcli bank send [to_address] [amount] [flags]
```

## 标志

| 命令，速记       | 类型    | 是否必须 | 默认值                | 描述                                                         |
| ---------------- | ------- | -------- | --------------------- | ------------------------------------------------------------ |
| --account-number | int     | 否       |                       | 发起交易的账户的编号                                         |
| --async          | bool | 否       | false                 | 是否异步广播交易                                             |
| --dry-run        | bool | 否       | false                 | 模拟执行交易，并返回消耗的`gas`。`--gas`指定的值会被忽略     |
| --fees           | string  | 是       |                       | 交易费，例如： 10stake,1atom                                 |
| --from           | string  | 是       |                       | 发送交易的账户名称                                           |
| --gas            | string  | 否       | 2000000               | 交易的gas上限; 设置为"auto"将自动计算相应的阈值              |
| --gas-adjustment | float   | 否       | 1                     | gas调整因子，这个值降乘以模拟执行消耗的`gas`，计算的结果返回给用户; 如果`--gas`的值不是`atuo`，这个标志将被忽略 |
| --gas-prices     | string  | 否       |                       | 交易费用单价，(例如： 0.00001stake)                          |
| --generate-only  | bool | 否       | false                 | 是否仅仅构建一个未签名的交易便返回。                         |
| -h, --help       |         | 否       |                       | 打印帮助                                                     |
| --indent         | bool | 否       | false                 | 格式化json字符串                                             |
| --ledger         | bool | 否       | false                 | 是否使用硬件钱包                                             |
| --memo           | string  | 否       |                       | 备注信息                                                     |
| --node           | string  | 否       | tcp://localhost:26657 | <主机>:<端口> tendermint节点的rpc地址。                      |
| --print-response | bool | 否       | true | 是否打印交易返回结果，仅在`async`为false的情况下有效  |
| --sequence       | int     | 否       |                       | 发起交易的账户的sequence                                     |
| --trust-node     | bool | 否       | true                  | 是否信任全节点返回的数据，如果不信任，客户端会验证查询结果的正确性 |

## 全局标志

| 命令，速记            | 默认值         | 描述                                | 是否必须 |
| --------------------- | -------------- | ----------------------------------- | -------- |
| | --chain-id | string | tendermint 节点网络ID | 是 |
| -e, --encoding string | hex            | 字符串二进制编码 (hex \|b64 \|btc ) | 否       |
| --home string         | /root/.hashgardcli | 配置和数据存储目录                  | 否       |
| -o, --output string   | text           | 输出格式 (text \|json)              | 否       |
| --trace               |                | 出错时打印完整栈信息                | 否       |

## 例子

### 发送通证到指定地址 

```
 hashgardcli bank send gard1c9vrvvz08hd4entr0y5kfrt43v6malv60qtjfl 10gard --from=hashgard --chain-id=hashgard --ind
 ent -o json
```

命令执行完成后，返回执行的细节信息

```
{
 "height": "21667",
 "txhash": "58110E97BD93CFA123B43B7C893386BA26F238570E1131A7B6E1E6ED5B7DA605",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "22344",
 "tags": [
  {
   "key": "action",
   "value": "send"
  },
  {
   "key": "sender",
   "value": "gard10tfnpxvxjh6tm6gxq978ssg4qlk7x6j9aeypzn"
  },
  {
   "key": "recipient",
   "value": "gard1c9vrvvz08hd4entr0y5kfrt43v6malv60qtjfl"
  }
 ]
}
PS

```