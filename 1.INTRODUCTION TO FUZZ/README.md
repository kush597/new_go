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


Crashes Triaging And Debugging:
This is a quick tutorial on how to fuzz a simple library using Dmitry Vyukov’s go-fuzz tool.

First, find a target. In this case, browsing /r/golang I saw a post announcing a package for decoding Adobe Swatch Exchange files. Parsing data is always a fraught endeavour, so it makes a good target for fuzzing.

Next, download the package.

$ go get github.com/arolek/ase
$ cd `go list -f '{{.Dir}}' github.com/arolek/ase`
$ git reset --hard b1bf7d7a70445821722b29395f07fcd13e940f8c

The third step is only needed if you want to play along at home. My fixes for the crashes have been merged in, so you’ll need to reset the git repository to just before that point.

Looking at the godoc link, we see

func Decode(r io.Reader) (ase ASE, err error)

which looks like a good entry point to fuzz.

In order to fuzz, we need two things

    a fuzzing function
    sample inputs

Finding sample inputs was pretty easy — the ase package ships with 3 in the samples directory. I also found a few more with Google, but this wasn’t necessary.

Next, write the fuzzing function. The function must have the signature

func Fuzz(data []byte) int

A return value of 0 from Fuzz() means the data wasn’t interesting — the parser detected an error. A return value of 1 means the data was parsed successfully, even though it had been modified. The fuzzer keeps these as more interesting inputs, since they appear to be valid.

Since our target function takes an io.Reader, but Fuzz provides us with a []byte, we can wrap it with bytes.NewReader. Here’s the complete fuzzing file, including the go-fuzz build tags. I put this into a file called ‘fuzz.go’.

// +build gofuzzpackage aseimport "bytes"func Fuzz(data []byte) int {
    if _, err := Decode(bytes.NewReader(data)); err != nil {
      return 0
    }
    return 1
}

If you don’t have go-fuzz already installed, do that now.

$ go get github.com/dvyukov/go-fuzz/go-fuzz
$ go get github.com/dvyukov/go-fuzz/go-fuzz-build

Next, build the package with go-fuzz:

$ go-fuzz-build github.com/arolek/ase

While this is building (it might take a while), create a work directory and put the sample files into the corpus. You can put this workdir right inside the github source directory for the package. We’re not going to be committing it.

$ mkdir -p workdir/corpus
$ cp samples/*.ase workdir/corpus

When go-fuzz-build finishes, it will have created a file called ‘ase-fuzz.zip’. This contains all the instrumented binaries go-fuzz is going to use while fuzzing.

Next, start the fuzzing process. We have to pass go-fuzz the path to the zip file it just created, and also the path to the workdir with the corpus.

$ go-fuzz -bin=ase-fuzz.zip -workdir=workdir

At this point, your machine will start to heat up. The fuzzing is mutating the samples and passing the resulting byte slices to our fuzzing function. If we reach new bits of the code with the mutated file, it’s added to the corpus directory. They’ll be named with the sha1 of the contents.

The fuzzer will start to print out log lines


Packages that need to be installed:
$ go get -d github.com/dvyukov/go-fuzz-corpus or $ go get  github.com/dvyukov/go-fuzz-corpus
$ go get -d github.com/dvyukov/go-fuzz or $ go get  github.com/dvyukov/go-fuzz
$ go get -d github.com/dvyukov/go-fuzz/go-fuzz or $ go get  github.com/dvyukov/go-fuzz/go-fuzz
$ go get -d github.com/dvyukov/go-fuzz/go-fuzz-build or $ go get  github.com/dvyukov/go-fuzz/go-fuzz-build


