use std::{collections::HashSet, str::FromStr};

use glam::IVec2;

use crate::direction::Direction;
use crate::p1::patrol_path;
use crate::world::World;

pub fn process(input: &str) -> miette::Result<String, String> {
    let world = World::from_str(input)?;
    let loop_positions = loop_positions(&world);

    Ok(loop_positions.len().to_string())
}

// given the state of the world,
// return a collection of positions
// where if an obstacle is added,
// the guard will enter an infinite loop
fn loop_positions(w: &World) -> HashSet<IVec2> {
    let mut patrol_path = patrol_path(w);
    patrol_path.remove(&w.guard.position);

    patrol_path
        .iter()
        .filter(|new_obstacle| {
            let mut guard_position = w.guard.position;
            let mut guard_direction = w.guard.direction;

            let mut visited: HashSet<(IVec2, Direction)> =
                HashSet::from([(guard_position, guard_direction)]);

            loop {
                let next_position = guard_position + guard_direction.to_ivec2();
                if &&next_position == new_obstacle || w.is_obstacle_at(next_position) {
                    guard_direction = guard_direction.turn_right();
                    continue;
                }

                if visited.contains(&(next_position, guard_direction)) {
                    break true;
                } else if w.is_in_bounds(next_position) {
                    guard_position = next_position;
                    visited.insert((guard_position, guard_direction));
                    continue;
                } else {
                    break false;
                }
            }
        })
        .cloned()
        .collect::<HashSet<IVec2>>()
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
        assert_eq!("6", process(INPUT)?);
        Ok(())
    }
}
