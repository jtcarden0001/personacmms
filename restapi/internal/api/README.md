The HTTP status codes of this application will follow standard REST guidelines to include:
200 – OK
Everything is working

201 – CREATED
A new resource has been created

204 – NO CONTENT
The resource was successfully deleted, no response body

304 – NOT MODIFIED
The date returned is cached data (data has not changed)

400 – BAD REQUEST
The request was invalid or cannot be served. The exact error should be explained in the error payload. E.g. „The JSON is not valid “.

401 – UNATHORIZED
The request requires user authentication.

403 – FORBIDDEN
The server understood the request but is refusing it or the access is not allowed.

404 – NOT FOUND
There is no resource behind the URI.

500 – INTERNAL SERVER ERROR
API developers should avoid this error. If an error occurs in the global catch blog, the stack trace should be logged and not returned as a response.