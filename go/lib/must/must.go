package must

import "log"

func Value[T any](out T, err error) T {
	if err != nil {
		log.Panicln(err)
	}
	return out
}
func Value2[T1 any, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	if err != nil {
		log.Panicln(err)
	}
	return t1, t2
}
func Nada(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
