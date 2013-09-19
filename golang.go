package snippet

import (
    "regexp"
    "strings"
)


// Eventually create the language package once you have more than one language to support
type Language interface {
    HandleNewLine(line string)   string
    IsNewBlock(blockIden string) bool
    IsInBlock()                  bool
    IsEndBlock()                 bool
}

type Golang struct {
    bracketCount int
    InBlock      bool
    Line         string
    MaxStrings   int
}

func (gl *Golang) HandleNewLine(line string) string {
    gl.Line = line
    if strings.Contains(line, "{") {
        gl.bracketCount += strings.Count(line, "{") - gl.countCharInString(line, "{")
    }
    if strings.Contains(line, "}") {
        gl.bracketCount -= strings.Count(line, "}") - gl.countCharInString(line, "}")
    }
    return line
}

func (gl *Golang) IsNewBlock(blockIden string) bool {
    //fmt.Println(gl.bracketCount, strings.HasPrefix(gl.Line, blockIden), blockIden, gl.Line)
    if gl.bracketCount == 1 && strings.HasPrefix(gl.Line, blockIden) {
        gl.InBlock = true
        return true    
    } else {
        return false    
    }
}

func (gl *Golang) IsInBlock() bool {
    return gl.InBlock    
}

func (gl *Golang) IsEndBlock() bool {
    if gl.bracketCount == 0 && gl.InBlock {
        gl.InBlock = false
        return true    
    } else {
        return false    
    }
}

func (gl *Golang) countCharInString(line, char string) int {
    r, _ := regexp.Compile(`"(`+char+`.*?)"`)
    s := r.FindAllString(line, gl.MaxStrings)
    count := 0
    for _, found := range(s) {
        count += strings.Count(found, char)
    }
    return count
}
