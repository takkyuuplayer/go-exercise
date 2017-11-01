package test

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"
)

func TestXMLUnmarshal(t *testing.T) {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}
	v := Result{Name: "none", Phone: "none"}

	data := `
	<Person>
		<FullName>Alice Takkyuu</FullName>
		<Company>Example Inc.</Company>
		<Email where="home">
			<Addr>alice@example.com</Addr>
		</Email>
		<Email where='work'>
			<Addr>alice@work.com</Addr>
		</Email>
		<Group>
			<Value>Friends</Value>
			<Value>Squash</Value>
		</Group>
		<City>Hanga Roa</City>
		<State>Easter Island</State>
	</Person>
	`

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	want := Result{
		XMLName: xml.Name{
			Space: "",
			Local: "Person",
		},
		Name:  "Alice Takkyuu",
		Phone: "none",
		Email: []Email{
			Email{
				Where: "home",
				Addr:  "alice@example.com",
			},
			Email{
				Where: "work",
				Addr:  "alice@work.com",
			},
		},
		Groups: []string{"Friends", "Squash"},
		Address: Address{
			City:  "Hanga Roa",
			State: "Easter Island",
		},
	}

	if !reflect.DeepEqual(v.XMLName, want.XMLName) {
		t.Errorf(`v.XMLName = %#v, want %#v`, v.XMLName, want.XMLName)
	}
	if !reflect.DeepEqual(v.Name, want.Name) {
		t.Errorf(`v.Name = %#v, want %#v`, v.Name, want.Name)
	}
	if !reflect.DeepEqual(v.Phone, "none") {
		t.Errorf(`v.Phone = %#v, want %#v`, v.Phone, "none")
	}
	if !reflect.DeepEqual(v.Email, want.Email) {
		t.Errorf(`v.Email = %#v, want %#v`, v.Email, want.Email)
	}
	if !reflect.DeepEqual(v.Groups, want.Groups) {
		t.Errorf(`v.Groups = %#v, want %#v`, v.Groups, want.Groups)
	}
	if !reflect.DeepEqual(v.Address, want.Address) {
		t.Errorf(`v.Address = %#v, want %#v`, v.Address, want.Address)
	}

	if !reflect.DeepEqual(v, want) {
		t.Errorf(`v = %#v, want %#v`, v, want)
	}

}
