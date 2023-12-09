package year2023

import utils.Puzzle

import scala.Console.println
import scala.util.matching.Regex

object Day02 extends Puzzle {
  trait Color {
    val point: Int
    val Max: Int
  }
  case class Blue(point: Int) extends Color {
    val Max = 14
  }
  case class Green(point: Int) extends Color {
    val Max = 13
  }
  case class Red(point: Int) extends Color {
    val Max = 12
  }
  case class Game(number: Int, colors: List[Color])
  object Game {
    def apply(string: String) = {
      new Game(
        number = string // Game 1: 1 blue
          .split(":")
          .head // Game 1
          .split(" ")
          .last // 1
          .toString
          .toInt,
        colors = string
          .split(":")
          .last // 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
          .split(";")
          .flatMap { // 3 blue, 4 red
            _.split(",") // 3 blue
          }
          .map { color =>
            color.trim match {
              case x if x.contains("blue") =>
                Blue(x.split(" ").head.toString.toInt)
              case x if x.contains("green") =>
                Green(x.split(" ").head.toString.toInt)
              case x if x.contains("red") =>
                Red(x.split(" ").head.toString.toInt)
            }
          }
          .toList
      )
    }
  }

  def part1(input: List[String]): Int = {
    input
      .map(line => Game.apply(line))
      .filterNot { game =>
        game.colors.find { c =>
          (c.point > c.Max)
        }.isDefined
      }
      .map(a => a.number)
      .sum
  }

  def part2(input: List[String]): Int = {
    input
      .map(line => Game.apply(line))
      .map(_.colors.groupBy(_.getClass).values.map(_.maxBy(_.point)))
      .map(_.map(_.point).product)
      .sum
  }

}
