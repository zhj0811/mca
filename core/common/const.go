package common

const (
	Success               = iota + 1 //1
	UserNameOrPasswordErr            //2
	TokenNilErr                      //3
	TokenInvalidErr                  //4
	RequestInfoErr                   //5
	InsertDBErr
	QueryDBErr
	UpdateDBErr
	DeleteDBErr
	FormatError
	LocalServiceErr
	RemoteServiceErr
	InvalidCryptoErr

	DBErr
	ExecErr
)
