use d06::p2::process;

fn main() -> miette::Result<()> {
    let file = include_str!("../../input.txt");
    let result = process(file).expect("failed to process input in part 2");
    println!("{}", result);
    Ok(())
}
