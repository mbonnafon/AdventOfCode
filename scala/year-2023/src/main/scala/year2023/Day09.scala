package year2023

import utils.Puzzle

import scala.annotation.tailrec

object Day09 extends Puzzle {
  def part1(input: List[String]): Long = {
    val ref = input.map(_.split(" ").map(_.toInt).toSeq)

    val toPrediction: Seq[Int] => Seq[Int] = input =>
      input.sliding(2).collect { case Seq(a, b) => b - a }.toList

    def predict(value: Seq[Int]): Int = {
      val in = toPrediction(value)
      if (in.forall(_ == 0)) {
        value.last
      } else {
        value.last + predict(in)
      }
    }

    ref.map(predict(_)).sum
  }

  def part2(input: List[String]): Long = {
    val ref = input.map(_.split(" ").map(_.toInt).toSeq)

    val toPrediction: Seq[Int] => Seq[Int] = input =>
      input.sliding(2).collect { case Seq(a, b) => b - a }.toList

    def predict(value: Seq[Int]): Int = {
      val in = toPrediction(value)
      if (in.forall(_ == 0)) {
        value.head
      } else {
        value.head - predict(in)
      }
    }

    ref.map(predict(_)).sum
  }

}
