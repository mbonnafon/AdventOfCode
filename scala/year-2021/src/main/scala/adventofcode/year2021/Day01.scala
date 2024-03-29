package adventofcode.year2021

import utils.Puzzle

object Day01 extends Puzzle {

  def part1(input: List[String]): Long =
    input
      .map(_.toInt)
      .sliding(2)
      .foldLeft(0) {
        case (acc, List(a, b)) => if (a < b) acc + 1 else acc
        case _                 => 0
      }

  def part2(input: List[String]): Long =
    input
      .map(_.toInt)
      .sliding(3)
      .foldLeft[(Option[List[Int]], Int)]((None, 0)) {
        case ((Some(previous), acc), current) =>
          if (previous.sum < current.sum) (Some(current), acc + 1)
          else (Some(current), acc)
        case (_, current) => (Some(current), 0)
      }
      ._2

}
