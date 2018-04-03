package main


import (
  "fmt"
  "os"

  "github.com/gdamore/tcell"
)

func main() {
  s, e := tcell.NewScreen()
  if e != nil {
    fmt.Fprintf(os.Stderr, "%v\n", e)
    os.Exit(1)
  }
  if e = s.Init(); e != nil {
    fmt.Fprintf(os.Stderr, "%v\n", e)
    os.Exit(1)
  }
  s.Clear()
  s.SetContent(0,0,'a',nil,tcell.StyleDefault)
  s.Show()
loop:
  for {
    ev := s.PollEvent()
    switch ev := ev.(type) {
    case *tcell.EventKey:
      switch ev.Key() {
      case tcell.KeyEscape:
        break loop
      }
    case *tcell.EventResize:
      s.Sync()
    }
  }
  s.Fini()
}
