use std::error::Error;
use std::io::{self, BufRead};
use std::str::FromStr;

pub fn stdin_parse<T: FromStr>() -> Result<Vec<T>, Box<dyn Error>>
where
    T: FromStr,
    <T as FromStr>::Err: Error,
    <T as FromStr>::Err: 'static,
{
    return parse(&mut io::stdin().lock());
}

pub fn parse<T>(reader: &mut dyn BufRead) -> Result<Vec<T>, Box<dyn Error>>
where
    T: FromStr,
    <T as FromStr>::Err: Error,
    <T as FromStr>::Err: 'static,
{
    return reader.lines().map(|line| Ok(line?.parse::<T>()?)).collect();
}
