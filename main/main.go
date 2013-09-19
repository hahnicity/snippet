package main

import (
    "flag"
    "github.com/hahnicity/snippet"
    "github.com/hahnicity/go-stringit"
)

var (
    language     string
    parseForType bool
    parseForFunc bool
    saveAll      bool
)

func getParseFile() *snippet.ParseFile {
    pf := new(snippet.ParseFile)
    parseArgs(pf)
    return pf
}

func parseArgs(pf *snippet.ParseFile) {
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
    flag.StringVar(
        &language,
        "lang",
        "golang",
        "Language to parse",
    )
    flag.BoolVar(
        &saveAll,
        "save-all",
        false,
        "Set to true to save all functions/types to file automatically",
    )
    flag.Parse()
}

func getLanguage() snippet.Language {
    if language == "golang" {
        return &snippet.Golang{InBlock: false, Line: "", MaxStrings: 10}
    } else {
        panic(stringit.Format("The following language {} is not supported", language))    
    }
}

func main() {
    pf := getParseFile()
    l := getLanguage()
    if parseForFunc {
        pf.ParseForFunc(l, saveAll)   
    }
    if parseForType {
        pf.ParseForType(l, saveAll)
    }
}
