# hashgardcli exchange query-order

## 描述

查看指定id的订单详情

## 使用方式

```
hashgardcli exchange query-order [order_id] [flags]
```

## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 例子

### 查询订单

```shell
hashgardcli exchange query-order 1 --chain-id hashgard --indent -o=json
```

下面是这个id为1的订单，seller是订单的创建者，supply是订单创建时提供的交换币种和数量，target是目标币种及数量，两者数量可计算交易的单价，remains则是目前订单中剩余的supply，create_time则是订单创建时间。

```txt
{
 "order_id": "1",
 "seller": "gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa",
 "supply": {
  "denom": "gard",
  "amount": "100"
 },
 "target": {
  "denom": "apple",
  "amount": "800"
 },
 "remains": {
  "denom": "gard",
  "amount": "100"
 },
 "create_time": "2019-04-16T07:12:39.254071Z"
}
```

如何同订单交易？

请点击下述链接：

[take-order](take-order.md)
