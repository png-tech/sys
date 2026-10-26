//go:build linux && e2k && gc

package unix

import (
	"syscall"
	"unsafe"
)

// The e2k port ships no per-arch assembly for this package (asm_linux_amd64.s
// and friends do not build for e2k), so the Syscall family declared by
// syscall_unix_gc.go for the other platforms is implemented here by
// delegating to the standard library syscall package, which provides working
// linux/e2k trampolines.

func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno) {
	return syscall.Syscall(trap, a1, a2, a3)
}

func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno) {
	return syscall.Syscall6(trap, a1, a2, a3, a4, a5, a6)
}

func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno) {
	return syscall.RawSyscall(trap, a1, a2, a3)
}

func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno) {
	return syscall.RawSyscall6(trap, a1, a2, a3, a4, a5, a6)
}

// SyscallNoError and RawSyscallNoError are declared by syscall_linux_gc.go
// for the platforms whose syscalls never fail; the e2k kernel follows the
// same ABI, so the no-error trampolines simply discard errno.
func SyscallNoError(trap, a1, a2, a3 uintptr) (r1, r2 uintptr) {
	r1, r2, _ = syscall.Syscall(trap, a1, a2, a3)
	return
}

func RawSyscallNoError(trap, a1, a2, a3 uintptr) (r1, r2 uintptr) {
	r1, r2, _ = syscall.RawSyscall(trap, a1, a2, a3)
	return
}

// gettimeofday backs the helpers of syscall_linux_e2k.go; on linux/amd64/gc
// it is declared in syscall_linux_amd64_gc.go and implemented in
// asm_linux_amd64.s, which does not build for e2k.
func gettimeofday(tv *Timeval) (err syscall.Errno) {
	_, _, e1 := RawSyscall(SYS_GETTIMEOFDAY, uintptr(unsafe.Pointer(tv)), 0, 0)
	return e1
}
