#[macro_use]
extern crate lazy_static;

use std::io::{self, BufRead};
use std::process;

use thiserror::Error;

use regex::Regex;

#[derive(Debug, PartialEq, Eq, Default)]
struct Record {
    byr: Option<String>,
    iyr: Option<String>,
    eyr: Option<String>,
    hgt: Option<String>,
    hcl: Option<String>,
    ecl: Option<String>,
    pid: Option<String>,
    cid: Option<String>,
}

impl Record {
    fn valid(&self) -> bool {
        self.byr.is_some()
            && self.iyr.is_some()
            && self.eyr.is_some()
            && self.hgt.is_some()
            && self.hcl.is_some()
            && self.ecl.is_some()
            && self.pid.is_some()
    }

    fn strict(&self) -> bool {
        if !self.valid() {
            return false;
        }

        if let Some(byr) = &self.byr {
            if let Ok(byr) = byr.parse::<usize>() {
                if byr < 1920 || byr > 2002 {
                    return false;
                }
            } else {
                return false;
            }
        }

        if let Some(iyr) = &self.iyr {
            if let Ok(iyr) = iyr.parse::<usize>() {
                if iyr < 2010 || iyr > 2020 {
                    return false;
                }
            } else {
                return false;
            }
        }

        if let Some(eyr) = &self.eyr {
            if let Ok(eyr) = eyr.parse::<usize>() {
                if eyr < 2010 || eyr > 2030 {
                    return false;
                }
            } else {
                return false;
            }
        }

        if let Some(hgt) = &self.hgt {
            match hgt.split_at(hgt.len() - 2) {
                (x, "cm") => {
                    if let Ok(x) = x.parse::<usize>() {
                        if x < 150 || x > 193 {
                            return false;
                        }
                    }
                }
                (x, "in") => {
                    if let Ok(x) = x.parse::<usize>() {
                        if x < 59 || x > 76 {
                            return false;
                        }
                    }
                }
                _ => return false,
            }
        }

        if let Some(hcl) = &self.hcl {
            lazy_static! {
                static ref RE: Regex = Regex::new("^#[0-9a-f]{6}$").unwrap();
            }

            if !RE.is_match(hcl) {
                return false;
            }
        }

        if let Some(ecl) = &self.ecl {
            match ecl.as_str() {
                "amb" => {}
                "blu" => {}
                "brn" => {}
                "gry" => {}
                "grn" => {}
                "hzl" => {}
                "oth" => {}
                _ => return false,
            }
        }

        if let Some(pid) = &self.pid {
            if pid.len() != 9 {
                return false;
            }
        }

        return true;
    }
}

#[derive(Error, Debug)]
#[error{"parse error on line {line:?}"}]
struct ParseError {
    line: usize,
}

fn parse_input(reader: &mut dyn BufRead) -> Result<Vec<Record>, ParseError> {
    let mut records = Vec::new();
    let mut cur = Record::default();

    for (i, line) in reader.lines().enumerate() {
        let line = line.map_err(|_| ParseError { line: i })?;

        if line == "" {
            records.push(cur);
            cur = Record::default();
            continue;
        }

        for pair in line.split(' ') {
            match pair.split(':').collect::<Vec<_>>().as_slice() {
                [key, value] => match &key[..] {
                    "byr" => cur.byr = Some(value.to_string()),
                    "iyr" => cur.iyr = Some(value.to_string()),
                    "eyr" => cur.eyr = Some(value.to_string()),
                    "hgt" => cur.hgt = Some(value.to_string()),
                    "hcl" => cur.hcl = Some(value.to_string()),
                    "ecl" => cur.ecl = Some(value.to_string()),
                    "pid" => cur.pid = Some(value.to_string()),
                    "cid" => cur.cid = Some(value.to_string()),
                    _ => return Err(ParseError { line: i }),
                },
                _ => return Err(ParseError { line: i }),
            }
        }
    }

    Ok(records)
}

fn main() {
    let records = parse_input(&mut io::stdin().lock()).unwrap_or_else(|err| {
        println!("failed to parse input: {}", err);
        process::exit(1);
    });

    println!("PART 1: {}", records.iter().filter(|r| r.valid()).count());
    println!("PART 2: {}", records.iter().filter(|r| r.strict()).count());
}
