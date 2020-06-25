package local

// func (l *local) Subscribe(route string) (chan *listener.Msg, string, error) {
// 	l.lock.Lock()
// 	defer l.lock.Unlock()

// 	var id, err = uuid.NewRandom()
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	var ch = make(chan *listener.Msg, 1000)

// 	lis, err := generic_lis.New(id.String(), route, func(msg *listener.Msg) error {
// 		ch <- msg

// 		return nil
// 	})

// 	if err != nil {
// 		return nil, "", err
// 	}

// 	err = l.r.Attach(lis)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	return ch, id.String(), nil
// }

// func (l *local) Close() error {

// 	l.r.Close()
// 	l.cancel()
// 	l.b.Close()

// 	return nil
// }
