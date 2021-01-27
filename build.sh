go get ./...
go mod vendor
cd module/list/main
go build -o pluto-list main.go
cd -
cd module/gm/main
go build -o pluto-gm main.go
cd -
cd module/stat/main
go build -o pluto-stat main.go
cd -
cd module/pay/main
go build -o pluto-pay main.go
cd -