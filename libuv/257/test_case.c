// Author: Ben Noordhuis <ben@strongloop.com>
#include <assert.h>
#include <stdlib.h>
#include "uv.h"

static void on_alloc(uv_handle_t* handle, size_t size, uv_buf_t* buf) {
  (void) &handle;
  buf->len = size;
  buf->base = malloc(size);
}

static void on_read(uv_stream_t* handle, ssize_t nread, const uv_buf_t* buf) {
  (void) &nread;
  (void) &buf;
  uv_close((uv_handle_t*) handle, NULL);
}

int main(void) {
  uv_loop_t loop;
  uv_tty_t tty;

  assert(0 == uv_loop_init(&loop));
  assert(0 == uv_tty_init(&loop, &tty, 0, 1));
  assert(0 == uv_tty_set_mode(&tty, UV_TTY_MODE_RAW));
  assert(0 == uv_read_start((uv_stream_t*) &tty, on_alloc, on_read));
  assert(0 == uv_run(&loop, UV_RUN_DEFAULT));
  assert(0 == uv_loop_close(&loop));

  return 0;
}
