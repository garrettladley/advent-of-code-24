use d06::p1::process;

fn main() -> miette::Result<()> {
    let file = include_str!("../../input.txt");
    let result = process(file).expect("failed to process input");
    println!("{}", result);
    Ok(())
}
