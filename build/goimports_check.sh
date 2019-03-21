#!/bin/bash
diff=$(goimports -d ./..)

if [ "$diff" != "" ]; then 
    echo "Goimports check has failed";  exit 1; 
else 
    echo "Goimports check is sucessful"; exit 0;
fi