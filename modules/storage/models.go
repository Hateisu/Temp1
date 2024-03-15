package storage

// Saved System File

type SavedConfig struct {
	ID   uint `gorm:"primarykey"`
	Data string
}

// Auth System

type Role struct {
	RoleID   uint `gorm:"primarykey"`
	RoleName string
}

type User struct {
	UserID      uint `gorm:"primarykey"`
	Firstname   string
	Lastname    string
	PhoneNumber string
	Email       string
	RoleId      int  `json:"-"`
	Role        Role `gorm:"constraint:OnUpdate:CASCADE,foreignKey:RoleId,OnDelete:SET NULL;"`

	Username string
	Password string `json:"-"`
}

// Write here all models for database
/*type Seller struct {
	SellerID  uint `gorm:"primarykey"`
	Firstname string
	Lastname  string
}

type Product struct {
	//Id          int
	ID          uint   `gorm:"primarykey"`
	Seller      Seller `gorm:"constraint:OnUpdate:CASCADE,foreignKey:SellerID,OnDelete:SET NULL;"`
	Name        string
	Description string
	Price       int
}*/

type Country struct {
	Name    string
	Lang    string
	Curency string
}
