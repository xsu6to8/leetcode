type Foo struct {
    firstDone  chan struct{}
    secondDone chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
        firstDone:  make(chan struct{}),
        secondDone: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
    printFirst()

	close(f.firstDone)
}

func (f *Foo) Second(printSecond func()) {
	/// Do not change this line
    <-f.firstDone

	printSecond()

    close(f.secondDone)
}

func (f *Foo) Third(printThird func()) {
	// Do not change this line
	<-f.secondDone

    printThird()
}