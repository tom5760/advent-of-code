#!/bin/sh

INPUT="day$1/input$2"

if test -f $INPUT
then
  cargo run --package "day$1" < "day$1/input$2"
else
  echo no input file
  cargo run --package "day$1"
fi
