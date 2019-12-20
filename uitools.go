package main

import(
  "strconv"
  "strings"
  "math"
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
  self.schedulesize = int(math.Max(float64(self.schedulesize), float64(self.GetPeriod())))
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

func (self *Schedule) GetPeriod() int {
  if len(self.procs) == 0 {
    return 0
  }
  if len(self.procs) == 1 {
    return self.procs[0].period
  }
  periods := make([]int, len(self.procs))
  for i, task := range self.procs {
    periods[i] = task.period
  }
  return LCM(periods)
}

func IsNumber(s string) bool {
  if _, err := strconv.Atoi(s); err == nil {
    return true
  }
  return false
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers []int) int {
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM([]int{result, integers[i]})
	}

	return result
}

func (self *Schedule) RMS_Scheduability() (float64, float64, bool) {
  total := float64(0)
  pass := false
  for _, task := range self.procs {
    total += (float64(task.ctime)/float64(task.period))
  }
  target := float64(len(self.procs)) * (math.Pow(2, 1/float64(len(self.procs))) - 1)
  if total <= target {
    pass = true
  }
  return total, target, pass
}

func (self *Schedule) EDF_Scheduability() (float64, bool) {
  total := float64(0)
  pass := false
  for _, task := range self.procs {
    total += (float64(task.ctime)/float64(task.period))
  }
  if total <= 1 {
    pass = true
  }
  return total, pass
}

func (self *Schedule) ScheduabilityAnalysis() (float64, float64, bool) {
  if self.is_edf {
    total, pass := self.EDF_Scheduability()
    return total, 1, pass
  } else {
    total, target, pass := self.RMS_Scheduability()
    return total, target, pass
  }
}
