// AUTOGENERATED FILE
package akaros

import . "github.com/google/syzkaller/prog"

func init() {
	RegisterTarget(&Target{OS: "akaros", Arch: "amd64", Revision: revision_amd64, PtrSize: 8, PageSize: 4096, NumPages: 4096, DataOffset: 536870912, Syscalls: syscalls_amd64, Resources: resources_amd64, Structs: structDescs_amd64, Consts: consts_amd64}, initTarget)
}

var resources_amd64 = []*ResourceDesc{
	{Name: "fd", Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", TypeSize: 4}}}, Kind: []string{"fd"}, Values: []uint64{18446744073709551615, 18446744073709551516}},
	{Name: "pid", Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", TypeSize: 4}}}, Kind: []string{"pid"}, Values: []uint64{0, 18446744073709551615}},
}

var structDescs_amd64 = []*KeyedStruct{
	{Key: StructKey{Name: "flock"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "flock", TypeSize: 32}, Fields: []Type{
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "flock_type", FldName: "type", TypeSize: 2}}, Vals: []uint64{0, 1, 2}},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "seek_whence", FldName: "whence", TypeSize: 2}}, Vals: []uint64{0, 1, 2}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 4}}, IsPad: true},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "start", TypeSize: 8}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "len", TypeSize: 8}}},
		&ResourceType{TypeCommon: TypeCommon{TypeName: "pid", FldName: "pid", TypeSize: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 4}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "timespec"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "timespec", TypeSize: 16}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "sec", TypeSize: 8}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "nsec", TypeSize: 8}}},
	}}},
	{Key: StructKey{Name: "timespec", Dir: 1}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "timespec", TypeSize: 16, ArgDir: 1}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "sec", TypeSize: 8, ArgDir: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "nsec", TypeSize: 8, ArgDir: 1}}},
	}}},
}

var syscalls_amd64 = []*Syscall{
	{NR: 33, Name: "abort_sysc_fd", CallName: "abort_sysc_fd", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
	}},
	{NR: 116, Name: "chdir", CallName: "chdir", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "dir", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
	}},
	{NR: 103, Name: "close", CallName: "close", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
	}},
	{NR: 124, Name: "fchdir", CallName: "fchdir", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
	}},
	{NR: 107, Name: "fcntl$F_DUPFD", CallName: "fcntl", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "cmd", TypeSize: 8}}},
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "arg", TypeSize: 4}},
	}, Ret: &ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "ret", TypeSize: 4, ArgDir: 1}}},
	{NR: 107, Name: "fcntl$F_DUPFD_CLOEXEC", CallName: "fcntl", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "cmd", TypeSize: 8}}, Val: 1030},
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "arg", TypeSize: 4}},
	}, Ret: &ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "ret", TypeSize: 4, ArgDir: 1}}},
	{NR: 107, Name: "fcntl$F_GETFD", CallName: "fcntl", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "cmd", TypeSize: 8}}, Val: 1},
	}},
	{NR: 107, Name: "fcntl$F_GETFL", CallName: "fcntl", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "cmd", TypeSize: 8}}, Val: 3},
	}},
	{NR: 107, Name: "fcntl$F_GETLK", CallName: "fcntl", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "cmd", TypeSize: 8}}, Val: 5},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lock", TypeSize: 8}, Type: &StructType{Key: StructKey{Name: "flock"}}},
	}},
	{NR: 107, Name: "fcntl$F_GETOWN", CallName: "fcntl", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "cmd", TypeSize: 8}}, Val: 9},
	}, Ret: &ResourceType{TypeCommon: TypeCommon{TypeName: "pid", FldName: "ret", TypeSize: 4, ArgDir: 1}}},
	{NR: 107, Name: "fcntl$F_SETFD", CallName: "fcntl", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "cmd", TypeSize: 8}}, Val: 2},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "fcntl_flags", FldName: "flags", TypeSize: 8}}, Vals: []uint64{1}},
	}},
	{NR: 107, Name: "fcntl$F_SETFL", CallName: "fcntl", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "cmd", TypeSize: 8}}, Val: 4},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "fcntl_status", FldName: "flags", TypeSize: 8}}, Vals: []uint64{1024, 8192, 2048}},
	}},
	{NR: 107, Name: "fcntl$F_SETLK", CallName: "fcntl", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "cmd", TypeSize: 8}}, Val: 6},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lock", TypeSize: 8}, Type: &StructType{Key: StructKey{Name: "flock"}}},
	}},
	{NR: 107, Name: "fcntl$F_SETLKW", CallName: "fcntl", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "cmd", TypeSize: 8}}, Val: 7},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "lock", TypeSize: 8}, Type: &StructType{Key: StructKey{Name: "flock"}}},
	}},
	{NR: 107, Name: "fcntl$F_SETOWN", CallName: "fcntl", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "cmd", TypeSize: 8}}, Val: 8},
		&ResourceType{TypeCommon: TypeCommon{TypeName: "pid", FldName: "pid", TypeSize: 4}},
	}},
	{NR: 104, Name: "fstat", CallName: "fstat", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "statbuf", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "array", ArgDir: 1, IsVarlen: true}}},
	}},
	{NR: 117, Name: "getcwd", CallName: "getcwd", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "buffer", FldName: "buf", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{ArgDir: 1, IsVarlen: true}}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "size", TypeSize: 8}}, Buf: "buf"},
	}},
	{NR: 112, Name: "link", CallName: "link", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "old", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "new", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
	}},
	{NR: 111, Name: "llseek", CallName: "llseek", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "offset_hi", TypeSize: 8}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "offset_lo", TypeSize: 8}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "result", TypeSize: 8}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", TypeSize: 8, ArgDir: 1}}}},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "seek_whence", FldName: "whence", TypeSize: 8}}, Vals: []uint64{0, 1, 2}},
	}},
	{NR: 106, Name: "lstat", CallName: "lstat", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "file", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "statbuf", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "array", ArgDir: 1, IsVarlen: true}}},
	}},
	{NR: 118, Name: "mkdir", CallName: "mkdir", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "path", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "open_mode", FldName: "mode", TypeSize: 8}}, Vals: []uint64{256, 128, 64, 32, 16, 8, 4, 2, 1}},
	}},
	{NR: 18, Name: "mmap", CallName: "mmap", Args: []Type{
		&VmaType{TypeCommon: TypeCommon{TypeName: "vma", FldName: "addr", TypeSize: 8}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "len", TypeSize: 8}}, Buf: "addr"},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "mmap_prot", FldName: "prot", TypeSize: 8}}, Vals: []uint64{4, 1, 2, 16777216, 33554432}},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "mmap_flags", FldName: "flags", TypeSize: 8}}, Vals: []uint64{1, 2, 64, 32, 2048, 4096, 0, 16, 256, 262144, 8192, 65536, 16384, 32768, 131072}},
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4, IsOptional: true}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "offset", TypeSize: 8}}},
	}},
	{NR: 20, Name: "mprotect", CallName: "mprotect", Args: []Type{
		&VmaType{TypeCommon: TypeCommon{TypeName: "vma", FldName: "addr", TypeSize: 8}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "len", TypeSize: 8}}, Buf: "addr"},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "mmap_prot", FldName: "prot", TypeSize: 8}}, Vals: []uint64{4, 1, 2, 16777216, 33554432}},
	}},
	{NR: 19, Name: "munmap", CallName: "munmap", Args: []Type{
		&VmaType{TypeCommon: TypeCommon{TypeName: "vma", FldName: "addr", TypeSize: 8}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "len", TypeSize: 8}}, Buf: "addr"},
	}},
	{NR: 36, Name: "nanosleep", CallName: "nanosleep", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "req", TypeSize: 8}, Type: &StructType{Key: StructKey{Name: "timespec"}}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "rem", TypeSize: 8, IsOptional: true}, Type: &StructType{Key: StructKey{Name: "timespec", Dir: 1}}},
	}},
	{NR: 102, Name: "openat", CallName: "openat", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "file", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "open_flags", FldName: "flags", TypeSize: 8}}, Vals: []uint64{0, 1, 2, 1024, 8192, 524288, 64, 65536, 128, 256, 131072, 2048, 1052672, 512, 4259840}},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "open_mode", FldName: "mode", TypeSize: 8}}, Vals: []uint64{256, 128, 64, 32, 16, 8, 4, 2, 1}},
	}, Ret: &ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "ret", TypeSize: 4, ArgDir: 1}}},
	{NR: 100, Name: "read", CallName: "read", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "buffer", FldName: "buf", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{ArgDir: 1, IsVarlen: true}}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "count", TypeSize: 8}}, Buf: "buf"},
	}},
	{NR: 115, Name: "readlink", CallName: "readlink", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "path", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "buffer", FldName: "buf", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{ArgDir: 1, IsVarlen: true}}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "siz", TypeSize: 8}}, Buf: "buf"},
	}},
	{NR: 123, Name: "rename", CallName: "rename", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "old", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "new", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
	}},
	{NR: 119, Name: "rmdir", CallName: "rmdir", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "path", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
	}},
	{NR: 105, Name: "stat", CallName: "stat", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "file", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "statbuf", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "array", ArgDir: 1, IsVarlen: true}}},
	}},
	{NR: 114, Name: "symlink", CallName: "symlink", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "old", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "new", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
	}},
	{NR: 113, Name: "unlink", CallName: "unlink", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "path", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "filename", IsVarlen: true}, Kind: 3}},
	}},
	{NR: 17, Name: "waitpid", CallName: "waitpid", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "pid", FldName: "pid", TypeSize: 4}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "status", TypeSize: 8}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", TypeSize: 4, ArgDir: 1}}}},
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "wait_options", FldName: "options", TypeSize: 8}}, Vals: []uint64{1, 2}},
	}},
	{NR: 101, Name: "write", CallName: "write", Args: []Type{
		&ResourceType{TypeCommon: TypeCommon{TypeName: "fd", FldName: "fd", TypeSize: 4}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "buffer", FldName: "buf", TypeSize: 8}, Type: &BufferType{TypeCommon: TypeCommon{IsVarlen: true}}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "count", TypeSize: 8}}, Buf: "buf"},
	}},
}

var consts_amd64 = []ConstValue{
	{Name: "AT_FDCWD", Value: 18446744073709551516},
	{Name: "FASYNC", Value: 8192},
	{Name: "FD_CLOEXEC", Value: 1},
	{Name: "F_DUPFD"},
	{Name: "F_DUPFD_CLOEXEC", Value: 1030},
	{Name: "F_GETFD", Value: 1},
	{Name: "F_GETFL", Value: 3},
	{Name: "F_GETLK", Value: 5},
	{Name: "F_GETOWN", Value: 9},
	{Name: "F_RDLCK"},
	{Name: "F_SETFD", Value: 2},
	{Name: "F_SETFL", Value: 4},
	{Name: "F_SETLK", Value: 6},
	{Name: "F_SETLKW", Value: 7},
	{Name: "F_SETOWN", Value: 8},
	{Name: "F_UNLCK", Value: 2},
	{Name: "F_WRLCK", Value: 1},
	{Name: "MAP_32BIT", Value: 64},
	{Name: "MAP_ANONYMOUS", Value: 32},
	{Name: "MAP_DENYWRITE", Value: 2048},
	{Name: "MAP_EXECUTABLE", Value: 4096},
	{Name: "MAP_FILE"},
	{Name: "MAP_FIXED", Value: 16},
	{Name: "MAP_GROWSDOWN", Value: 256},
	{Name: "MAP_HUGETLB", Value: 262144},
	{Name: "MAP_LOCKED", Value: 8192},
	{Name: "MAP_NONBLOCK", Value: 65536},
	{Name: "MAP_NORESERVE", Value: 16384},
	{Name: "MAP_POPULATE", Value: 32768},
	{Name: "MAP_PRIVATE", Value: 2},
	{Name: "MAP_SHARED", Value: 1},
	{Name: "MAP_STACK", Value: 131072},
	{Name: "O_APPEND", Value: 1024},
	{Name: "O_CLOEXEC", Value: 524288},
	{Name: "O_CREAT", Value: 64},
	{Name: "O_DIRECTORY", Value: 65536},
	{Name: "O_EXCL", Value: 128},
	{Name: "O_NOCTTY", Value: 256},
	{Name: "O_NOFOLLOW", Value: 131072},
	{Name: "O_NONBLOCK", Value: 2048},
	{Name: "O_RDONLY"},
	{Name: "O_RDWR", Value: 2},
	{Name: "O_SYNC", Value: 1052672},
	{Name: "O_TRUNC", Value: 512},
	{Name: "O_WRONLY", Value: 1},
	{Name: "PROT_EXEC", Value: 4},
	{Name: "PROT_GROWSDOWN", Value: 16777216},
	{Name: "PROT_GROWSUP", Value: 33554432},
	{Name: "PROT_READ", Value: 1},
	{Name: "PROT_WRITE", Value: 2},
	{Name: "SEEK_CUR", Value: 1},
	{Name: "SEEK_END", Value: 2},
	{Name: "SEEK_SET"},
	{Name: "SYS_abort_sysc_fd", Value: 33},
	{Name: "SYS_chdir", Value: 116},
	{Name: "SYS_close", Value: 103},
	{Name: "SYS_fchdir", Value: 124},
	{Name: "SYS_fcntl", Value: 107},
	{Name: "SYS_fork", Value: 15},
	{Name: "SYS_fstat", Value: 104},
	{Name: "SYS_getcwd", Value: 117},
	{Name: "SYS_link", Value: 112},
	{Name: "SYS_llseek", Value: 111},
	{Name: "SYS_lstat", Value: 106},
	{Name: "SYS_mkdir", Value: 118},
	{Name: "SYS_mmap", Value: 18},
	{Name: "SYS_mprotect", Value: 20},
	{Name: "SYS_munmap", Value: 19},
	{Name: "SYS_nanosleep", Value: 36},
	{Name: "SYS_openat", Value: 102},
	{Name: "SYS_read", Value: 100},
	{Name: "SYS_readlink", Value: 115},
	{Name: "SYS_rename", Value: 123},
	{Name: "SYS_rmdir", Value: 119},
	{Name: "SYS_stat", Value: 105},
	{Name: "SYS_symlink", Value: 114},
	{Name: "SYS_unlink", Value: 113},
	{Name: "SYS_waitpid", Value: 17},
	{Name: "SYS_write", Value: 101},
	{Name: "S_IRGRP", Value: 32},
	{Name: "S_IROTH", Value: 4},
	{Name: "S_IRUSR", Value: 256},
	{Name: "S_IWGRP", Value: 16},
	{Name: "S_IWOTH", Value: 2},
	{Name: "S_IWUSR", Value: 128},
	{Name: "S_IXGRP", Value: 8},
	{Name: "S_IXOTH", Value: 1},
	{Name: "S_IXUSR", Value: 64},
	{Name: "WNOHANG", Value: 1},
	{Name: "WUNTRACED", Value: 2},
	{Name: "__O_TMPFILE", Value: 4259840},
}

const revision_amd64 = "830b4d26672d98c7e59a44635a52c120cbb66866"
