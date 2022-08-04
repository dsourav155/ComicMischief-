package main
import "fmt"


func main() {

	fmt.Println("This is the list of some best comic books in Comic Mischief:")

    var publisher, writer, artist, title string
    var year, pageNumber int
    var grade float32

//First book info.
    fmt.Println("First Comic Book: ")
    title = "Mr. GoToSleep"
    writer = "Tracey Hatchet"
    artist = "Jewel Tampson"
    publisher = "DizzyBooks Publishing Inc."
    year = 1997
    pageNumber = 14
    grade = 6.5

    fmt.Println(title, "written by one of the best comic writer", writer, "and drawn by talented", artist, "in", year, "is published by", publisher, "It is graded", grade, "out of 10 and it only have", pageNumber, "pages.")

//Second book info.
    fmt.Println("Second Comic Book:")
    title = "Epic Vol. 1"
    writer = "Ryan N. Shwan"
    artist= "Phoebe Paperclips"
    year = 2013
    pageNumber = 160
    grade = 9.0

fmt.Println(title, "written by one of the best comic writer", writer, "and drawn by very talented", artist, "in", year, "is published by", publisher, "It is graded", grade, "out of 10 and it only have", pageNumber, "pages.")

}
