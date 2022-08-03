#!/usr/bin/env bash
set -euxo pipefail
wget -O /usr/local/bin/gintest "https://github.com/thom-vend/dummyapi/raw/main/mini-api-benchmark"
chmod 755 /usr/local/bin/gintest
cat > /etc/systemd/system/gintest.service <<EOF
[Unit]
Description=gintest

[Service]
Type=simple
Restart=always
Environment=LBURL=http://internal-thomas-lb-202024779.us-west-2.elb.amazonaws.com/ping
ExecStart=/usr/local/bin/gintest
StandardOutput=null
StandardError=null

[Install]
WantedBy=default.target
EOF
systemctl daemon-reload
systemctl start gintest.service
systemctl enable gintest.service
systemctl stop docker.service containerd.service datacollector.service datadog-agent.service ||true
systemctl disable docker.service containerd.service datacollector.service datadog-agent.service ||true
echo "done"
