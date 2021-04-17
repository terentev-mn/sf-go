//Создайте словарь (небольшую базу данных), где ключом является имя человека, а значением — экземпляр структуры Man.
// Добавьте в словарь от 5 до 10 элементов.
//Пример словаря: var people map[string]Man.
//Далее создайте слайс — suspects, в котором перечислены имена подозреваемых.
//Допишите программу так, чтобы на экран выводился подозреваемый,
// у которого наибольшее количество совершённых преступлений.
// Если ни одного подозреваемого нет в базе данных, нужно вывести на экран сообщение:
// «В базе данных нет информации по запрошенным подозреваемым».

package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {

	people := map[string]Man{
		"John":  {"John", "Wuick", 33, "male", 100},
		"Vasya": {"Vasya", "Dudkin", 22, "male", 10},
		"Tia":   {"Tia", "Karerra", 25, "female", 13},
		"Johns": {"Johns", "Indiana", 44, "male", 0},
	}

	var suspects []string
	var maxCrime int
	var maxCrimeName string

	for k, v := range people {
		//fmt.Println(name.Name, name.Crimes)
		//fmt.Println(k, v)
		if v.Crimes > 0 {
			suspects = append(suspects, k)
		}
		if v.Crimes > maxCrime {
			maxCrime = v.Crimes
			maxCrimeName = k
		}

	}

	if maxCrime > 0 {
		fmt.Println("Все подозреваемые:")
		for _, s := range suspects {
			fmt.Println(people[s])
		}
		fmt.Println("Самый злодей:", people[maxCrimeName])
	} else {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	}
}
