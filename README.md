snippet
=======

A code snippet manager for go

## Purpose
Use this manager when you want to parse functions and types in a .go file into
separate files so that you can manage them as a resposity of code snippets in the
future.

## Usage
The manager can be run through the `main.go` console script

        main/main.go -p path/to/my/file.go

There are additional options that can be specified as well

 * --find-funcs => boolean you can set to false if you do not want to parse for functions
 * --find-types => boolean you can set to false if you do not want to parse for types
 * --func-file => The name of the file we can write all functions to
 * --type-file => The name of the file we can write all types to
