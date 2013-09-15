package main

import (
    "flag"
    "github.com/hahnicity/snippet"
)

var (
    parseForType bool
    parseForFunc bool
)

func initializeParseFile() *snippet.ParseFile {
    return new(snippet.ParseFile)
}

func ParseArgs(pf *snippet.ParseFile) {
    flag.StringVar(
        &pf.FilePath, 
        "p", 
        "", 
        "The abs/relative path to the file we wish to parse",
    )
    flag.StringVar(
        &pf.FuncOutFile, 
        "func-file", 
        "funcfile.txt", 
        "The name of the file we wish to write functions to",
    )
    flag.StringVar(
        &pf.TypeOutFile,
        "type-file",
        "typefile.txt",
        "The name of the file we wish to write types to",
    )
    flag.BoolVar(
        &parseForFunc,
        "find-funcs",
        true,
        "Parse for functions in our file",
    )
    flag.BoolVar(
        &parseForType,
        "find-types",
        true,
        "Parse for types in our file",
    )
    flag.Parse()
}

func main() {
    pf := initializeParseFile()
    ParseArgs(pf)
    if parseForFunc {
        pf.ParseForFunc()   
    }
    if parseForType {
        pf.ParseForType()
    }
}
