package main

import (
	"fmt"

	"app/controller"
	"app/models"
)

func main() {

	// Get User By Page
	var usersNumber = 11
	controller.GenerateUser(usersNumber)
	// overallPages := 0
	// if usersNumber/10 != 0 {
	// 	overallPages = usersNumber/10 + 1
	// }

	// fmt.Println("Pagelar soni:", overallPages)
	// fmt.Println("Page raqamini kiriting:")
	// pageNumber := 0
	// fmt.Scan(&pageNumber)

	// pageOffset := pageNumber*10 - 10
	// // limit := 10
	// if pageNumber == overallPages {
	// 	limit = usersNumber % 10
	// }

	users, err := controller.GetListUser(models.GetListRequest{
		Offset: 1,
		Limit:  2,
		// Search: "xumoyun",
		FromDate: "2002-06-11",
		ToDate:   "2010-07-09",
	})

	fmt.Println(users)

	if err != nil {
		fmt.Println(err, users)
	}

	for _, user := range users {
		fmt.Println(user)
	}

}
