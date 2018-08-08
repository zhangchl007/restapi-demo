Golang learning 

I will spend some time to learn golang everyday. keep on it

restapi-demo

```
cd $GOPATH
go get github.com/zhangchl007/restapi

go build 
go install github.com/zhangchl007/restapi

ls -l $GOPATH/bin
total 9028
-rwxr-xr-x. 1 root root 7224691 Aug  8 02:17 restapi

./restapi

GRUD support 
# curl http://192.168.233.134:9000/people
[{"id":"1","firstname":"zhang","lastname":"jimmy","address":{"city":"shenzhen","state":"Guangdong"}},{"id":"2","firstname":"liu","lastname":"tony","address":{"city":"shenzhen","state":"Guangdong"}}]
# curl http://192.168.233.134:9000/people/1
{"id":"1","firstname":"zhang","lastname":"jimmy","address":{"city":"shenzhen","state":"Guangdong"}}
# curl http://192.168.233.134:9000/people/2
{"id":"2","firstname":"liu","lastname":"tony","address":{"city":"shenzhen","state":"Guangdong"}}

```
