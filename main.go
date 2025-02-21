package main

import (
	"fmt"
	"io/ioutil"
	"log"

	addressbookpb "github.com/maaaashin324/protobuf-example-go/src/addressbook"
	complexpb "github.com/maaaashin324/protobuf-example-go/src/complex"
	enumpb "github.com/maaaashin324/protobuf-example-go/src/enum_example"
	simplepb "github.com/maaaashin324/protobuf-example-go/src/simple"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()

	readAndWriteDemo(sm)
	jsonDemo(sm)

	doEnum()

	doComplex()

	addressBook := addPerson()

	getPerson(addressBook)
}

func addPerson() *addressbookpb.AddressBook {
	person := &addressbookpb.Person{
		Name:  "Masataka Shintoku",
		Id:    1,
		Email: "masataka.shintoku@gmail.com",
		Phones: []*addressbookpb.Person_PhoneNumber{
			{
				Number: "0120123456",
				Type:   addressbookpb.Person_MOBILE,
			},
			{
				Number: "0120123456",
				Type:   addressbookpb.Person_HOME,
			},
		},
	}

	addressbook := &addressbookpb.AddressBook{}
	addressbook.People = append(addressbook.People, person)

	fmt.Println(addressbook)

	return addressbook
}

func getPerson(addressBook *addressbookpb.AddressBook) {
	persons := addressBook.GetPeople()

	for i, p := range persons {
		fmt.Printf("No. %v person: %v\n", i+1, p)
	}
}

func doComplex() {
	cm := &complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{
				Id:   2,
				Name: "Second message",
			},
			{
				Id:   1,
				Name: "Third message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	em := &enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY

	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)
}

func toJSON(pb proto.Message) string {
	out, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	}
	return string(out)
}

func fromJSON(in string, pb proto.Message) error {
	if err := protojson.Unmarshal([]byte(in), pb); err != nil {
		return err
	}
	return nil
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Couldn;t put the bytes into thbe protocol buffers strcut", err)
		return err
	}

	return nil
}

func doSimple() proto.Message {
	sm := &simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"
	fmt.Println(sm)

	fmt.Println("The ID is:", sm.GetId())

	return sm
}
