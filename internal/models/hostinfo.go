package models

type HostInfo struct {
	CurrentTime       string
	Hostname          string
	Uptime            string
	OS                string
	Platform          string
	PlatformVersion   string
	CPUP              int
	CPUV              int
	TotalMemory       string
	CacheMemory       string
	FreeMemory        string
	TotalDiskSpace    string
	FreeDiskSpace     string
	CPUTemperature    string
	CPUUsage          string
	LoadAverage       string
	TotalSwap         string
	FreeSwap          string
	NetworkInterfaces []string
	RunningProcesses  int
	KernelVersion     string
	BootTime          string
}