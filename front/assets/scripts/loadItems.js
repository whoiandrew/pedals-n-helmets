function deleteUser(userId){
    const req = new XMLHttpRequest();
    req.open("POST", `http://localhost:8081/deleteUser/${userId}`, true);
    req.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    req.send();
};



