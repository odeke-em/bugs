#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

int main(int argc, char **argv) {
    const char *command = "ls";
    if (argc >= 2)
      command = argv[1];

    const int status = system(command);
    if (status != 0)
      fprintf(stderr, "errno=%d desc=%s status=%d\n", errno, strerror(errno), status);
    return status;
}
