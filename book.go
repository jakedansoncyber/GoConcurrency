package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished,
	)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "Harry Potter",
		Author:        "JK Rowling",
		YearPublished: 1980,
	},
	Book{
		ID:            2,
		Title:         "King Athur",
		Author:        "Ham Hames",
		YearPublished: 1940,
	},
	Book{
		ID:            3,
		Title:         "Sloppy Shlop",
		Author:        "Shooter Mcgavin",
		YearPublished: 1950,
	},
	Book{
		ID:            4,
		Title:         "Drive Home",
		Author:        "Randy Moolton",
		YearPublished: 1960,
	},
	Book{
		ID:            5,
		Title:         "Cops",
		Author:        "Officer Toodles",
		YearPublished: 1990,
	},
	Book{
		ID:            6,
		Title:         "Screen Dismount",
		Author:        "Hiya Byeya",
		YearPublished: 2022,
	},
	Book{
		ID:            7,
		Title:         "Potter Lofter",
		Author:        "Rowling Money",
		YearPublished: 1924,
	},
	Book{
		ID:            8,
		Title:         "Running Low",
		Author:        "Oil Gasoline",
		YearPublished: 1830,
	},
	Book{
		ID:            9,
		Title:         "Second Coming",
		Author:        "Lord Christ",
		YearPublished: 1709,
	},
	Book{
		ID:            10,
		Title:         "Bunion Onion",
		Author:        "Strawberry Pie",
		YearPublished: 1410,
	},
}
