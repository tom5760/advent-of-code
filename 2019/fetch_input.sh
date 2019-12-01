#!/bin/sh

# The directory where this script is located.
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

SESSION_FILE="$SCRIPT_DIR/session-token"

if [ ! -f "$SESSION_FILE" ]
then
  echo "session file not found"
fi

SESSION=$(cat "$SESSION_FILE")

if [ -z "$SESSION" ]
then
  echo "session file empty"
fi

DAY=$1

if [ -z "$DAY" ]
then
  echo "must specify day"
fi

if [ "$DAY" -lt 10 ]
then
  OUTFILE="day0$DAY"
else
  OUTFILE="day$DAY"
fi

curl \
  --fail \
  --silent \
  --cookie "session=$SESSION" \
  "https://adventofcode.com/2019/day/$DAY/input" > "$OUTFILE/input"
