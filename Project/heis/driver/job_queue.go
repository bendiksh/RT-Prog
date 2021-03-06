package driver
import(
	"sync"
)

type Job_queue_t struct {
	sync.Mutex
	Floor []int
	Dir []int
	head uint8
	tail uint8
	count uint8
}


func Make_job_queue() *Job_queue_t{
	return &Job_queue_t{Floor: make([]int,256), Dir: make([]int,256), head: 0, tail: 0, count: 0}
}

func Pop(queue *Job_queue_t) (int, int){ //Floor, Dir
	queue.Lock() 
	if(queue.count > 0){
		queue.head += 1
		queue.count -= 1
	}
	queue.Unlock()
	return queue.Floor[queue.head], queue.Dir[queue.head]
}

func Push(queue *Job_queue_t, Floor int, Dir int) {
	queue.Lock() 
	queue.count += 1
	queue.tail += 1
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
		Push(q, i*2, i*3)
	}
	for i := 0; i < 10; i++ {
		p, r := Pop(q)
		g := Event_t{p,r}
		fmt.Println(p)
		fmt.Println(r)
		fmt.Println(g)
	}
}*/
