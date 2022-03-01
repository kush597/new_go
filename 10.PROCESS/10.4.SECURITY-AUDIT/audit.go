
	
package main
 
import (
    "fmt"
    "reflect"
    "unsafe"
)
 
func Typelinks() (sections []unsafe.Pointer, offset [][]int32) {
    return typelinks()
}
 
//go:linkname typelinks reflect.typelinks
func typelinks() (sections []unsafe.Pointer, offset [][]int32)
 
func Add(p unsafe.Pointer, x uintptr, whySafe string) unsafe.Pointer {
    return add(p, x, whySafe)
}
 
//go:linkname add reflect.add
func add(p unsafe.Pointer, x uintptr, whySafe string) unsafe.Pointer
 
func main() {
    sections, offsets := Typelinks()
    for i, base := range sections {
        for _, offset := range offsets[i] {
            typeAddr := Add(base, uintptr(offset), ";")
            typ := reflect.TypeOf(*(*interface{})(unsafe.Pointer(&typeAddr)))
            fmt.Println(typ)
        }
    }
}