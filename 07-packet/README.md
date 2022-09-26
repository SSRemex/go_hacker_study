## 原始数据包操作

linux需安装 `apt install libpcap-dev`

windows需要安装 `winpcap`

所用go依赖 `go get github.com/google/gopacket`

运行时需要管理员权限，WSL无法使用

## syn 翻红保护进行端口骚猫
TCP标志位

7   6   5   4   3   2   1   0

CWR ECE URG ACK PSH RST SYN FIN

检查指定位置的标志位的值，0代表关闭 1代表打开
 - ACK和FIN：00010001(0x11)
 - ACK:00010000(0x10)
 - ACK和PSH：00011000(0x18)
以上代表开放