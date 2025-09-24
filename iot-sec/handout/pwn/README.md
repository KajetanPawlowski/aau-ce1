# PWN introduction!

This folder contains 3 challenges which increase in difficulty. It is supposed to be a soft introduction to Stack buffer overflows. Hints are available, but try to reason about what is going on first!

Some problems can occur if the libc and linker used when developing and testing the challenge is not identical, as a lot of functionality is dependent on both libc and linker version. To mitigate this, run `patchelf` like so: 

`patchelf --set-interpreter </path/to/ld.so> --replace-needed libc.so.6 </path/to/libc.so> /path/to/target`

Replace the `</path/to/ld.so>` and `</path/to/libc.so>` to the linker and libc in lib folder respectively. The `path/to/target` is the challenge you want to patch. Happy hacking!

## One

Simple buffer overflow challenge. This challenge includes the C source code as well.

## Two

This challenge will introduce a simple way to jump to a different function in the binary. Use Ghidra to decompile the binary to check out what you have to work with! It might be a stretch to call this a Return Oriented Programming (ROP) challenge, but you are actually going to do some ROP to make this work!

## Three

There is no flag in this challenge, as you will have to spawn a shell. Again you can use Ghidra to understand the binary. Here you might want to check out `ROPgadget` or `ropper` as well.