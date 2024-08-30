package synchro

func FilterOutId(maps map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{})

	for k, v := range maps {
		if k == "id" {
			continue
		}
		out[k] = v
	}

	return out
}
