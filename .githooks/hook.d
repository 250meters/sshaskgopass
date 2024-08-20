#!/bin/bash

dirname="`dirname \"$0\"`"
hooktype="`basename \"$0\"`"
problem=0
for script in "${dirname}"/"${hooktype}.d"/*; do
  "${script}" "$@" || problem=1
done

if [ $problem -eq 1 ]; then
    exit 1
fi
