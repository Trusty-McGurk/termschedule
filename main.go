package main

import (
	"log"
	ui "github.com/gizak/termui/v3"
	//"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

  sched := NewSchedule()
  sched.procs = append(sched.procs, NewProc("proc1"))
	sched.procs = append(sched.procs, NewProc("proc2"))
	sched.procs[0].sched = []int{1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0}
	sched.procs[1].sched = []int{0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0}
  sched.SetRect(0, 0, 80, 20)

  ui.Render(sched)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
