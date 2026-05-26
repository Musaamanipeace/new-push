#!/bin/bash
# Generate unique randomized sequence lists for large array evaluation
exercise_Count=${1:-100}
if [[ "$OSTYPE" == "darwin"* ]]; then
    jot -r "$exercise_Count" -2000 2000 | awk '!x[$0]++' | tr '\n' ' '
else
    shuf -i 1-2000 -n "$exercise_Count" | tr '\n' ' '
fi
echo ""