
in a system service file (/etc/systemd/system/YOUR_SERVICE_)
```sh

[Unit]
Description=VPSRedployDaemon

[Service]
Type=simple
Restart=always
User=root
ExecStart=/root/tools/systemd-webhook/bin/gogithook

[Install]
WantedBy=multi-user.target

```


```sh

systemctl status gogithook.service

```
