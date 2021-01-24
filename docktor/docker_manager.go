package docktor

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"time"
)

const (
	T_CONTAINERS          = "T_CONTAINERS"
	T_CONTAINERS_FILTERED = "T_CONTAINERS_FILTERED"
	T_IMAGES              = "T_IMAGES"
	T_IMAGES_FILTERED     = "T_IMAGES_FILTERED"
)

type DockerManager struct {
	Cli            *client.Client
	UpdateInterval time.Duration
	Context        context.Context
	Type           string
}

func (dm *DockerManager) GetContainerList() []types.Container {
	list, err := dm.Cli.ContainerList(dm.Context, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	return list
}

func (dm *DockerManager) GetFilteredContainerList(filters filters.Args) []types.Container {
	list, err := dm.Cli.ContainerList(dm.Context, types.ContainerListOptions{
		All:     false,
		Filters: filters,
	})
	if err != nil {
		panic(err)
	}
	return list
}

func (dm *DockerManager) GetImageList() []types.ImageSummary {
	list, err := dm.Cli.ImageList(dm.Context, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	return list
}
func (dm *DockerManager) GetFilteredImageList(filters filters.Args) []types.ImageSummary {
	list, err := dm.Cli.ImageList(dm.Context, types.ImageListOptions{
		All:     false,
		Filters: filters,
	})
	if err != nil {
		panic(err)
	}
	return list
}

func NewDockerManager() *DockerManager {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	self := DockerManager{
		cli,
		500 * time.Millisecond,
		context.Background(),
		T_CONTAINERS,
	}
	return &self
}
