package main

// how to build: go build -buildmode=c-shared -o goumdf.dll -ldflags '-v -extldflags -Wl,--allow-multiple-definition,--export-all-symbols'
// cmd .....
// Visual Studio でビルド
// goumdf.dll を C:\Windows\System32\Drivers\UMDF\ にコピー
// deploy

// #cgo CFLAGS: -Wall -Wextra
// #cgo LDFLAGS: -Lthirdparty
// extern void* WdfFunctions_02015;
// extern void* WdfDriverGlobals;
// //extern unsigned long _imp_DbgPrintEx(unsigned long ComponentId, unsigned long Level, void* Format,...) { return 0;}
// /*
//  __declspec(dllexport) void* _guard_check_icall_nop(unsigned long ptr) { return 0;}
//  __declspec(dllexport) void* __guard_check_icall_fptr(unsigned long ptr) { return 0;}
//  __declspec(dllexport) void* __guard_dispatch_icall_fptr(unsigned long ptr) {return 0;}
//  __declspec(dllexport) void* _guard_xfg_table_dispatch_icall_fptr(unsigned long ptr) { return 0;}
//  __declspec(dllexport) void* _guard_xfg_dispatch_icall_fptr(unsigned long ptr) { return 0;}
// __declspec(dllexport)  void* _castguard_check_failure_os_handled_fptr(unsigned long ptr) {return 0;}
//  __declspec(dllexport)  void* _guard_xfg_check_icall_fptr(unsigned long ptr) { return 0;} */
// extern __declspec(dllexport) int FxDriverEntryUm(void* a,void* b,void* c,void* d);
import "C"
import (
	"syscall"
	"unsafe"

	//"github.com/TKMAX777/RemoteRelativeInput/winapi"
	"golang.org/x/sys/windows"
)

type HANDLE uintptr
type size_t uint64

const WdfDriverCreateTableIndex uint32 = 57
const WDF_NO_HANDLE uintptr = 0

// /* void _guard_check_icall_nop() {
// __asm__  inline ("jmp %rax; jmp %rax;jmp %rax; jmp %rax;jmp %rax; jmp %rax;jmp %rax; jmp %rax;jmp %rax; jmp %rax");
// } */
// /* void __imp_DbgPrintEx() {} */
var wdf_driver_globals uintptr = 0

var wdf_functions uintptr = 0
var attributes WDF_OBJECT_ATTRIBUTES
var config WDF_DRIVER_CONFIG

func main() {}

//export DriverEntry
func DriverEntry(DriverObject uintptr, RegistryPath uintptr) int32 {
	wdf_functions = uintptr(C.WdfFunctions_02015)
	wdf_driver_globals = uintptr(C.WdfDriverGlobals)
	attributes = WDF_OBJECT_ATTRIBUTES{
		Size:                 uint32(56),
		ExecutionLevel:       1,
		SynchronizationScope: 1,
	}
	config = WDF_DRIVER_CONFIG{
		Size:               uint32(32),
		EvtDriverDeviceAdd: syscall.NewCallback(EvtDeviceAdd),
	}
	//wdfObjectAttributesInit(&attributes)
	ptr := (*uintptr)(unsafe.Add(unsafe.Pointer(wdf_functions), WdfDriverCreateTableIndex*8))
	r1, _, _ := syscall.SyscallN(
		uintptr(*ptr),
		wdf_driver_globals,
		DriverObject,
		RegistryPath,
		uintptr(unsafe.Pointer(&attributes)),
		uintptr(unsafe.Pointer(&config)),
		0)
	PreventFromUnload()
	return int32(r1)
}

//export PreventFromUnload
func PreventFromUnload() int {
	var h windows.Handle
	// set GET_MODULE_HANDLE_EX_FLAG_PIN to ignore FreeLibrary()
	windows.GetModuleHandleEx(windows.GET_MODULE_HANDLE_EX_FLAG_PIN, MustUTF16PtrFromString("truegoumdf.dll"), &h)
	return 0

}

func MustUTF16PtrFromString(str string) *uint16 {
	ptr, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		panic(err)
	}
	return ptr
}

//export EvtDeviceAdd
func EvtDeviceAdd(driver uintptr, deviceinit uintptr) int {
	var hDevice uintptr
	ptr := (*uintptr)(unsafe.Add(unsafe.Pointer(wdf_functions), WdfDeviceCreateTableIndex*8))

	r, _, _ := syscall.SyscallN(
		*ptr,
		wdf_driver_globals,
		uintptr(unsafe.Pointer(&deviceinit)),
		0,
		uintptr(unsafe.Pointer(&hDevice)),
	)
	return int(r)
}

func wdfObjectAttributesInit(attributes *WDF_OBJECT_ATTRIBUTES) {
	// memset(attributes,0, sizeof((WDF_OBJECT_ATTRIBUTES)))
	//C.memset(attributes, 0, )
	/*
			    Attributes->Size = sizeof(WDF_OBJECT_ATTRIBUTES);
		    Attributes->ExecutionLevel = WdfExecutionLevelInheritFromParent;
		    Attributes->SynchronizationScope = WdfSynchronizationScopeInheritFromParent;

	*/
}
