an attmept to get a bit more profficient in Golang by building a secure backend for a podcasting site. Essentially  the rest api will be consumed by a website and an app in the future. It will allow an admin to upload audio files (podcast) and administer your own podcast data. the app will just be for listening to podcasts. rating them etc. for the moment im using tls to secure communicaton and jwt tokens to verify admin users. Trying to make the code as clean as possible. contemplated using the echo framework as if helps make the routing look a bit tidier. this is something ill get back to. also could of probably included a "use case" layer as the routes.go file contains methods that have a bit of repeating code and are doing way to much. this is a work in progress.
