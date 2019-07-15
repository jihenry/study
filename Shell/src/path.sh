#!/bin/bash

echo $(cd `dirname $0`; pwd)

WDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo $WDIR

echo ${BASH_SOURCE[0]}