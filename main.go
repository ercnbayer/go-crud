package main

import (
	"fmt"
	"log"
	"os"
	//"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type Person struct {
	ID	 uint   `gorm:"primaryKey<-:create"`
	Name     string `gorm:"name"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
}

  
  // TableName overrides the table name used by User to `profiles`
  func (table *Person) TableName() string {
	return "persons";
  }

func InsertPerson( person *Person, db *gorm.DB) {

	db.Save(&person)

	fmt.Println(person,"inserted");

}

func DeletePerson(person *Person,db *gorm.DB){	
	
	db.Delete(&person); 
	
	fmt.Println(person,"deleted");
	
}

func UpdatePerson(person *Person,db *gorm.DB){

	db.Save(person);
	
	fmt.Println("updated:",person);

}

func ReadPerson(db *gorm.DB,person *Person)(){

	db.Find(person);
	 
	fmt.Println("found",person);
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("POSTGRES_DB")
	

	userName := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s database=%s port=%s sslmode=disable TimeZone=Etc/UTC", userName, dbPassword,dbName,dbPort)
	
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	

	var example_person Person=Person{ID:1,Name:"ornek",Password:"topsecret",Email:"secret@secret.com"};



	ReadPerson(db,&example_person);

	DeletePerson(&example_person,db);

	InsertPerson(&example_person,db);
	
	example_person.Name="new name"

	UpdatePerson(&example_person,db);



	// Example of inserting a person

}
