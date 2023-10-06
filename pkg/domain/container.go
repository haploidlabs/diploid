package domain

type ContainerStatus string

const (
	ContainerStatusRunning ContainerStatus = "running"
	ContainerStatusStopped ContainerStatus = "stopped"
)

type ContainerPortType string

const (
	ContainerPortTypeTCP ContainerPortType = "tcp"
	ContainerPortTypeUDP ContainerPortType = "udp"
)

type ContainerPort struct {
	PrivatePort int64  `json:"private_port"`
	PublicPort  int64  `json:"public_port"`
	Type        string `json:"type"`
}

type ContainerVolume struct {
	Host      string `json:"host"`
	Container string `json:"container"`
}

type Container struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Image   string            `json:"image"`
	ImageID string            `json:"imageId"`
	Status  ContainerStatus   `json:"status"`
	Ports   []ContainerPort   `json:"ports"`
	Volumes []ContainerVolume `json:"volumes"`
}
