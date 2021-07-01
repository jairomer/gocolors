package gtestcoloring

import (
  "go-cmake-colors/color"
  "regexp"
  "io"
  "log"
)

// OK : Regular expression to match for success cases.
const OK      string = "OK"
// ERROR : Regular expression to match for error cases.
const ERROR   string = "ERROR"
// OkSub : Substitution for success cases.
var OkSub    string = color.Green + OK + color.Reset
// ErrorSub : Substitution for error cases.
var ErrorSub string = color.Red + ERROR + color.Reset

// GTestColorer : Container class for compiled regular expressions.
type GTestColorer struct {
  okRegex  *regexp.Regexp
  errRegex *regexp.Regexp
}

// CreateGTestColorer : Create an instance of a colorer for gTest.
func CreateGTestColorer() *GTestColorer {
  var (
    errRegex = regexp.MustCompile(ERROR)
    okRegex  = regexp.MustCompile(OK)
  )
  return &GTestColorer {
    okRegex: okRegex,
    errRegex: errRegex,
  }
}

// Colorize : Apply colorings to a received string and feed it to a writer.
func (writer* GTestColorer) Colorize(b string, w io.Writer) {
  // Apply regular expressions
  b = writer.okRegex.ReplaceAllString(b, OkSub)
  b = writer.errRegex.ReplaceAllString(b, ErrorSub)
  _, err := io.WriteString(w,b)
  if err != nil {
		log.Fatalln(err)
  }
}
