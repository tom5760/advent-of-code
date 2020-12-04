#![feature(test)]

use std::process;

#[path = "../input.rs"]
mod input;

const TARGET: i32 = 2020;

fn main() {
    let expenses = input::stdin_parse().unwrap_or_else(|err| {
        println!("failed to parse input: {}", err);
        process::exit(1);
    });

    let x = part1(&expenses).unwrap_or_else(|| {
        println!("PART 1: target not found");
        process::exit(1);
    });

    println!("PART 1: {}", x);

    let x = part2(&expenses).unwrap_or_else(|| {
        println!("PART 2: target not found");
        process::exit(1);
    });

    println!("PART 2: {}", x,);
}

fn part1(expenses: &Vec<i32>) -> Option<i32> {
    for x in expenses {
        for y in expenses {
            if x + y == TARGET {
                return Some(x * y);
            }
        }
    }

    return None;
}

fn part2(expenses: &Vec<i32>) -> Option<i32> {
    for x in expenses {
        for y in expenses {
            for z in expenses {
                if x + y + z == TARGET {
                    return Some(x * y * z);
                }
            }
        }
    }

    return None;
}

#[cfg(test)]
mod tests {

    extern crate test;

    use std::fs::File;
    use std::io::BufReader;
use std::error::Error;

    use test::Bencher;

    use itertools::Itertools;

    use super::*;

    fn ex1() -> Vec<i32> {
        return vec![1721, 979, 366, 299, 675, 1456];
    }

    fn input() -> Result<Vec<i32>, Box<dyn Error>> {
        let f = File::open("inputs/day01")?;
        let mut bf = BufReader::new(f);

        return input::parse(&mut bf);
    }

    #[test]
    fn test_ex1() {
        let rs = part1(&ex1()).unwrap();
        assert_eq!(rs, 514579)
    }

    #[test]
    fn test_ex2() {
        let rs = part2(&ex1()).unwrap();
        assert_eq!(rs, 241861950)
    }

    #[test]
    fn test_part1() {
        let rs = part1(&input().unwrap()).unwrap();
        assert_eq!(rs, 55776)
    }

    #[test]
    fn test_part2() {
        let rs = part2(&input().unwrap()).unwrap();
        assert_eq!(rs, 223162626)
    }

    #[bench]
    fn bench_part1_loop(b: &mut Bencher) {
        b.iter(|| part1(&input().unwrap()));
    }

    #[bench]
    fn bench_part2_loop(b: &mut Bencher) {
        b.iter(|| part2(&input().unwrap()));
    }

    // testing out using fancy iterators.

    fn find_target_sum(expenses: &Vec<i32>, n: usize) -> Option<i32> {
        Some(
            expenses
                .into_iter()
                .combinations(n)
                .find(|v| v.iter().map(|x| *x).sum::<i32>() == TARGET)?
                .into_iter()
                .product(),
        )
    }

    fn part1_iter(expenses: &Vec<i32>) -> Option<i32> {
        find_target_sum(expenses, 2)
    }

    fn part2_iter(expenses: &Vec<i32>) -> Option<i32> {
        find_target_sum(expenses, 3)
    }

    #[test]
    fn test_part1_iter() {
        let rs = part1_iter(&input().unwrap()).unwrap();
        assert_eq!(rs, 55776)
    }

    #[test]
    fn test_part2_iter() {
        let rs = part2_iter(&input().unwrap()).unwrap();
        assert_eq!(rs, 223162626)
    }

    #[bench]
    fn bench_part1_iter(b: &mut Bencher) {
        b.iter(|| part1_iter(&input().unwrap()));
    }

    #[bench]
    fn bench_part2_iter(b: &mut Bencher) {
        b.iter(|| part2_iter(&input().unwrap()));
    }

    // testing out using a map

    fn part1_map(expenses: &Vec<i32>) -> Option<i32> {
        let mut map = std::collections::HashSet::new();

        for x in expenses {
            let diff = TARGET - x;

            if map.contains(&diff) {
                return Some(x * diff);
            }

            map.insert(*x);
        }

        return None;
    }

    #[test]
    fn test_part1_map() {
        let rs = part1_map(&input().unwrap()).unwrap();
        assert_eq!(rs, 55776)
    }

    #[bench]
    fn bench_part1_map(b: &mut Bencher) {
        b.iter(|| part1_map(&input().unwrap()));
    }
}
