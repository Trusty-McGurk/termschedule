package main

import (
	"log"
	ui "github.com/gizak/termui/v3"
	//"fmt"
	"github.com/gizak/termui/v3/widgets"
	//term "github.com/nsf/termbox-go"
)

func main() {

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

  sched := NewSchedule()
	t1 := NewProc("T3", 16, 8)
	t2 := NewProc("T2", 12, 3)
	t3 := NewProc("T1", 12, 3)
	sched.AddTask(t1)
	sched.AddTask(t2)
	sched.AddTask(t3)
  sched.SetRect(0, 0, 64, 20)
	if sched.is_edf {
		sched.Title = " EDF Schedule "
	} else {
		sched.Title = " RMS Schedule "
	}

	infobox := widgets.NewParagraph()
	infobox.Title = " Task Info "
	infobox.Text = sched.TaskInfoString()
	infobox.SetRect(65, 0, 80, 20)

	terminal := widgets.NewParagraph()
	terminal.Text = ""
	terminal.SetRect(0, 20, 80, 23)

	tb := NewTextBuffer()

  ui.Render(sched, infobox, terminal)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "<C-c>":
			return
		case "<Tab>":
			sched.is_edf = !sched.is_edf
			if sched.is_edf {
				sched.Title = " EDF Schedule "
			} else {
				sched.Title = " RMS Schedule "
			}
		case "<Backspace>":
			tb.Del()
		case "<Space>":
			tb.AddChar(" ")
		case "<Enter>":
			sched.ParseCommand(tb.buffer)
			tb.Clear()
		default:
			tb.AddChar(e.ID)
		}
		terminal.Text = tb.buffer
		infobox.Text = sched.TaskInfoString()
		ui.Render(sched, infobox, terminal)
	}

}
