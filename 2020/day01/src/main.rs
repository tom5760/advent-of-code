use std::error::Error;
use std::io::{self, BufRead};
use std::process;

const TARGET: i32 = 2020;

fn main() {
    let expenses = parse_input(&mut io::stdin().lock()).unwrap_or_else(|err| {
        println!("failed to parse input: {}", err);
        process::exit(1);
    });

    let (x, y) = part1(&expenses, TARGET).unwrap_or_else(|| {
        println!("PART 1: target not found");
        process::exit(1);
    });

    println!("PART 1: {} * {} = {}", x, y, x * y);

    let (x, y, z) = part2(&expenses, TARGET).unwrap_or_else(|| {
        println!("PART @: target not found");
        process::exit(1);
    });

    println!("PART 2: {} * {} * {} = {}", x, y, z, x * y * z);
}

fn parse_input(reader: &mut dyn BufRead) -> Result<Vec<i32>, Box<dyn Error>> {
    return reader
        .lines()
        .map(|line| Ok(line?.parse::<i32>()?))
        .collect();
}

fn part1(expenses: &Vec<i32>, target: i32) -> Option<(i32, i32)> {
    for x in expenses {
        for y in expenses {
            if x + y == target {
                return Some((*x, *y));
            }
        }
    }

    return None;
}

fn part2(expenses: &Vec<i32>, target: i32) -> Option<(i32, i32, i32)> {
    for x in expenses {
        for y in expenses {
            for z in expenses {
                if x + y + z == target {
                    return Some((*x, *y, *z));
                }
            }
        }
    }

    return None;
}
