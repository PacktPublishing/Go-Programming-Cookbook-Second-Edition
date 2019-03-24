package pool

import "golang.org/x/crypto/bcrypt"

func hashWork(wr WorkRequest) WorkResponse {
	val, err := bcrypt.GenerateFromPassword(wr.Text, bcrypt.DefaultCost)
	return WorkResponse{
		Result: val,
		Err:    err,
		Wr:     wr,
	}
}

func compareWork(wr WorkRequest) WorkResponse {
	var matched bool
	err := bcrypt.CompareHashAndPassword(wr.Compare, wr.Text)
	if err == nil {
		matched = true
	}
	return WorkResponse{
		Matched: matched,
		Err:     err,
		Wr:      wr,
	}
}
