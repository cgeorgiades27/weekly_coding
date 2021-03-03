#!/bin/bash

echo "g++ -std=c++11 -Wall -Wextra -o $1.x $1.cpp"
g++ -std=c++11 -Wall -Wextra -o $1.x $1.cpp