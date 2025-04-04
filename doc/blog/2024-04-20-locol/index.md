---
slug: locol
title: 内网穿透
authors: [uif]
tags: [内网穿透, VPS, VPN转发]
---

---

有属于自己的一台 VPS 有什么好处？假如你肉身在中国大陆，你在美国的家里有一台常年开机的服务器，通常为了可以像在美国家里一样访问网络的任何资源（也就是俗称翻墙），你可以通过 VPN 技术，把所有的网络流量转发到美国。

同样的，你也可以在美国时，想要使用网易云、QQ音乐等有版权隔离的应用，也需要IP”回国“才能正常使用。

## 为什么需要内网穿透？

如果你的需要时“翻墙”、“回国”，通常只需购买具有公网 IP 的 VPS 即可，并不需要内网穿透。但有些人就喜欢折腾。

## 最佳实践

必要的：

- 一台具有公网的VPS
- 两台电脑客户端（Client）和服务端（Server）

### 使用 [frp](https://github.com/fatedier/frp)

#### 将 frp 安装在公网 VPS

#### 将 frp 安装在服务端

#### 在服务端安装 UIF

#### 开启 UIF 入站

#### 在客户端安装 UIF

#### 添加 UIF 出站

#### 客户端启用 Tun 模式或者 HTTP

启用 Tun，那跟 VPN 无差异，使用 HTTP 就可以满足绝大部分人群。
