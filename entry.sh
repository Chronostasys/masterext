#!/bin/bash
cd /fastgithub_linux-x64
export DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=1
./fastgithub>/dev/null &
sleep 3s
cp cacert/fastgithub.cer /usr/local/share/ca-certificates/my-cert.crt
cat cacert/fastgithub.cer >> /etc/ssl/certs/ca-certificates.crt
update-ca-certificates
export HTTP_PROXY="http://127.0.0.1:38457"
export HTTPS_PROXY="http://127.0.0.1:38457"

