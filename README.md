# file-processor-lab
Repo to test file processing code 

## Sha256 calculation

### Rust
* sha2: 673ms (12mb dpx)

### Python
* hashlib: 1.5ms (12mb dpx)
* x20 m4a files, 326ms

### Go
* crypto/sha256: 11ms (12mb dpx)
* x20 m4a files, single core, 295ms
* x20 m4a files, go routine, 47ms

## File transform

### Python
* x20 m4a files, checksum, mp3 transcode, single core, 96sec
* x20 m4a files, checksum, mp3 transcode, multiprocessing, 21sec
* x6 wav files, checksum, multiprocessing, 1sec 79ms
* x6 wav files, checksum, mp3 transcode, multiprocessing, 1min 1sec

### Go
* x20 m4a files, checksum, mp3 transcode, go routine, 20sec
* x6 wav files, checksum, multiprocessing, 965ms
* x6 wav files, checksum, mp3 transcode, go routine, 1min 4sec