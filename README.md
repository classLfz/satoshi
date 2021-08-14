# Satoshi

## 项目不再维护

**很遗憾，该项目将不再维护更新。介于作者在市面上找到了很多类似，甚至比该项目好的多的项目，其代码质量远胜该项目，作者已决定不再维护当前项目。**

## 方案替代

[Home Assistant](https://www.home-assistant.io/) 是一个很成熟的方案。

当然，如果你跟作者类似，也是将设备接入到苹果的生态里边去的话，[homebridge](https://homebridge.io/)或许会是你的不二选择。

## 简介

Satoshi（宝可梦主角小智）目标是成为一个用于构建简易物联网系统的小工具。

Satoshi 作用于树莓派上，仅需几个简单的命令与配置，即可轻松接入语音助手（目前仅支持了Siri）。同时，Satoshi也提供了HTTP接口，在不方便接入语音助手的情形，也可通过接口访问操控设备。

## 使用说明

**项目初期，配置，接口可能会有较大的改动**

可通过 `help` 命令查看使用说明

```bash
$ satoshi --help
```

### 配置

#### 配置命令说明

```bash
$ satoshi config
```

#### 配置文件初始化

默认配置文件生成于 `~/.satoshi/config.yaml`

```bash
$ satoshi config init
```

#### 配置文件说明

```yaml
lircd:
  lircd_path: /var/run/lirc/lircd
satoshi:
  # 接口设备
  http:
    port: '8234' # 服务端口号
    devices:
      - id: '123' # 设备ID，调用时唯一标识
      type: switch # 设备类型
      name: device_name
      switch_config: # 设备对应的配置信息
        on_pin: 17 # GPIO口
        on_lircd_cmd: xxx on # lircd 控制开命令
        off_lircd_cmd: xxx off # lircd 控制关命令

  # Siri设备列表
  siri:
    pin_code: '00102003'   # satoshi birdge 8位HomeKit设置代码
    devices:
      - id: 2 # accessory设备ID，数字且大于等于2
      type: switch  # accessory设备类型
      name: device1  # 设备名
      switch_config: # 开关设备对应的配置信息
        on_pin: 17   # GPIO口
        update_interval: 5  # 定时更新状态间隔，单位秒
        on_lircd_cmd: xxx on # lircd 控制开命令
        off_lircd_cmd: xxx off # lircd 控制关命令
      - id: 3
      type: lightbulb
      name: device2
      lightbulb_config: # 灯泡设备对应的配置信息
        on_pin: 26   # GPIO口
        update_interval: 5  # 定时更新状态间隔，单位秒
        on_lircd_cmd: xxx on # lircd 控制开命令
        off_lircd_cmd: xxx off # lircd 控制关命令
```

### 开启服务

```bash
$ satoshi serve
```

#### http接口说明

- 调用设备

```bash
$ curl -X POST http://satoshi-host/v1/deviceCallers -d '{"id": "1", "toggle": true}'
```

## 计划

### 短期计划

- [x] 定时更新Siri设备状态
- [x] 接入lircd控制设备
- [ ] 接入更多Siri设备类型
- [ ] 完善http接口
