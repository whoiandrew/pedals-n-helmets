$(document).ready(function(){
    $('#articles .article').sort(function(a,b) {
        return new Date(b.dataset.date) - new Date(a.dataset.date);
    }).appendTo('#articles');

    $("#loginBtn").click(function(){
        $("#loginModal").show();
    });
    $("#closeLogin").click(function(){
        $("#loginModal").hide();
    });
    $("#registerBtn").click(function(){
        $("#registerModal").show();
    });
    $("#closeReg").click(function(){
        $("#registerModal").hide();
    });
    $(".deleteArticleButton").click(function(){
        const id = $(this).data("id");
        const article = $(this).parent().parent();
        $.ajax({
            type: "POST",
            url: "http://localhost:8081/deleteArticle",
            data: {
                id: id
            },
            success: function(){
                article.remove()
            }
        });
    });
    $(".editArticleButton").click(function(){
        $("#editModal").show();
        $("#submitEditionButton").attr("data-id", $(this).data("id"));
    });
    $("#closeEdit").click(function(){
        $("#editModal").hide();
    });

    $("#submitEditionButton").click(function(){
        const id = $(this).data("id");
        const content = $(`#edition`).val();
        $.ajax({
            type: "POST",
            url: "http://localhost:8081/editArticle/" + id,
            data: {
                content: content
            },
            success: function(){
                $("#content"+id).text(content)
            }
        });
    });

    $("#addArticleButton").click(function(){
        const title = $("#titleInput").val()
        if (!title.replace(/\s/g, '').length){
            alert("Empty title!")
        }
        else {
            $.ajax({
                type: "POST",
                url: "http://localhost:8081/addArticle",
                data: {
                    author: $("#loggedUserNickname").attr("data-nickname"),
                    title: $("#titleInput").val(),
                    content: $("#contentInput").val()
                },
                dataType: "json",
                success: function(data) {
                    $("#articles").prepend(`<div class="article">
                                                <a href="/article/${data.id}"><h2>${$("#titleInput").val()}</h2></a>
                                                <p>${data.prettyTime}</p>
                                                <p>created by ${$("#loggedUserNickname").attr("data-nickname")}</p>
                                                <p>${$("#contentInput").val()}</p>
                                                <div class="editArticle">
                                                    <input class="editArticleButton" type="button" value="editArticle" data-id=${data.id}>
                                                    <input class="deleteArticleButton" type="button" value="deleteArticle" data-id=${data.id}>
                                                </div>
                                            </div>`);
                }
            });           
        }
    });
    function rateAJAXReq(rateType, id){
        $.ajax({
                type: "POST",
                url: "http://localhost:8081/rateArticle",
                data: {
                    id: id,
                    rate: rateType
                }
            });          
    }
    $(".likeButton").click(function(){
        rateAJAXReq("like", $(this).attr("data-id"));
    });
    $(".dislikeButton").click(function(){
        rateAJAXReq("dislike", $(this).attr("data-id"));
    });

    $("#pwd1").change(function(){
        submitRegController();
    });
    $("#pwd2").change(function(){
        submitRegController();
    });
    $("#regName").change(function(){
        submitRegController();
    });

    function submitRegController(){
        const match = $('#pwd1').val() == $('#pwd2').val();
        const name = $("#regName").val();
        if (match && name.length > 3) {
            $("#submitReg").removeAttr("disabled")
        }
        else {
            $("#submitReg").attr("disabled", "disabled")
        }
    };

    $("#loginName").change(function(){
        submitLoginController();
    });

    $("#loginPwd").change(function(){
        submitLoginController();
    });

    function submitLoginController(){
        const name = $("#loginName").val();
        const pwd = $("#loginPwd").val();
        if (name.length > 3 && pwd.length > 5)  {
            $("#submitLogin").removeAttr("disabled")
        }
        else {
            $("#submitLogin").attr("disabled", "disabled")
        }

    };

    
});