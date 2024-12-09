use std::{collections::HashSet, str::FromStr};

use glam::IVec2;

use crate::world::World;
use crate::world::OBSTACLE;
use crate::world::PATROL;

pub fn process(input: &str) -> miette::Result<String, String> {
    let world = World::from_str(input)?;
    let patrol_path = patrol_path(&world);
    let len = patrol_path.len() - 1;

    Ok(len.to_string())
}

fn patrol_path(w: &World) -> HashSet<IVec2> {
    let mut guard_position = w.guard.position;
    let mut guard_direction = w.guard.direction;
    let mut visited: HashSet<IVec2> = HashSet::from([guard_position]);

    while w.is_in_bounds(guard_position) {
        let next_position = guard_position + guard_direction.to_ivec2();
        if w.is_obstacle_at(next_position) {
            guard_direction = guard_direction.turn_right();
        } else {
            guard_position = next_position;
            visited.insert(guard_position);
        }
    }

    visited
}

#[allow(dead_code)]
fn render(w: &World, patrol_path: HashSet<IVec2>) -> String {
    let mut output = String::new();
    for y in 0..w.height {
        for x in 0..w.width {
            let pos = IVec2::new(x as i32, y as i32);
            let c = if w.obstacle_positions.contains(&pos) {
                OBSTACLE
            } else if patrol_path.contains(&pos) {
                PATROL
            } else if w.guard.position == pos {
                w.guard.direction.to_string().chars().next().unwrap()
            } else {
                '.'
            };
            output.push(c);
        }
        if y < w.height - 1 {
            output.push('\n');
        }
    }

    output
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "\
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...";

    #[test]
    fn test_process() -> miette::Result<(), String> {
        assert_eq!("41", process(INPUT)?);
        Ok(())
    }

    #[test]
    fn test_render() -> miette::Result<(), String> {
        let world = World::from_str(INPUT)?;
        let patrol_path = patrol_path(&world);
        let expected = "\
....#.....
....XXXXX#
....X...X.
..#.X...X.
..XXXXX#X.
..X.X.X.X.
.#XXXXXXX.
.XXXXXXX#.
#XXXXXXX..
......#X..";

        assert_eq!(expected, render(&world, patrol_path));
        Ok(())
    }
}
