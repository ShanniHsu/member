package gun

import "testing"

// 客戶端只會看到使用代碼後，拿到相關產品的詳細資料
func TestAK47(t *testing.T) {
	ak47, _ := GetGun("ak47")

	name := ak47.GetName()
	if name != "AK47 gun" {
		t.Fatal("AK47 name test error")
	}

	power := ak47.GetPower()
	if power != 4 {
		t.Fatal("AK47 power test error")
	}
}

func TestMusket(t *testing.T) {
	musket, _ := GetGun("musket")

	name := musket.GetName()
	if name != "Musket gun" {
		t.Fatal("Musket name test error")
	}

	power := musket.GetPower()
	if power != 1 {
		t.Fatal("Musket power test error")
	}
}
