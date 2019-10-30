1.
File structure after unzipping:
crawler>

├── Copyright.txt
├── Readme.txt
├── crawler.iml
├── run.sh
├── server.go
└── src
    └── com.crawler
        ├── service
        │   └── crawler.go
        └── util
            └── http_handler.go

2.
Install golang:

if using MACOS:
crawler>brew install go

verify installation:

crawler> go version
go version go1.9 darwin/amd64

3.
Assumptions:

https_url-> "https://www.abc.com"
http_url-> "http://www.abc.com"

Url domain check: if url starts with above http/https url.

Non-Domain: any url starts with "http" and not a valid domain as above

4.
run crawler as:

crawler> ./run.sh

Type ^C to stop the process

5.
Sample run:

crawler> ./run.sh
GOPATH is unset
Setting GOPATH as '/Users/nkumar/gowork'
