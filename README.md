# GoLang-REST-api-demonstration #
<br />
It is a sample project showing the use of Go Lang in creating restfull apis.This project demonstrated all the operations performed during CRUD .We have taken a sample project of Movie listing and also we created a model instead of instantiating a database connection.
</br>
**Port Used in all the api is 8000**

<br />Model of Movies that we have used<br />
{MovieName:"Test",MovieRating:"5",MovieGenre:"testing"},<br />
{MovieName:"Test2",MovieRating:"4",MovieGenre:"testing2"},<br />
{MovieName:"Test3",MovieRating:"3",MovieGenre:"testing3"}<br />
<br />
**To read all the movies from the model<br />
API : http://localhost:8000/movie**<br />
<br />
**To create a new Movie object. This movie object is needded to be inserted in the body and needed to be sent with API<br />
API : http://localhost:8000/movie/createMovie**<br />
{<br />
    "MovieName":"Test4",<br />
    "MovieRating":"2",<br />
    "MovieGenre":"testing4"<br />
}<br />
<br />
**To Update the rating of the movie on the basis of genre we will pass the name of the movie as a query parameter <br />
API :  http://localhost:8000/movie/updateMovie?movie_name=Test** <br />
{<br />
	"MovieName":"Test4",<br />
	"MovieRating":"3",<br />
	"MovieGenre":"testing4"<br />
}<br />
<br />
**To perform the deletion successfully we will be requiring the movie name as the deletion is being performed on the basis of Movie Name<br />
API : http://localhost:8000/movie/deleteMovie?movie_name=Test2**<br />
<br />
**All the Endpoints are listed below<br />
 <br />All the api endpoints require postman  
 <br />http://localhost:8000/movie   method=GET
 <br />http://localhost:8000/movie/createMovie{it will require postman} method=POST
 <br />http://localhost:8000/movie/updateMovie?movie_name=Test method=PUT
 <br />http://localhost:8000/movie/deleteMovie?movie_name=Test2 method=DELETE
 <br />
 <br />//These two are Filter endpoints 
 <br />http://localhost:8000/movie/filter/genre?genre=testing method=GET
 <br />http://localhost:8000/movie/filter/rating/2 method=GET
**
