#!/bin/bash
diff=$(gofmt -d $(find . -type f -name '*.go' -not -path "./vendor/*"))

if [ "$diff" != "" ]; then 
    echo "Gofmt check has failed";  exit 1; 
else 
    echo "Gofmt check is sucessful"; exit 0;
fi