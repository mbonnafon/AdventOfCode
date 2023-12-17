package adventofcode.year2021

import utils.Puzzle

object Day04 extends Puzzle {

  def part1(input: List[String]): Long = {
    val draws = input.head.split(",").map(_.toInt).toList
    val boards = input.drop(2)
    print(boards)
    // val game = new Game(draws, )
    return 0
  }

  def part2(input: List[String]): Long = {
    return 0
  }

}

class Game(draws: List[Int], boards: List[Board]) {}

class Board(lines: List[String]) {}
