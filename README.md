# Distributed microservice written in Golang for counting Twitter hashtags 

## Tech used:
- NSQ (Message bus written in Go)
- MongoDB
- Golang (of course)
- Docker 
- Microservices


## Steps:
- `brew install nsq`
- nsq depends on gpm, so need to install that too: `brew install gpm`
- `go get github.com/bitly/go-nsq`
- `brew install mongodb`
- `go get gopkg.in/mgo.v2`
- start nsqlookupd `nsqlookupd` (window 1)
- start nsqd and point it to nsqlookupd port `nsdq --lookupd-tcp-address=localhost:4160` (window 2)
- `mkdir db` (window 3)
- `mongod --dbpath ./db` (window 3)
- go to apps.twitter.com, and create a new app, then create an access token
- note down API key, API secrete, Access token, Access token secrete and set it as `ENV` vars (see setup.sh)
