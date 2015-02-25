#include <stdio.h>

int main() {
    char *fmt = "%6s%5d";
    //           "ssssssddddd"
    char *line = "some      3";
    char str[6];
    int numb;

    int n = sscanf(line, fmt, str, &numb);
    printf("n: %d str: %s numb: %d\n", n, str, numb);
    return 0;
}
