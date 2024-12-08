use glam::IVec2;
use std::collections::HashMap;

const A: char = 'A';

const MS: [char; 2] = ['M', 'S'];

const DIRECTIONS: [[IVec2; 2]; 4] = [
    [IVec2::new(-1, -1), IVec2::new(1, 1)], // diagonal up left, diagonal down right
    [IVec2::new(-1, 1), IVec2::new(1, -1)], // diagonal down left, diagonal up right
    [IVec2::new(1, 1), IVec2::new(-1, -1)], // diagonal down right, diagonal up left
    [IVec2::new(1, -1), IVec2::new(-1, 1)], // diagonal up right, diagonal down left
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

    let result: usize = positions
        .iter()
        .filter(|(_position, value)| **value == A)
        .filter(|(position, _value)| {
            DIRECTIONS
                .iter()
                .map(|ms_positions| {
                    ms_positions
                        .iter()
                        .map(|pos| positions.get(&(*position + pos)))
                        .enumerate()
                        .all(|(index, value)| MS.get(index) == value)
                })
                .filter(|b| *b)
                .count()
                == 2
        })
        .count();

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
        assert_eq!("9", process(input)?);
        Ok(())
    }
}
