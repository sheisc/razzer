// AUTOGENERATED FILE


#if defined(__x86_64__) || 0
#define GOARCH "amd64"
#define SYZ_REVISION "830b4d26672d98c7e59a44635a52c120cbb66866"
#define SYZ_PAGE_SIZE 4096
#define SYZ_NUM_PAGES 4096
#define SYZ_DATA_OFFSET 536870912
unsigned syscall_count = 35;
call_t syscalls[] = {
	{"abort_sysc_fd", 33},
	{"chdir", 116},
	{"close", 103},
	{"fchdir", 124},
	{"fcntl$F_DUPFD", 107},
	{"fcntl$F_DUPFD_CLOEXEC", 107},
	{"fcntl$F_GETFD", 107},
	{"fcntl$F_GETFL", 107},
	{"fcntl$F_GETLK", 107},
	{"fcntl$F_GETOWN", 107},
	{"fcntl$F_SETFD", 107},
	{"fcntl$F_SETFL", 107},
	{"fcntl$F_SETLK", 107},
	{"fcntl$F_SETLKW", 107},
	{"fcntl$F_SETOWN", 107},
	{"fstat", 104},
	{"getcwd", 117},
	{"link", 112},
	{"llseek", 111},
	{"lstat", 106},
	{"mkdir", 118},
	{"mmap", 18},
	{"mprotect", 20},
	{"munmap", 19},
	{"nanosleep", 36},
	{"openat", 102},
	{"read", 100},
	{"readlink", 115},
	{"rename", 123},
	{"rmdir", 119},
	{"stat", 105},
	{"symlink", 114},
	{"unlink", 113},
	{"waitpid", 17},
	{"write", 101},

};
#endif
