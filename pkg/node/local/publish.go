package local

// // func (l *local) Publish(route string) (chan<- []byte, <-chan error) {
// func (l *Local) Publish(route string) (chan<- []byte, error) {
// 	l.lock.Lock()
// 	defer l.lock.Unlock()

// 	var (
// 		ch = make(chan []byte, 1000)
// 		// errChan = make(chan error)
// 	)

// 	go func() {
// 		for {
// 			select {
// 			case <-l.ctx.Done():
// 				return

// 			case msg := <-ch:
// 				var err = l.b.Insert(&listener.Msg{
// 					Route:    route,
// 					Contents: msg,
// 				})

// 				if err != nil {
// 					log.Fatalln("err l.b.Insert:", err)
// 				}
// 			}
// 		}
// 	}()

// 	return ch, nil
// }
