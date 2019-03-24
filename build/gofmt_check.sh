#!/bin/bash
diff=$(gofmt -d ./..)

if [ "$diff" != "" ]; then 
    echo "Gofmt check has failed";  exit 1; 
else 
    echo "Gofmt check is sucessful"; exit 0;
fi