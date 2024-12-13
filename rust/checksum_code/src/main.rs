use std::time::Instant;
use std::fs::File;
use std::io::{Read, BufReader};
use sha2::{Sha256, Digest};

fn generate_file_checksum(filepath: &str) -> Result<String, std::io::Error> {
    let start = Instant::now();

    let file = File::open(filepath)?;
    let mut reader = BufReader::new(file);
    let mut hasher = Sha256::new();
    let mut buffer = [0; 1024];
    
    loop {
        let bytes_read = reader.read(&mut buffer)?;
        if bytes_read == 0 {
            break;
        }
        hasher.update(&buffer[0..bytes_read]);
    }

    let hash = hasher.finalize();
    let duration = start.elapsed();
    println!("{:?}", duration);
    Ok(format!("{:x}", hash))
}

fn main() {
    match generate_file_checksum("test_file.dpx") {
        Ok(checksum) => println!("{}", checksum),
        Err(e) => eprintln!("{}", e),
    }
}

