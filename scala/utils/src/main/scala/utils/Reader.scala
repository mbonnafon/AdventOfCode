package utils

import scala.io.Source

case class Reader(filename: String) {
  def read(): List[String] = Source.fromResource(filename).getLines().toList
}
