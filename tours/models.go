package tours

import "github.com/jinzhu/gorm"

type Activity struct {
	gorm.Model
	Name string
	Description string
}

type Country struct {
	gorm.Model
	Name string
}

type Tour struct {
	gorm.Model
	Title string
	Description string
	Features []Feature
	Activities []Activity
}


type Feature struct {
	gorm.Model
	Name string
	Description string
}


type Lodge struct {
	gorm.Model
	Name string
	Description string
	Website string
}

type Accommodation struct {
	gorm.Model
	Lodge *Lodge
	Meals string
}