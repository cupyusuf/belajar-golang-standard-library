package main

import "container/list"

func main() {
	data := list.New()
	data.PushBack("Yusuf")
	data.PushBack("Supriadi")
	data.PushBack("Cimahi")

	for e := data.Front(); e != nil; e = e.Next() {
		println(e.Value.(string))
	}
}
