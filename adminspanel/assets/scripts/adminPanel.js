$(document).ready(function(){
    $(".deleteUser").click(function(){
        const elem = $(this);
        const id = elem.val()
        $.ajax({
            type: "POST",
            url: "http://localhost:8083/deleteUser/"+id,
            success: function(){
                elem.parent().remove();
            }
        });
    });
});