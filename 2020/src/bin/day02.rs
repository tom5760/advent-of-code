#![feature(try_trait)]

#[macro_use]
extern crate lazy_static;

use std::char::ParseCharError;
use std::num::ParseIntError;
use std::process;
use std::str::FromStr;

use thiserror::Error;

use regex::Regex;

#[path = "../input.rs"]
mod input;

#[derive(Error, Debug)]
enum ParseError {
    #[error{"failed to parse number: {0}"}]
    Int(#[from] ParseIntError),

    #[error{"failed to parse letter: {0}"}]
    Char(#[from] ParseCharError),

    #[error{"invalid line format"}]
    Regex(),
}

struct Policy {
    low: usize,
    high: usize,
    letter: char,
    password: String,
}

impl Policy {
    fn valid1(&self) -> bool {
        let mut count = 0;
        for letter in self.password.chars() {
            if self.letter == letter {
                count = count + 1;
            }

            if count > self.high {
                return false;
            }
        }

        return count >= self.low;
    }

    fn valid2(&self) -> bool {
        let pw: Vec<char> = self.password.chars().collect();
        let a = pw[self.low - 1];
        let b = pw[self.high - 1];

        return ((a == self.letter) || (b == self.letter))
            && !((a == self.letter) & (b == self.letter));
    }
}

impl FromStr for Policy {
    type Err = ParseError;

    fn from_str(line: &str) -> Result<Self, Self::Err> {
        lazy_static! {
            static ref RE: Regex =
                Regex::new(r"(?P<low>\d+)-(?P<high>\d+) (?P<letter>[a-z]): (?P<password>[a-z]+)")
                    .unwrap();
        }

        let captures = match RE.captures(line) {
            Some(captures) => captures,
            None => return Err(ParseError::Regex()),
        };

        Ok(Policy {
            low: captures
                .name("low")
                .ok_or(ParseError::Regex())?
                .as_str()
                .parse()?,
            high: captures
                .name("high")
                .ok_or(ParseError::Regex())?
                .as_str()
                .parse()?,
            letter: captures
                .name("letter")
                .ok_or(ParseError::Regex())?
                .as_str()
                .parse()?,
            password: captures
                .name("password")
                .ok_or(ParseError::Regex())?
                .as_str()
                .to_string(),
        })
    }
}

fn main() {
    let pwdb: Vec<Policy> = input::stdin_parse().unwrap_or_else(|err| {
        println!("failed to parse input: {}", err);
        process::exit(1);
    });

    let part1 = pwdb.iter().filter(|p| p.valid1()).count();
    println!("PART 1: {}", part1);

    let part2 = pwdb.iter().filter(|p| p.valid2()).count();
    println!("PART 2: {}", part2);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part2() {
        let a = Policy {
            low: 1,
            high: 3,
            letter: 'a',
            password: String::from("abcde"),
        };
        assert_eq!(a.valid2(), true);

        let b = Policy {
            low: 1,
            high: 3,
            letter: 'b',
            password: String::from("cdefg"),
        };
        assert_eq!(b.valid2(), false);

        let c = Policy {
            low: 2,
            high: 9,
            letter: 'c',
            password: String::from("ccccccccc"),
        };
        assert_eq!(c.valid2(), false)
    }
}
