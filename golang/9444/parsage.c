#include <stdio.h>

int main() {
    char s1[10], s2[10];
    char *src = "22333";

    sscanf(src, "%2s%3s", s1, s2);
    printf("src1: '%s' '%s'\n", s1, s2);
    return 0;
}
