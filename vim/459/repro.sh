#!/bin/bash
mkdir "with spaces"
ln -s $(which sh) "with spaces"
SHELL=$(pwd)/"with spaces"/sh vim
