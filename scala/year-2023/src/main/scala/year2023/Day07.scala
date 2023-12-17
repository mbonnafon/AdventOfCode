package year2023

import utils.Puzzle

import scala.collection.immutable.{AbstractMap, SeqMap, SortedMap}

object Day07 extends Puzzle {

  val Ranker = Map(
    "High card" -> 0,
    "One pair" -> 1,
    "Two pair" -> 2,
    "Three of a kind" -> 3,
    "Full house" -> 4,
    "Four of a kind" -> 5,
    "Five of a kind" -> 6
  )
  case class Player(
      hand: String,
      bid: Int,
      private val cardScores: Map[String, Int],
      private val handTypeAttribution: Map[Char, Int] => String
  ) {
    private def handScore() = {
      hand.toList
        .groupBy(identity)
        .mapValues(_.size)
        .toMap
    }

    val handScoreByCard: List[Int] = {
      hand
        .split("")
        .map(cardScores.get(_).get)
        .toList
    }

    val handType: String = handTypeAttribution(handScore())
  }
  object Player {
    def apply(
        string: String,
        cardScores: Map[String, Int],
        handTypeAttribution: Map[Char, Int] => String
    ) = {
      new Player(
        hand = string
          .split(" ")
          .head,
        bid = string
          .split(" ")
          .last
          .toInt,
        cardScores = cardScores,
        handTypeAttribution = handTypeAttribution
      )
    }
  }

  def part1(input: List[String]): Long = {
    val CardScores: Map[String, Int] = Map(
      "A" -> 14,
      "K" -> 13,
      "Q" -> 12,
      "J" -> 11,
      "T" -> 10,
      "9" -> 9,
      "8" -> 8,
      "7" -> 7,
      "6" -> 6,
      "5" -> 5,
      "4" -> 4,
      "3" -> 3,
      "2" -> 2
    )
    val handTypeAttribution = { handScore: Map[Char, Int] =>
      handScore match {
        case map if map.exists(_._2 == 5) => "Five of a kind"
        case map if map.exists(_._2 == 4) => "Four of a kind"
        case map if map.exists(_._2 == 3) && map.exists(_._2 == 2) =>
          "Full house"
        case map if map.exists(_._2 == 3) && map.exists(_._2 != 2) =>
          "Three of a kind"
        case map if map.count(_._2 == 2) == 2 => "Two pair"
        case map if map.count(_._2 == 2) == 1 => "One pair"
        case map                              => "High card"
      }
    }
    input
      .map(Player(_, CardScores, handTypeAttribution))
      .groupBy(_.handType)
      .toList
      .sortBy(rankType =>
        Ranker.get(rankType._1).getOrElse(0)
      ) // sort by type of hand
      .map { handType =>
        handType._2.sortWith { (listA, listB) =>
          listA.handScoreByCard.zip(listB.handScoreByCard).foldLeft(0) {
            case (result, (a, b)) if result == 0 && a != b =>
              if (a < b) -1 else 1
            case (result, _) => result
          } < 0
        }
      }
      .flatten
      .zip(Stream.from(1)) // add index, e.g (Player(32T3K,765),1)
      .map(g => g._2 * g._1.bid)
      .sum
  }

  def part2(input: List[String]): Long = {
    val CardScores: Map[String, Int] = Map(
      "A" -> 13,
      "K" -> 12,
      "Q" -> 11,
      "T" -> 10,
      "9" -> 9,
      "8" -> 8,
      "7" -> 7,
      "6" -> 6,
      "5" -> 5,
      "4" -> 4,
      "3" -> 3,
      "2" -> 2,
      "J" -> 1
    )
    val handTypeAttribution = { handScore: Map[Char, Int] =>
      val normalizedHandByJ = handScore match {
        case map if map.exists(c => c._1 == 'J' && c._2 == 5) => Map('A' -> 5)
        case map
            if map.contains('J') && map.exists(c => c._1 != 'J' && c._2 > 1) =>
          val maxValue = map.removed('J').values.max
          val maxKey = map
            .filter((v) => v._2 == maxValue) // keep only max values
            .removed('J') // remove J
            .keys // collect keys Set(A, Q)
            .toList
            .map(_.toString)
            .sortBy(CardScores.get(_))
            .reverse // desc order
            .head
          map
            .updated(
              maxKey.toCharArray.head,
              map.get(maxKey.toCharArray.head).get + map.get('J').get
            )
            .removed('J')
        case map if map.contains('J') =>
          val maxKey = map
            .removed('J') // remove J
            .keys // collect keys Set(A, Q)
            .toList
            .map(_.toString)
            .sortBy(CardScores.get(_))
            .reverse // desc order
            .head
          map
            .updated(
              maxKey.toCharArray.head,
              map.get(maxKey.toCharArray.head).get + map.get('J').get
            )
            .removed('J')
        case _ => handScore
      }
      normalizedHandByJ match {
        case map if map.exists(_._2 == 5) => "Five of a kind"
        case map if map.exists(_._2 == 4) => "Four of a kind"
        case map if map.exists(_._2 == 3) && map.exists(_._2 == 2) =>
          "Full house"
        case map if map.exists(_._2 == 3) && map.exists(_._2 != 2) =>
          "Three of a kind"
        case map if map.count(_._2 == 2) == 2 => "Two pair"
        case map if map.count(_._2 == 2) == 1 => "One pair"
        case map                              => "High card"
      }
    }
    input
      .map(Player(_, CardScores, handTypeAttribution))
      .groupBy(_.handType)
      .toList
      .sortBy(rankType =>
        Ranker.get(rankType._1).getOrElse(0)
      ) // sort by type of hand
      .map { handType =>
        handType._2.sortWith { (listA, listB) =>
          listA.handScoreByCard.zip(listB.handScoreByCard).foldLeft(0) {
            case (result, (a, b)) if result == 0 && a != b =>
              if (a < b) -1 else 1
            case (result, _) => result
          } < 0
        }
      }
      .flatten
      .zip(Stream.from(1)) // add index, e.g (Player(32T3K,765),1)
      .map(g => g._2 * g._1.bid)
      .sum

  }

}
