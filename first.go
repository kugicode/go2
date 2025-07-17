package main

import "fmt"

func checkPassword(password string){
    fmt.Println("Checking password...", password)
    if password == "anas123"{
        fmt.Println("Correct password, You logged in!")
    } else{
        fmt.Println("Incorrect password, try again.")
    }
}
//Blueprint for student.
type Student struct {
    Name string
    Age int
}

func main(){
    fmt.Println("...")

    checkPassword("123")
    checkPassword("anas123")

    //Example of Slices
    Students := []string {"Anas", "Uknown", "David"}

    //Printing out the example
    fmt.Println("Students:", Students)

    //len of example
    fmt.Println("Length:", len(Students))

    //example of maps
    ages := map[string]int{
        "anas": 22,
        "unknown": 21, 
    }
    //Print out the map
    fmt.Println("map:", ages)
    
    //Access single value
    fmt.Println("Anas's age is:", ages["anas"])

    //Add or update an item
    ages ["unknown"] = 25
    ages ["david"] = 35
    fmt.Println("Updated vales:", ages)

    //Delete an item

    delete(ages, "unknown")
    fmt.Println("After deletion:", ages)

    age, ok := ages["unknown"]

    if ok{
        fmt.Println("unknown's age is", age)
    } else{
        fmt.Println("unknown isn't in the map!")
    }
    
    //Instance of student
    s1 := Student{
        Name: "anas",
        Age: 22,
    }

    fmt.Println("Name:", s1.Name)
    fmt.Println("Age:", s1.Age)

    s2 := Student{
        Name: "unknown",
        Age: 25,
    }

    fmt.Println("Name:", s2.Name)
}