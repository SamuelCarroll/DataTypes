package dataTypes

//Data Basic data type that can hold any supervised learner type
type Data struct {
	Class        int
	UID          string
	FeatureSlice []interface{}
}
