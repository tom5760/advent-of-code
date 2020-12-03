#!/bin/sh

INPUT="inputs/day$1$2"

if test -f $INPUT
then
  cargo run --release --bin "day$1" < "$INPUT"
else
  echo no input file
  cargo run --release --bin "day$1"
fi
