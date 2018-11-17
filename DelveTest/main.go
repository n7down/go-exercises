package main

type Person struct {
	Name string
	Age  int
}

func main() {
	var me Person
	me.Name = "Nelvin"
	me.Age = 41

	var wife Person
	wife.Name = "Raye"
	wife.Age = 36

	var daughter Person
	daughter.Name = "Katherine"
	daughter.Age = 5

	var son Person
	son.Name = "Vito"
	son.Age = 2

	var family []Person

	family = append(family, me, wife, daughter, son)

	birthdayToday(&daughter)
}

func birthdayToday(person *Person) {
	person.Age = person.Age + 1
}
