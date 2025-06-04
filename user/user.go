package user

import (
	"app/db"
	"fmt"
)

func RegisterMember(name string) {
	_, err := db.DB.Exec("INSERT INTO members (username) VALUES ($1)", name)
	if err != nil {
		fmt.Println("❌ Failed to register member:", err)
		return
	}
	fmt.Println("✅ Member registered:", name)
}

func BuyProduct(productName string) {
	fmt.Println("🛒 Member bought product:", productName)
	// يمكنك توسيع هذا لتسجيل عملية الشراء في جدول لاحقًا
}
