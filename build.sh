#!/usr/bin/env bash

cd cmd/tandemo/ && go build -v && cd ../../
mv cmd/tandemo/tandemo ./
