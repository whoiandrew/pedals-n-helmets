$(document).ready(function(){
    $("#addCommentButton").on("click", function(){
        $.ajax({
            type: "POST",
            url: "http://localhost:8081/addComment",
            data: {
                articleId: $("#articleId").val(),
                author: $("#author").val(),
                content: $("#content").val()
            },
            dataType: "json",
            success: function(data) {
                $("#comments").prepend(`<div><p>${data.author}</p><p>${data.content}</p><p>${data.prettyTime}</p></div>`);
            }
        });  
                 
    });
});