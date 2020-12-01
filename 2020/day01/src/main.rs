use std::error::Error;
use std::io::{self, BufRead};
use std::process;

const TARGET: i32 = 2020;

fn main() {
    let expenses = parse_input(&mut io::stdin().lock()).unwrap_or_else(|err| {
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

fn parse_input(reader: &mut dyn BufRead) -> Result<Vec<i32>, Box<dyn Error>> {
    return reader
        .lines()
        .map(|line| Ok(line?.parse::<i32>()?))
        .collect();
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
    use std::fs::File;
    use std::io::BufReader;

    use super::*;

    fn ex1() -> Vec<i32> {
        return vec![ 1721, 979, 366, 299, 675, 1456 ];
    }

    fn input() -> Result<Vec<i32>, Box<dyn Error>> {
        let f = File::open("input")?;
        let mut bf = BufReader::new(f);

        return parse_input(&mut bf);
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
}

// #[cfg(test)]
// mod tests {
//     use super::*;
//     use test::Bencher;

//     use std::fs::File;
//     use std::io::BufReader;

//     use itertools::Itertools;

//     fn get_input() -> Result<Vec<i32>, Box<dyn Error>> {
//         let f = File::open("input")?;
//         let mut bf = BufReader::new(f);

//         return parse_input(&mut bf);
//     }

//     fn find_target_sum(expenses: &Vec<i32>, target: i32, n: usize) -> Option<i32> {
//         Some(expenses
//             .into_iter()
//             .combinations(n)
//             .find(|v| {
//                 let sum: i32 = v.iter().map(|x| *x).sum();
//                 sum == target
//             })?
//             .into_iter()
//             .fold(1, |a, &b| a * b))
//     }

//     // fn find_target_sum(expenses: &Vec<i32>, target: i32, n: usize) -> Option<i32> {
//     //     let x = expenses.iter().combinations(n).find(|v| {
//     //         let sum: i32 = v.iter().sum();
//     //         sum == target
//     //     })?;

//     //     Some(itertools::fold(x, 1i32, |a, b| a * b))
//     // }

//     fn part1_iter(expenses: &Vec<i32>, target: i32) -> Option<i32> {
//         find_target_sum(expenses, target, 2)
//     }

//     fn part2_iter(expenses: &Vec<i32>, target: i32) -> Option<i32> {
//         find_target_sum(expenses, target, 3)
//     }

//     #[bench]
//     fn bench_part1_iter(b: &mut Bencher) {
//         let expenses = get_input().unwrap();
//         b.iter(|| part1_iter(&expenses, TARGET));
//     }

//     #[bench]
//     fn bench_part1_loop(b: &mut Bencher) {
//         let expenses = get_input().unwrap();
//         b.iter(|| part1(&expenses, TARGET));
//     }

//     #[bench]
//     fn bench_part2_iter(b: &mut Bencher) {
//         let expenses = get_input().unwrap();
//         b.iter(|| part2_iter(&expenses, TARGET));
//     }

//     #[bench]
//     fn bench_part2_loop(b: &mut Bencher) {
//         let expenses = get_input().unwrap();
//         b.iter(|| part2(&expenses, TARGET));
//     }
// }
