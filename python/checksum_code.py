import os
import hashlib
import subprocess
import glob
import multiprocessing as mp
from functools import partial
import time

def generate_checksum(file):
    # start_time = time.time()

    hash_algorithm = hashlib.sha256()

    try:
        with open(file, "rb") as f:
            while buffer := f.read(8192):
                hash_algorithm.update(buffer)
    except FileNotFoundError:
        print("File not found")

    # duration = (time.time() - start_time) * 1000 # in ms
    # print(duration)

    return hash_algorithm.hexdigest()

def file_transform(location, file):
    new_file = os.path.basename(file) + ".mp3"
    target = os.path.join(location, new_file)
    try:
        subprocess.call(["ffmpeg", "-hide_banner", "-i", file, "-c:a", "libmp3lame", "-b:a", "256k", "-ar", "44100", target])
        print("Transformed file: " + target)
    except Exception as e:
        print(e)

def cleanup():
    cleanup_location = "/workspaces/file-processor-lab/_test_files"
    path = os.path.join(cleanup_location, "*.mp3")
    files = glob.glob(path)
    for file in files:
        os.remove(file)
    print("Cleaned up files")


def process_file(location, file):
    try:
        checksum = generate_checksum(file)
        print(checksum)

        file_transform(location, file)
        return True
    except Exception as e:
        print(e)
        return False

def main():
    start_time = time.time()

    location = "/workspaces/file-processor-lab/_test_files"
    path = os.path.join(location, "*.wav")
    files = glob.glob(path)

    num_processes = mp.cpu_count()
    with mp.Pool(processes=num_processes) as pool:
        process = partial(process_file, location)
        results = pool.map(process, files)


    duration = (time.time() - start_time) * 1000 # in ms
    print(duration)

if __name__ == "__main__":
    main()
    cleanup()