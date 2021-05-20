
function login(){
    
    let account = document.getElementsByName("account")
    let pwd = document.getElementsByTagName("password")

    let xhr = new XMLHttpRequest();
    xhr.open("GET", "http://localhost:8080/login")
    xhr.onreadystatechange = function(){
        alert("xhr state: ", xhr.readyState)
    }
    
    xhr.send(null)

}
