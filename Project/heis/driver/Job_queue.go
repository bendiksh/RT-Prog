package driver
import(
	"sync"
	"fmt"
)

type Job_queue_t struct {
	sync.Mutex
	IP []string
	Type []int
	Floor []int
	Dir []int
	head uint8
	tail uint8
	count uint8
}

func Make_job_queue() *Job_queue_t{
	return &Job_queue_t{IP: make([]string,256), Type: make([]int,256), Floor: make([]int,256), Dir: make([]int,256) head: 0, tail: 0, count: 0}
}

func Pop(queue *Job_queue_t) (string, int, int, int){ //IP, Button_ID
	queue.Lock() 
	if(queue.count > 0){
		queue.head += 1
		queue.count -= 1
	}
	queue.Unlock()
	return queue.IP[queue.head], queue.Type[queue.head], queue.Floor[queue.head]
}

func Push(queue *Job_queue_t, IP string, Type int, Floor int, Dir int) {
	queue.Lock() 
	queue.count += 1
	queue.tail += 1
	queue.IP[queue.tail] = IP
	queue.Type[queue.tail] = Type
	queue.Floor[queue.tail] = Floor
	queue.Dir[queue.tail] = Dir
	queue.Unlock()
}

func Empty(queue *Job_queue_t) bool {
	if(queue.count == 0){
		return true
	}
	return false
}

/*func main(){
	q := Make_job_queue()
	for i := 0; i < 10; i++ {
		Push(q, "a", i*2, i*3)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(Pop(q))
	}
}*/