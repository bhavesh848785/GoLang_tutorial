package main

import "fmt"

//single struct.......................................................
type person struct {
	fisrt string
	last  string
	age   int
}


//multiple struct.......................................................

type personal struct{
	fname string
	lname string
	p_age int
}

type contact struct{
	email string
	phone int
}

type Address struct{
	house int
	area string
	state string
}

type employee struct{
	personal_detail personal
	personal_contact contact
	personal_address Address
}


//main function...............................................
func main() {

	//first method to use struct..................................................

	var bhavesh person
	fmt.Println("bhavesh person:", bhavesh) // " " " " 0 // __0
	bhavesh.fisrt = "bhavesh"
	bhavesh.last = "vaghela"
	bhavesh.age = 20
    fmt.Println("bhavesh person:", bhavesh)

	//second method to use struct.............................................................

	person1 := person{
		fisrt: "bhavesh2",
		last: "vaghela2",
		age: 20,
	}
	fmt.Println("person1:", person1)

	//third method to use struct.............................................................
	//in this output: person2: &{bhavesh3 vaghela3 20}
	//& beacuse we use here new keyword and this concept of pointer and & is also called a pointer

	var person2 = new(person)
	person2.fisrt = "bhavesh3"
	person2.last = "vaghela3"
	person2.age = 20

	fmt.Println("person2:", person2)



	//multiple struct demo...........................................................

	employee1 := employee{
		personal_detail : personal{
			fname: "bhavesh",
			lname: "vaghela",
			p_age: 20,
		},

		personal_contact : contact{
			email: "b@gmail.com",
			phone: 9898989898,
		},

		personal_address : Address{
			house: 555,
			area: "shantigram",
			state: "gujarat",
		},
	}
	fmt.Printf("%+v\n", employee1)
	fmt.Printf("%+v\n", employee1.personal_address.house) //specific
}

