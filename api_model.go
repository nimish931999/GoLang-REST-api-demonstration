package main
import ("fmt"
"net/http"
"log"
"github.com/gorilla/mux"
"encoding/json"
)

type Movie struct{
	 MovieName string `json:"MovieName"`
	 MovieRating string `json:"MovieRating"`
	 MovieGenre string `json:"MovieGenre"`
}

var movie []Movie
func getAll(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(movie)
}

func createMovie(w http.ResponseWriter,r *http.Request){
	var mov Movie
	_ = json.NewDecoder(r.Body).Decode(&mov)
	movie = append(movie,mov)
	json.NewEncoder(w).Encode(mov)
}
func getByRating(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	url:=r.URL.Query()
	id := url["id"]
	//url.set("key","value")
	fmt.Println(id)
	//r.URL.RawQuery = url.Encode()
	for _,mov := range movie{
		if mov.MovieRating >= params["rating"]{
			json.NewEncoder(w).Encode(mov)
			
		}else{
			return
		}
	}
	//json.NewEncoder(w).Encode(&Movie{})
}

func getMovieByGenre (w http.ResponseWriter,r *http.Request){
	url := r.URL.Query()
	mv_gen := url["genre"]
	for _,mov:=range movie{
		if mov.MovieGenre == mv_gen[0]{
			json.NewEncoder(w).Encode(mov)
		}
	}
}

func updateMovie(w http.ResponseWriter,r *http.Request){
	url:= r.URL.Query()
	mv_nm := url["movie_name"][0]
	for i,mov := range movie{
		if mov.MovieName == mv_nm {
			movie = append(movie[:i],movie[i+1:]...)
			var mov_update Movie
			_ =json.NewDecoder(r.Body).Decode(&mov_update)
			movie = append(movie,mov_update)
			json.NewEncoder(w).Encode(mov_update)
		}
	}
}

func deleteMovie(w http.ResponseWriter,r *http.Request){
	url := r.URL.Query()
	mv_nm := url["movie_name"][0]
	for i,mov := range movie{
		if mov.MovieName==mv_nm{
			movie =append(movie[:i],movie[i+1:]...)
			json.NewEncoder(w).Encode(&movie)
		}
	}
}

func api_start(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w,"Welcome to sample route for Go Lang restful Api")
}

func handleRouter(){
	r:= mux.NewRouter()
	r.HandleFunc("/",api_start).Methods("GET")
	r.HandleFunc("/movie",getAll).Methods("GET")
	r.HandleFunc("/movie/createMovie",createMovie).Methods("POST")
	r.HandleFunc("/movie/updateMovie",updateMovie).Methods("PUT")
	r.HandleFunc("/movie/deleteMovie",deleteMovie).Methods("DELETE")
	//The below api endpoints are filters to filter the movie on the basis of movie rating and movie genre//
	r.HandleFunc("/movie/filter/rating/{rating}",getByRating).Methods("GET")
	r.HandleFunc("/movie/filter/genre",getMovieByGenre).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000",r))
}
func main(){
	movie = append(movie,Movie{MovieName:"Test",MovieRating:"5",MovieGenre:"testing"},Movie{MovieName:"Test2",MovieRating:"4",MovieGenre:"testing2"},Movie{MovieName:"Test3",MovieRating:"3",MovieGenre:"testing"})
	fmt.Println(" All the api endpoints require postman  \n http://localhost:8000/movie \n http://localhost:8000/movie/createMovie{it will require postman} \n http://localhost:8000/movie/updateMovie?movie_name=Test \n http://localhost:8000/movie/deleteMovie?movie_name=Test2")
	fmt.Println(" http://localhost:8000/movie/filter/genre?genre=testing \n http://localhost:8000/movie/filter/rating/2")
	handleRouter()
}
