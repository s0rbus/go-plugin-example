#!/bin/bash
CGO_ENABLED=1 go build -buildmode=plugin -o pg1/pg1.so pg1/pg1.go
