# file-processor-lab
Repo to test file processing code 

## Sha256 comparisons

### Rust
* sha2: 673ms (12mb dpx)

### Python
* hashlib: 1.5ms

### Go
* crypto/sha256: 11ms (12mb dpx)
* x20 m4a files, single core 295ms
* x20 m4a files, go routine 47ms