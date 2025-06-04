package admin

import (
	"app/db"
	"fmt"
)

func AddAdmin(name string) {
	_, err := db.DB.Exec("INSERT INTO admins (username) VALUES ($1)", name)
	if err != nil {
		fmt.Println("❌ Failed to add admin:", err)
		return
	}
	fmt.Println("✅ Admin added:", name)
}

func AddStaff(name string) {
	_, err := db.DB.Exec("INSERT INTO staffs (username) VALUES ($1)", name)
	if err != nil {
		fmt.Println("❌ Failed to add staff:", err)
		return
	}
	fmt.Println("✅ Staff added:", name)
}

func RemoveStaff(name string) {
	_, err := db.DB.Exec("DELETE FROM staffs WHERE username = $1", name)
	if err != nil {
		fmt.Println("❌ Failed to remove staff:", err)
		return
	}
	fmt.Println("✅ Staff removed:", name)
}

func AddMember(name string) {
	_, err := db.DB.Exec("INSERT INTO members (username) VALUES ($1)", name)
	if err != nil {
		fmt.Println("❌ Failed to add member:", err)
		return
	}
	fmt.Println("✅ Member added:", name)
}

func RemoveMember(name string) {
	_, err := db.DB.Exec("DELETE FROM members WHERE username = $1", name)
	if err != nil {
		fmt.Println("❌ Failed to remove member:", err)
		return
	}
	fmt.Println("✅ Member removed:", name)
}
