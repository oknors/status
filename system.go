package main

type System struct {
	Host string      `json:"host"`
	Disk interface{} `json:"disk"`
	Cpu  interface{} `json:"cpuinfo"`
	Load interface{} `json:"load"`
	Ram  interface{} `json:"ram"`
	Time string      `json:"time"`
}

func system(s string) System {
	return System{host(), disk(), cpuinfo(), load(), ram(), now()}
}
