package main

import (
	"log"
	ui "github.com/gizak/termui/v3"
	"strconv"
	"fmt"
	"github.com/gizak/termui/v3/widgets"
	//term "github.com/nsf/termbox-go"
)

// func main() {
// 	sched := NewSchedule()
// 	sched.AddTask(NewProc("t1", 8, 4, *sched))
// 	sched.AddTask(NewProc("t2", 8, 3, *sched))
// 	sched.AddTask(NewProc("t3", 8, 1, *sched))
// 	sched.ExactAnalysis()
// 	fmt.Println(sched.ExactAnalysisString())
// }

func main() {

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

  sched := NewSchedule()
  sched.SetRect(0, 0, 64, 20)
	if sched.is_edf {
		sched.Title = " EDF Schedule "
	} else {
		sched.Title = " RMS Schedule "
	}

	eainfo := false
	infobox := widgets.NewParagraph()
	infobox.Title = " Task Info "
	infobox.Text = sched.TaskInfoString()
	infobox.SetRect(65, 0, 80, 20)

	terminal := widgets.NewParagraph()
	terminal.Text = "_"
	terminal.SetRect(0, 20, 50, 23)

	ctxswitchbox := widgets.NewParagraph()
	ctxswitchbox.SetRect(65, 20, 80, 24)
	ctxswitchbox.Text = "Period: " + strconv.Itoa(sched.GetPeriod()) + "\nCTXS: " + strconv.Itoa(sched.ctxswitches)

	schedanalysisbox := widgets.NewParagraph()
	schedanalysisbox.SetRect(50, 20, 65, 24)
	schedanalysisbox.Text = ""

	tb := NewTextBuffer()

  ui.Render(sched, infobox, terminal, ctxswitchbox, schedanalysisbox)

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
		case "<C-e>":
			eainfo = !eainfo
		case "<Backspace>":
			tb.Del()
		case "<Space>":
			tb.AddChar(" ")
		case "<Enter>":
			sched.ParseCommand(tb.buffer)
			tb.Clear()
		case "<Up>":
			if sched.windowbase < len(sched.procs) - 4 {
				sched.windowbase++
			}
		case "<Down>":
			if sched.windowbase > 0 {
				sched.windowbase--
			}
		case "<Left>":
			if sched.schedulebase > 0 {
				sched.schedulebase -= 5
			}
		case "<Right>":
			sched.schedulebase += 5
		default:
			tb.AddChar(e.ID)
		}
		terminal.Text = tb.buffer + "_"
		if eainfo {
			infobox.Title = " EA Info "
			infobox.Text = sched.ExactAnalysisString()
		} else {
			infobox.Title = " Task Info "
			infobox.Text = sched.TaskInfoString()
		}
		ui.Render(sched, infobox, terminal)
		ctxswitchbox.Text = "Period: " + strconv.Itoa(sched.GetPeriod()) + "\nCTXS: " + strconv.Itoa(sched.ctxswitches)
		ui.Render(ctxswitchbox)
		analysisstring := ""
		if len(sched.procs) > 0 {
			total, target, pass := sched.ScheduabilityAnalysis()
			analysisstring = fmt.Sprintf("%.2f ≤ %.2f\n", total, target)
			if pass {
				analysisstring += "[PASS](fg:green)"
			} else {
				analysisstring += "[FAIL](fg:red)"
			}
		}
		schedanalysisbox.Text = analysisstring
		ui.Render(schedanalysisbox)
	}

}
