#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <sys/errno.h>
#include <stdio.h>

int main() {
    int fd = open("foo.txt", O_CREAT|O_TRUNC|O_WRONLY, 0777);
    if (fd < 0) {
        fprintf(stderr, "opening file: %s\n", strerror(errno));
        return -1;
    }
    close(fd);
    return 0;
}
