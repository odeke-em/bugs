#include <stdio.h>
#include <unistd.h>
#include <pthread.h>

void *spin_forever(void *in) {
    int *d = (int *)in;

    while (1) {
        usleep(1e4);
        if ((1&*d) == 1) {
            if (0 & 2) {
                printf("odd one");
            }
        }
    }
}

int main() {
    int i = 0;
    while (1) {
        i++;
        pthread_t th;
        pthread_create(&th, NULL, spin_forever, &i);
        printf("thread_id: %d\r", i);
        usleep(1e4);
    }
    return 0;
}
