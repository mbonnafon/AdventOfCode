package adventofcode.year2021

import utils.Puzzle

object Day02 extends Puzzle {

  def part1(input: List[String]): Int = {
    val (horizontal, depth) = {
      input
        .foldLeft((0, 0)) { case ((horizontal, depth), command) =>
          command match {
            case s"forward $unit" => (horizontal + unit.toInt, depth)
            case s"down $unit"    => (horizontal, depth + unit.toInt)
            case s"up $unit"      => (horizontal, depth - unit.toInt)
          }
        }
    }
    horizontal * depth
  }

  def part2(input: List[String]): Int = {
    val (horizontal, depth, _) = {
      input
        .foldLeft((0, 0, 0)) { case ((horizontal, depth, aim), command) =>
          command match {
            case s"down $unit" => (horizontal, depth, aim + unit.toInt)
            case s"up $unit"   => (horizontal, depth, aim - unit.toInt)
            case s"forward $unit" =>
              (horizontal + unit.toInt, depth + aim * unit.toInt, aim)
          }
        }
    }
    horizontal * depth
  }

}
