#!/usr/bin/env bash

mkdir /opt/socks5-server

cp socks5-server /opt/socks5-server

cp .env /opt/socks5-server

cp socks5-server.service /etc/systemd/system/

systemctl daemon-reload

service socks5-server start