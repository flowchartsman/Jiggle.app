package main

import (
	"time"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/go-vgo/robotgo"
)

var toggleChan = make(chan bool)

func main() {
	robotgo.MouseSleep = 100
	go dojiggle()
	onExit := func() {
		// now := time.Now()
		// ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
	}

	systray.Run(onReady, onExit)
}

const (
	txtJiggleStart = `Start Jigglin'`
	txtJiggleStop  = `Stop Jigglin'`
	speedX         = 1.0
	speedY         = 10.0
	startX         = 10
	startY         = 20
)

func dojiggle() {
	t := time.NewTicker(5 * time.Second)
	amJigglin := false
	x := startX
	y := startY
	for {
		select {
		case state := <-toggleChan:
			amJigglin = state
		case <-t.C:
			if !amJigglin {
				break
			}
			robotgo.MoveSmoothRelative(x, y, speedX, speedY)
			x = -x
			y = -y
		}
	}
}

func onReady() {
	systray.SetTemplateIcon(iconMacInactive, iconOtherInactive)
	// systray.SetTooltip("Lantern")
	mJiggle := systray.AddMenuItemCheckbox(txtJiggleStart, "", false)
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "")
	mQuit.SetIcon(icon.Data)
	jigglin := false

	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
			case <-mJiggle.ClickedCh:
				jigglin = !jigglin
				if jigglin { // mi.Checked()
					// mi.Check()
					systray.SetTemplateIcon(iconMac, iconOther)
					toggleChan <- true
					mJiggle.SetTitle(txtJiggleStop)
					// mi.SetTemplateIcon(icon.Data, icon.Data)
				} else {
					// mi.Uncheck()
					systray.SetTemplateIcon(iconMacInactive, iconOtherInactive)
					toggleChan <- false
					mJiggle.SetTitle(txtJiggleStart)
				}
			}
		}
	}()
}
