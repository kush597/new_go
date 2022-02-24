										INTRODUCTION TO FUZZING


Fuzzing :Fuzzing is a testing technique with randomized inputs that is used to find problematic edge cases or security problems in code that accepts user input.

Coverage guided fuzzing: (also known as greybox fuzzing) uses program instrumentation to trace the code coverage reached by each input fed to a fuzz target. Fuzzing engines use this information to make informed decisions about which inputs to mutate to maximize coverage.


Requirements
Below are rules that fuzz tests must follow.

A fuzz test must be a function named like FuzzXxx, which accepts only a *testing.F, and has no return value.
Fuzz tests must be in *_test.go files to run.
A fuzz target must be a method call to (*testing.F).Fuzz which accepts a *testing.T as the first parameter, followed by the fuzzing arguments. There is no return value.
There must be exactly one fuzz target per fuzz test.
All seed corpus entries must have types which are identical to the fuzzing arguments, in the same order. This is true for calls to (*testing.F).Add and any corpus files in the testdata/fuzz directory of the fuzz test.
The fuzzing arguments can only be the following types:
string, []byte
int, int8, int16, int32/rune, int64
uint, uint8/byte, uint16, uint32, uint64
float32, float64
bool


Running fuzz test:

To run fuzz test create file and test file in same directory and there are two ways to run the test

First mode is through "go test" Command in command promt
Second mode is through "go test -fuzz=FuzzTestName" in Command promt


Packages that need to be installed:
$ go get -d github.com/dvyukov/go-fuzz-corpus or $ go get  github.com/dvyukov/go-fuzz-corpus



