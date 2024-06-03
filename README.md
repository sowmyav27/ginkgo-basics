# ginkgo-basics

A few basic tests using Ginkgo testing framework.

## Pre-requisites

- Install Go
- Install Ginkgo by following instructions [here](https://onsi.github.io/ginkgo/#installing-ginkgo).
- After installing, ensure GO bin is in your PATH using: 
```
go env GOPATH # /home/user/go
export PATH=$PATH:$(go env GOPATH)/bin
```

## To run tests
To run tests, run `ginkgo`
> ginkgo
```
% ginkgo
Running Suite: 
=================================================================================

Will run 3 of 3 specs
•••

Ran 3 of 3 Specs in 0.002 seconds
SUCCESS! -- 3 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 1.657041292s
Test Suite Passed
```