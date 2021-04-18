package versions

const Version = "1.0"
type newOrganisationStruct struct{
	Version string
	Name string `bson:"name"`
	Desc string `bson:"desc"`
	Email string `bson:"email"`
	Contact string `bson:"email"`
	Image string `bson:"image"`
	Video string `bson:"video"`
}

