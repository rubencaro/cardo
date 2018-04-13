# Cardo

## Setup

See [Setup a polite Golang environment](https://gist.github.com/rubencaro/5ce32fb30bbfa70e7db6be14cf42a35c) if you need to work with other Go projects with different dependencies.

If you have a single Go workspace, `cd` into it and you can go with:

```bash
go get -d github.com/rubencaro/cardo
cd src/github.com/rubencaro/cardo
go get -u github.com/golang/dep/cmd/dep
dep ensure -v
```

You will need to [install ArangoDB](https://www.arangodb.com/) on your system. Follow their instructions to get there.

## Development

`go run main.go`

## TODOs

* Stop using `viper` to read config
