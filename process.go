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
  ctime int
  ready bool
  complete bool
}

func NewProc(n string, p int, c int) Process {
  p := Process{
    name:     n,
    period:   p,
    ctime:    c,
    ready:    true,
    complete: false,
  }
}
