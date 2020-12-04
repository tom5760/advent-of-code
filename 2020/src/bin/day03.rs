use std::fmt::{self, Display};
use std::io::{self, Read};
use std::process;

use thiserror::Error;

#[derive(Debug, PartialEq, Eq)]
enum Tile {
    Empty { x: usize, y: usize },
    Tree { x: usize, y: usize },
}

impl Display for Tile {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            Tile::Empty { .. } => write!(f, "."),
            Tile::Tree { .. } => write!(f, "#"),
        }
    }
}

#[derive(Debug)]
struct Grid {
    width: usize,
    height: usize,

    tiles: Vec<Tile>,
}

impl Grid {
    fn get(&self, x: usize, y: usize) -> Option<&Tile> {
        // Wrap horizontally
        let x = x % self.width;
        return Some(self.tiles.get(y * self.width + x)?);
    }

    fn count(&self, right: usize, down: usize) -> usize {
        let mut x = 0;
        let mut y = 0;
        let mut count = 0;

        while y < self.height {
            match self.get(x, y).unwrap() {
                Tile::Tree { .. } => count = count + 1,
                _ => {}
            }

            x = x + right;
            y = y + down;
        }

        count
    }
}

impl Display for Grid {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        for y in 0..self.height {
            for x in 0..self.width {
                let tile = self.get(x, y).unwrap();
                write!(f, "{}", tile)?;
            }
            write!(f, "\n")?;
        }

        return Ok(());
    }
}

#[derive(Error, Debug)]
#[error{"parse error"}]
struct ParseError {}

fn parse_input(reader: &mut dyn Read) -> Result<Grid, ParseError> {
    let mut grid = Grid {
        width: 0,
        height: 0,
        tiles: Vec::new(),
    };

    let mut x: usize = 0;
    let mut y: usize = 0;

    for b in reader.bytes() {
        match b {
            Ok(b'\n') => {
                if x > grid.width {
                    grid.width = x;
                }

                x = 0;
                y = y + 1;
            }

            Ok(b'.') => {
                x = x + 1;
                grid.tiles.push(Tile::Empty { x, y });
            }
            Ok(b'#') => {
                x = x + 1;
                grid.tiles.push(Tile::Tree { x, y })
            }
            _ => return Err(ParseError {}),
        }
    }

    grid.height = y;

    Ok(grid)
}

fn main() {
    let grid = parse_input(&mut io::stdin().lock()).unwrap_or_else(|err| {
        println!("failed to parse input: {}", err);
        process::exit(1);
    });

    println!("PART 1: {}", grid.count(3, 1));
    println!(
        "PART 2: {}",
        grid.count(1, 1)
            * grid.count(3, 1)
            * grid.count(5, 1)
            * grid.count(7, 1)
            * grid.count(1, 2)
    );
}
