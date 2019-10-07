#include <stdio.h>
#include <string.h>
#include <sys/time.h>

int main() {
    struct timeval t1;
    gettimeofday(&t1, NULL);
    char b1[20] = "abcdefghijklmnopqrs\0";
    char b2[40];
    memcpy(b2, b1, 20);
    memcpy(b2+19, b1, 20);
    printf("b2: %s\n", b2);
    struct timeval t2;
    gettimeofday(&t2, NULL);
    printf("time spent: %lds %dus\n", t2.tv_sec - t1.tv_sec, t2.tv_usec - t1.tv_usec);
    return 0;
}
