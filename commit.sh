#!/bin/bash
gofmt -s -w .
git add "$1"
git commit