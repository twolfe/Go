package individual

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Individual struct {
	Gene       string
	Pos        int
	Ref        string
	Alt        string
	Ch         string
	P1         string
	P2         string
	Diagnostic bool //for string organised code
}
