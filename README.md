# Light Go Container

This is a very lightweight IoC base container written in what I consider a go centric manor, for use as a starting point in any go application

There will be additions to the functionality over time but it will remain a very light weight package

Given the go architecture is assuming that users will be able to manage and understand their own thread safety, I am making the same assumptions and allowing access to the pointers of the internals of the container, this is done for efficiency and should not be abused.