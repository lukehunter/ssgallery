#!/bin/sh
gox -output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"
