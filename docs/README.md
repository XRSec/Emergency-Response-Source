# Emergency-Response

![icon](/assert/icon.jpg)

## TODO
- [ ] Add Linux SH to [SH](/sh)
- [ ] Add Exec General module
- [ ] ...


```go
type App struct {}
var Api App
```

### Build

```bash
# https://github.com/XRSec/xgo
# [国内] sudo wget "https://ghproxy.com/https://github.com/XRSec/xgo/releases/download/$(curl -sL "https://api.github.com/repos/crazy-max/xgo/releases/latest" | grep tag_name | awk '{print $2}' | tr -d '"' | tr -d ',')/xgo-$(uname -s)-$(uname -m)" -O /usr/local/bin/xgo
sudo wget "https://ghproxy.com/https://github.com/XRSec/xgo/releases/download/$(curl -sL "https://api.github.com/repos/crazy-max/xgo/releases/latest" | grep tag_name | awk '{print $2}' | tr -d '"' | tr -d ',')/xgo-$(uname -s)-$(uname -m)" -O /usr/local/bin/xgo
sudo chmod +x /usr/local/bin/xgo
xgo .
cd github.com/XRSec/
```

- [Emergency-Response-Notes](https://github.com/Bypass007/Emergency-Response-Notes)
- [Emergency-response-notes](https://github.com/wpsec/Emergency-response-notes)
- [EmergencyResponse](https://github.com/yaunsky/EmergencyResponse)
- [Linux-emergency-response-script](https://github.com/looosooo/Linux-emergency-response-script)
- [LinuxEmergency](https://github.com/b0bac/LinuxEmergency)
- [GScan](https://github.com/grayddq/GScan)
- [zhihu-gopsutil](https://zhuanlan.zhihu.com/p/126362239#%E4%BD%BF%E7%94%A8%E7%8E%87)
- [soiiy-gopsutil](http://soiiy.com/go/16228.html)
- [jianshu–gopsutil](https://www.jianshu.com/p/914acdb5d7c2)
