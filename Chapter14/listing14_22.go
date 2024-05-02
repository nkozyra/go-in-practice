package main

// #include <stdio.h>
// void delimiter(int num, int limit) {
//	if (num < limit) {
//		printf(", ");
//		return;
//	}
// 	printf("\n");
// }
// void fizz_buzz(int limit) {
// for (int i = 0; i <= limit; i++) {
// 	if (i % 3 == 0 || i % 5 == 0) {
// 		printf("FizzBuzz");
// 	}
// 	else if (i % 3 == 0) {
// 		printf("Fizz");
// 	}
// 	else if (i % 5 == 0) {
// 		printf("Buzz");
// 	} else {
// 		printf("%d", i);
// }
// 	delimiter(i, limit);
// }
// }
import "C"

func main() {
	if _, err := C.fizz_buzz(15); err != nil {
		panic(err)
	}
}
