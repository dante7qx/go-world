<h2 align="center">睿阳视频监控解析平台</h2>

## 一. 概述
睿阳视频监控解析平台基于开源项目 [RTSPtoWeb](https://github.com/deepch/RTSPtoWeb)，进行二次研发。平台用于将您的 RTSP 流转换为可在 Web 浏览器中使用的格式，
例如：MSE、WebRTC 或 HLS。 它是完全原生的 Golang，无需使用 FFmpeg 或 GStreamer！

## 二. 使用说明
### 1. 下载平台
睿阳视频监控解析平台已上传SVN，地址：`svn://116.176.33.76:8443/videomonitor`。目录结构如下：
```shell
├── config.json                   ## 配置文件
├── config.json.example.json      ## 配置样例
├── examples                      ## 页面集成样例
│   ├── hls                       ## hls 集成样例
│   ├── hlsll                     ## hlsll 集成样例
│   ├── mse                       ## mse 集成样例
│   └── webrtc                    ## webrtc 集成样例
├── run                           ## Linux 可执行文件
├── run.exe                       ## Windows 可执行文件
├── runMAC                        ## MacOS 可执行文件
└── web                           ## 平台控制台页面
    ├── static
    └── templates
```

### 2. 配置并运行
```json
{
  "channel_defaults": {},
  "server": {
    "debug": true,
    "http_debug": false,
    "http_demo": true,
    "http_dir": "",
    "http_login": "user",
    "http_password": "userpassword",
    "http_port": ":8083",
    "https": false,
    "https_auto_tls": false,
    "https_auto_tls_name": "",
    "https_cert": "",
    "https_key": "",
    "https_port": "",
    "ice_credential": "",
    "ice_servers": [
      "stun:stun.l.google.com:19302"
    ],
    "ice_username": "",
    "log_level": "info",
    "rtsp_port": ":5541",
    "token": {
      "backend": "",
      "enable": false
    },
    "webrtc_port_max": 0,
    "webrtc_port_min": 0
  },
  "streams": {
    "test": {
      "channels": {
        "0": {
          "name": "摄像头1",
          "on_demand": true,
          "url": "rtsp://admin:risun8768@192.168.1.249:554/Streaming/Channels/101"
        }
      },
      "name": "测试"
    }
  }
}
```

- 配置说明
  
  使用时只可修改下面的配置，其他的配置不要进行修改

  - **Server 配置**
  
  | 配置项          | 配置值                                                            | 说明                                                                 |
  |----------------|--------------------------------------------------------------------| --------------------------- |
  | debug           | true (false)，默认：true                                           | 启用调试输出                                                             |
  | log_level       | trace, debug, info, warning, error, fatal, or panic，默认：info    | 日志级别                                                               |
  | http_login      |                                                                | 认证用户名                                                              |
  | http_password   |                                                                | 认证密码                                                               |
  | http_port       | 默认:8083（注意，要带:，例如:8000）                                        | HTTP端口                                                             |
  | webrtc_port_min | 默认0，Linux环境可设置为 60000                                          | WebRTC使用的最小端口（UDP）                                                 |
  | webrtc_port_max | 默认0，Linux环境可设置为 61000                                          | WebRTC使用的最大端口（UDP）                                                 |
  
  **注意：** 修改`http_login`或`http_password`后，要同步修改`index.js`(位于`web/static/js/index.js`)
  ```javascript
  // 替换
  xhr.setRequestHeader("Authorization", "Basic " + btoa("<原始用户名>:<原始密码>"));
  // 为
  xhr.setRequestHeader("Authorization", "Basic " + btoa("<新用户名>:<新密码>"));
  ```
  - **Stream 配置**
  
  服务启动可不配置`streams`，通过平台控制台进行添加。对示例配置进行说明
  ```properties
  test：视频流解析标识（重要）
  0：   视频流通道标识（重要，按顺序配置0、1、2、3、4）
  url： RTSP地址
  ```

- 运行方式
  - Windows：双击`run.exe`
  - Linux：`./run`
  - MacOS：`./runMac`
  
  运行后，打开浏览器，输入地址：`http://<ip>:<port>`（本地启动的地址默认是：http://localhost:8083）

### 3. 集成Web页面

平台已经提供了集成Web页面的示例，详情请查看`examples`目录下的文件。

### 4. 限制
支持的视频编解码器： 所有`H264`配置（不支持`H265`）

支持的音频编解码器：否

## 三. 延时对比

| 采用技术 | 延迟    | 监控路数 |
| -------- | ------- | -------- |
| WebRTC   | < 1s    | 不限     |
| MSE      | 1s 左右 | < 6路    |
| LL-HLS   | 2～3s   | 不限     |
| HLS      | > 20s   | 不限     |

## 四. 对接实践
设置完视频流后，可以通过API获取对应的信息。

`GET /stream/{STREAM_ID}/info`

鉴权方式：Basic Auth

示例：`curl -u <username>:<password> http:<ip>:<port>/stream/{STREAM_ID}/info`

响应：
```json
{
    "status": 1,
    "payload": {
        "name": "test video",
        "channels": {
            "0": {
                "name": "ch1",
                "url": "rtsp://admin:admin@{YOUR_CAMERA_IP}/uri",
                "on_demand": true,
                "debug": false,
                "status": 0
            },
            "1": {
                "name": "ch2",
                "url": "rtsp://admin:admin@{YOUR_CAMERA_IP}/uri",
                "on_demand": true,
                "debug": false,
                "status": 0
            }
        }
    }
}
```

## 五. 技术原理

### 1. 视频监控
监控系统中摄像头的作用是将所采集的图像数据进行编码，而硬盘录像机的作用就是对摄像头编码的数据进行解码，然后输出到显示器，这样就能看到完整
和清晰的监控图像。
- 监控路数：视频监控系统里面的路数指的是你安装了多少摄像头（硬盘录像机），一般规格有4路、8路、16路和32路四种规格。
- IPC：IP Camera，网络摄像头的意思

### 2. RTSP协议
RTSP(Real Time Streaming Protocol)，实时流传输协议，是TCP/IP协议体系中的一个应用层协。RTSP可以是双向的，客户机和服务器都可以发出请求。
RTSP是用来控制声音或影像的多媒体串流协议，并允许同时多个串流需求控制，传输时所用的网络通讯协定并不在其定义的范围内，服务器端可以自行选择使用TCP或UDP来传送串流内容，它的语法和运作跟HTTP 1.1类似，但并不特别强调时间同步，所以比较能容忍网络延迟。

参考：https://www.jianshu.com/p/4e3925f98e84

### 3. HLS和LL-HLS

**HLS**
>HLS（HTTP Live Streaming）协议是由苹果提出的基于流媒体的网络传输协议。通过将整条流切割成一个个小的可以通过 HTTP 下载的媒体文件，然后提供一个配套的媒体列表文件，
提供给客户端，让客户端顺序地拉取这些媒体文件播放，来实现看上去是在播放一条流的效果。由于传输层协议只需要标准的 HTTP 协议，HLS 可以方便的透过防火墙或者代理服务器，而且可以很方便的利用 CDN 进行分发加速，并且客户端实现起来也很方便。
> 
> HLS 把整个流分成一个个小的基于 HTTP 的文件来下载，每次只下载一些。HLS 协议由三部分组成：HTTP、M3U8、TS。这三部分中，HTTP 是传输协议，M3U8 是索引文件，TS 是音视频的媒体信息。
> 
>HLS的延迟：HLS的一个分片大概5-10s,一般加载3个分片播放。所以延迟大概15-30s。

**LL-HLS**
> Low-Latency HLS。相比较于HLS，延迟更低，官方说明，最低延迟在3s左右。
> 
> 更小的分片划分
> 
> LL-HLS将大的分片且分为一个个较小的分片, 这种切分方式不是简单的将源分片等分, 而是结合fMP4封装和#EXT-X-MAP规则, 将整视频的头部和内容分开, 而且内容源被划分的很细, 例如原来一个分片6s左右, 可能被切分为30个200ms的fMP4分片, 这些分片使用#EXT-X-PART来标注.
> 
> 增量更新
> 
> 不发送完整的播放列表，而是发送播放列表的增量（默认的播放列表被保存，然后只在出现时发送增量，而不是发送完整的播放列表）
> 
> 服务器阻塞，消除轮询
> 
> 服务器的责任是保留请求（阻塞），直到包含新片段的播放列表版本可用。阻断播放列表的重新加载消除了轮询
> 
> 预加载支持
> 
> 通过EXT-X-PRELOAD-HINT标签，来声明即将产生的片段。客户端可以提前发起请求，当该片段生成时立即回复给客户端。

### 4. MSE

MSE（Media Source Extensions）媒体源扩展 API 提供了实现无插件且基于 Web 的流媒体的功能。使用 MSE，媒体串流能够通过 JavaScript 创建， 并且能通过使用`<audio>`和`<video>`元素进行播放。

### 5. WebRTC

WebRTC (Web Real-Time Communications)，由Google发起的实时通讯解决方案，其中包含视频音频采集，编解码，数据传输，音视频展示等功能。

参考资料：https://zhuanlan.zhihu.com/p/86751078

## 六. 常见问题

1. Linux环境无法播放WebRTC
   
   WebRTC端口被防火墙拦截，需要进行如下操作：
   - 修改`config.json`，设置
   ```json
    "webrtc_port_max": 61000,
    "webrtc_port_min": 60000
   ```
   - 服务器开放端口，在终端执行命令
   ```shell
    firewall-cmd --add-port=60000-61000/udp --permanent
    firewall-cmd --reload
   ```
