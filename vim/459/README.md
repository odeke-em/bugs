# 459

The shell itself needs to backquote any spaces

### Repro
- Firstly run `./repro.sh`
- Next compile the minimalist system caller.
```shell
$ gcc sys.c -o syso
```

* Case 1 (Unquoted space):
```shell
$ ./syso "$(PWD)/with spaces/sh"
sh: /Users/emmanuelodeke/Desktop/openSrc/bugs/vim/459/with: No such file or directory
errno=0 desc=Undefined error: 0 status=32512
```

* Case 2 (Properly quoted spaces):
```shell
$ ./syso "$(PWD)/with\\ spaces/sh"
sh-3.2$
```
