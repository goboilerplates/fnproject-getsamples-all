# Boilerplate for GetSamplesAll FnProject Function in Go
[![Build Status](https://travis-ci.org/goboilerplates/fnproject-getsamples-all.svg?branch=master)](https://travis-ci.org/goboilerplates/fnproject-getsamples-all)
[![codecov](https://codecov.io/gh/goboilerplates/fnproject-getsamples-all/branch/master/graph/badge.svg)](https://codecov.io/gh/goboilerplates/fnproject-getsamples-all)
[![Go Report Card](https://goreportcard.com/badge/github.com/goboilerplates/fnproject-getsamples-all)](https://goreportcard.com/report/github.com/goboilerplates/fnproject-getsamples-all)
[![GoDoc](https://godoc.org/github.com/goboilerplates/fnproject-getsamples-all?status.svg)](https://godoc.org/github.com/goboilerplates/fnproject-getsamples-all)
[![](https://images.microbadger.com/badges/image/goboilerplates/fnproject-getsamples-all.svg)](https://microbadger.com/images/goboilerplates/fnproject-getsamples-all)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/goboilerplates/fnproject-getsamples-all/blob/master/LICENSE)

## Features
- FnProject Function
- Run anywhere
    - Public, private and hybrid cloud
    - Import Lambda functions and run them anywhere
- CI with Travis
- Docker Build

## Installation

Get the fnproject-getsamples-all repository

```
go get github.com/goboilerplates/fnproject-getsamples-all

cd echo $GOPATH/src/github.com/goboilerplates/fnproject-getsamples-all
```
## Running the tests

Run all tests

```
go test ./...
```

Or run all tests with coverage

```
bash script/coverage.sh
```

## Docker support 

Run and build function docker image

```
bash script/Dockerbuild.sh
```

Deploy function to local FnProject Server

```
bash script/deploy.sh
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/goboilerplates/fnproject-getsamples-all/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details