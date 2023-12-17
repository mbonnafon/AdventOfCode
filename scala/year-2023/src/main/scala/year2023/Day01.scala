package year2023

import utils.Puzzle

import scala.Console.println
import scala.util.matching.Regex

object Day01 extends Puzzle {
  def part1(input: List[String]): Long = {
    input
      .map(a => a.filter(_.isDigit))
      .map { a =>
        a.size match {
          case 1 => a ++ a
          case 2 => a
          case _ => a.head.toString ++ a.last.toString
        }
      }
      .map(_.toInt)
      .sum
  }

  def part2(input: List[String]): Long = {
    val mapping = Map(
      "one" -> "1",
      "two" -> "2",
      "three" -> "3",
      "four" -> "4",
      "five" -> "5",
      "six" -> "6",
      "seven" -> "7",
      "eight" -> "8",
      "nine" -> "9"
    )

    input
      .map { line =>
        line.zipWithIndex
          .map { case (c, idx) =>
            if (c.isDigit) Some(c)
            else {
              mapping
                .find { number =>
                  line.slice(idx, line.length).startsWith(number._1)
                }
                .map(_._2)
            }
          }
          .collect { case v: Some[Any] =>
            v.value.toString
          }
      }
      .map(line => (line.head ++ line.last).toInt)
      .sum
  }

}
