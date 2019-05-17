package main

import (
	"fmt"
	"sort"
	"strings"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u user) String() string {
	return fmt.Sprintf("%v %v is age %v\n%v", u.First, u.Last, u.Age, u.formatedSayings())
}

func (u user) formatedSayings() string {
	var sb strings.Builder
	for _, v := range u.Sayings {
		sb.WriteString(fmt.Sprintf("%v%v\n", strings.Repeat(" ", 8), v))
	}
	return sb.String()
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"331. Shaken, not stirred",
			"001. Youth is no guarantee of innovation",
			"000. In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	arr := []user{u1, u2, u3}

	//covert to type users which has method Println
	users(arr).SortSayings().SortByAge().Println()
}

type users []user

func (u users) Println() *users {
	for _, v := range u {
		fmt.Println(v)
	}

	return &u
}

func (u users) SortSayings() *users {
	for _, v := range u {
		sort.Strings(v.Sayings)
	}

	return &u
}

func (u users) SortByAge() *users {
	sort.Sort(byAge(u))
	return &u
}

func (u users) SortByLastName() users {
	sort.Sort(byLastName(u))
	return u
}

type byAge users

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byLastName users

func (a byLastName) Len() int           { return len(a) }
func (a byLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLastName) Less(i, j int) bool { return a[i].Last < a[j].Last }
