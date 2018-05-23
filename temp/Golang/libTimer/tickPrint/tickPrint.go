package tickPrint

import (
	"log"
	"sync"
	"time"
)

//获取ticket的channel
var ticketCh <-chan struct{}
var once sync.Once

func init() {
	log.SetFlags(log.Lmicroseconds)
	makeTicketCh(time.Second)
}

//按照参数指定的时间，生成获取ticket的实际channel
func makeTicketCh(d time.Duration) {
	//会有多处import本库，使用once可以保证ticketCh不被改写。
	once.Do(
		func() {
			result := make(chan struct{})

			go func() {
				tc := time.Tick(d)
				log.Println("tick is ready")
				for {
					//此处的运行逻辑有一个小小的瑕疵
					//只能保证任意连续三个出票的时间点的时间长度，大于两个访问周期。
					//如果result<-struct{}{}被阻塞大于一个时间周期后，后一张票就有可能跟着一起发出。
					result <- struct{}{}
					<-tc
				}
			}()

			ticketCh = result
		})
}

func waitTicket() {
	<-ticketCh
}

//X 会输出X=i
func X(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("\t准备输出, X=", i)
	waitTicket()
	log.Println("\tX=", i)
}

//Y 会输出Y=i
func Y(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("准备输出, Y=", i)
	waitTicket()
	log.Println("Y=", i)
}
