package no2

import "errors"

// ネストが深すぎて期待する実行の流れを見分けるのが困難
func NGJoin(s1, s2 string, max int) (string, error) {
	if s1 == "" {
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
	return "", nil
}

// 早期リターンを活用してネストを解消
func Solution(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}

	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}

	return concat, nil
}

func concatenate(s1, s2 string) (string, error) {
	return s1 + s2, nil
}
