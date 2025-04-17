package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

var globalCloser = New()

func Add(fs ...func() error) {
	globalCloser.Add(fs...)
}

func CloseAll() {
	globalCloser.CloseAll()
}

func Wait() {
	globalCloser.Wait()
}

type Closer struct {
	mu    *sync.Mutex
	once  *sync.Once
	wg    *sync.WaitGroup
	done  chan struct{}
	funcs []func() error
}

func New(sig ...os.Signal) *Closer {
	c := &Closer{
		mu:   &sync.Mutex{},
		once: &sync.Once{},
		wg:   &sync.WaitGroup{},
		done: make(chan struct{}),
	}
	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
			c.CloseAll()
		}()
	}
	return c
}

func (c *Closer) Add(fs ...func() error) {
	c.mu.Lock()
	c.funcs = append(c.funcs, fs...)
	c.mu.Unlock()
}

func (c *Closer) CloseAll() {
	c.once.Do(func() {
		defer close(c.done)

		c.mu.Lock()
		funcs := c.funcs
		c.funcs = nil
		c.mu.Unlock()

		errs := make(chan error, len(funcs))
		c.wg.Add(len(funcs))
		for _, f := range funcs {
			go func() {
				defer c.wg.Done()

				errs <- f()
			}()
		}

		go func() {
			c.wg.Wait()
			close(errs)
		}()

		for err := range errs {
			if err != nil {
				log.Println("error returned from Closer")
			}
		}
	})
}

func (c *Closer) Wait() {
	<-c.done
}
