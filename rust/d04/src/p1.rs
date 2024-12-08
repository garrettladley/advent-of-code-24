use glam::IVec2;
use std::collections::HashMap;

const X: char = 'X';

const MAS: [char; 3] = ['M', 'A', 'S'];

const DIRECTIONS: [[IVec2; 3]; 8] = [
    [IVec2::new(0, 1), IVec2::new(0, 2), IVec2::new(0, 3)], // vertical down
    [IVec2::new(0, -1), IVec2::new(0, -2), IVec2::new(0, -3)], // vertical up
    [IVec2::new(1, 1), IVec2::new(2, 2), IVec2::new(3, 3)], // diagonal down right
    [IVec2::new(1, -1), IVec2::new(2, -2), IVec2::new(3, -3)], // diagonal up right
    [IVec2::new(-1, 1), IVec2::new(-2, 2), IVec2::new(-3, 3)], // diagonal down left
    [IVec2::new(-1, -1), IVec2::new(-2, -2), IVec2::new(-3, -3)], // diagonal up left
    [IVec2::new(1, 0), IVec2::new(2, 0), IVec2::new(3, 0)], // horizontal right
    [IVec2::new(-1, 0), IVec2::new(-2, 0), IVec2::new(-3, 0)], // horizontal left
];

pub fn process(input: &str) -> miette::Result<String> {
    let positions = input
        .lines()
        .enumerate()
        .flat_map(|(y, line)| {
            line.chars()
                .enumerate()
                .map(move |(x, value)| (IVec2::new(x as i32, y as i32), value))
        })
        .collect::<HashMap<IVec2, char>>();

    let result = positions
        .iter()
        .filter(|(_position, value)| **value == X)
        .map(|(position, _value)| {
            let count = DIRECTIONS
                .iter()
                .map(|mas_positions| {
                    mas_positions
                        .iter()
                        .map(|offset| positions.get(&(position + offset)))
                        .enumerate()
                        .all(|(index, value)| MAS.get(index) == value)
                })
                .filter(|b| *b)
                .count();
            count
        })
        .sum::<usize>();

    Ok(result.to_string())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_process() -> miette::Result<()> {
        let input = "MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX";
        assert_eq!("18", process(input)?);
        Ok(())
    }
}
