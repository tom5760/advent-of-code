#!/bin/sh

INPUT="day$1/input$2"

if test -f $INPUT
then
  go run ./day$1 < day$1/input$2
else
  echo no input file
  go run ./day$1
fi
