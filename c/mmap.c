#include <sys/stat.h>
#include <sys/mman.h>
#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <unistd.h>

int fd;
char *data;

#define handle_error(msg) \
    do { perror(msg); exit(EXIT_FAILURE); } while (0)

int main(int argc, char *argv[]) {
	char *addr;
	struct stat sb;
	fd = open("data.bin", O_RDONLY);
	if (fd == -1)
	  handle_error("open");
  if (fstat(fd, &sb) == -1)           /* To obtain file size */
	  handle_error("fstat");
	
	int pagesize = 4000;
	addr = mmap(NULL, pagesize, PROT_READ, MAP_SHARED, fd, pagesize);
	
	if (addr == MAP_FAILED)
		handle_error("mmap");
	
	printf("data:\n%s\n", data);
	
	return 0;
}