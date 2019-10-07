#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/time.h>

int main() {
    struct timeval t1;
    gettimeofday(&t1, NULL);
    char *b1 = malloc(sizeof(char) * 10);
    strcpy(b1, "abcdefghi\0");
    char b2[20];
    memcpy(b2, b1, 10);
    memcpy(b2+9, b1, 10);
    printf("b2: %s\n", b2);
    struct timeval t2;
    gettimeofday(&t2, NULL);
    printf("time spent: %lds %dus\n", t2.tv_sec - t1.tv_sec, t2.tv_usec - t1.tv_usec);
    return 0;
}
