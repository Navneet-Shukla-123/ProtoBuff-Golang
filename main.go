package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	name := &Person{
		Name: "Navneet",
		Age:  24,
		SocialFollowers: &SocialFollowers{
			Youtube: 2500,
			Twitter: 1400,
		},
	}

	data, err := proto.Marshal(name)
	if err != nil {
		log.Println("Error in marshalling to proto")
		return
	}

	fmt.Println("Data is ", data)

	newPerson:=&Person{}

	err=proto.Unmarshal(data,newPerson)

	if err!=nil{
		log.Println("Error in unmarshalling from proto : ",err)
		return
	}

	fmt.Println("Person is ",newPerson)

	fmt.Println("Age is ",newPerson.GetAge())
	fmt.Println("Name is ",newPerson.GetName())
	fmt.Println("Twitter followers are ",newPerson.SocialFollowers.GetTwitter())
	fmt.Println("Youtube followers are ",newPerson.SocialFollowers.GetYoutube())

}
