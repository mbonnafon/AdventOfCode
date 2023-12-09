package com.mbonnafon.adventofcode.year2021

import com.mbonnafon.adventofcode.{Puzzle, Reader}

import scala.collection.mutable.Stack

object Day10 extends Puzzle {

  def part1(input: List[String]): Int = {
    input.foreach { line =>
      val s = Stack[Char]()
      line.foreach { c =>
        c.toString match {
          case "(" | "[" | "{" | "<" => s.push(c)
          case _                     => print(s.pop)
        }
      }
    }
    9
  }

  def part2(input: List[String]): Int =
    0

}
