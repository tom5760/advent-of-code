use std::collections::HashSet;
use std::io::{self, BufRead};
use std::process;

use anyhow::{Context, Result};

#[derive(Debug, Default)]
struct Group {
    people: Vec<HashSet<char>>,
}

fn intersect(sets: &[HashSet<char>]) -> HashSet<char> {
    let mut rv = HashSet::new();

    for set in sets {
        for value in set {
            rv.insert(*value);
        }
    }

    return rv;
}

fn union(sets: &[HashSet<char>]) -> HashSet<char> {
    let mut rv = HashSet::new();
    let values = intersect(sets);

    'outer: for value in values {
        for set in sets {
            if !set.contains(&value) {
                continue 'outer;
            }
        }

        rv.insert(value);
    }

    return rv;
}

fn parse_input(reader: &mut dyn BufRead) -> Result<Vec<Group>> {
    let mut groups = Vec::new();
    let mut cur = Group::default();

    for (i, line) in reader.lines().enumerate() {
        let line = line.with_context(|| format!("failed to read line {}", i))?;

        if line == "" {
            groups.push(cur);
            cur = Group::default();
            continue;
        }

        let mut person = HashSet::new();

        for c in line.chars() {
            person.insert(c);
        }

        cur.people.push(person);
    }

    groups.push(cur);

    return Ok(groups);
}

fn main() {
    let groups = parse_input(&mut io::stdin().lock()).unwrap_or_else(|err| {
        println!("failed to parse input: {}", err);
        process::exit(1);
    });

    println!(
        "PART 1: {}",
        groups
            .iter()
            .map(|g| intersect(&g.people).len())
            .sum::<usize>()
    );
    println!(
        "PART 2: {}",
        groups.iter().map(|g| union(&g.people).len()).sum::<usize>()
    );
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_intersect() {
        let mut h1 = HashSet::new();
        h1.insert('a');
        h1.insert('b');

        let mut h2 = HashSet::new();
        h2.insert('a');
        h2.insert('c');

        let h3 = intersect(&vec![h1, h2]);

        assert_eq!(h3.contains(&'a'), true);
        assert_eq!(h3.contains(&'b'), true);
        assert_eq!(h3.contains(&'c'), true);
    }

    #[test]
    fn test_union() {
        let mut h1 = HashSet::new();
        h1.insert('a');
        h1.insert('b');

        let mut h2 = HashSet::new();
        h2.insert('a');
        h2.insert('c');

        let h3 = union(&vec![h1, h2]);

        assert_eq!(h3.contains(&'a'), true);
        assert_eq!(h3.contains(&'b'), false);
        assert_eq!(h3.contains(&'c'), false);
    }
}
