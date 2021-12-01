use std::error::Error;
use std::fmt::Debug;
use std::io::{self, BufRead};
use std::str::FromStr;

use anyhow::Result;

use thiserror::Error;

#[derive(Error, Debug)]
pub enum InputError<T: FromStr + Debug>
where
    T::Err: Error,
    T::Err: 'static,
{
    #[error{"i/o error on line {line:?}: {source}"}]
    IO { source: io::Error, line: usize },

    #[error{"parse error on line {line:?}: {source}"}]
    Parse { source: T::Err, line: usize },
}

pub fn stdin_parse<T: FromStr + Debug>() -> Result<Vec<T>, InputError<T>>
where
    T::Err: Error,
    T::Err: 'static,
{
    parse(&mut io::stdin().lock())
}

pub fn parse<T: FromStr + Debug>(reader: &mut dyn BufRead) -> Result<Vec<T>, InputError<T>>
where
    T::Err: Error,
{
    reader
        .lines()
        .enumerate()
        .map(|(i, line)| match line {
            Ok(line) => match line.parse::<T>() {
                Ok(value) => return Ok(value),
                Err(err) => {
                    return Err(InputError::Parse {
                        source: err,
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
