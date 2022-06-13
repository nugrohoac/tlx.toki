package astro

func BubbleSort(items []int) []int {
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[i] {
				temp := items[j]

				items[j] = items[i]
				items[i] = temp
			}
		}
	}

	return items
}

func MainSort(items []int32) []int32 {
	mapNumberExist := map[int32]struct {
		exist bool
		count int
	}{}

	for i := 0; i < len(items); i++ {
		if mapNumberExist[items[i]].exist {
			mapNumberExist[items[i]] = struct {
				exist bool
				count int
			}{exist: true, count: mapNumberExist[items[i]].count + 1}

			continue
		}

		mapNumberExist[items[i]] = struct {
			exist bool
			count int
		}{exist: true, count: 1}
	}

	var holders []struct {
		freq  int
		value int32
	}

	for k, v := range mapNumberExist {
		holders = append(holders, struct {
			freq  int
			value int32
		}{freq: v.count, value: k})
	}

	for i := 0; i < len(holders); i++ {
		for j := i + 1; j < len(holders); j++ {
			if holders[j].freq < holders[i].freq {
				temp := struct {
					freq  int
					value int32
				}{freq: holders[j].freq, value: holders[j].value}

				holders[j] = holders[i]
				holders[i] = temp

				continue
			}

			if holders[j].freq == holders[i].freq {
				if holders[j].value < holders[i].value {
					temp := struct {
						freq  int
						value int32
					}{freq: holders[j].freq, value: holders[j].value}

					holders[j] = holders[i]
					holders[i] = temp
				}
			}
		}
	}

	var result []int32

	for _, holder := range holders {
		for i := 0; i < holder.freq; i++ {
			result = append(result, holder.value)
		}
	}

	return result
}
