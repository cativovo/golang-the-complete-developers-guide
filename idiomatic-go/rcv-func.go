//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Health float64
type Energy float64

type Player struct {
	name              string
	health, maxHealth Health
	energy, maxEnergy Energy
}

func (player *Player) updatePlayerHealth(health Health) {
	newHealth := player.health + health

	if newHealth < 0 {
		newHealth = 0
	} else if newHealth > player.maxHealth {
		newHealth = player.maxHealth
	}

	player.health = newHealth
}

func (player *Player) updatePlayerEnergy(energy Energy) {
	newEnergy := player.energy + energy

	if newEnergy < 0 {
		newEnergy = 0
	} else if newEnergy > player.maxEnergy {
		newEnergy = player.maxEnergy
	}

	player.energy = newEnergy
}

func (player *Player) applyDamage(damage Health) {
	player.updatePlayerHealth(-(damage))

	fmt.Println(player.name, "receives", damage, "damage health -> ", player.health)
}

func (player *Player) heal(health Health) {
	player.updatePlayerHealth(health)

	fmt.Println(player.name, "heals", health, "health -> ", player.health)
}

func (player *Player) consumeEnergy(energy Energy) {
	player.updatePlayerEnergy(-(energy))

	fmt.Println(player.name, "consumes", energy, "energy -> ", player.energy)
}

func (player *Player) restoreEnergy(energy Energy) {
	player.updatePlayerEnergy(energy)

	fmt.Println(player.name, "restores", energy, "energy -> ", player.energy)
}

func main() {
	player := Player{
		name:      "John doe",
		health:    100.0,
		maxHealth: 100.0,
		energy:    100.0,
		maxEnergy: 100.0,
	}

	player.applyDamage(99)
	player.heal(10)
	player.consumeEnergy(999)
	player.restoreEnergy(100)
}
