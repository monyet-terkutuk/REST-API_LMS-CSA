package helper

func IfErrReturnBasic(err error) (any, error) {
	if err != nil {
		return nil, err
	}
	return nil, nil
}
