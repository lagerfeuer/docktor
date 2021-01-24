package tui

import (
	"docktor/util"
	"fmt"
	"github.com/docker/docker/api/types"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"image"
	"log"
	"strings"
)

type Dashboard struct {
	*ui.Grid
	List          *widgets.List
	Data          interface{}
	ImageLenRatio float64
	NameLenRatio  float64
}

func (dboard *Dashboard) update(rows []string) {
	if rows != nil {
		dboard.List.Rows = rows
	}
}

func (dboard *Dashboard) Dx() int {
	return dboard.Grid.Block.Dx()
}

func (dboard *Dashboard) Dy() int {
	return dboard.Grid.Block.Dy()
}

func (dboard *Dashboard) SetRect(x1, y1, x2, y2 int) {
	log.Printf("Dasboard Size: %d %d %d %d", x1, y1, x2, y2)
	dboard.Grid.SetRect(x1, y1, x2, y2)
}

func (dboard *Dashboard) init(size image.Rectangle) {
	dboard.List.Title = "Containers"
	dboard.List.TitleStyle = util.STYLE_TITLE
	dboard.Grid.SetRect(size.Min.X, size.Min.Y, size.Max.X, size.Max.Y)
	dboard.Grid.Set(
		ui.NewRow(1,
			ui.NewCol(1, dboard.List)))
	dboard.update(nil)
}

func (dboard *Dashboard) ParseContainers(containers []types.Container) {
	var data []string
	idLen := 15
	nameLen := int(float64(dboard.Dx()) * dboard.NameLenRatio)
	imageLen := int(float64(dboard.Dx()) * dboard.ImageLenRatio)

	rowFormat := " %-*s  %-*s %-*s"

	header := fmt.Sprintf(rowFormat,
		idLen,
		"CONTAINER ID",
		imageLen,
		"IMAGE",
		nameLen,
		"NAMES")[1:]
	dboard.List.Title = header

	for _, container := range containers {
		names := strings.Join(container.Names[:], ",")
		tmp := fmt.Sprintf(rowFormat,
			idLen,
			container.ID[:12],
			imageLen,
			container.Image[:util.Min(imageLen, len(container.Image))],
			nameLen,
			names[:util.Min(nameLen, len(names))],
		)
		data = append(data, tmp)
	}
	dboard.Data = containers
	dboard.update(data)
}

func (dboard *Dashboard) ParseImages(images []types.ImageSummary) {

}

func NewDashboard(size image.Rectangle) *Dashboard {
	self := Dashboard{
		ui.NewGrid(),
		widgets.NewList(),
		nil,
		0.5,
		0.3,
	}
	self.init(size)
	return &self
}
