package goutils

const (
	MERGE_OPTION_OVERRIDE = 1
	MERGE_OPTION_DISCARD  = 2
)

func MergeStringMap(m1, m2 map[string]string) map[string]string {
	result := map[string]string{}

	for k, v := range m1 {
		result[k] = v
	}

	for k, v := range m2 {
		result[k] = v
	}
	return result
}
