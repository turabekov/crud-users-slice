package controller

import (
	"app/models"
	"fmt"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
)

var Users []models.User

func CreateUser(data models.User) {
	Users = append(Users, data)
}

// getbyid
func GetUserById(id int) (models.User, bool) {
	for _, val := range Users {
		if val.Id == id {
			return val, false
		}
	}

	return models.User{}, true
}

// update
func UpdateUser(id int) (models.User, bool) {
	for i, val := range Users {
		if val.Id == id {
			Users[i].Name = "changed Name"
			Users[i].Surname = "changed Surname"
			return Users[i], false
		}
	}

	return models.User{}, true
}

// delete
func DeleteUser(id int) bool {
	for i, val := range Users {
		if val.Id == id {
			Users = append(Users[:i], Users[i+1:]...)
			return false
		}
	}

	return true
}

// Search and filtiring by Date
func GetListUser(req models.GetListRequest) (resp []models.User, err error) {

	// if searched not given or empty
	if len(req.Search) <= 0 {

		// if dates are given
		if len(req.FromDate) > 0 && len(req.ToDate) > 0 {
			users, _ := filteringByDate(Users, req.FromDate, req.ToDate)
			if req.Limit+req.Offset > len(users) {
				if req.Offset > len(users) {
					return []models.User{}, nil
				}
				return users[req.Offset:], nil
			}

			return Users[req.Offset : req.Limit+req.Offset], nil
		}
		if req.Limit+req.Offset > len(Users) {
			if req.Offset > len(Users) {
				return []models.User{}, nil
			}
			return Users[req.Offset:], nil
		}

		return Users[req.Offset : req.Limit+req.Offset], nil

	}

	// Search
	res := []models.User{}
	for _, v := range Users {
		fullName := strings.ToLower(v.Name + v.Surname)
		if strings.Contains(fullName, strings.ToLower(strings.Replace(req.Search, " ", "", -1))) {
			res = append(res, v)
		}
	}

	// if user not found
	if len(res) <= 0 {
		return res, nil
	}
	// fmt.Println("Result", res[req.Offset:req.Limit+req.Offset])

	// filtering By Date
	if len(req.FromDate) > 0 && len(req.ToDate) > 0 {
		users, err := filteringByDate(res, req.FromDate, req.ToDate)
		fmt.Println("Filtered", users)
		if req.Limit+req.Offset > len(users) {
			if req.Offset > len(res) {
				return []models.User{}, nil
			}
			return users[req.Offset:], nil
		}

		return users[req.Offset : req.Limit+req.Offset], err
	} else {
		// if offset and limit out of range
		if req.Limit+req.Offset > len(res) {
			if req.Offset > len(res) {
				return []models.User{}, nil
			}
			return res[req.Offset:], nil
		}
		return res[req.Offset : req.Limit+req.Offset], nil
	}
}

// filter Function
func filteringByDate(arr []models.User, fromD, toD string) ([]models.User, error) {
	fromDate, err1 := time.Parse("2006-01-02", fromD)
	toDate, err2 := time.Parse("2006-01-02", toD)

	if err1 != nil {
		return arr, err1
	}
	if err2 != nil {
		return arr, err2
	}

	filterRes := []models.User{}

	for _, v := range arr {
		userDate, _ := time.Parse("2006-01-02", v.Birthday)

		if (userDate.Before(toDate) || userDate == toDate) && (userDate.After(fromDate) || userDate == fromDate) {
			filterRes = append(filterRes, v)
		}
	}

	return filterRes, nil
}

func GenerateUser(count int) {
	Users = append(Users, models.User{
		Id:       1,
		Name:     "Xumoyun",
		Surname:  "Turabekov",
		Birthday: "2002-06-12",
	})
	for i := 1; i < count; i++ {
		Users = append(Users, models.User{
			Id:       i + 1,
			Name:     faker.FirstName(),
			Surname:  faker.LastName(),
			Birthday: faker.Date(),
		})
	}
}

// 1.
// GetList(
// offset,
// limit,
// search: "Julian Dooley"
// "Dooley Julian"
// "dooley julian"
// "do"
//   --"Dooley Julian"
//   --"Don Jack"
// )

// 2.
// User
//   birthay: faker.Date()

// GetListRequest
//   FromDate: 2000-10-01
//   ToDate: 2023-10-12
