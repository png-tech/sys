//go:build e2k

package cpu

const cacheLineSize = 64

const IsBigEndian = false

// e2k: CPU feature detection is not used; assume all flags are disabled
func initOptions() {
}

// e2k: no CPU feature detection (no HWCAP, no CPUID). The generic archInit from
// cpu_linux.go calls readHWCAP/hostByteOrder, which don't know about Elbrus
// and panic with "unknown architecture" — so e2k arch-init is empty.
func archInit() {
}
