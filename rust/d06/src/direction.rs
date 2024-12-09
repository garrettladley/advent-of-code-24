use glam::IVec2;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Direction {
    Up,
    Down,
    Left,
    Right,
}

impl Direction {
    pub fn to_ivec2(&self) -> IVec2 {
        match self {
            Direction::Up => IVec2::NEG_Y,
            Direction::Down => IVec2::Y,
            Direction::Left => IVec2::NEG_X,
            Direction::Right => IVec2::X,
        }
    }

    pub fn turn_right(&self) -> Direction {
        match self {
            Self::Up => Self::Right,
            Self::Right => Self::Down,
            Self::Down => Self::Left,
            Self::Left => Self::Up,
        }
    }
}

impl TryFrom<char> for Direction {
    type Error = String;
    fn try_from(value: char) -> Result<Self, Self::Error> {
        match value {
            '^' => Ok(Self::Up),
            'v' => Ok(Self::Down),
            '<' => Ok(Self::Left),
            '>' => Ok(Self::Right),
            _ => Err(format!("Invalid direction: {}", value)),
        }
    }
}

impl ToString for Direction {
    fn to_string(&self) -> String {
        match self {
            Self::Up => "^".to_string(),
            Self::Down => "v".to_string(),
            Self::Left => "<".to_string(),
            Self::Right => ">".to_string(),
        }
    }
}
