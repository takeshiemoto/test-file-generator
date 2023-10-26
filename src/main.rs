use rand::Rng;
use std::fs;
use std::fs::File;
use std::io::Write;
use std::path::Path;

const MB_BLOCK: usize = 1_048_576; // 1MB
const GB4: usize = 4_294_967_296; // 4GB

fn random_string(length: usize) -> String {
    const CHARSET: &[u8] = b"ABCDEFGHIJKLMNOPQRSTUVWXYZ\
                             abcdefghijklmnopqrstuvwxyz\
                             0123456789";
    let mut rng = rand::thread_rng();

    let s: String = (0..length)
        .map(|_| {
            let idx = rng.gen_range(0..CHARSET.len());
            CHARSET[idx] as char
        })
        .collect();

    s
}

fn main() {
    let args: Vec<String> = std::env::args().collect();

    if args.len() < 2 || args.len() > 3 {
        eprintln!("使い方: {} [生成するファイルの個数] [--large]", args[0]);
        std::process::exit(1);
    }

    let count: usize = match args[1].parse() {
        Ok(n) => n,
        Err(_) => {
            eprintln!("エラー: 数字を指定してください");
            std::process::exit(1);
        }
    };

    let large_files = args.contains(&"--large".to_string());
    let large4gb_files = args.contains(&"--large4gb".to_string());

    // output フォルダを作成（既に存在する場合は無視）
    let output_dir = Path::new("output");
    if output_dir.exists() {
        for entry in fs::read_dir(output_dir).expect("ディレクトリの読み込みに失敗しました")
        {
            let entry = entry.expect("エントリの取得に失敗しました");
            fs::remove_file(entry.path()).expect("ファイルの削除に失敗しました");
        }
    } else {
        fs::create_dir(output_dir).expect("outputディレクトリの作成に失敗しました");
    }

    for i in 0..count {
        let filename = format!("output/file_{}.txt", i);
        let path = Path::new(&filename);
        let mut file = File::create(&path).expect("ファイルの作成に失敗しました");

        if large4gb_files {
            let block = random_string(MB_BLOCK);
            let iterations = GB4 / MB_BLOCK;

            for _ in 0..iterations {
                file.write_all(block.as_bytes())
                    .expect("書き込みに失敗しました");
            }
        } else if large_files {
            let content = random_string(1_048_576 + rand::thread_rng().gen_range(1..=20));
            file.write_all(content.as_bytes())
                .expect("書き込みに失敗しました");
        } else {
            let content = random_string(rand::thread_rng().gen_range(1..=20));
            file.write_all(content.as_bytes())
                .expect("書き込みに失敗しました");
        }

        println!("{} を生成しました", filename);
    }
}
