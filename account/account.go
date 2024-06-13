package account

import "errors"

type Account struct {
	balance float64
	name    string
}

func New(name string) (a *Account, err error) {
	if name == "" {
		err := errors.New("account must have a name")

		return &Account{
			name:    "default",
			balance: 0.0,
		}, err
	}

	return &Account{
		name:    name,
		balance: 0.0,
	}, nil
}
