use std::error::Error;
use std::io::{self, BufRead};
use std::str::FromStr;

use thiserror::Error;

#[derive(Error, Debug)]
pub enum InputError {
    #[error{"i/o error on line {line:?}: {source}"}]
    IO { source: io::Error, line: usize },

    #[error{"parse error on line {line:?}: {source}"}]
    Parse { source: Box<dyn Error>, line: usize },
}

pub fn stdin_parse<T: FromStr>() -> Result<Vec<T>, InputError>
where
    T: FromStr,
    <T as FromStr>::Err: Error,
    <T as FromStr>::Err: 'static,
{
    parse(&mut io::stdin().lock())
}

pub fn parse<T>(reader: &mut dyn BufRead) -> Result<Vec<T>, InputError>
where
    T: FromStr,
    <T as FromStr>::Err: Error,
    <T as FromStr>::Err: 'static,
{
    reader
        .lines()
        .enumerate()
        .map(|(i, line)| match line {
            Ok(line) => match line.parse::<T>() {
                Ok(value) => return Ok(value),
                Err(err) => {
                    return Err(InputError::Parse {
                        source: Box::new(err),
                        line: i,
                    })
                }
            },
            Err(err) => {
                return Err(InputError::IO {
                    source: err,
                    line: i,
                })
            }
        })
        .collect()
}
