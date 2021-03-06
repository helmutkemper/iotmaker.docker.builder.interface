package iotmaker_docker_builder_interface

import (
	"github.com/docker/go-connections/nat"
	network "github.com/helmutkemper/iotmaker.docker.builder.network.interface"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"time"
)

type ContainerBuilderInterface interface {
	GetImageName() (name string)
	GetNetworkInterface() (network network.ContainerBuilderNetworkInterface)
	GetNetworkIPV4() (IPV4 string)
	GetNetworkGatewayIPV4() (IPV4 string)
	GetNetworkIPV4ByNetworkName(networkName string) (IPV4 string)
	GetNetworkGatewayIPV4ByNetworkName(networkName string) (IPV4 string)
	SetGitCloneToBuild(url string)
	SetGitCloneToBuildWithUserPassworh(url, user, password string)
	SetGitCloneToBuildWithPrivateToken(url, privateToken string)
	RemoveAllByNameContains(value string) (err error)
	GetIdByContainerName() (err error)
	ContainerInspect() (inspect iotmakerdocker.ContainerInspect, err error)
	ImageRemove() (err error)
	ContainerRemove(removeVolumes bool) (err error)
	ContainerStop() (err error)
	ContainerPause() (err error)
	ContainerStart() (err error)
	FindTextInsideContainerLog(value string) (contains bool, err error)
	GetContainerLog() (log []byte, err error)
	ContainerBuildFromImage() (err error)
	ImageBuildFromServer() (err error)
	ImageBuildFromFolder() (err error)
	WaitForTextInContainerLog(value string) (dockerLogs string, err error)
	ImagePull() (err error)
	GetChannelEvent() (channel *chan iotmakerdocker.ContainerPullStatusSendToChannel)
	GetChannelOnContainerInspect() (channel *chan bool)
	GetChannelOnContainerReady() (channel *chan bool)
	Init() (err error)
	SetInspectInterval(value time.Duration)
	SetDoNotOpenContainersPorts()
	AddPortToChange(imagePort string, newPort string)
	AddPortToOpen(value string)
	SetEnvironmentVar(value []string)
	SetNetworkDocker(network network.ContainerBuilderNetworkInterface)
	SetWaitString(value string)
	SetWaitStringWithTimeout(value string, timeout time.Duration)
	SetContainerName(value string)
	SetImageName(value string)
	SetGitCloneToBuildWithPrivateSshKey(url, privateSshKeyPath, password string)
	SetBuildFolderPath(value string)
	GetLastLogs() (logs string)
	GetImageID() (ID string)
	GetLastInspect() (inspect iotmakerdocker.ContainerInspect)
	GetContainerID() (ID string)
	AddFiileOrFolderToLinkBetweenConputerHostAndContainer(computerHostPath, insideContainerPath string) (err error)
	ImageRemoveByName(name string) (err error)
	ImageListExposedPorts() (list []nat.Port, err error)
	ImageListExposedVolumes() (list []string, err error)
	WaitForTextInContainerLogWithTimeout(value string, timeout time.Duration) (dockerLogs string, err error)
}
