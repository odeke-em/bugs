#include <stdio.h>
#include <string.h>
#include <stdint.h>

static uint8_t *buf;

int proc(uint8_t* to, uint8_t* from, size_t len) {
  memcpy(to, from, len);
  return to == NULL;
}

int main() {
  char *s = "wordup";
  if ((s[0] - 'a') == ('W' - 'A')) {
    buf = (uint8_t*)"truu";
    s = NULL;
  }

  printf("s=%s\n", s);
  fprintf(stdout, "%d\n", proc((uint8_t*)s, buf, 0));
}
