package main

/*
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>

void sigHandler(const int sig) {
  printf("got this signal: %d\n", sig);
  exit(0);
}

void printIt(void *ptr, int handleSig) {
  if (handleSig) {
    struct sigaction sa;
    sa.sa_handler = sigHandler;
    sigaction(SIGSEGV, &sa, NULL);
  }
  char *s = (char *)ptr;
  printf(":%c\n", *s);
}
*/
import "C"

func main() {
    C.printIt(nil, 1)
}
