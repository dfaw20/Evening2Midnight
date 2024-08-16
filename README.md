
以下の内容で`systemd`のサービスを記述する

/usr/local/bin/shutdown_service
```
[Unit]
Description=Lid Monitor for Shutdown
After=multi-user.target

[Service]
#ExecStart=/usr/local/bin/lid-monitor.sh
ExecStart=/usr/local/bin/shutdown_service
Restart=always

[Install]
WantedBy=multi-user.target
```

起動する

```
sudo systemctl enable lid-monitor.service
```