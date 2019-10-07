#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export LUA_PATH="$DIR/mac/?.lua;;"
php "$DIR/lib/compile_scripts.php" $*
