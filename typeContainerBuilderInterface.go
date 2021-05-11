package iotmaker_docker_builder_interface

import (
	"github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"time"
)

type ContainerBuilderInterface interface {
	GetLastLogs() (logs string)
	SetBuildFolderPath(value string)
	SetImageName(value string)
	SetContainerName(value string)
	SetWaitString(value string)
	SetNetworkDocker(network Network)
	SetEnvironmentVar(value []string)
	AddPortToOpen(value string)
	AddPortToChange(imagePort string, newPort string)
	SetDoNotOpenContainersPorts()
	SetInspectInterval(value time.Duration)
	Init() (err error)
	GetChannelOnContainerReady() (channel *chan bool)
	GetChannelOnContainerInspect() (channel *chan bool)
	GetChannelEvent() (channel *chan iotmakerdocker.ContainerPullStatusSendToChannel)
	ImagePull() (err error)
	WaitFortextInContainerLog(value string) (dockerLogs string, err error)
	BuildFromFolder() (err error)
	BuildFromImage() (err error)
	GetContainerLog() (log []byte, err error)
	ContainerStart() (err error)
	ContainerStop() (err error)
	ContainerRemove() (err error)
	ImageRemove() (err error)
	ContainerInspect() (inspect iotmakerdocker.ContainerInspect, err error)
	GetFindIdByContainerName() (err error)
	RemoveAllByNameContains(name string) (err error)
}
