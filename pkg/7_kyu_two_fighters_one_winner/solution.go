package string_ends_with

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func solution(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	if fighter1.Name == firstAttacker {
		fighter1.Attack(&fighter2)

		if !fighter2.IsAlive() {
			return fighter1.Name
		}
	}

	return solution(fighter2, fighter1, fighter2.Name)
}

func (f *Fighter) Attack(aim *Fighter) {
	aim.Health -= f.DamagePerAttack
}

func (f *Fighter) IsAlive() bool {
	return f.Health > 0
}
