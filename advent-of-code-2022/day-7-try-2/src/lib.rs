use std::str::Lines;

use regex::Regex;


#[derive(Debug, PartialEq)]
pub enum FSObj {
    Folder {
        name: String,
        children: Vec<FSObj>
    },
    File {
        name: String,
        size: usize
    }
}

pub fn reconstruct_file_system(input: Lines<'_>) -> FSObj {
    let mut root = FSObj::Folder { name: String::from("root"), children: vec![] };

    for line in input {
        let obj = parse_line(line);
        match obj {
            Some(fs) => {
                match fs {
                    FSObj::File { name, size } => {
                        println!("File: {} {}", name, size);
                    },
                    FSObj::Folder { name, children } => {
                        println!("Folder: {}, children: {:?}", name, children);
                    },
                }

            }            
            None => todo!(),
        }
    }


    root
}

pub fn parse_line(line: &str) -> Option<FSObj> {
    let file_re = Regex::new(r"\d+(\s|\w|.)+").unwrap();
    let cd_re = Regex::new(r"^\$\scd\s+.+").unwrap();

    if file_re.is_match(line) {
        let mut split = line.split_whitespace();
        let size = split.next().unwrap().parse::<usize>().unwrap();
        let name = split.next().unwrap().to_string();
        Some(FSObj::File { name, size })
    } else if cd_re.is_match(line) {
        let mut split = line.split_whitespace();
        split.next();
        let name = split.next().unwrap().to_string();
        Some(FSObj::Folder { name, children: vec![] })
    } else {
        None
    }
}

pub fn find_small_directories(input: Lines<'_>) -> usize {
    println!("{:?}", input);
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_input() {
        // arrange
        let test_input = include_str!("../test-input").lines();
        let expected = 95437;

        // act
        let actual = find_small_directories(test_input);

        // assert
        assert_eq!(
            expected, actual,
            "Expected: {} should be equal to actual: {}",
            expected, actual
        );
    }
}
