#!/bin/bash

dir="day$1"

mkdir -p "$dir"
touch "$dir/input1.txt"
touch "$dir/input2.txt"

echo "Created $dir with input1.txt and input2.txt"
