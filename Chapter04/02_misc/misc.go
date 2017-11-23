package main



//import "theirpkg"

//func MyFunction(t *theirpkg.AType)
//
//func MyOtherFunction(i theirpkg.AnInterface)



type errorBehavior interface {
	Retryable() bool
}


func IsRetryable(err error) bool {
	eb, ok := err.(errorBehavior)
	return ok && eb.Retryable()
}


//type writeFlusher interface {
//     io.Writer
//     http.Flusher
//}





func main() {

}

/*

- Foot, paddle!
- Foot, paddle!
- Foot, paddle!
- Bill, eat a bug!
- Foot, paddle!
- Foot, paddle!
- Bill, eat a bug!
----------------------
Ponds Processed:
{BugSupply:1 StrokesRequired:3}
{BugSupply:1 StrokesRequired:2}
Strokes remaining: 2
----------------------



- Foot, paddle!
- Foot, paddle!
- Foot, paddle!


2017/05/12 19:11:51 Our duck died!
exit status 1




 */