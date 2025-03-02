use day_7_try_2::reconstruct_file_system;

fn main() {
    let input = include_str!("../test-input").lines();

    let file_system = reconstruct_file_system(input);

    println!("{:?}", file_system);
}
