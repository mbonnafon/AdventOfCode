package year2023

import utils.Puzzle

import java.io.File
import java.util.Scanner
import scala.io.Source
import scala.math.pow

object Day06 extends Puzzle {

  def part1(input: List[String]): Int = {
    input
      .map(line => line.split(" ").drop(1).filter(_.nonEmpty))
      .transpose // List(7, 9)
      .map(line => (line.head.toInt, line.last.toInt)) // (7,9)
      .map { case (time, distance) =>
        var winningSpeeds: List[Int] = List.empty
        for (holdSpeed <- 0 to time) {
          val remainingTime = time - holdSpeed
          val myDistance = remainingTime * holdSpeed
          if (myDistance > distance) {
            winningSpeeds = winningSpeeds ::: List(holdSpeed)
          }
        }
        winningSpeeds
      }
      .map(_.size)
      .product

  }

  def part2(input: List[String]): Int = {
    val race = input
      .map(_.split(":").last.trim.split(" ").mkString)

    (race.head.toLong, race.last.toLong) match {
      case (time, distance) =>
        val min = (0L to time)
          .find { holdSpeed =>
            val remainingTime = time - holdSpeed
            val myDistance = remainingTime * holdSpeed
            myDistance > distance
          }
        val max = (time to 0L by -1L)
          .find { holdSpeed =>
            val remainingTime = time - holdSpeed
            val myDistance = remainingTime * holdSpeed
            myDistance > distance
          }
        (max.get - min.get + 1L).toInt
    }
  }

}
