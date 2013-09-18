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

 Upon running the manager will go through all the code blocks and pick out all
 functions/types. Upon finding a function the user will be queried:

        Is this chunk of code important? (y/N):
        
If `y` is input the code will be saved to file. If `N` then the opposite. We can
input a function/type description after being querried if the chunk of code is 
important to us. The query should look like this;

        Please enter a brief description for this chunk of code:       
