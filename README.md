#create an http.Handler that will look at the path of any incoming web request and determine if it should redirect the user to a new page.
For instance, if we have a redirect setup for /dogs to https://www.somesite.com/a-story-about-dogs we would look for any incoming web requests with the path /dogs and redirect them.
In main.go
```
         yaml := `
          - path: /urlshort
            url: https://github.com/anushkamittal20/urlshort
          - path: /link
            url: https://github.com/anushkamittal20/link
          - path: /urlshort-final
            url: https:`
  ```          
One can add their own urls by changing    ``` 
                                                  - path: /shortpathname
                                                   url: url_to_be redirected_to 
                                           ```
or in main.go

```  
       pathsToUrls := map[string]string{
              "/w3":        "https://www.w3schools.com/html/",
              "/html-wiki": "https://en.wikipedia.org/wiki/HTML",
            } 
```
            
add by adding elements to the mapp in form :    
```    
"/shortpathname": "url_to_be redirected_to",  
```
To execute clone the repository and 
```
go run main/main.go
```
in the terminal
