#!/bin/bash
diff=$(goimports -d $(find . -type f -name '*.go' -not -path "./vendor/*"))

if [ "$diff" != "" ]; then 
    echo "Goimports check has failed";  exit 1; 
else 
    echo "Goimports check is sucessful"; exit 0;
fi