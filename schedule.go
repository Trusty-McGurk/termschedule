//custom widget for terminal schedule.

package main

import(
  "fmt"
  "image"
  . "github.com/gizak/termui/v3"
  "strconv"
)

type Schedule struct {
  Block

  procs []Process
  MaxVal float64
  HorizontalScale int
  is_edf bool
  windowbase int
  schedulebase int
  schedulesize int
}

const (
	xAxisLabelsHeight = 1
	yAxisLabelsWidth  = 6
	xAxisLabelsGap    = 1
	yAxisLabelsGap    = 3
)

func NewSchedule() *Schedule {
  return &Schedule{
    Block:            *NewBlock(),
    HorizontalScale:  2,
    is_edf:           true,
    windowbase:       0,
    schedulebase:     0,
    schedulesize:     100,
  }
}
/*
func (self *Schedule) drawLine(buf *Buffer, maxVal float64){
  barWidth := int((float64(50) / 100) * float64(self.Inner.Dx()))
	buf.Fill(
		NewCell(' ', NewStyle(ColorClear, ColorGreen)),
		image.Rect(self.Inner.Min.X, self.Inner.Min.Y, self.Inner.Min.X+barWidth, self.Inner.Max.Y),
	)
}
*/

func (self *Schedule) plotAxes(buf *Buffer, maxVal float64) {
	// draw x axis line
	for i := yAxisLabelsWidth; i < self.Inner.Dx(); i++ {
		buf.SetCell(
			NewCell(HORIZONTAL_DASH, NewStyle(ColorWhite)),
			image.Pt(i+self.Inner.Min.X - 1, self.Inner.Max.Y-xAxisLabelsHeight-2),
		)
	}
	// draw y axis line
	for i := 0; i < self.Inner.Dy()-xAxisLabelsHeight-1; i++ {
		buf.SetCell(
			NewCell('⎸', NewStyle(ColorWhite)),
			image.Pt(self.Inner.Min.X+yAxisLabelsWidth, i+self.Inner.Min.Y),
		)
	}

	// draw x axis labels
	// draw 0
	buf.SetString(
		strconv.Itoa(self.schedulebase),
		NewStyle(ColorBlue),
		image.Pt(self.Inner.Min.X+yAxisLabelsWidth, self.Inner.Max.Y-2),
	)
	// draw rest
  mark := 1
	for x := self.Inner.Min.X + yAxisLabelsWidth + 5; x < self.Inner.Max.X-1; {
    label := fmt.Sprintf(
			"%d",
			mark*5 + self.schedulebase,
		)
		buf.SetString(
			label,
			NewStyle(ColorBlue),
			image.Pt(x, self.Inner.Max.Y-2),
		)
		x += 5
    mark++
	}
	// draw y axis labels
	//verticalScale := maxVal / float64(self.Inner.Dy()-xAxisLabelsHeight-1)
  for i := 0; i < len(self.procs); i++ {
    self.procs[i].curheight = -1
  }
	for i := 0; i < len(self.procs) && i < 4; i++ {
    self.procs[i + self.windowbase].curheight = self.Inner.Max.Y-(i*(yAxisLabelsGap+1))-5
		buf.SetString(
			fmt.Sprintf(self.procs[i + self.windowbase].name),
			NewStyle(ColorWhite),
			image.Pt(self.Inner.Min.X, self.procs[i + self.windowbase].curheight),
		)
	}
}

func (self *Schedule) plotSchedules(buf *Buffer) {
  for i := self.windowbase; i < len(self.procs) && i < self.windowbase + 4; i++ {
    if self.procs[i].curheight != -1 {
      mark := self.schedulebase
      for j := yAxisLabelsWidth + 1; j < self.Inner.Dx(); j++ {
        if mark >= len(self.procs[i].sched) {
          buf.SetCell(
              NewCell(' ', NewStyle(ColorClear)),
              image.Pt(j, self.procs[i].curheight),
          )
        } else {
          if self.procs[i].sched[mark] == 1 {
            buf.SetCell(
                NewCell('⎸', NewStyle(ColorWhite, ColorGreen)),
		            image.Pt(j, self.procs[i].curheight),
	          )
          } else if self.procs[i].sched[mark] == 2{
            buf.SetCell(
                NewCell('⎸', NewStyle(ColorWhite, ColorRed)),
		            image.Pt(j, self.procs[i].curheight),
	          )
          } else if self.procs[i].sched[mark] == 3{
            buf.SetCell(
                NewCell('⎸', NewStyle(ColorMagenta, ColorClear)),
		            image.Pt(j, self.procs[i].curheight),
	          )
          } else {
            buf.SetCell(
                NewCell('⎸', NewStyle(ColorWhite, ColorClear)),
		            image.Pt(j, self.procs[i].curheight),
	          )
          }
        }
        if mark != self.schedulebase && mark % 5 == 0 {
          buf.SetCell(
              NewCell('⎸', NewStyle(ColorBlue, ColorClear)),
              image.Pt(j, self.procs[i].curheight + 1),
          )
          buf.SetCell(
              NewCell('⎸', NewStyle(ColorBlue, ColorClear)),
              image.Pt(j, self.procs[i].curheight - 1),
          )
        }
        if mark == self.schedulebase || mark % self.procs[i].period == 0 {
          buf.SetCell(
              NewCell('⎸', NewStyle(ColorYellow, ColorClear)),
              image.Pt(j, self.procs[i].curheight + 1),
          )
          buf.SetCell(
              NewCell('⎸', NewStyle(ColorYellow, ColorClear)),
              image.Pt(j, self.procs[i].curheight - 1),
          )
        }
        mark++
      }
    }
  }
}

func (self *Schedule) IncreaseScheduleLength() {
  for i, task := range self.procs {
    tmp := make([]int, self.schedulesize*2)
    copy(tmp, task.sched)
    self.procs[i].sched = tmp
    self.schedulesize *= 2
  }
}

func (self *Schedule) Draw(buf *Buffer){
  self.Block.Draw(buf)
  self.plotAxes(buf, 10)
  if len(self.procs) != 0 {
    if self.schedulebase + (self.Inner.Dx() - (yAxisLabelsWidth + 1)) > len(self.procs[0].sched) - 10{
      self.IncreaseScheduleLength()
    }
    if self.is_edf {
      self.generateEDF()
    } else {
      self.generateRMS()
    }
    self.plotSchedules(buf)
    self.ExactAnalysis()
  }
}
