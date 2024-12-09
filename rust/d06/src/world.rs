use glam::IVec2;
use std::{collections::HashSet, str::FromStr};

use crate::direction::Direction;

pub const OBSTACLE: char = '#';
pub const GUARDS: [char; 4] = ['^', 'v', '<', '>'];
pub const PATROL: char = 'X';

pub struct Guard {
    pub position: IVec2,
    pub direction: Direction,
}

pub struct World {
    pub obstacle_positions: HashSet<IVec2>,
    pub guard: Guard,
    pub width: usize,
    pub height: usize,
}

impl World {
    pub fn is_obstacle_at(&self, position: IVec2) -> bool {
        self.obstacle_positions.contains(&position)
    }

    #[inline(always)]
    pub fn is_in_bounds(&self, position: IVec2) -> bool {
        position.x >= 0
            && position.x < self.width as i32
            && position.y >= 0
            && position.y < self.height as i32
    }
}

impl FromStr for World {
    type Err = String;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let lines: Vec<&str> = s.lines().collect();
        let height = lines.len() as usize;
        let width = lines.first().map(|line| line.len()).unwrap_or(0) as usize;

        let obstacle_positions = s
            .lines()
            .enumerate()
            .flat_map(|(y, line)| {
                line.chars()
                    .enumerate()
                    .filter(|&(_, value)| value == OBSTACLE)
                    .map(move |(x, _)| IVec2::new(x as i32, y as i32))
            })
            .collect::<HashSet<IVec2>>();

        let mut guard_position = None;
        let mut guard_direction = None;
        for (y, line) in s.lines().enumerate() {
            for (x, ch) in line.chars().enumerate() {
                if GUARDS.contains(&ch) {
                    guard_position = Some(IVec2::new(x as i32, y as i32));
                    guard_direction = Some(Direction::try_from(ch)?);
                }
            }
        }

        let guard = match (guard_position, guard_direction) {
            (Some(position), Some(direction)) => Guard {
                position,
                direction,
            },
            _ => return Err("Guard not found".to_string()),
        };

        Ok(World {
            obstacle_positions,
            guard,
            width,
            height,
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::collections::HashSet;

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
    fn test_from_str() -> miette::Result<(), String> {
        let world = World::from_str(INPUT)?;
        assert_eq!(world.width, 10);
        assert_eq!(world.height, 10);

        let expected_obstacles: HashSet<IVec2> = HashSet::from([
            IVec2::new(4, 0),
            IVec2::new(9, 1),
            IVec2::new(2, 3),
            IVec2::new(7, 4),
            IVec2::new(1, 6),
            IVec2::new(8, 7),
            IVec2::new(0, 8),
            IVec2::new(6, 9),
        ]);

        assert_eq!(
            world.obstacle_positions, expected_obstacles,
            "Found unexpected obstacles"
        );

        assert_eq!(
            world.obstacle_positions.len(),
            expected_obstacles.len(),
            "Found unexpected number of obstacles"
        );

        assert_eq!(world.guard.position, IVec2::new(4, 6));
        assert!(matches!(world.guard.direction, Direction::Up));
        Ok(())
    }
}
