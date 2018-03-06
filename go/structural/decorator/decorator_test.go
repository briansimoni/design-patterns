// When you think about extending legacy code without the risk of breaking
// something, you should think of the Decorator pattern first.

// When to use the decorator pattern:
// - When you need to add functionality to some code that you don't have access to,
//   or you don't want to modify to avoid a negative effect on the code, and follow
//   the open/closed principle (like legacy code)
// - When you want the functionality of an object to be created or altered dynamically,
//   and the number of features is unknown and could grow fast.

package decorator

import (
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("Expected: '%s', Got: '%s'", expectedText, pizzaResult)
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	onion := &Onion{}
	onionResult, err := onion.AddIngredient()
	if err == nil {
		t.Errorf("When the ingredient field of onion is nil, an error should be returned. Got %s", onionResult)
	}

	onion = &Onion{&PizzaDecorator{}}
	onionResult, err = onion.AddIngredient()

	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(onionResult, "onion") {
		t.Errorf("Expected onion. Got %s", onionResult)
	}
}

func TestMeat_AddIngredient(t *testing.T) {
	meat := &Meat{}
	meatResult, err := meat.AddIngredient()
	if err == nil {
		t.Errorf("When the ingredient field of meat is nil, an error should be returned. Got %s", meatResult)
	}

	meat = &Meat{&PizzaDecorator{}}
	meatResult, err = meat.AddIngredient()

	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(meatResult, "meat") {
		t.Errorf("Expected meat. Got %s", meatResult)
	}
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &Onion{&Meat{&PizzaDecorator{}}}
	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	expectedText := "Pizza with the following ingredients: meat, onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("Expected: '%s' but got '%s'", expectedText, pizzaResult)
	}
	t.Log(pizzaResult)
}
