package snippet

import (
    "bufio"
    "fmt"
    "github.com/hahnicity/snippet/config"
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
    file := pf.OpenFile()
    defer file.Close()
    c := Code{blockName: "func", file: file}
    c.GetCodeBlocks()
    c.WriteLines(pf.FuncOutFile)
}

func (pf *ParseFile) ParseForType() {
    file := pf.OpenFile()
    defer file.Close()
    c := Code{blockName: "type", file: file}
    c.GetCodeBlocks()
    c.WriteLines(pf.TypeOutFile)
}

type Code struct {
    blockName string
    chunk     []string
    file      *os.File
    allLines  []byte
}

func (c *Code) GetCodeBlocks() {
    inBlock := false  // Boolean for determining if we are in our block of code
    scanner := bufio.NewScanner(c.file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, c.blockName) {
            c.resetChunk()
            c.chunk = append(c.chunk, line + "\n")
            inBlock = true
        } else if inBlock && strings.HasPrefix(line, "}") {
            c.handleLastLine(line)
            inBlock = false
        } else if inBlock {
            c.chunk = append(c.chunk, line + "\n")
        }
    }
}

func (c *Code) handleLastLine(line string) {
    c.chunk = append(c.chunk, line + config.EndBlockSuffix)
    fmt.Printf("%s \n", c.chunk)
    if !c.isChunkImportant() {
        return 
    }
    desc := c.scanForDescription()
    c.insertDescription(desc)
    c.transferCodeToLines()
}

func (c *Code) insertDescription(desc string) {
    c.chunk = append(c.chunk, "")
    copy(c.chunk[0+1:], c.chunk[0:])
    c.chunk[0] = desc
}

func (c *Code) isChunkImportant() bool {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println(config.ImportantQuery)
    for scanner.Scan() {
        if scanner.Text() == config.ImportantNegative {
            return false    
        } else if scanner.Text() == config.ImportantAffirmative {
            return true    
        } else {
            fmt.Println(config.ImportantRetryQuery)    
        }
    }
    return true
}

func (c *Code) resetChunk() {
    c.chunk = make([]string, 0)    
}

func (c *Code) scanForDescription() string {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    return "// " + strings.ToUpper(c.blockName) + " " + config.DescPrefix + 
        scanner.Text() + config.DescSuffix
}

func (c *Code) transferCodeToLines() {
    for _, line := range c.chunk {
        c.allLines = append(c.allLines, line...)   
    }
}

func (c *Code) WriteLines(path string) {
    file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, os.ModePerm)
    if err != nil {
        file, err = os.Create(path)
        if err != nil {
            panic(err)    
        }
    }
    defer file.Close()
    _, err = file.Write(c.allLines)
    if err != nil {
        panic(err)    
    }
}
