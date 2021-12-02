package com.mbonnafon.adventofcode

trait Puzzle {

  def solve(resource: String): (Int, Int) = {
    val input = Reader(resource).read(_.toInt)
    (part1(input), part2(input))
  }

  def part1(input: List[Int]): Int

  def part2(input: List[Int]): Int

}
