package cli

import (
	"fmt"
	"time"

	"github.com/pterm/pterm"
	"github.com/yogs696/skilltest/app"
	"github.com/yogs696/skilltest/config"
)

var runApp bool

var appCommands = cli{
	argVar:   &runApp,
	argName:  "run",
	argUsage: "--run To run the App services",
	run:      printInfo,
	cb:       printUsage,
}

const (
	// Year and copyright
	yc     = "(c) %v skilltest"
	banner = `
       _     _ _ _                        
     | |   (_) | |   _              _    
  ___| |  _ _| | |  | |_  ____  ___| |_  
 /___) | / ) | | |  |  _)/ _  )/___)  _) 
|___ | |< (| | | |  | |_( (/ /|___ | |__ 
(___/|_| \_)_|_|_|   \___)____|___/ \___)
  %s %s`
)

func printInfo() {
	// App name
	pyc := fmt.Sprintf(yc, time.Now().Year())
	header := fmt.Sprintf(pterm.LightGreen(banner), pterm.Red(app.Version), pterm.LightGreen(pyc))
	pterm.DefaultCenter.Println(header)

	additional := config.Of.App.Desc

	// App version and last build info
	lastBuild := "N/A"
	if app.LastBuildAt != "" && app.LastBuildAt != " " {
		lastBuild = app.LastBuildAt
	}
	additional += fmt.Sprintf("\nLast Build Binary at: %v", lastBuild)

	// Print additional info
	pterm.DefaultCenter.WithCenterEachLineSeparately().Println(pterm.Cyan(additional))

	// Comand list and usage headers
	fmt.Println(" Usage: --<argument>...")
	fmt.Println("")
	fmt.Println(" Arguments:")
}

func printUsage(commands map[string]cli) {
	var lists []pterm.BulletListItem
	for _, c := range commands {
		text := fmt.Sprintf("%v  [%v]", c.argName, c.argUsage)
		lists = append(lists, pterm.BulletListItem{
			Level: 2,
			Text:  text,
		})

		for _, v := range c.boolOptions {
			lists = append(lists, pterm.BulletListItem{
				Level: 4,
				Text:  fmt.Sprintf("%v  [%v]", v.optionName, v.optionUsage),
			})
		}
		for _, v := range c.float64Options {
			lists = append(lists, pterm.BulletListItem{
				Level: 4,
				Text:  fmt.Sprintf("%v  [%v]", v.optionName, v.optionUsage),
			})
		}
		for _, v := range c.intOptions {
			lists = append(lists, pterm.BulletListItem{
				Level: 4,
				Text:  fmt.Sprintf("%v  [%v]", v.optionName, v.optionUsage),
			})
		}
		for _, v := range c.stringOptions {
			lists = append(lists, pterm.BulletListItem{
				Level: 4,
				Text:  fmt.Sprintf("%v  [%v]", v.optionName, v.optionUsage),
			})
		}
		for _, v := range c.uintOptions {
			lists = append(lists, pterm.BulletListItem{
				Level: 4,
				Text:  fmt.Sprintf("%v  [%v]", v.optionName, v.optionUsage),
			})
		}
	}

	pterm.DefaultBulletList.WithItems(lists).Render()
}
