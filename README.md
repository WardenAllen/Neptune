
# Neptune is go tcp/http server.

cd /home/down
wget https://dl.google.com/go/go1.15.linux-amd64.tar.gz
rm -rf /usr/local/go
tar -zxf go1.15.linux-amd64.tar.gz -C /usr/local
vim /etc/profile

#golang env config
export GO111MODULE=on
export GOROOT=/usr/local/go
export GOPATH=/home/gopath
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

cd /home
mkdir gopath

source /etc/profile
go version

go env -w GOPROXY=https://goproxy.cn,direct

alias goroot='cd $GOROOT/src'
alias gopath='cd $GOPATH'

goroot
git clone https://gitee.com/WardenAllen/pluto_go.git
cd pluto_go
sh build.sh
sh run.sh