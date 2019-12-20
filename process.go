package main

// import (
//   "fmt"
// )

type Process struct {
  name string
  shortname string

  sched []int
  curheight int
  period int
  nextdeadline int
  ctime int
  workdone int
  passEE bool
}

func NewProc(n string, p int, c int, s Schedule) Process {
  proc := Process{
    name:     n,
    period:   p,
    ctime:    c,
    workdone: 0,
  }
  proc.sched = make([]int, s.schedulesize)
  proc.shortname = ShortName(n)
  return proc
}
