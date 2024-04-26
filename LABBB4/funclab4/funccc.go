package funclab4

import "fmt"

func (r *RingBufer) IsEmpty() bool {
	return r.count == 0
}

func (r *RingBufer) IsFull() bool {
	return r.count == r.size
}

// запись в буфер если есть свободные ячейки
func (r *RingBufer) Writer(data, id int) {

	r.mu.Lock()
	defer r.mu.Unlock()

	for r.IsFull() {
		r.notFull.Wait()
	}

	if r.IsFull() {
		fmt.Println("Buf IsFull")
	}
	r.buf[r.writePoint] = data
	r.writePoint = (r.writePoint + 1) % r.size
	fmt.Println(id, "Запись в буфер", data, r.buf)
	r.count++

	r.notEmpty.Broadcast()

}

// Чтение из буфера если есть заполненные ячейки
func (r *RingBufer) Reader(id int) {

	r.mu.Lock()
	defer r.mu.Unlock()

	for r.IsEmpty() {
		r.notEmpty.Wait()
	}

	if r.IsEmpty() {
		fmt.Println("Buf IsEmpty")
	}
	data := r.buf[r.readPoint]
	fmt.Println("Прочитано", data, id)
	r.readPoint = (r.readPoint + 1) % r.size
	r.count--

	r.notFull.Broadcast()

}
