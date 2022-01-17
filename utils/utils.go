package utils

func DetectError(err error) interface{} {
	if err != nil {
		return err.Error()
	}
	return nil
}
