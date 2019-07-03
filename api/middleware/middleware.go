package middleware

import (
	"fmt"

	"github.com/Curicaveri/proxy-app/api/repository"
	"github.com/kataras/iris"
)

var q []*repository.Queue
var lowPriority int
var mediumPriority int
var highPriority int

// InitQueue will initialize the queue
func InitQueue() {
	q = append(q, &repository.Queue{})
	lowPriority = 1
	mediumPriority = 2
	highPriority = 3
}

// ProxyMiddleware will handle the Queue
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}

	var repo repository.Repository
	repo = &repository.Queue{}
	for _, row := range repo.Read() {
		fmt.Println("Domain: ", row.Domain)
		fmt.Println("Weigth: ", row.Weigth)
		fmt.Println("Priority: ", row.Priority)
		if domain == row.Domain {
			priority, messageObj := prioritize(domain, row)
			enqueue(priority, messageObj)
		}
	}
	c.JSON(iris.Map{"Queue": q})
	c.Next()
}

func prioritize(domain string, queueObject *repository.Queue) (int, *repository.Queue) {
	priority := 0
	messageObj := &repository.Queue{}

	if queueObject.Priority < 5 && queueObject.Weigth < 5 {
		priority = lowPriority
	} else if (queueObject.Priority >= 5 && queueObject.Weigth < 5) || (queueObject.Priority < 5 && queueObject.Weigth >= 5) {
		priority = mediumPriority
	} else if queueObject.Priority >= 5 && queueObject.Weigth >= 5 {
		priority = highPriority
	}
	messageObj.Priority = queueObject.Priority
	messageObj.Weigth = queueObject.Weigth

	return priority, messageObj
}

func enqueue(priority int, messageObj *repository.Queue) {
	for ind, msg := range q {
		if msg.Priority < messageObj.Priority {
			tmp := msg
			remove(ind)
			q = append(q, messageObj)
			q = append(q, tmp)
		}
	}
}

func remove(i int) {
	q[i] = q[len(q)-1]
	q = q[:len(q)-1]
}
