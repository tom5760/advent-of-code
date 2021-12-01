use std::collections::HashMap;

use anyhow::{Context, Result};

#[path = "../input.rs"]
mod input;

#[derive(Debug, Default)]
struct Rule {
    name: String,
    edges: HashMap<String, usize>
}

static TARGET: &str = "shiny gold";

fn parse_rules(rule_strs: &[String]) -> Result<HashMap<String, Rule>> {
    let mut map: HashMap<String, Rule> = HashMap::new();

    for rule_str in rule_strs {
        let parts: Vec<&str> = rule_str.split(" ").collect();
        let name = parts.get(0..2).context("name missing")?.join(" ");

        let rule = map.entry(name.clone()).or_insert_with(|| Rule{
            name,
            ..Rule::default()
        });

        for edge in parts.get(4..).context("parts missing")?.chunks(4) {
            if edge == vec!["no", "other", "bags."] {
                continue
            }

            let weight: usize = edge
                .get(0).context("edge weight missing")?
                .parse().context("failed to parse edge weight")?;
            let edge_name = edge.get(1..3).context("edge name missing")?.join(" ");

            rule.edges.insert(edge_name, weight);
        }
    }

    return Ok(map);
}

fn has_bag(rules: &HashMap<String, Rule>, rule: &Rule, target: &str) -> bool {
    if rule.name == target {
        return true;
    }

    for key in rule.edges.keys() {
        let rule = match rules.get(key) {
            Some(rule) => rule,
            _ => continue,
        };

        if has_bag(rules, rule, target) {
            return true;
        }
    }

    return false;
}

fn part1(rules: &HashMap<String, Rule>) -> usize {
    return rules
        .values()
        .map(|rule| has_bag(rules, rule, TARGET))
        .filter(|has| *has)
        .count();
}

fn main() -> Result<()> {
    let rules = parse_rules(&input::stdin_parse()
        .context("failed to parse input")?)?;

    println!("PART 1: {}", part1(&rules));

    return Ok(());
}