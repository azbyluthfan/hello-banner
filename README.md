## Building and running tests

The software is designed to be able to run both on the machine on top of docker. You can see `Makefile` to see what
kind of targets supported.

### Running Binary

Before running the binary, you must build the binary via `Makefile` and run the argument based on the operating system that you run:

- building and running binary for linux operating system
    ```
    make hello-banner-linux && ./hello-banner-linux
    ```

- building and running binary for Mac OSX operating system
    ```
    make hello-banner-osx && ./hello-banner-osx
    ``` 

### Integration Test


```
make test
```