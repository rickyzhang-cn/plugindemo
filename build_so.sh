#!/bin/bash
projectpath=`pwd`
export GOPATH="${projectpath}:$GOPATH"
export GOBIN="${projectpath}/bin"
pluginname="logic_plugin$1"
pluginpath="logic_plugin_version_$1"
go build -buildmode plugin -o $pluginname.so --ldflags="-pluginpath=$pluginpath" logic
