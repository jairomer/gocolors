package gtestcoloring

import (
  "go-cmake-colors/color"
  "regexp"
  "log"
)

const OK      string = "OK"
const ERROR   string = "ERROR"
var OK_sub    string = color.Green + OK + color.Reset
var ERROR_sub string = color.Red + ERROR + color.Reset

type GTestColorer struct {
  okRegex  *regexp.Regexp
  errRegex *regexp.Regexp
}

func CreateGTestColorer() *GTestColorer {
    errRegex, err := regexp.Compile(ERROR)
    okRegex, err := regexp.Compile(OK)
    if  err != nil {
      log.Fatalln(err)
    }
    gwriter := GTestColorer {
      okRegex: okRegex,
      errRegex: errRegex,
    }
    return &gwriter
}

func (writer* GTestColorer) Colorize(b string) string {
  // Apply regular expressions
  b = writer.okRegex.ReplaceAllString(b, OK_sub)
  b = writer.errRegex.ReplaceAllString(b, ERROR_sub)
  return b
}
