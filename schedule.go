//custom widget for terminal schedule.

package main

import(
  "fmt"
  "image"
  . "github.com/gizak/termui/v3"
)

type Schedule struct {
  Block

  procs []Process
  MaxVal float64
  HorizontalScale int
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
	// draw origin cell
	buf.SetCell(
		NewCell(BOTTOM_LEFT, NewStyle(ColorWhite)),
		image.Pt(self.Inner.Min.X+yAxisLabelsWidth, self.Inner.Max.Y-xAxisLabelsHeight-1),
	)
	// draw x axis line
	for i := yAxisLabelsWidth + 1; i < self.Inner.Dx(); i++ {
		buf.SetCell(
			NewCell(HORIZONTAL_DASH, NewStyle(ColorWhite)),
			image.Pt(i+self.Inner.Min.X, self.Inner.Max.Y-xAxisLabelsHeight-1),
		)
	}
	// draw y axis line
	for i := 0; i < self.Inner.Dy()-xAxisLabelsHeight-1; i++ {
		buf.SetCell(
			NewCell(VERTICAL_DASH, NewStyle(ColorWhite)),
			image.Pt(self.Inner.Min.X+yAxisLabelsWidth, i+self.Inner.Min.Y),
		)
	}
	// draw x axis labels
	// draw 0
	buf.SetString(
		"0",
		NewStyle(ColorWhite),
		image.Pt(self.Inner.Min.X+yAxisLabelsWidth, self.Inner.Max.Y-1),
	)
	// draw rest
	for x := self.Inner.Min.X + yAxisLabelsWidth + (xAxisLabelsGap)*self.HorizontalScale + 1; x < self.Inner.Max.X-1; {
		label := fmt.Sprintf(
			"%d",
			(x-(self.Inner.Min.X+yAxisLabelsWidth)-1)/(self.HorizontalScale)+1,
		)
		buf.SetString(
			label,
			NewStyle(ColorWhite),
			image.Pt(x, self.Inner.Max.Y-1),
		)
		x += (len(label) + xAxisLabelsGap) * self.HorizontalScale
	}
	// draw y axis labels
	//verticalScale := maxVal / float64(self.Inner.Dy()-xAxisLabelsHeight-1)
  for i := 0; i < len(self.procs); i++ {
    self.procs[i].curheight = -1
  }
	for i := 0; i < len(self.procs); i++ {
    self.procs[i].curheight = self.Inner.Max.Y-(i*(yAxisLabelsGap+1))-5
		buf.SetString(
			fmt.Sprintf(self.procs[i].name),
			NewStyle(ColorWhite),
			image.Pt(self.Inner.Min.X, self.procs[i].curheight),
		)
	}
}

func (self *Schedule) plotSchedules(buf *Buffer) {
  for i := 0; i < len(self.procs); i++ {
    if self.procs[i].curheight != -1 {
      mark := 0
      for j := yAxisLabelsWidth + 2; j < self.Inner.Dx(); j++ {
        if mark >= len(self.procs[i].sched) {
          buf.SetCell(
              NewCell(' ', NewStyle(ColorClear)),
              image.Pt(j, self.procs[i].curheight),
          )
        } else {
          if self.procs[i].sched[mark] == 1 {
            buf.SetCell(
                NewCell(' ', NewStyle(ColorClear, ColorGreen)),
		            image.Pt(j, self.procs[i].curheight),
	          )
          } else {
            buf.SetCell(
                NewCell(' ', NewStyle(ColorClear)),
		            image.Pt(j, self.procs[i].curheight),
	          )
          }
        }
        mark++
      }
    }
  }
}

func (self *Schedule) generateEDF() {
  for time := 1; time < 100; time++ {
    for proc := 0; proc < len(self.procs); proc++ {
      
    }
  }
}

func (self *Schedule) generateRMS() {

}

func (self *Schedule) Draw(buf *Buffer){
  self.Block.Draw(buf)

  self.plotAxes(buf, 10)

  self.plotSchedules(buf)
}
