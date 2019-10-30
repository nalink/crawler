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

https_url-> "https://www.qubole.com"
http_url-> "http://www.qubole.com"

Both are valid domains for qubole.
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

[Domain] https://www.qubole.com
[Non-Domain] https://qubole.zendesk.com/hc/en-us
[Domain] https://www.qubole.com/education/
[Domain] https://www.qubole.com/products/autonomous-data-platform/
[Domain] https://www.qubole.com/solutions/by-job-roles/#teams
[Domain] https://www.qubole.com/solutions/by-job-roles/#teams
[Domain] https://www.qubole.com/solutions/by-job-roles/#users
[Domain] https://www.qubole.com/solutions/by-job-roles/#users
[Non-Domain] https://twitter.com/@qubole^C

Exiting Crawler...
