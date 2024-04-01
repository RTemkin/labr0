package main

import (
	"fmt"
	"sync"
	"time"
)

type mapRedis struct {
	sync.Mutex
	maps    map[int]interface{}
	mapsExp map[int]time.Time
}

func NewMapRedis(i int, v interface{}) *mapRedis {
	return &mapRedis{
		maps: map[int]interface{}{
			i: v,
		},
		mapsExp: map[int]time.Time{},
	}
}

func (m *mapRedis) Get(key int) interface{} {
	m.Lock()
	defer m.Unlock()
	return m.maps[key]
}

func (m *mapRedis) Set(key int, val string) {
	m.Lock()
	defer m.Unlock()
	m.maps[key] = val
}

func (m *mapRedis) addDel(key int, val interface{}) {
	m.Lock()
	defer m.Unlock()
	m.maps[key] = val
	myTime := time.Now()
	m.mapsExp[key] = myTime.Add(time.Second * 2)

	//fmt.Println(re.mapsExp)
}

func initC(maps map[int]interface{}, mapsExp map[int]time.Time) map[int]interface{} {

	go func() {
		for key, val := range mapsExp {

			if val.Unix() <= time.Now().Unix() {
				delete(maps, key)
				delete(mapsExp, key)
			}
		}
	}()
	return maps
}

func generationMap(n int) map[int]interface{} {

	mapsGen := make(map[int]interface{}, n)
	for i := 2; i < n; i++ {
		i := i
		mapsGen[i] = i
	}

	return mapsGen
}

func main() {
	start := time.Now()
	//var wg sync.WaitGroup

	mr := NewMapRedis(1, "one")
	fmt.Println(mr)

	genMaps := generationMap(1000)

	for i, v := range genMaps {
		mr.addDel(i, v)
	}

	time.Sleep(2 * time.Second)

	initC(mr.maps, mr.mapsExp)

	time.Sleep(1 * time.Millisecond)

	fmt.Println(mr.maps)
	fmt.Println(mr.mapsExp)

	duration := time.Since(start)
	fmt.Println(duration)
}
