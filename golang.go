package snippet

//import "fmt"
import "strings"


type Language interface {
    HandleNewLine(line string)   string
    IsNewBlock(blockIden string) bool
    IsEndBlock()                 bool
}

type Golang struct {
    bracketCount int
    InBlock      bool
    Line         string
}

func (gl *Golang) HandleNewLine(line string) string {
    gl.Line = line
    if strings.Contains(line, "{") && !gl.InBlock {
        // XXX This behavior is ok if our style is fairly idiomatic and brackets are not in strings. Otherwise it is untenable.
        gl.bracketCount = 1  
    } else if strings.Contains(line, "{") && gl.InBlock {
        gl.bracketCount += 1
    }
    if strings.Contains(line, "}") {
        gl.bracketCount -= 1    
    }
    return line
}

func (gl *Golang) IsNewBlock(blockIden string) bool {
    //fmt.Println(gl.bracketCount, strings.HasPrefix(gl.Line, blockIden), blockIden, gl.Line)
    // XXX The behavior here needs to be de-coupled from the function above
    if gl.bracketCount == 1 && strings.HasPrefix(gl.Line, blockIden) {
        gl.InBlock = true
        return true    
    } else {
        return false    
    }
}

func (gl *Golang) IsEndBlock() bool {
    if gl.bracketCount == 0 && gl.InBlock {
        gl.InBlock = false
        return true    
    } else {
        return false    
    }
}
