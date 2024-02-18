#### 使用文档详细说明
###### 功能说明
- 获取动态公网ip，支持ipv4或者ipv6
- 支持ipv6需要修改配置文件中的：update_ipv4_url的url为ipv6的地址，地址从http://dynv6.com获取，需要修改配置文件中的：get_wan_ip_url的url为：http://v6.ipinfo.io/
- ipv4获取公网ip地址：http://ipinfo.io/
- ipv6获取公网ip地址：http://v6.ipinfo.io/
###### 配置文件说明

- 根据配置文件：config.ini里面的注释修改

###### 编译说明
- go版本：go 1.20
- git clone 主分支
- go build

###### 程序启动说明
- 添加运行用户：useradd ddns
- 使用systemd管理
- 创建ddns.service
```shell
vim /etc/systemd/system/ddns.service 
[Unit]
Description=Prometheus
After=network.target
[Service]
Type=simple
WorkingDirectory=/Data/ddns
User=ddns
Group=ddns
ExecStart=/Data/ddns/tools -inifilename /Data/ddns/config.ini
Restart=on-failure
RestartSec=5
LimitNOFILE=65536
[Install]
WantedBy=multi-user.target

systemctl daemon-reload

# 启动ddns
systemctl start ddns.service
# 开启开机自动启动
systemctl enable ddns.service
# 查看运行状态
systemctl status ddns.service
```

##### 依赖免费的解析提供商：http://dynv6.com
- 具体注册使用不再说明


## =============================English==================================


#### Documentation Details
###### Feature Description
- Acquire dynamic public IP, supporting both IPv4 and IPv6
- To enable IPv6 support, modify the update_ipv4_url in the configuration file to an IPv6 address, which can be obtained from http://dynv6.com. Also, change the get_wan_ip_url in the configuration file to: http://v6.ipinfo.io/
- For IPv4, acquire the public IP address from:http://ipinfo.io/
- For IPv6, acquire the public IP address from: http://v6.ipinfo.io/
###### Configuration File Description
- Modify according to the comments in the configuration file: config.ini

###### Compilation Instructions
- Go version:go 1.20
- Git clone the master branch
- Go build

###### Program Start-up Instructions
- Add a runtime user: useradd ddns
- Managed using systemd
- Create ddns.service
```shell
vim /etc/systemd/system/ddns.service 
[Unit]
Description=Prometheus
After=network.target
[Service]
Type=simple
WorkingDirectory=/Data/ddns
User=ddns
Group=ddns
ExecStart=/Data/ddns/tools -inifilename /Data/ddns/config.ini
Restart=on-failure
RestartSec=5
LimitNOFILE=65536
[Install]
WantedBy=multi-user.target

systemctl daemon-reload
# Start ddns
systemctl start ddns.service
# Enable auto-start at boot
systemctl enable ddns.service
# Check the running status
systemctl status ddns.service
```

##### Depends on the free DNS provider: http://dynv6.com
- Specific registration and usage will not be explained further