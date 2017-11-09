package main

func main() {
	m := make(map[string]string)

	m[""] = ""

	v, ok := m[""]

	println(m)
	println(len(m))
	println(v)
	println(ok)
}
