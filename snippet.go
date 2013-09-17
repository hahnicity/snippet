package snippet

import (
    "bufio"
    "fmt"
    "github.com/hahnicity/go-stringit"
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

func (pf *ParseFile) ParseFor(block string) {
    file := pf.OpenFile()
    defer file.Close()
    c := Code{blockName: block, file: file}
    c.GetCodeBlocks()
    c.WriteLines(pf.FuncOutFile)
}

func (pf *ParseFile) ParseForFunc() {
    pf.ParseFor("func")
}

func (pf *ParseFile) ParseForType() {
    pf.ParseFor("type")
}

type Code struct {
    blockName string
    chunk     []string
    file      *os.File
    allLines  []byte
}

func (c *Code) GetCodeBlocks() {
    // Eventually create interface to swap Golang with diff. languages
    gl := &Golang{0, false, ""}
    scanner := bufio.NewScanner(c.file)
    fmt.Println(c.blockName)
    for scanner.Scan() {
        line := gl.HandleNewLine(scanner.Text())
        if gl.IsNewBlock(c.blockName) {
            c.resetChunk()
            c.chunk = append(c.chunk, line + "\n")
        } else if gl.IsEndBlock() {
            c.handleLastLine(line)
        } else if gl.InBlock {
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
    fmt.Println(config.DescQuery)
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
    fmt.Print(config.ImportantQuery)
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
    return stringit.Format(
        "// {} {}{}{}", 
        strings.ToUpper(c.blockName), 
        config.DescPrefix, 
        scanner.Text(), 
        config.DescSuffix,
    )
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
