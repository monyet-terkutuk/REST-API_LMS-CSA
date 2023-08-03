package helper

func IfErrReturnBasic(err error) error {
	if err != nil {
		return err
	}
	return nil
}
