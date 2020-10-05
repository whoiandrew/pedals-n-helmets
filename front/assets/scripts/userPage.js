$(document).ready(function(){
    $("#modifyUser").click(function(){
        $.ajax({
            type: "POST",
            url: "http://localhost:8081/editUserInfo",
            data: {
                firstname: $("#firstname").val(),
                lastname: $("#lastname").val(),
                age: $("#age").val(),
                bicycle: $("#bicycle").val()
            },
            success: function(){
                $("#inputStatus")
                    .text("Successfully updated")
                    .css("color", "green");
                setTimeout(function(){
                    $("#inputStatus").css("visibility", "hidden");
                }, 2000); 
            }
        });           
    });
});