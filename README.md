
Description
It is a sample project showing the use of Go Lang in creating restfull apis.This project demonstrated all the operations performed during CRUD .We have taken a sample project of Movie listing and also we created a model instead of instantiating a database connection.

Model of Movies that we have used
{MovieName:"Test",MovieRating:"5",MovieGenre:"testing"},
{MovieName:"Test2",MovieRating:"4",MovieGenre:"testing2"},
{MovieName:"Test3",MovieRating:"3",MovieGenre:"testing3"}

To read all the movies from the model
API : http://localhost:8000/movie

To create a new Movie object. This movie object is needded to be inserted in the body and needed to be sent with API
API : http://localhost:8000/movie/createMovie
{
    "MovieName":"Test4",
    "MovieRating":"2",
    "MovieGenre":"testing4"
}

To Update the rating of the movie on the basis of genre we will pass the name of the movie as a query parameter \n
API :  http://localhost:8000/movie/updateMovie?movie_name=Test method=PUT \n
{\n
	"MovieName":"Test4",\n
	"MovieRating":"3",\n
	"MovieGenre":"testing4"\n
}

To perform the deletion successfully we will be requiring the movie name as the deletion is being performed on the basis of Movie Name
API : http://localhost:8000/movie/deleteMovie?movie_name=Test2

All the Endpoints are listed below
 All the api endpoints require postman  
 http://localhost:8000/movie   method=GET
 http://localhost:8000/movie/createMovie{it will require postman} method=POST
 http://localhost:8000/movie/updateMovie?movie_name=Test method=PUT
 http://localhost:8000/movie/deleteMovie?movie_name=Test2 method=DELETE
 
 //These two are Filter endpoints 
 http://localhost:8000/movie/filter/genre?genre=testing method=GET
 http://localhost:8000/movie/filter/rating/2 method=GET

