use std::process;
use std::str::FromStr;

use thiserror::Error;

#[path = "../input.rs"]
mod input;

#[derive(Error, Debug, Clone, Copy)]
enum ParseError {
    #[error{"line is incorrect length"}]
    Length,

    #[error{"column definition invalid"}]
    Col,

    #[error{"row definition invalid"}]
    Row,
}

enum SearchDir {
    Left,
    Right,
}

#[derive(Debug)]
struct Seat {
    row: usize,
    col: usize,
}

impl Seat {
    fn id(&self) -> usize {
        self.row * 8 + self.col
    }
}

const TOTAL_ROWS: usize = 128;
const TOTAL_COLS: usize = 8;

trait BinSearch {
    fn bin_search(&mut self, total: usize) -> usize;
}

impl BinSearch for std::slice::Iter<'_, SearchDir> {
    fn bin_search(&mut self, total: usize) -> usize {
        let (rv, _) = self.fold((0, total), |(left, right), dir| {
            let range = (right - left) / 2;
            match dir {
                SearchDir::Left => (left, right - range),
                SearchDir::Right => (left + range, right),
            }
        });

        return rv;
    }
}

impl FromStr for Seat {
    type Err = ParseError;

    fn from_str(line: &str) -> Result<Self, Self::Err> {
        let rowsize = (TOTAL_ROWS as f64).log2() as usize;
        let colsize = (TOTAL_COLS as f64).log2() as usize;
        let defsize = rowsize + colsize;

        if line.len() != defsize {
            return Err(ParseError::Length);
        }

        let row = line
            .get(..rowsize)
            .ok_or(ParseError::Length)?
            .chars()
            .map(|dir| match dir {
                'F' => Ok(SearchDir::Left),
                'B' => Ok(SearchDir::Right),
                _ => Err(ParseError::Col),
            })
            .collect::<Result<Vec<SearchDir>, ParseError>>()?
            .iter()
            .bin_search(TOTAL_ROWS);

        let col = line
            .get(rowsize..)
            .ok_or(ParseError::Length)?
            .chars()
            .map(|dir| match dir {
                'L' => Ok(SearchDir::Left),
                'R' => Ok(SearchDir::Right),
                _ => Err(ParseError::Row),
            })
            .collect::<Result<Vec<SearchDir>, ParseError>>()?
            .iter()
            .bin_search(TOTAL_COLS);

        return Ok(Seat { row, col });
    }
}

fn part2(sorted_ids: Vec<usize>) -> Option<usize> {
    let first = sorted_ids.first().unwrap();
    let last = sorted_ids.last().unwrap();

    for i in *first..*last {
        if let None = sorted_ids.iter().find(|&x| *x == i) {
            return Some(i);
        }
    }

    return None;
}

fn main() {
    let seats: Vec<Seat> = input::stdin_parse().unwrap_or_else(|err| {
        println!("failed to parse input: {}", err);
        process::exit(1);
    });

    let mut ids: Vec<usize> = seats.iter().map(|seat| seat.id()).collect();

    println!("PART 1: {}", ids.iter().max().unwrap_or(&0));

    ids.sort_unstable();

    println!("PART 2: {}", part2(ids).unwrap());
}
