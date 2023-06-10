package model

type Amenitie struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null;unique"`

	Hotels []*Hotel `gorm:"many2many:hotel_amenities;"`
}

type Amenities []Amenitie
