package main

import (
	"fmt"
)

const (
	Ana    = "Ana"
	Baldur = "Baldur"
)

type character struct {
	name     string
	position int32
	health   int32
	power    int32
	isDead   bool
}
type hero struct {
	character
	deffend int32
}

type enemy struct {
	character
	poison int32
}

func (h hero) HeroAttack(e *enemy) {
	e.character.health = e.character.health - h.character.power
}

func (e enemy) EnemyAttack(h *hero) {
	h.character.health = h.character.health - ((e.character.power * e.poison) - h.deffend)
}

func (c *character) Move(step int32) int32 {
	c.position += step
	return c.position
}

func (c *character) Refresh() {
	if c.health <= 0 {
		c.isDead = true
		*c = character{}
		fmt.Println(c)
		fmt.Println("Character is dead")
	}
}

func main() {
	var hero1 = hero{
		character: character{
			name:     Ana,
			position: 0,
			health:   100,
			power:    30,
			isDead:   false,
		},
		deffend: 20,
	}

	var enemy1 = enemy{
		character: character{
			name:     Baldur,
			position: 10,
			health:   100,
			power:    20,
			isDead:   false,
		},
		poison: 5,
	}

	hero1.Move(5)
	enemy1.Move(10)

	hero1.HeroAttack(&enemy1)
	enemy1.EnemyAttack(&hero1)
	enemy1.EnemyAttack(&hero1)

	enemy1.Refresh()
	hero1.Refresh()

	hero1.Move(5)

	fmt.Println(hero1)
	fmt.Println(enemy1)
}
