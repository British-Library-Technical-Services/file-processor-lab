import threading
import time
from concurrent.futures import ThreadPoolExecutor
import hashlib
import glob
import os

def main():
    start = time.perf_counter()
    test_files = "_test_files"
    file_list = glob.glob(os.path.join(test_files, "*.*"))
    concurrent_session(file_list) # concurrent processing
    # sequential_processing(file_list) # single core processing
    duration = (time.perf_counter() - start)
    print(duration)

def checksum_creation(file):
    hash_algorithm = hashlib.sha256()

    try:
        with open(file, "rb") as f:
            while buffer := f.read(8196):
                hash_algorithm.update(buffer)
    except FileNotFoundError as e:
        print(e)
    
    print(hash_algorithm.hexdigest())

def concurrent_session(file_list):
    with ThreadPoolExecutor(max_workers=4) as executor:
        executor.map(checksum_creation, file_list)

def sequential_processing(file_list):
    for file in file_list:
        checksum_creation(file)

if __name__ == "__main__":
    main()