package main

import (
	"docktor/lib"
	"docktor/tui"
	ui "github.com/gizak/termui/v3"
	"image"
	"log"
	"os"
	"time"
)

var dm *lib.DockerManager
var dashboard *tui.Dashboard

func initDocktor() {
	width, height := ui.TerminalDimensions()
	log.Printf("Termsize: %d %d\n", width, height)

	dm = lib.NewDockerManager()
	dashboard = tui.NewDashboard(image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: width, Y: height},
	})
	render()

	go func() {
		for _ = range time.NewTicker(dm.UpdateInterval).C {

			switch dm.Type {
			case lib.T_CONTAINERS:
				dashboard.ParseContainers(dm.GetContainerList())
				break
			case lib.T_IMAGES:
				dashboard.ParseImages(dm.GetImageList())
				break
			}
		}

	}()
}

func render() {
	ui.Render(dashboard)
}

func loop() int {
	defer ui.Clear()

	ticker := time.NewTicker(dm.UpdateInterval).C
	uiEvents := ui.PollEvents()

	for {
		select {
		// TODO handle signals

		case <-ticker:
			render()

		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return 0
			}
		}
	}
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()

	file, err := os.OpenFile("/tmp/docktor.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Could not open log file: %v", err)
	}
	defer file.Close()

	if err == nil {
		log.SetOutput(file)
	}
	log.Println("Starting...")
	log.Println("Log initialized")

	initDocktor()

	os.Exit(loop())
}
