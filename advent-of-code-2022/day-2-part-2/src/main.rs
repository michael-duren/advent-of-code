use itertools::Itertools;
use std::str::FromStr;

use color_eyre::eyre::Result;

#[derive(Debug, Clone, Copy)]
enum Move {
    Rock,
    Paper,
    Scissors,
}

#[derive(Debug, Clone, Copy)]
struct Round {
    theirs: Move,
    ours: Move,
}

impl Round {
    fn outcome(self) -> Outcome {
        self.ours.outcome(self.theirs)
    }
    fn our_score(self) -> usize {
        self.ours.inherit_points() + self.outcome().inherent_points()
    }
}

// Mood parser
impl TryFrom<char> for Move {
    type Error = color_eyre::Report;

    fn try_from(c: char) -> Result<Self, Self::Error> {
        match c {
            'A' => Ok(Move::Rock),
            'B' => Ok(Move::Paper),
            'C' => Ok(Move::Scissors),
            _ => Err(color_eyre::eyre::eyre!("not a valid move: {c:?}")),
        }
    }
}

impl TryFrom<char> for Outcome {
    type Error = color_eyre::Report;

    fn try_from(c: char) -> Result<Self, Self::Error> {
        match c {
            'X' => Ok(Outcome::Loss),
            'Y' => Ok(Outcome::Draw),
            'Z' => Ok(Outcome::Win),
            _ => Err(color_eyre::eyre::eyre!("not a valid outcome: {c:?}")),
        }
    }
}

impl FromStr for Round {
    type Err = color_eyre::Report;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut chars = s.chars();
        let (Some(theirs), Some(' '), Some(outcome), None) =
            (chars.next(), chars.next(), chars.next(), chars.next())
        else {
            return Err(color_eyre::eyre::eyre!(
                "expected <theirs>SP<ours>EOF, got {s:?}"
            ));
        };

        let theirs = Move::try_from(theirs)?;
        let outcome = Outcome::try_from(outcome)?;
        let ours = outcome.matching_move(theirs);

        Ok(Self { theirs, ours })
    }
}

impl Move {
    fn inherit_points(self) -> usize {
        match self {
            Move::Rock => 1,
            Move::Paper => 2,
            Move::Scissors => 3,
        }
    }
}

impl Move {
    const ALL_MOVES: [Move; 3] = [Move::Rock, Move::Paper, Move::Scissors];

    fn winning_move(self) -> Self {
        Self::ALL_MOVES
            .iter()
            .copied()
            .find(|m| m.beats(self))
            .expect("at least one move beatus us")
    }

    fn losing_move(self) -> Self {
        Self::ALL_MOVES
            .iter()
            .copied()
            .find(|&m| self.beats(m))
            .expect("we beat at least one move")
    }

    fn drawing_move(self) -> Self {
        self
    }
}

#[derive(Debug, Clone, Copy)]
enum Outcome {
    Loss,
    Draw,
    Win,
}

impl Outcome {
    fn inherent_points(self) -> usize {
        match self {
            Outcome::Win => 6,
            Outcome::Draw => 3,
            Outcome::Loss => 0,
        }
    }
}

impl Outcome {
    fn matching_move(self, theirs: Move) -> Move {
        match self {
            Outcome::Win => theirs.winning_move(),
            Outcome::Draw => theirs.drawing_move(),
            Outcome::Loss => theirs.losing_move(),
        }
    }
}

impl Move {
    fn beats(self, other: Move) -> bool {
        matches!(
            (self, other),
            (Self::Rock, Self::Scissors)
                | (Self::Paper, Self::Rock)
                | (Self::Scissors, Self::Paper)
        )
    }

    fn outcome(self, theirs: Move) -> Outcome {
        if self.beats(theirs) {
            Outcome::Win
        } else if theirs.beats(self) {
            Outcome::Loss
        } else {
            Outcome::Draw
        }
    }
}

fn main() -> color_eyre::Result<()> {
    color_eyre::install()?;

    let total_score: usize = itertools::process_results(
        include_str!("input.txt")
            .lines()
            .map(Round::from_str)
            .map_ok(|r| r.our_score()),
        |it| it.sum(),
    )?;
    dbg!(total_score);

    Ok(())
}
