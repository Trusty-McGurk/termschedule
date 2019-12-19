package main

func (self *Schedule) ExactAnalysis() {
  tasklist := make([]*Process, len(self.procs))
  for i := 0; i < len(self.procs); i++ {
    tasklist[i] = &(self.procs[i])
  }
  for _, EETarget := range tasklist {
    t := 0
    tnext := 0
    Round:
      for {
        for _, task := range tasklist {
          if EETarget.period >= task.period {
            tnext = tnext + ((1 + (t/task.period)) * task.ctime)
          }
        }
        if tnext == t{
          EETarget.passEE = true
          break Round
        } else if tnext > (1 + (t/EETarget.period)) * EETarget.period {
          EETarget.passEE = false
          break Round
        }
        t = tnext
        tnext = 0
      }
  }
}
