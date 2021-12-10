package com.mbonnafon.adventofcode.year2021

import com.mbonnafon.adventofcode.Puzzle

object Day04 extends Puzzle {

  def part1(input: List[String]): Int = {
    val draws = input.head.split(",").map(_.toInt).toList
    val boards = input.drop(2).print(boards)
    // val game = new Game(draws, )
    return 0
  }

  def part2(input: List[String]): Int = {
    return 0
  }

}

class Game(draws: List[Int], boards: List[Board]) {}

class Board(lines: List[String]) {}
