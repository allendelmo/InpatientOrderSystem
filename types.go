package main

import "time"

// TODO: deprecate when rest api refactor is finished
type User_OLD struct {
	Username string
	Password string
}

// TODO: deprecate when rest api refactor is finished
type Medication_Order_OLD struct {
	Order_Number     int32
	File_Number      int32
	Nurse_Name       string
	Ward             string
	Bed              string
	Medication       string
	UOM              string
	Request_time     time.Time
	Nurse_Remarks    string
	Status           string
	PHARMACY_REMARKS string
}
