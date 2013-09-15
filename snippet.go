package snippet

import (
    "bufio"
    "os"
    "strings"
)


type ParseFile struct {
    FilePath, FuncOutFile, TypeOutFile string
}

func (pf *ParseFile) OpenFile() *os.File {
    f, err := os.OpenFile(pf.FilePath, os.O_RDONLY, os.ModePerm)
    if err != nil {
        panic(err)    
    }
    return f
}

func (pf *ParseFile) ParseForFunc() {
    f := pf.OpenFile()
    defer f.Close()
    blocks := GetCodeBlocks(f, "func")
    WriteBlocksToFile(blocks, pf.FuncOutFile)
}

func (pf *ParseFile) ParseForType() {
    f := pf.OpenFile()
    defer f.Close()
    blocks := GetCodeBlocks(f, "type")
    WriteBlocksToFile(blocks, pf.TypeOutFile)
}

func GetCodeBlocks(f *os.File, block string) []byte {
    lines := make([]byte, 0)
    inBlock := false  // Boolean for determining if we are in our block of code
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, block) {
            lines = append(lines, line+"\n"...)
            inBlock = true
        } else if inBlock && strings.HasPrefix(line, "}") {
            lines = append(lines, line+"\n"...)
            inBlock = false
        } else if inBlock {
            lines = append(lines, line+"\n"...)
        }
    }
    return lines
}

func WriteBlocksToFile(blocks []byte, path string) {
    file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, os.ModePerm)
    if err != nil {
        file, err = os.Create(path)
        if err != nil {
            panic(err)    
        }
    }
    defer file.Close()
    _, err = file.Write(blocks)
    if err != nil {
        panic(err)    
    }
}
