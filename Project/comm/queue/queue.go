package queue
import(
	"sync"
)

type Btn_queue_t struct {
	sync.Mutex
	IP []string
	Button_id []string
	head uint8
	tail uint8
	count uint8
}

func Make_btn_queue() *Btn_queue_t{
	return &Btn_queue_t{IP: make([]string,256), Button_id: make([]string,256), head: 0, tail: 0, count: 0}
}

func Pop(queue *Btn_queue_t) (string, string){ //IP, Button_ID
	queue.Lock() 
	if(queue.count > 0){
		queue.head += 1
		queue.count -= 1
	}
	queue.Unlock()
	return queue.IP[queue.head], queue.Button_id[queue.head]
}

func Push(queue *Btn_queue_t, IP string, Button_id string) {
	queue.Lock() 
	queue.count += 1
	queue.tail += 1
	queue.IP[queue.tail] = IP
	queue.Button_id[queue.tail] = Button_id
	queue.Unlock()
}

func Empty(queue *Btn_queue_t) bool {
	if(queue.count == 0){
		return true
	}
	return false
}