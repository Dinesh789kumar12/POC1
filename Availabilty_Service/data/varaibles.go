package data

var LocationName = [][]string{
	{"New Delhi", "Mandoli", "Tughlaqabad", "Yammuna Vihar", "Bhajanpura"},
	{"Ambala", "Gurugram ", "Faridabad", "Bhiwani", "Fatehabad"},
	{"Marathahalli", "BTM Layout", "Banashankari", "Domlur", " Rajajinagar"},
	{"Bhopura", "Mohan Nagar", "Ghaziabad", "Anand Vihar", "Noida"},
}

type Car struct {
	CarId     int32
	Available bool
	Model     string
}

var Carpool = []*Car{
	{111, false, "Hyundai"}, {112, false, "Sedan"}, {113, true, "Suzuki"}, {114, true, "Hatchback"}, {115, true, "Hyundai Xcent"}, {116, true, "Suzuki"}, {118, true, "Suzuki"}, {117, true, "Suzuki"}, {118, true, "Hatchback"}, {119, true, "Hyundai Xcent"},
	{120, true, "Sedan"},
}
