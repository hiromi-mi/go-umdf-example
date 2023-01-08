package main

type WDF_OBJECT_ATTRIBUTES struct {
	Size                 uint32
	EvtCleanupCallback   uintptr
	EvtDestroyCallback   uintptr
	ExecutionLevel       int32  // enum
	SynchronizationScope int32  //enum
	ParentObject         HANDLE // handle
	ContextSizeOverride  size_t
	ContextTypeInfo      uintptr
}

type WDF_DRIVER_CONFIG struct {
	Size               uint32
	EvtDriverDeviceAdd uintptr
	EvtDriverUnload    uintptr
	DriverInitFlags    uint32
	DriverPoolTag      uint32
}

const NULL uintptr = 0
const WdfDeviceCreateTableIndex = 25

/*go:generate go run golang.org/x/sys/windows/mkwinsyscall -output wdf_generate.go wdf.go*/
