package year2023

import scala.math.pow
import utils.Puzzle

object Day04 extends Puzzle {
  case class Card(
      number: Int,
      winningNumbers: List[Int],
      myNumbers: List[Int]
  )

  object Card {
    def apply(string: String) = {
      // Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
      new Card(
        number = string
          .split(":")
          .head // Card 1
          .split(" ")
          .last
          .toInt,
        winningNumbers = string
          .split(":")
          .last // 41 48 83 86 17 | 83 86  6 31 17  9 48 53
          .split('|')
          .head
          .split(" ")
          .filter(_.nonEmpty)
          .map(_.toInt)
          .toList,
        myNumbers = string
          .split(":")
          .last // 41 48 83 86 17 | 83 86  6 31 17  9 48 53
          .split('|')
          .last
          .split(" ")
          .filter(_.nonEmpty)
          .map(_.toInt)
          .toList
      )
    }
  }

  def part1(input: List[String]): Int = {
    input
      .map(Card.apply(_))
      .map(card => card.myNumbers.filter(x => card.winningNumbers.contains(x)))
      .map { winningNumbers => pow(2, winningNumbers.size - 1) }
      .map(_.toInt)
      .sum
  }

  def part2(input: List[String]): Int = {
    var acc: Map[Int, Int] = Map.empty
    input
      .map(Card.apply(_))
      .foreach { card =>
        val winningNumbers =
          card.myNumbers.filter(x => card.winningNumbers.contains(x))

        acc = acc.updatedWith(card.number) { // original set is always valid
          case Some(v) => Some(v + 1)
          case None    => Some(1)
        }

        if (winningNumbers.nonEmpty) {
          for (i <- (1 + card.number) to (winningNumbers.size + card.number)) {
            val winningFactor = acc.get(card.number).getOrElse(0)
            acc = acc.updatedWith(i) {
              case Some(v) => Some(v + winningFactor)
              case None    => Some(winningFactor)
            }
          }
        }
      }
    acc.values.sum
  }

}
