package main

import(
  "strconv"
  "strings"
)

type TextBuffer struct {
  maxsize int
  cursize int
  buffer string
}

func NewTextBuffer() TextBuffer {
  return TextBuffer{
    maxsize:  50,
    cursize:  0,
    buffer:   "",
  }
}

func (self *TextBuffer) AddChar(s string) {
  if self.cursize != self.maxsize {
    self.buffer += s
    self.cursize += len(s)
  }
}

func (self *TextBuffer) Del() {
  if self.cursize != 0 {
    self.buffer = self.buffer[0:len(self.buffer) - 1]
    self.cursize--
  }
}

func (self *TextBuffer) Clear() {
  self.cursize = 0
  self.buffer = ""
}

func ShortName(s string) string {
  if len(s) < 7 {
    return s
  } else {
    return s[0:3] + "..."
  }
}

func (self *Schedule) TaskInfoString() string {
  retstring := ""
  for i, task := range self.procs {
    retstring += task.shortname + ":\n"
    retstring += "(" + strconv.Itoa(task.ctime) + ", " + strconv.Itoa(task.period) + ")\n"
    if i < len(self.procs) - 1{
      retstring += "\n"
    }
  }
  return retstring
}

func (self *Schedule) AddTask(p Process){
  self.procs = append(self.procs, p)
}

func (self *Schedule) ExactAnalysisString() string {
  retstring := ""
  for i, task := range self.procs {
    retstring += task.shortname + ":\n"
    if task.passEE {
      retstring += "[PASS](fg:green)\n"
    } else {
      retstring += "[FAIL](fg:red)\n"
    }
    if i < len(self.procs) - 1{
      retstring += "\n"
    }
  }
  return retstring
}

func (self *Schedule) ParseCommand(s string){
  args := strings.Fields(s)
  if len(args) == 0 {
    return
  }

  switch args[0] {
  case "add":
    if len(args) == 4 && IsNumber(args[2]) && IsNumber(args[3]){
      arg1, _ := strconv.Atoi(args[3])
      arg2, _ := strconv.Atoi(args[2])
      self.AddTask(NewProc(args[1], arg1, arg2, *self))
    }
  case "del":
    if len(args) == 2 {
      for i, task := range self.procs {
        if task.name == args[1] {
          copy(self.procs[i:], self.procs[i+1:])
          self.procs = self.procs[:len(self.procs)-1]
          if self.windowbase != 0 {
            self.windowbase--
          }
          break
        }
      }
    }
  }
}

func IsNumber(s string) bool {
  if _, err := strconv.Atoi(s); err == nil {
    return true
  }
  return false
}
