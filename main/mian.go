package main

import (
	"fmt"
	"headfirstgo/magazine"
)

func showSubInfo(s *magazine.Subsriber) {
	fmt.Println("Name:", s.Name)
	fmt.Printf("Rate: %0.2f\n", s.Rate)
	fmt.Println("Active?", s.Active)
}

func defSub(name string) *magazine.Subsriber {
	s := magazine.Subsriber{Name: name, Rate: 5.99, Active: true}
	return &s
}

func applyDiscount(s *magazine.Subsriber) {
	s.Rate = 4.99
}

func main() {
	s1 := defSub("gao")
	applyDiscount(s1)
	showSubInfo(s1)
	e1 := magazine.Employee{Name: "zhao", Salary: 18000, HomeAddr: magazine.Address{Street: "guli", City: "NanJing", State: "Jiangsu", PostalCode: "123456"}}
	fmt.Println(e1)
	s1.Address.City = "BeiJing"
	fmt.Println(s1.City)
}
