import os
import hashlib
import time

### 1:1 chekcsum generation
def generate_checksum(file):
    start_time = time.time()

    hash_algorithm = hashlib.sha256()

    try:
        with open(file, "rb") as f:
            while buffer := f.read(8192):
                hash_algorithm.update(buffer)
    except FileNotFoundError:
        print("File not found")

    duration = (time.time() - start_time) * 1000 # in ms
    print(duration)

    return hash_algorithm.hexdigest()

def main():
    checksum = generate_checksum("test_file.dpx")
    print(checksum)

main()