package scheme

type Response struct {
	LscpuOutput 
	// processList []Process
	// diskSpace DiskSpace
}

type LscpuOutput struct{
	Architecture string `json:"Architecture"`
	CPUMode string `json:"CPU op-mode(s)"`
	ByteOrder string `json:"Byte Order"`
	CPUs string `json:"CPU(s)"`
	OnlineCPUsList string `json:"On-line CPU(s) list"`
	ThreadPerCore string `json:"Thread(s) per core"`
	Sockets string `json:"Socket(s)"`
	NUMANodes string `json:"NUMA node(s)"`
}

// type Process struct{

// }

// type DiskSpace struct {
// 	totalDiskSpace string
// 	usedSpace string
// 	freeSpace string
// }
