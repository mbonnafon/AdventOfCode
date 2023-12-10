package year2023

import utils.Puzzle

import scala.annotation.tailrec

object Day08 extends Puzzle {

  case class Game(
      instructions: List[String],
      mapping: Map[String, (String, String)]
  )

  object Game {
    def apply(input: List[String]) = {
      new Game(
        instructions = input.head.toString.toList.map(_.toString),
        mapping = input
          .drop(1)
          .drop(1)
          .flatMap { line =>
            val parts = line.split("=")
            val key = parts(0).trim
            val values = "\\((.+)\\)".r
              .findFirstMatchIn(parts(1))
              .map(_.group(1))
              .getOrElse("")
              .split(", ")
            Map(key -> (values(0), values(1)))
          }
          .toMap
      )
    }
  }
  def part1(input: List[String]): Int = {
    val game = Game.apply(input)
    val start = "AAA"

    def countGameMovement(
        acc: Int,
        movement: List[String],
        currentPosition: String
    ): Int = {
      val refreshMovement = if (movement.size == 0) {
        game.instructions
      } else {
        movement
      }
      val side = refreshMovement.head.toString match {
        case "L" => game.mapping.get(currentPosition).get._1
        case "R" => game.mapping.get(currentPosition).get._2
      }
      if (side == "ZZZ") {
        acc + 1
      } else {
        countGameMovement(acc + 1, refreshMovement.drop(1), side)
      }
    }

    countGameMovement(0, game.instructions, start)
  }

  def part2(input: List[String]): Int = {
    val game = Game.apply(input)

    val startWith = game.mapping.filterKeys(_.endsWith("A")).keySet.toList

    def countGameMovement(
        acc: Int,
        movement: List[String],
        currentPosition: String
    ): Int = {
      val refreshMovement = if (movement.size == 0) {
        game.instructions
      } else {
        movement
      }
      val newPosition = refreshMovement.head.toString match {
        case "L" => game.mapping.get(currentPosition).get._1
        case "R" => game.mapping.get(currentPosition).get._2
      }

      if (newPosition.endsWith("Z")) {
        return acc + 1
      } else {
        countGameMovement(acc + 1, refreshMovement.drop(1), newPosition)
      }
    }

    val endsWith = startWith.map(countGameMovement(0, game.instructions, _))

    def lcm(nums: Seq[Long]): Long = {
      def _gcd(a: Long, b: Long): Long =
        if (a == b) a else _gcd(Math.max(a, b) - Math.min(a, b), Math.min(a, b))
      def _lcm(a: Long, b: Long): Long =
        if (a != 0 || b != 0) a * b / _gcd(a, b) else 0
      nums.reduce(_lcm)
    }

    println(lcm(endsWith.toSeq.map(_.toLong))) // it's a Long
    1
  }

}
