package main

import(
	"encoding/csv"
	"github.com/gocolly/colly"
	"os"
	"log"
	"strings"
)

//Scrap all airports around the world
func main(){
	fName := "airports.csv"
	static_domain := "http://www.nationsonline.org/oneworld/IATA_Codes/"

	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	//Write CSV header for airport information
	writer.Write([]string{"iata_code","city","airport","country"})
	//Instantiate the colly scrapper
	c := colly.NewCollector()

	for i :=1; i<27 ; i++ {
		scrapWebpage(static_domain+"IATA_Code_"+string(toChar(i))+".htm",writer,c,string(toChar(i)))
	}
}
/*
This function is to scrap one instance of a webpage
*/
func scrapWebpage(url string, writer *csv.Writer, c *colly.Collector,alphabet string) {
	
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		var airportInfo [4]string
		checkValid := true
		e.ForEach("td.border1", func(count int, el *colly.HTMLElement) {
			if count == 0{
				checkValid = true
			}

			if count == 0 && len(el.Text) != 3 {
				checkValid = false
			}
			
			if checkValid && len(el.Text) > 0 && !strings.Contains(el.Text,"Code") && !strings.Contains(el.Text,"City") && !strings.Contains(el.Text,"Airport") && !strings.Contains(el.Text,"Country"){
				airportInfo[count] = el.Text
			}
			
			//Only adds if the line of information is valid
			if count == 3 && checkValid{
				writer.Write([]string{airportInfo[0],airportInfo[1],airportInfo[2],strings.Replace(airportInfo[3],",",";",-1)})
			}
		})
		
		
	})

	c.Visit(url)
}
//Gets A to Z
func toChar(i int) rune {
    return rune('A' - 1 + i)
}