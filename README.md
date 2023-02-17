
<img src="/lll.jpg" width="800">

![build](https://github.com/eigenhombre/lll/actions/workflows/build.yml/badge.svg)

Experiments in LLVM, probably leading to some sort of simple,
low-level Lisp language implementation.

Right now I'm working on building a small compiler that parses
Lisp expressions to do integer math only.

Eventually, it might become something like a backend for [this other
Lisp](https://github.com/eigenhombre/l1/).

For the moment, see `compiler_test.go` to get an idea of the direction.

<!-- The following examples are autogenerated, do not change by hand! -->
<!-- BEGIN EXAMPLES -->

    
    $ go build .
    
    $ echo 42 | ./lll > answer.ll
    
    $ cat answer.ll
    declare void @_print_int(i32 %x)
    
    define i32 @main() {
    0:
    	call void @_print_int(i32 42)
    	ret i32 0
    }
    
    $ clang answer.ll _print.c -o answer
    
    $ ls -l answer
    -rwxr-xr-x  1 jacobsen  staff  49464 Feb 17 07:47 answer
    
    $ ./answer
    42
    
    
    
<!-- END EXAMPLES -->

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
