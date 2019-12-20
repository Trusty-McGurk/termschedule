package main

import(
//  "fmt"
  "math"
)


func (self *Schedule) ExactAnalysis() {
  tasklist := make([]*Process, len(self.procs))
  for i := 0; i < len(self.procs); i++ {
    tasklist[i] = &(self.procs[i])
  }
  for targetindex, EETarget := range tasklist {
    t := 0
    for taskindex, task := range tasklist {
      if EETarget.period > task.period || (EETarget.period == task.period && targetindex >= taskindex) {
        t += task.ctime
      }
    }
    tnext := 0
    Round:
      for {
        for taskindex, task := range tasklist {
          if EETarget.period > task.period || (EETarget.period == task.period && targetindex <= taskindex) {
            tnext = tnext + (int(math.Ceil(float64(t)/float64(task.period))) * task.ctime)
          }
        }
        if tnext == t {
          EETarget.passEE = true
          break Round
        } else if tnext > int(math.Ceil(float64(t)/float64(EETarget.period))) * EETarget.period {
          EETarget.passEE = false
          break Round
        }
        t = tnext
        tnext = 0
      }
  }
}
