#!/bin/bash
go build -tags "yq_nolua yq_notoml yq_noxml yq_nojson" -ldflags "-s -w" .