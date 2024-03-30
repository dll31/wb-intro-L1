Чем отличаются RWMutex от Mutex?

from https://pkg.go.dev/sync#RWMutex
A RWMutex is a reader/writer mutual exclusion lock. The lock can be held by an arbitrary number of readers or a single writer.
То есть неограниченное число читаталей и один редактирующий.