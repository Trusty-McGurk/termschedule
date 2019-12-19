package main

// import(
//   "fmt"
// )

func (self *Schedule) wipeSchedules() {
  tasklist := make([]*Process, len(self.procs))
  for i := 0; i < len(self.procs); i++ {
    tasklist[i] = &(self.procs[i])
  }
  for _, task := range tasklist {
    task.workdone = 0
    for i := 0; i < len(task.sched); i++{
      task.sched[i] = 0
    }
  }
}

func (self *Schedule) getDeadlineSortedProcs() []*Process {
  numprocs := len(self.procs)
  proclist := make([]*Process, numprocs)
  for i := 0; i < len(self.procs); i++ {
    proclist[i] = &(self.procs[i])
  }
  return mergeSortByParameter(proclist, "deadline")
}

func (self *Schedule) getPeriodSortedProcs() []*Process {
  numprocs := len(self.procs)
  proclist := make([]*Process, numprocs)
  for i := 0; i < len(self.procs); i++ {
    proclist[i] = &(self.procs[i])
  }
  return mergeSortByParameter(proclist, "period")
}

func mergeSortByParameter(items []*Process, param string) []*Process {
    var num = len(items)

    if num == 1 {
        return items
    }

    middle := int(num / 2)
    var (
        left = make([]*Process, middle)
        right = make([]*Process, num-middle)
    )
    for i := 0; i < num; i++ {
        if i < middle {
            left[i] = items[i]
        } else {
            right[i-middle] = items[i]
        }
    }
    if param == "deadline" {
      return mergeByDeadline(mergeSortByParameter(left, param), mergeSortByParameter(right, param))
    } else {
      return mergeByPeriod(mergeSortByParameter(left, param), mergeSortByParameter(right, param))
    }
}

func mergeByDeadline(left []*Process, right []*Process) (result []*Process) {
    result = make([]*Process, len(left) + len(right))

    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0].nextdeadline < right[0].nextdeadline {
            result[i] = left[0]
            left = left[1:]
        } else {
            result[i] = right[0]
            right = right[1:]
        }
        i++
    }

    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }

    return
}

func mergeByPeriod(left []*Process, right []*Process) (result []*Process) {
    result = make([]*Process, len(left) + len(right))

    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0].period < right[0].period {
            result[i] = left[0]
            left = left[1:]
        } else {
            result[i] = right[0]
            right = right[1:]
        }
        i++
    }

    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }

    return
}

func (self *Schedule) generateEDF() {
  self.wipeSchedules()
  //initialize the processes.
  for proc := 0; proc < len(self.procs); proc++ {
    self.procs[proc].nextdeadline = self.procs[proc].period
  }

  for time := 0; time < 100; time++ {
    tasklist := make([]*Process, len(self.procs))
    for i := 0; i < len(self.procs); i++ {
      tasklist[i] = &(self.procs[i])
    }
    for _, task := range tasklist {
      if time == task.nextdeadline {
        if task.workdone < task.ctime {
          for i := 1; i <= task.period; i++ {
            if task.sched[time - i] == 1 {
              task.sched[time - i] = 2
            } else {
              task.sched[time - i] = 3
            }
          }
        }
        task.nextdeadline = task.nextdeadline + task.period
        task.workdone = 0
      }
    }

    tasks_sorted := self.getDeadlineSortedProcs()
    //initially, set all tasks to not running.
    for _, task := range tasks_sorted {
      task.sched[time] = 0
    }

    for _, task := range tasks_sorted {
      if task.workdone < task.ctime {
        task.sched[time] = 1
        task.workdone++
        break
      }
    }

  }
}

func (self *Schedule) generateRMS() {
  self.wipeSchedules()
  //initialize the processes.
  for proc := 0; proc < len(self.procs); proc++ {
    self.procs[proc].nextdeadline = self.procs[proc].period
  }

  for time := 0; time < 100; time++ {
    tasklist := make([]*Process, len(self.procs))
    for i := 0; i < len(self.procs); i++ {
      tasklist[i] = &(self.procs[i])
    }
    for _, task := range tasklist {
      if time == task.nextdeadline {
        if task.workdone < task.ctime {
          for i := 1; i <= task.period; i++ {
            if task.sched[time - i] == 1 {
              task.sched[time - i] = 2
            } else {
              task.sched[time - i] = 3
            }
          }
        }
        task.nextdeadline = task.nextdeadline + task.period
        task.workdone = 0
      }
    }

    tasks_sorted := self.getPeriodSortedProcs()
    //initially, set all tasks to not running.
    for _, task := range tasks_sorted {
      task.sched[time] = 0
    }

    for _, task := range tasks_sorted {
      if task.workdone < task.ctime {
        task.sched[time] = 1
        task.workdone++
        break
      }
    }

  }
}
