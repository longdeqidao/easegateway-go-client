package pdu

type ResourceUsage struct {
	Utime    Time  `json:"Utime"`
	Stime    Time  `json:"Stime"`
	Maxrss   int64 `json:"Maxrss"`
	Ixrss    int64 `json:"Ixrss"`
	Idrss    int64 `json:"Idrss"`
	Isrss    int64 `json:"Isrss"`
	Minflt   int64 `json:"Minflt"`
	Majflt   int64 `json:"Majflt"`
	Nswap    int64 `json:"Nswap"`
	Inblock  int64 `json:"Inblock"`
	Oublock  int64 `json:"Oublock"`
	Msgsnd   int64 `json:"Msgsnd"`
	Msgrcv   int64 `json:"Msgrcv"`
	Nsignals int64 `json:"Nsignals"`
	Nvcsw    int64 `json:"Nvcsw"`
	Nivcsw   int64 `json:"Nivcsw"`
}
