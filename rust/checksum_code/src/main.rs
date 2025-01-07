use std::time::Instant;
use std::fs::File;
use std::io::{BufReader, Read};

use sha2::{Sha256, Digest};
use glob::glob;

fn get_file_list(pattern: &str) -> Vec<String> {
    let mut file_list = Vec::new();

    for entry in glob(pattern).expect("failed to read glob pattern") {
        match entry {
            Ok(entry) => {
                let path = entry.as_path();
                file_list.push(path.display().to_string());
            }
            Err(e) => {
                println!("{:?}", e)
            }
        }
    }
    println!("{}", file_list.len());
    file_list
}

fn generate_checksum(file_path: &str) -> String {
    let file = File::open(file_path).expect("Failed to open file");
    let mut reader = BufReader::new(file);
    let mut hasher = Sha256::new();
    let mut buffer = [0; 1024];

    loop {
        let n = reader.read(&mut buffer).expect("Failed to read file");
        if n == 0 {
            break;
        } else {
            hasher.update(&buffer[..n]);
        }
    }
    format!("{:x}", hasher.finalize())
}

fn main() {
    let file_path = "/workspaces/file-processor-lab/_test_files/*.*";
    let files = get_file_list(file_path);
    let start = Instant::now();

    for file in files {
        let checksum = generate_checksum(&file);
        println!("{}: {}", file, checksum);
    }
    println!("Elapsed time: {:?}", start.elapsed());
}