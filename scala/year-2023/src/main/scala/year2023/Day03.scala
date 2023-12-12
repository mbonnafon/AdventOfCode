package year2023

import utils.Puzzle

import scala.collection.immutable.{AbstractSeq, LinearSeq}

object Day03 extends Puzzle {

  val sequenceToCheck: Seq[(Int, Int)] = List(
    (-1, -1),
    (-1, 0),
    (-1, 1),
    (1, -1),
    (1, 0),
    (1, 1),
    (0, -1),
    (0, 1)
  )

  def part1(input: List[String]): Int = {
    val matrix: Vector[Vector[String]] =
      input.toVector.map(_.map(_.toString).toVector)

    var sum: Int = 0
    for (x <- 0 to matrix.size - 1) {
      var number: String = ""
      var symbolHasBeenFound: Boolean = false
      for (y <- 0 to matrix(0).size - 1) {
        matrix(x)(y) match {
          case value if value.toIntOption.isDefined =>
            number = number ++ value
            val symbolHasBeenFound0 = sequenceToCheck.filterNot { seq =>
              (x + seq._1 < 0) || (y + seq._2 < 0) ||
              (seq._1 + x > matrix.size - 1) ||
              (seq._2 + y > matrix(0).size - 1) ||
              matrix(x + seq._1)(y + seq._2).toIntOption.isDefined ||
              matrix(x + seq._1)(y + seq._2) == "."
            }
            if (symbolHasBeenFound0.nonEmpty) {
              symbolHasBeenFound = true
            }
            if (symbolHasBeenFound && y == matrix(0).size - 1) {
              sum = sum + number.toInt
              number = ""
              symbolHasBeenFound = false
            }
          case _ =>
            if (symbolHasBeenFound) {
              sum = sum + number.toInt
            }
            number = ""
            symbolHasBeenFound = false
        }
      }
    }
    sum
  }

  def part2(input: List[String]): Int = {

    val matrix: Vector[Vector[String]] =
      input.toVector.map(_.map(_.toString).toVector)

    var sum: Int = 0

    for (x <- 0 to matrix.size - 1) {
      for (y <- 0 to matrix(0).size - 1) {
        matrix(x)(y) match {
          case "*" =>
            val positionOfNumbersCloseToStar: List[(Int, Int)] = sequenceToCheck
              .filterNot { seq =>
                (x + seq._1 < 0) || (y + seq._2 < 0) ||
                (seq._1 + x > matrix.size - 1) ||
                (seq._2 + y > matrix(0).size - 1) ||
                matrix(x + seq._1)(y + seq._2).toIntOption.isEmpty
              }
              .groupBy(_._1)
              .map { seq =>
                seq._2 match { // collect only non adjacents numbers
                  case seq if seq.size == 2 =>
                    val a = seq.head
                    val b = seq.last
                    if (a._1 == 0 && b._1 == 0) {
                      List(a, b)
                    } else if (a._1 == b._1 && (a._2 + b._2) == 0) {
                      List(a, b)
                    } else if (a._2 - b._2 < 0) {
                      Seq(a)
                    } else {
                      seq
                    }
                  case _ => Seq(seq._2.head)
                }
              }
              .flatten
              .toList

            if (positionOfNumbersCloseToStar.size == 2) {
              // keep only valid *
              val eligibleNumbers = positionOfNumbersCloseToStar.map { ref =>
                def getLeftSymbol(xTarget: Int, yTarget: Int): String = {
                  if (
                    yTarget - 1 > -1 &&
                    matrix(xTarget)(yTarget - 1).toIntOption.isDefined
                  ) {
                    getLeftSymbol(xTarget, yTarget - 1) ++ matrix(xTarget)(
                      yTarget - 1
                    )
                  } else ""
                }
                def getRightSymbol(xTarget: Int, yTarget: Int): String = {
                  if (
                    yTarget + 1 < matrix(xTarget).size &&
                    matrix(xTarget)(yTarget + 1).toIntOption.isDefined
                  ) {
                    matrix(xTarget)(yTarget + 1) ++ getRightSymbol(
                      xTarget,
                      yTarget + 1
                    )
                  } else ""
                }

                val num = (getLeftSymbol(x + ref._1, y + ref._2) ++
                  matrix(x + ref._1)(y + ref._2)
                  ++ getRightSymbol(x + ref._1, y + ref._2)).toInt
                num
              }
              sum = sum + eligibleNumbers.product
            }
          case _ => ""
        }
      }
    }

    sum
  }

}
