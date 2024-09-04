package gun

import "fmt"

// method
type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() (name string)
	GetPower() (power int)
}

// Product
// 槍的名字跟火力
// 這邊是讓客戶端輸入資料後拿到相對應的回應
type Gun struct {
	Name  string
	Power int
}

func (g *Gun) SetName(name string) {
	if name != "" {
		g.Name = name
	}
}

func (g *Gun) SetPower(power int) {
	if power != 0 {
		g.Power = power
	}
}

func (g *Gun) GetName() (name string) {
	return g.Name
}

func (g *Gun) GetPower() (power int) {
	return g.Power
}

// 有一把AK47的產品
type ak47 struct {
	Gun
}

func newAk47() IGun {
	return &ak47{
		Gun: Gun{
			Name:  "AK47 gun",
			Power: 4,
		},
	}
}

// 有一把Musket的產品
type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			Name:  "Musket gun",
			Power: 1,
		},
	}
}

// Factory
// 裡面回傳有哪些產品(但在這不會看到詳細產品的細節，已經被包裝在裡面)
func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}

	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
