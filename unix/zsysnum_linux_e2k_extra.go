//go:build linux && e2k

package unix

const (
    SYS_PREAD64  = SYS_PREAD
    SYS_PWRITE64 = SYS_PWRITE
)
