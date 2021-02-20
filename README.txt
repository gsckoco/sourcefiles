This is a simple program I made because I am lazy.
It grabs all the source files in a given directory
recursively, grabbing any source files in child
directories and outputting them to be taken as the
input into gcc.

Usage:
    sourcefiles <path> <File extension 1> [File extension 2 ...]

Example:
    $ sourcefiles ~/Projects/lmc/ h cpp
    > /home/ben/Projects/lmc/src/cpu.h /home/ben/Projects/lmc/src/main.cpp /home/ben/Projects/lmc/src/main.h 

Peak lazyness but it works