package main

<<<<<<< HEAD

func main() {

}
=======
import (
	"Editor_NaN/controller"
	"net/http"
)




func main (){

	mc := controller.InitMainController()

	http.HandleFunc("/", mc.Index)
	http.ListenAndServe(":8080", nil)
}

>>>>>>> fdfe077cb069498fefae3f385a9239ea9a7f1358
