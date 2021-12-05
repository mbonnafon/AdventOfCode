package com.mbonnafon.adventofcode.year2021

import com.mbonnafon.adventofcode.Puzzle

object Day03 extends Puzzle {
  implicit class StringToInt(s: String) {
    def toInt(radix: Int) = Integer.parseInt(s, radix)
  }

  def part1(input: List[String]): Int = {
    val gammaRate =
      input.transpose
        .map(_.groupBy(identity).maxBy(_._2.size)._1)
        .mkString("")
        .toInt(2)

    val epsilonRate: Int =
      input.transpose
        .map(_.groupBy(identity).minBy(_._2.size)._1)
        .mkString("")
        .toInt(2)

    gammaRate * epsilonRate
  }

  def part2(input: List[String]): Int = {
    // TODO
    print(input.transpose)
    return 0
  }
}
