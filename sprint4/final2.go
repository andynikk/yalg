package sprint4

import (
	"strconv"
	"strings"
)

type pair struct {
	key   int
	value int
}

type HashTable struct {
	size    int
	buckets [][]pair
}

func newHashTable(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([][]pair, size),
	}
}

func hashFinal2(key string) int {
	h := 0
	for i := 0; i < len(key); i++ {
		h = (h << 5) + h + int(key[i])
	}
	return h
}

func (ht *HashTable) put(key string, value int) {
	index := hashFinal2(key) % ht.size
	for i, p := range ht.buckets[index] {
		if p.key == hashFinal2(key) {
			ht.buckets[index][i] = pair{hashFinal2(key), value}
			return
		}
	}
	ht.buckets[index] = append(ht.buckets[index], pair{hashFinal2(key), value})
}

func (ht *HashTable) get(key string) (int, bool) {
	index := hashFinal2(key) % ht.size
	for _, p := range ht.buckets[index] {
		if p.key == hashFinal2(key) {
			return p.value, true
		}
	}
	return 0, false
}

func (ht *HashTable) delete(key string) (int, bool) {
	index := hashFinal2(key) % ht.size
	for i, p := range ht.buckets[index] {
		if p.key == hashFinal2(key) {
			//ht.buckets[index] = append(ht.buckets[index][:i], ht.buckets[index][i+1:]...)
			ht.buckets[index][i], ht.buckets[index][len(ht.buckets[index])-1] = ht.buckets[index][len(ht.buckets[index])-1], ht.buckets[index][i]
			ht.buckets[index] = ht.buckets[index][:len(ht.buckets[index])-1]

			return p.value, true
		}
	}

	return -1, false
}

func Final2(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	commands := lines[1:]
	result := []string{}
	//q := Queue10{}

	ht := newHashTable(1<<12 + 1)
	for c := 0; c < len(commands); c++ {
		command := strings.Split(commands[c], " ")
		if command[0] == "" {
			continue
		}

		switch command[0] {
		case "get":
			v, found := ht.get(command[1])
			if !found {
				result = append(result, "None")
				continue
			}
			result = append(result, strconv.Itoa(v))
		case "put":
			//k, _ := strconv.Atoi(command[1])
			v, _ := strconv.Atoi(command[2])

			ht.put(command[1], v)
		case "delete":
			v, found := ht.delete(command[1])
			if !found {
				result = append(result, "None")
				continue
			}
			result = append(result, strconv.Itoa(v))
		default:
			continue
		}
	}

	return strings.Join(result, "\n") + "\n"
}
