package main

//
//import "fmt"
//
//type User struct {
//	ID        int
//	FirstName string
//	LastName  string
//	email     string
//}
//
//func (user User) displayProfile() {
//	fmt.Printf("Personal Info of Account ID %d:\nName: %s %s\nEmail: %s\n\n", user.ID, user.FirstName, user.LastName, user.email)
//}
//
//type Party struct {
//	ID          int
//	Name        string
//	Leader      User
//	Members     []User
//	IsAvailable bool
//}
//
//func (party Party) displayPartyProfile() {
//	fmt.Printf("Info of Group ID %d:\nName: %s\nLeader: %s\nMember: %d\nIs Available: %t", party.ID, party.Name, party.Leader.FirstName, len(party.Members), party.IsAvailable)
//}
//
//func main() {
//	user1 := User{1, "Ivan", "Budianto", "ivanb0603@gmail.com"}
//	user2 := User{2, "Jessica", "Maya", "jessicamaya1702@gmail.com"}
//	user3 := User{3, "Raras", "Dwistian", "rarasdw1901@gmail.com"}
//
//	party1_member := []User{user1, user2, user3}
//	party1 := Party{1, "Home", user1, party1_member, true}
//
//	user1.displayProfile()
//	party1.displayPartyProfile()
//}
