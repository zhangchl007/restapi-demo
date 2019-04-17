
It's a demo to deploy restapi-demo on  cloud foundry

```
 cf api https://api.run.haas-421.pez.pivotal.io
 cf login
 cd  restapi-demo
 cf push

for example:

 cf apps
 Getting apps in org redhat / space test as jimmy...
 OK

 name       requested state   instances   memory   disk   urls
 api-demo   started           1/1         64M      1G     api-demo.cfapps.haas-421.pez.pivotal.io

curl http://api-demo.cfapps.haas-421.pez.pivotal.io/people

OS/Other deployment

cd $GOPATH
go get github.com/zhangchl007/restapi-demo

go build 
go install github.com/zhangchl007/restapi-demo

ls -l $GOPATH/bin
total 9028
-rwxr-xr-x. 1 root root 7224691 Aug  8 02:17 restapi-demo

./restapi-demo

GRUD support 
# curl http://192.168.233.134:9000/people
[{"id":"1","firstname":"zhang","lastname":"jimmy","address":{"city":"shenzhen","state":"Guangdong"}},{"id":"2","firstname":"liu","lastname":"tony","address":{"city":"shenzhen","state":"Guangdong"}}]
# curl http://192.168.233.134:9000/people/1
{"id":"1","firstname":"zhang","lastname":"jimmy","address":{"city":"shenzhen","state":"Guangdong"}}
# curl http://192.168.233.134:9000/people/2
{"id":"2","firstname":"liu","lastname":"tony","address":{"city":"shenzhen","state":"Guangdong"}}

```
