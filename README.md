# FSS
一个简易的HTTP-Server

# Useage
* need a mogondb instance:
    see scriptes/mongo.sh
* ./fss run
    -- default addr :8888
    -- default metrics addr: 9999
    -- default dsn(data source name) : mongodb://localhost:27017
curl localhost:8888/api/v1/text/get
curl localhost:9999/metrics


