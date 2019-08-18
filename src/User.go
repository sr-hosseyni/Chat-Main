package main

type User struct {
	Name string `json:"name"`
	Id int		`json:"id"`
	Password string
}

type IUser struct{
	Name string `json:"name"`
	Id int `json:"id"`
}

type UserList struct {
	UserList []IUser `json:"userList"`
}

var users = map[int]User{
	1: {
		Name:     "rassoul",
		Id:       1,
		Password: "manche3ter",
	},
	2: {
		Name:     "ali",
		Id:       2,
		Password: "ali",
	},
	3: {
		Name:     "roya",
		Id:       3,
		Password: "roya",
	},
	4: {
		Name:     "mohammad",
		Id:       4,
		Password: "mohammad",
	},
	5: {
		Name:     "misagh",
		Id:       5,
		Password: "misagh",
	},
	6: {
		Name:     "jon",
		Id:       6,
		Password: "jon",
	},
}

var usersNameIndex = map[string]int{
	"rassoul": 1,
	"ali": 2,
	"roya": 3,
	"mohammad": 4,
	"misagh": 5,
	"jon": 6,
}

func (user User) checkPassword(password string) bool {
	return user.Password == password
}

func (user User) transform() IUser {
	return IUser{
		Id: user.Id,
		Name: user.Name,
	}
}

