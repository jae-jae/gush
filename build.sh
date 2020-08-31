#!/bin/bash

rm -rf dist/*
gox -os="linux darwin" -arch="amd64 386" -output="dist/gush_{{.OS}}_{{.Arch}}"