package main

import "fmt"

// Define an interface called Speaker
// Any type that has a method called 'Speak'
// which takes no arguments and returns a string
// satisfies this interface.
type Speaker interface {
	Speak() string
}

// --- A Type that satisfies Speaker ---
// Dog has a Speak() method
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

// --- Another Type that satisfies Speaker ---
// Cat has a Speak() method
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

// --- A function that takes the Interface type ---
// This function can accept *any* type that satisfies the Speaker interface
func MakeItSpeak(s Speaker) {
	// We don't know if 's' is a Dog or a Cat, but we know
	// it MUST have a Speak() method because it's a Speaker.
	fmt.Println(s.Speak())
}

func main() {
	// Create instances of our types
	myDog := Dog{Name: "Buddy"}
	myCat := Cat{Name: "Whiskers"}

	// Pass the Dog and Cat instances to the function
	// This works because both Dog and Cat satisfy the Speaker interface
	MakeItSpeak(myDog)
	MakeItSpeak(myCat)

	// You can also declare variables of interface type
	var s Speaker // 's' can now hold any type that implements Speaker

	s = myDog // Assign a Dog (it satisfies Speaker)
	fmt.Println(s.Speak())

	s = myCat // Assign a Cat (it satisfies Speaker)
	fmt.Println(s.Speak())
}
/*• Use Interfaces quando você deseja definir um contrato que será
seguido por diferentes classes que não têm relação direta entre si.
• Exemplo: Uma interface MotorCombustao pode ser implementada por várias
classes (Carro, Moto, Avião, Motoserra etc.), sem necessidade de herança
comum.
• Use Classes Abstratas quando você quer fornecer implementações
parciais e definir métodos concretos que serão compartilhados por
classes relacionadas.
• Exemplo: Uma classe abstrata Animal pode fornecer um método concreto
respirar() que será herdado por subclasses como Cachorro, Gato, Onça,
Macaco...*/

// ou seja interface seria um metodo abstrato que dá pra modificar
// mas tem coisas em comum