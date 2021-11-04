package storage

var data = 0

func GetData() int {
	return data
}

func SetData(num int) int {
	data = num
	return data
}
